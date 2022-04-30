package lasagna

import (
	"strings"
)

func PreparationTime(layers []string, prepTime int) int {
	if prepTime == 0 {
		prepTime = 2
	}
	return len(layers) * prepTime
}

func Quantities(layers []string) (int, float64) {
	noodleLayer := 50
	noodleCount := 0

	sauceLayer := 0.2
	sauceCount := 0

	for i := 0; i < len(layers); i++ {
		if strings.Count(layers[i], "noodles") > 0 {
			noodleCount++
		}
		if strings.Count(layers[i], "sauce") > 0 {
			sauceCount++
		}
	}
	return noodleCount * noodleLayer, float64(sauceCount) * sauceLayer
}

func AddSecretIngredient(friendRecipe, myRecipe []string) {
	myRecipe[len(myRecipe)-1] = friendRecipe[len(friendRecipe)-1]
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	scaledQuantities := make([]float64, len(quantities))
	copy(scaledQuantities, quantities)

	for i := 0; i < len(scaledQuantities); i++ {
		scaledQuantities[i] = scaledQuantities[i] * float64(portions) / 2
	}
	return scaledQuantities
}
