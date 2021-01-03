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

func checkAllergens(valueMap map[string]bool, allergens []string) map[string]bool {
	for allergen := range valueMap {
		if inSlice(allergen, allergens) {
			delete(valueMap, allergen)
		}
	}
	return valueMap
}

func main() {
	file, _ := os.Open("day21_ex.in")
	scanner := bufio.NewScanner(file)

	ingredientMap := map[string]map[string]bool{}
	all_I := [][]string{}
	all_A := [][]string{}

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "(contains ")
		ingredients := strings.Split(line[0][:len(line[0])-1], " ")
		all_I = append(all_I, ingredients)
		allergens := strings.Split(line[1][:len(line[1])-1], ", ")
		all_A = append(all_A, allergens)

		for _, ingredient := range ingredients {
			valueMap := ingredientMap[ingredient]
			if valueMap == nil {
				valueMap = map[string]bool{}
			}
			for _, allergen := range allergens {
				valueMap[allergen] = true
			}
			ingredientMap[ingredient] = valueMap
		}
	}
	fmt.Println(ingredientMap)

	for _, ingredients := range all_I {
		for i := range ingredients {
			valueMap := ingredientMap[ingredients[i]]
			ingredientMap[ingredients[i]] = checkAllergens(valueMap, all_A[i])
		}
	}

	fmt.Println(ingredientMap)
	containNoAllergen := []string{}
	for k, v := range ingredientMap {
		if len(v) == 0 {
			containNoAllergen = append(containNoAllergen, k)
		}
	}
	fmt.Println(containNoAllergen)
}
