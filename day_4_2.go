package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./day_4_1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	foundFields := make(map[string]string, 0)
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	sort.Strings(requiredFields)

	validPassportCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			keys := make([]string, 0, len(foundFields))
			for k := range foundFields {
				keys = append(keys, k)
			}
			sort.Strings(keys)
			if testEq(keys, requiredFields) || len(foundFields) == 8 {
				if testValidFields(foundFields) {
					validPassportCount += 1
				}
			}
			foundFields = make(map[string]string, 0)
		} else {
			fields := strings.Split(line, " ")
			for _, field := range fields {
				keyValPair := strings.Split(field, ":")
				foundFields[keyValPair[0]] = keyValPair[1]
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Passport Count: %v", validPassportCount)
}

func testValidFields(passportFields map[string]string) bool {
	for k, v := range passportFields {
		if k == "byr" {
			byr, err := strconv.Atoi(v)
			if err != nil {
				log.Printf("%v %v", k, v)

				return false
			}
			if byr < 1920 || byr > 2002 {
				log.Printf("%v %v", k, v)
				return false
			}
		} else if k == "iyr" {
			iyr, err := strconv.Atoi(v)
			if err != nil {
				log.Printf("%v %v", k, v)

				return false
			}

			if iyr < 2010 || iyr > 2020 {
				log.Printf("%v %v", k, v)
				return false
			}
		} else if k == "eyr" {
			eyr, err := strconv.Atoi(v)
			if err != nil {
				log.Printf("%v %v", k, v)

				return false
			}

			if eyr < 2020 || eyr > 2030 {
				log.Printf("%v %v", k, v)

				return false
			}
		} else if k == "hgt" {
			var val int
			var unit string
			_, err := fmt.Sscanf(v, "%d%s", &val, &unit)
			if err != nil || unit == "" {
				log.Printf("%v %v", k, v)

				return false
			}

			if unit == "cm" && (val < 150 || val > 193) {
				log.Printf("%v %v", k, v)

				return false
			}

			if unit == "in" && (val < 59 || val > 76) {
				log.Printf("%v %v", k, v)

				return false
			}
		} else if k == "hcl" {
			var val int
			_, err := fmt.Sscanf(v, "#%x", &val)
			if err != nil {
				log.Printf("parse %v %v", k, v)

				return false
			}

			if len(v) != 7 {
				log.Printf("%v %v", k, v)

				return false
			}
		} else if k == "ecl" {
			validColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
			if !stringInSlice(v, validColors) {
				log.Printf("%v %v", k, v)

				return false
			}
		} else if k == "pid" {
			var val int
			_, err := fmt.Sscanf(v, "%d", &val)
			if err != nil {
				log.Printf("%v %v", k, v)

				return false
			}

			if len(v) != 9 {
				log.Printf("%v %v", k, v)

				return false
			}
		}

	}
	return true

}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func testEq(a, b []string) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
