package mounts

import (
	"fmt"
	"path"
	"strings"
)

// OpenShell-aligned guest reserved roots (whaleshell control plane + OCI runtime mounts).
// Intentionally not a general Linux system-path denylist for host sources —
// host workspace secrets use DenyBasenames / ResolveWorkspace instead.
var (
	// ControlRoots are in-guest paths owned by whaleshell (must not be user-mounted over).
	ControlRoots = []string{
		"/whaleshell",
		"/etc/whaleshell",
		"/run/whaleshell",
		"/run/netns",
		"/var/run/netns",
	}
	// OCIRuntimeMountRoots must not be used as workspace or user bind targets.
	OCIRuntimeMountRoots = []string{
		"/proc", "/sys", "/dev",
	}
)

// ValidateContainerMountTarget rejects user mounts that overlap reserved guest paths.
func ValidateContainerMountTarget(target string) error {
	normalized, err := normalizeAbsoluteContainerPath(target, "mount target")
	if err != nil {
		return err
	}
	p := path.Clean(normalized)
	for _, reserved := range append(append([]string{}, ControlRoots...), OCIRuntimeMountRoots...) {
		if pathsOverlap(p, path.Clean(reserved)) {
			return fmt.Errorf("mount target %q conflicts with reserved path %q", target, reserved)
		}
	}
	return nil
}

// ValidateUploadDest rejects uploads into reserved control paths.
// Destinations under /workspace are always allowed.
func ValidateUploadDest(dest string) error {
	normalized, err := normalizeAbsoluteContainerPath(dest, "upload dest")
	if err != nil {
		return err
	}
	p := path.Clean(normalized)
	ws := path.Clean(WorkdirInContainer)
	if p == ws || strings.HasPrefix(p, ws+"/") {
		return nil
	}
	return ValidateContainerMountTarget(dest)
}

// PathsOverlap reports whether a and b are equal or one contains the other.
func PathsOverlap(a, b string) bool {
	return pathsOverlap(path.Clean(a), path.Clean(b))
}

func pathsOverlap(a, b string) bool {
	if a == b {
		return true
	}
	if a == "/" || b == "/" {
		return true
	}
	return strings.HasPrefix(a, b+"/") || strings.HasPrefix(b, a+"/")
}

func normalizeAbsoluteContainerPath(value, field string) (string, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return "", fmt.Errorf("%s must not be empty", field)
	}
	if value != strings.TrimSpace(value) {
		return "", fmt.Errorf("%s must not contain surrounding whitespace", field)
	}
	for _, r := range value {
		if r < 0x20 {
			return "", fmt.Errorf("%s must not contain control characters", field)
		}
	}
	if !strings.HasPrefix(value, "/") {
		return "", fmt.Errorf("%s must be an absolute container path", field)
	}
	segs := strings.Split(strings.TrimPrefix(value, "/"), "/")
	for i, s := range segs {
		if s == "." || s == ".." {
			return "", fmt.Errorf("%s must be normalized without '.' or '..' segments", field)
		}
		if s == "" && i < len(segs)-1 {
			return "", fmt.Errorf("%s must be normalized without empty path segments", field)
		}
	}
	normalized := strings.TrimRight(value, "/")
	if normalized == "" {
		return "", fmt.Errorf("%s must not be the container root", field)
	}
	return normalized, nil
}
