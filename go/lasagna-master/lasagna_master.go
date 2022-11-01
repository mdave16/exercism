package lasagna

const (
	defaultPreparationTimePerLayer = 2
	noodlesPerLayer                = 50
	saucePerLayer                  = 0.2
	defaultServingInRecipe         = 2
)

// PreperationTime calculates total preperation time given layers and how long an average layer takes.
func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer <= 0 {
		timePerLayer = defaultPreparationTimePerLayer
	}
	return len(layers) * timePerLayer
}

// Quantities calculates total quantities of noodles and sauce needed, given the layers.
func Quantities(layers []string) (noodles int, sauce float64) {
	noodleLayers := 0
	sauceLayers := 0
	for _, l := range layers {
		switch l {
		case "noodles":
			noodleLayers++
		case "sauce":
			sauceLayers++
		}
	}
	return noodleLayers * noodlesPerLayer, float64(sauceLayers) * saucePerLayer
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
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// ScaleRecipe scales a generic recipe to one for n people.
func ScaleRecipe(quantities []float64, guests int) (scaled []float64) {
	if len(quantities) == 0 {
		return scaled
	} else {
		for _, v := range quantities {
			scaled = append(scaled, v*float64(guests)/2.0)
		}
		return scaled
	}
}
