package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type Foods struct {
	ingredients []string
	allergens   []string
}

func inSlice(s string, slice []string) bool {
	for _, se := range slice {
		if s == se {
			return true
		}
	}
	return false
}

func preprocess(filename string) ([]Foods, map[string]bool, map[string]bool, map[string]map[string]bool) {
	data, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(data), "\n")
	foods := []Foods{}
	allIngredients := map[string]bool{}
	allAllergens := map[string]bool{}
	for _, line := range lines {
		food := strings.Split(line, "(contains ")
		ingredients := strings.Split(food[0][:len(food[0])-1], " ")
		for _, ingredient := range ingredients {
			allIngredients[ingredient] = true
		}
		allergens := strings.Split(food[1][:len(food[1])-1], ", ")
		for _, allergen := range allergens {
			allAllergens[allergen] = true
		}
		foods = append(foods, Foods{ingredients, allergens})
	}
	ingredientAllergens := map[string]map[string]bool{}
	for ingredient := range allIngredients {
		allergens := make(map[string]bool, len(allAllergens))
		for k, v := range allAllergens {
			allergens[k] = v
		}
		ingredientAllergens[ingredient] = allergens
	}
	return foods, allIngredients, allAllergens, ingredientAllergens
}

func main() {
	foods, allIngredients, allAllergens, ingredientAllergens := preprocess("day21.in")

	occurance := map[string]int{}
	for _, food := range foods {
		for _, i := range food.ingredients {
			occurance[i]++
		}

		for _, a := range food.allergens {
			for i := range allIngredients {
				if !inSlice(i, food.ingredients) {
					delete(ingredientAllergens[i], a)
				}
			}
		}
	}

	p1 := 0
	for i := range allIngredients {
		if len(ingredientAllergens[i]) == 0 {
			p1 += occurance[i]
		}
	}
	fmt.Println("P1:", p1)

	allergenMapping := map[string]string{}
	used := []string{}

	for len(allergenMapping) < len(allAllergens) {
		for i := range allIngredients {
			p := []string{}
			for a := range ingredientAllergens[i] {
				if !inSlice(a, used) {
					p = append(p, a)
				}
			}
			if len(p) == 1 {
				allergenMapping[p[0]] = i
				used = append(used, p[0])
			}
		}
	}

	keys := make([]string, 0, len(allergenMapping))
	for k := range allergenMapping {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	p2 := ""
	for _, k := range keys {
		p2 += allergenMapping[k] + ","
	}
	fmt.Println("P2:", p2)
}
