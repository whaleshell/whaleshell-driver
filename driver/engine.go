package driver

import "context"

// Probe is a host-side Docker/Podman readiness report for `whaleshell health`.
type Probe struct {
	OK              bool
	ServerVersion   string
	APIVersion      string
	OperatingSystem string
	Architecture    string
	Context         string
	Isolation       string
	HostGOOS        string
	Error           string
}

// Engine extends ComputeDriver with Docker Engine API ops used by the CLI
// (health, policy path, container IP). Implemented by docker and podman backends.
type Engine interface {
	ComputeDriver
	Health(ctx context.Context) Probe
	ImagePresent(ctx context.Context, ref string) bool
	RunProbe(ctx context.Context, initBin string) (string, error)
	PolicyHostPath(ctx context.Context, nameOrID string) (string, error)
	ContainerIP(ctx context.Context, containerID, networkName string) (string, error)
}
