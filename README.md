<h1 align="center">osg-driver</h1>

<p align="center">
  <strong>Compute drivers</strong><br>
  Docker-first sandbox runtime — mounts, ExtraHosts, egress sidecar wiring.
</p>
<p align="center">
  <a href="https://github.com/zorneth/osg-driver/actions/workflows/ci.yml"><img src="https://github.com/zorneth/osg-driver/actions/workflows/ci.yml/badge.svg" alt="CI"></a>
  <a href="https://pkg.go.dev/github.com/zorneth/osg-driver"><img src="https://pkg.go.dev/badge/github.com/zorneth/osg-driver.svg" alt="Go Reference"></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-yellow.svg" alt="License"></a>
  <a href="https://github.com/zorneth/osg-driver"><img src="https://img.shields.io/badge/Go-1.27+-00ADD8?logo=go" alt="Go Version"></a>
</p>
<p align="center">
  <sub>Part of the <a href="https://github.com/zorneth">zorneth / osg</a> ecosystem</sub>
</p>

---

## Overview

**osg-driver** implements `ComputeDriver` for osg: create/start/exec/delete containers, attach the egress sidecar, validate bind mounts, and inject OpenShell-style host-gateway aliases (`host.osg.internal`).

### Key Features

| Category | Capabilities |
|----------|--------------|
| **Docker** | Engine API create/exec/logs (Podman-compatible socket) |
| **Sidecar** | Proxy container on dual-home network + CA bundle env |
| **Hosts** | `HostGatewayExtraHosts()` → `host.osg.internal` / `host.docker.internal` |
| **Mounts** | Workdir + reserved system paths validation |
| **Stubs** | VM / K8s placeholders for future drivers |

---

## Installation

```bash
go get github.com/zorneth/osg-driver@latest
```

**Requirements:** Go 1.27+, Docker Engine API access.

---

## Quick Start

```go
import dockerdriver "github.com/zorneth/osg-driver/driver/docker"

d, err := dockerdriver.New()
_ = d
_ = err
hosts := dockerdriver.HostGatewayExtraHosts()
// []string{"host.osg.internal:host-gateway", "host.docker.internal:host-gateway"}
_ = hosts
```

---

## Package Structure

| Path | Purpose |
|------|---------|
| `driver/` | Driver interface + stubs |
| `driver/docker/` | Docker Engine implementation |
| `mounts/` | Bind-mount policy |
| `sidecar/` | CA / env helpers for the proxy sidecar |


---

## Related

| Resource | Link |
|----------|------|
| Organization | [https://github.com/zorneth](https://github.com/zorneth) |
| Organization overview | [github.com/zorneth](https://github.com/zorneth) |
| pkg.go.dev | [`github.com/zorneth/osg-driver`](https://pkg.go.dev/github.com/zorneth/osg-driver) |

## License

[MIT](./LICENSE) © zorneth
