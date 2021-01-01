package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func inSlice(s string, slice []string) bool {
	for _, se := range slice {
		if s == se {
			return true
		}
	}
	return false
}

func checkAllergens(ingredientAllergens map[string]bool, allergens []string) map[string]bool {
	for allergen := range ingredientAllergens {
		if inSlice(allergen, allergens) {
			delete(ingredientAllergens, allergen)
		}
	}
	return ingredientAllergens
}

func main() {
	file, _ := os.Open("day21_ex.in")
	scanner := bufio.NewScanner(file)

	ingredientMap := map[string]map[string]bool{}

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "(")
		ingredients := strings.Split(line[0][:len(line[0])-1], " ")
		allergens := strings.Split(line[1][9:len(line[1])-1], ",")

		for _, ingredient := range ingredients {
			valueMap := map[string]bool{}
			for _, allergen := range allergens {
				a := strings.Trim(allergen, " ")
				valueMap[a] = true
			}

			ingredientMap[ingredient] = valueMap
		}
		fmt.Println(ingredients)
		fmt.Println(allergens)
		for _, ingredient := range ingredients {
			valueMap := ingredientMap[ingredient]
			ingredientMap[ingredient] = checkAllergens(valueMap, allergens)
		}
	}

	fmt.Println(ingredientMap)
}
