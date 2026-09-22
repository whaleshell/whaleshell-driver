package docker

import "github.com/whaleshell/whaleshell-driver/driver"

func init() {
	driver.Register("docker", func() (driver.ComputeDriver, error) {
		return New()
	})
	driver.RegisterEngine("docker", func() (driver.Engine, error) {
		return New()
	})
}
