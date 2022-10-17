package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
    return true
	}
	return false
}

func better(option1, option2 string) string {
	if option1 > option2 {
    return option2
	}	else {
		return option1
	}
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	return better(option1, option2) + " is clearly the better choice."
}

func markdown(age float64) float64 {
	if age >= 10 {
    return 0.5
	}	else if age >= 3 {
		return 0.7
	} else {
		return 0.8
	}
}
// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
// For a rough estimate, assume if the vehicle is less than 3 years old, it costs 80% of the original price it had when it was brand new.
// If it is at least 10 years old, it costs 50%.
// If the vehicle is at least 3 years old but less than 10 years, it costs 70% of the original price.
	return originalPrice * markdown(age)
}
