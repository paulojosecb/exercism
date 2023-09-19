package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, preparationPerLayer int) int {
	if preparationPerLayer == 0 {
		return len(layers) * 2
	} else {
		return len(layers) * preparationPerLayer
	}
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noddles := 0
	sauce := 0.0

	for _, element := range layers {
		switch element {
		case "noodles":
			noddles += 50
		case "sauce":
			sauce += 0.2
		}
	}

	return noddles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(ingredients []float64, numberOfPortions int) []float64 {
	newIngredients := []float64{}

	for _, ingredient := range ingredients {
		newIngredients = append(newIngredients, ingredient*float64(numberOfPortions)/2)
	}

	return newIngredients
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
