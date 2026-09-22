package podman

import "github.com/whaleshell/whaleshell-driver/driver"

func init() {
	driver.Register("podman", func() (driver.ComputeDriver, error) {
		return New()
	})
	driver.RegisterEngine("podman", func() (driver.Engine, error) {
		return New()
	})
}
