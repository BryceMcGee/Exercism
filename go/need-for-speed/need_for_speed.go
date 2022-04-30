package speed

type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

type Track struct {
	distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	battery := 100
	distance := 0

	car := Car{
		battery:      battery,
		batteryDrain: batteryDrain,
		speed:        speed,
		distance:     distance,
	}

	return car
}

// NewTrack created a new track
func NewTrack(distance int) Track {
	raceTrack := Track{
		distance: distance,
	}
	return raceTrack
}

// Drive drives the car one time. If there is not enough battery to drive on more time,
// the car will not move.
func Drive(car Car) Car {
	if car.batteryDrain > car.battery {
		return car
	}
	car.distance = car.distance + car.speed
	car.battery = car.battery - car.batteryDrain
	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	num1 := track.distance / car.speed
	num2 := num1 * car.batteryDrain

	if car.battery-num2 >= 0 {
		return true
	}
	return false
}
