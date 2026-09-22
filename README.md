<h1 align="center">whaleshell-driver</h1>

<p align="center">
  <strong>Compute drivers</strong><br>
  Docker-first sandbox runtime — mounts, ExtraHosts, egress sidecar wiring.
</p>
<p align="center">
  <a href="https://github.com/whaleshell/whaleshell-driver/actions/workflows/ci.yml"><img src="https://github.com/whaleshell/whaleshell-driver/actions/workflows/ci.yml/badge.svg" alt="CI"></a>
  <a href="https://pkg.go.dev/github.com/whaleshell/whaleshell-driver"><img src="https://pkg.go.dev/badge/github.com/whaleshell/whaleshell-driver.svg" alt="Go Reference"></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-yellow.svg" alt="License"></a>
  <a href="https://github.com/whaleshell/whaleshell-driver"><img src="https://img.shields.io/badge/Go-1.27+-00ADD8?logo=go" alt="Go Version"></a>
</p>
<p align="center">
  <sub>Part of the <a href="https://github.com/whaleshell">whaleshell / whaleshell</a> ecosystem</sub>
</p>

---

## Overview

**whaleshell-driver** implements `ComputeDriver` for whaleshell: create/start/exec/delete containers, attach the egress sidecar, validate bind mounts, and inject OpenShell-style host-gateway aliases (`host.whaleshell.internal`).

### Key Features

| Category | Capabilities |
|----------|--------------|
| **Docker** | Engine API create/exec/logs (Podman-compatible socket) |
| **Sidecar** | Proxy container on dual-home network + CA bundle env |
| **Hosts** | `HostGatewayExtraHosts()` → `host.whaleshell.internal` / `host.docker.internal` |
| **Mounts** | Workdir + reserved system paths validation |
| **Stubs** | VM / K8s placeholders for future drivers |

---

## Installation

```bash
go get github.com/whaleshell/whaleshell-driver@latest
```

**Requirements:** Go 1.27+, Docker Engine API access.

---

## Quick Start

```go
import (
    "github.com/whaleshell/whaleshell-driver/driver"
    _ "github.com/whaleshell/whaleshell-driver/driver/all" // register backends
)

d, err := driver.OpenEngine("docker")
_ = d
_ = err
hosts := driver.HostGatewayExtraHosts()
// []string{"host.whaleshell.internal:host-gateway", "host.docker.internal:host-gateway"}
_ = hosts
```

---

## Package Structure

| Path | Purpose |
|------|---------|
| `driver/` | Public API: `ComputeDriver`, `Engine`, `Open` / `OpenEngine`, helpers |
| `driver/all/` | Blank-import to register backends |
| `internal/docker/` | Docker Engine implementation |
| `internal/podman/` | Podman socket discovery → Docker API client |
| `internal/vm/`, `internal/kubernetes/` | Stub drivers |
| `internal/mounts/` | Bind-mount policy |
| `internal/sidecar/` | CA / env helpers for the proxy sidecar |


---

## Related

| Resource | Link |
|----------|------|
| Roadmap | [ROADMAP.md](./ROADMAP.md) |
| Organization | [https://github.com/whaleshell](https://github.com/whaleshell) |
| Organization overview | [github.com/whaleshell](https://github.com/whaleshell) |
| pkg.go.dev | [`github.com/whaleshell/whaleshell-driver`](https://pkg.go.dev/github.com/whaleshell/whaleshell-driver) |

## License

[MIT](./LICENSE) © whaleshell
