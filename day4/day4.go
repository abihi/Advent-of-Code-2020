package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func preprocessPassports(filename string) []string {
	file, _ := os.Open(filename)

	passports := []string{}
	passport := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			passport += line
		} else {
			passports = append(passports, passport)
			passport = ""
		}
	}
	passports = append(passports, passport)
	return passports
}

func contains(v string, slice []string) bool {
	for _, s := range slice {
		if v == s {
			return true
		}
	}
	return false
}

func validateData(passportDict map[string]string, keys []string) bool {
	pdKeys := make([]string, len(passportDict))
	i := 0
	for pdKey := range passportDict {
		pdKeys[i] = pdKey
		i++
	}

	for _, key := range keys {
		if !contains(key, pdKeys) {
			return false
		}
	}

	byr, _ := strconv.Atoi(passportDict["byr"])
	if byr < 1920 || byr > 2002 {
		return false
	}
	iyr, _ := strconv.Atoi(passportDict["iyr"])
	if iyr < 2010 || iyr > 2020 {
		return false
	}
	eyr, _ := strconv.Atoi(passportDict["eyr"])
	if eyr < 2020 || eyr > 2030 {
		return false
	}

	metric := passportDict["hgt"][len(passportDict["hgt"])-2:]
	height, _ := strconv.Atoi(passportDict["hgt"][:len(passportDict["hgt"])-2])
	am := make([]string, 2)
	am[0] = "cm"
	am[1] = "in"
	if !contains(metric, am) {
		return false
	}
	if metric == "cm" && (height < 150 || height > 193) {
		return false
	}
	if metric == "in" && (height < 59 || height > 76) {
		return false
	}
	if len(passportDict["hcl"]) != 7 || passportDict["hcl"][0] != '#' {
		return false
	}

	hexDigits := strings.Split("0123456789abcdef", "")
	for _, c := range passportDict["hcl"][1:] {
		if !contains(string(c), hexDigits) {
			return false
		}
	}

	eyeColors := strings.Split("amb blu brn gry grn hzl oth", " ")
	if !contains(passportDict["ecl"], eyeColors) {
		return false
	}
	if len(passportDict["pid"]) != 9 {
		return false
	}

	return true
}

func splitKeyValue(s, sep string) (string, string) {
	kv := strings.Split(s, sep)
	return kv[0], kv[1]
}

func validPassports(passports []string, keys []string) int {
	valid := 0
	for _, x := range passports {
		pl := strings.Split(x, " ")
		pd := make(map[string]string)
		for _, y := range pl {
			k, v := splitKeyValue(y, ":")
			pd[k] = v
		}

		if validateData(pd, keys) {
			valid++
		}
	}

	return valid
}

func main() {
	passportKeys := strings.Split("byr iyr eyr hgt hcl ecl pid", " ")
	passports := preprocessPassports("day4_ex.in")
	answer := validPassports(passports, passportKeys)
	fmt.Println(answer)
}
