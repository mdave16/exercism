package elon

import "fmt"

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func (car *Car) Drive() {
	if car.battery < car.batteryDrain {
		return
	}
	car.battery -= car.batteryDrain
	car.distance += car.speed
}

// DisplayDistance will display the distance currently travelled by the car.
func (car Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

// DisplayBattery() will display the current battery level of the car.
func (car Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}

// CanFinish checks if a car is able to finish a certain track.
func (car Car) CanFinish(trackDistance int) bool {
	return (car.speed * (car.battery / car.batteryDrain)) >= trackDistance
}
