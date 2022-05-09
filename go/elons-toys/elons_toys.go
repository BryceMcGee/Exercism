package elon

import "fmt"

func (c *Car) Drive() {
	if c.battery-c.batteryDrain >= 0 {
		c.distance += c.speed
		c.battery -= c.batteryDrain
	}
}

func (c *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %v meters", c.distance)
}

func (c *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %v%%", c.battery)
}

func (c *Car) CanFinish(trackDistance int) bool {
	if (c.battery/c.batteryDrain)*c.speed >= trackDistance {
		return true
	}
	return false
}
