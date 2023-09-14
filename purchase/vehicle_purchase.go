package purchase

import (
	"fmt"
	"strings"
)

type Vehicle uint8

const (
	Car Vehicle = iota
	Truck
)

func (v Vehicle) String() string {
	switch v {
	case Car:
		return "car"
	case Truck:
		return "truck"
	}

	return "unknown"
}

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	kindLowered := strings.ToLower(kind)
	return kindLowered == Car.String() || kindLowered == Truck.String()
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	bestChoice := ""
	if option1 < option2 {
		bestChoice = option1
	} else {
		bestChoice = option2
	}

	return fmt.Sprintf("%s is clearly the better choice.", bestChoice)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	percentage := 0.0

	switch {
	case age < 3:
		percentage = 0.8
	case age >= 10:
		percentage = 0.5
	default:
		percentage = 0.7
	}

	return originalPrice * percentage
}
