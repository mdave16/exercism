package lasagna

// PreperationTime calculates total preperation time given layers and how long an average layer takes.
func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer == 0 {
		return PreparationTime(layers, 2)
	}
	return len(layers) * timePerLayer
}

// Quantities calculates total quantities of noodles and sauce needed, given the layers.
func Quantities(layers []string) (noodles int, sauce float64) {
	noodles = 0
	sauce = 0
	for _, v := range layers {
		if v == "noodles" {
			noodles += 50
		}
		if v == "sauce" {
			sauce += 0.2
		}
	}
	return
}

// inList searches a haystack for a needle and tells us if it's in the haystack.
func inList(haystack []string, needle string) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}

// AddSecretIngredient will add any secret ingredient from friends recipe into my recipe
func AddSecretIngredient(friendsList, myList []string) []string {
	for _, f := range friendsList {
		if !inList(myList, f) {
			myList[len(myList)-1] = f
		}
	}
	return myList
}

// scale scales a vector by a scalar.
func scale(slice []float64, scalar float64) []float64 {
	for i := range slice {
		slice[i] = slice[i] * scalar
	}
	return slice
}

// ScaleRecipe scales a generic recipe to one for n people.
func ScaleRecipe(quantities []float64, guests int) []float64 {
	return scale(
		append(make([]float64, 0), quantities...),
		float64(guests)/2)
}
