// Package all blank-imports compute backends so they register with driver.Open.
//
//	import _ "github.com/whaleshell/whaleshell-driver/driver/all"
package all

import (
	_ "github.com/whaleshell/whaleshell-driver/internal/docker"
	_ "github.com/whaleshell/whaleshell-driver/internal/kubernetes"
	_ "github.com/whaleshell/whaleshell-driver/internal/podman"
	_ "github.com/whaleshell/whaleshell-driver/internal/vm"
)
