package kubernetes

import "github.com/whaleshell/whaleshell-driver/driver"

func init() {
	driver.Register("kubernetes", func() (driver.ComputeDriver, error) {
		return New(), nil
	})
}
