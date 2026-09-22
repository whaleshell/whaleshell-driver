package vm

import "github.com/whaleshell/whaleshell-driver/driver"

func init() {
	driver.Register("vm", func() (driver.ComputeDriver, error) {
		return New(), nil
	})
}
