// Package mounts validates host bind mounts and guest mount targets.
//
// Host workspace sources use a deny-list (DenyBasenames / $HOME) unless --i-know.
// Guest targets use OpenShell-style reserved roots (ControlRoots) so users cannot
// overwrite /whaleshell control state — not a general Linux system-path denylist.
package mounts

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// WorkdirInContainer is where the host workspace is mounted.
const WorkdirInContainer = "/workspace"

// DenyBasenames are never auto-mounted as workspace (secrets / cloud CLIs).
// OpenShell relies more on --upload; whaleshell still bind-mounts workspace, so this
// host-side check remains as hardening.
var DenyBasenames = []string{
	".ssh", ".aws", ".gnupg", ".kube", ".docker", ".config",
	".cursor", ".codex", ".claude",
}

// ResolveWorkspace returns an absolute path and checks the deny-list.
// iKnow skips deny-list (still requires the path to exist as a directory).
func ResolveWorkspace(path string, iKnow bool) (string, error) {
	if strings.TrimSpace(path) == "" {
		path = "."
	}
	abs, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("workspace: %w", err)
	}
	fi, err := os.Stat(abs)
	if err != nil {
		return "", fmt.Errorf("workspace: %w", err)
	}
	if !fi.IsDir() {
		return "", fmt.Errorf("workspace: not a directory: %s", abs)
	}
	if !iKnow {
		if err := checkDeny(abs); err != nil {
			return "", err
		}
	}
	return abs, nil
}

func checkDeny(abs string) error {
	home, _ := os.UserHomeDir()
	home = filepath.Clean(home)
	abs = filepath.Clean(abs)

	if home != "" && abs == home {
		return fmt.Errorf("workspace: refusing to mount $HOME (%s); use a project dir, --upload, or --i-know", abs)
	}
	base := filepath.Base(abs)
	for _, d := range DenyBasenames {
		if strings.EqualFold(base, d) {
			return fmt.Errorf("workspace: refusing to mount %q (%s); pass --i-know to override", d, abs)
		}
	}
	// Also refuse if path is under ~/.ssh etc.
	if home != "" {
		for _, d := range DenyBasenames {
			prefix := filepath.Join(home, d)
			if abs == prefix || strings.HasPrefix(abs, prefix+string(os.PathSeparator)) {
				return fmt.Errorf("workspace: refusing path under ~/%s (%s); pass --i-know to override", d, abs)
			}
		}
	}
	return nil
}
