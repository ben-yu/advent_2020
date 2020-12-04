package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("./day_4_1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	foundFields := make([]string, 0)
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	sort.Strings(requiredFields)

	validPassportCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			sort.Strings(foundFields)
			if testEq(foundFields, requiredFields) || len(foundFields) == 8 {
				validPassportCount += 1
			}
			foundFields = make([]string, 0)
		} else {
			fields := strings.Split(line, " ")
			for _, field := range fields {
				keyValPair := strings.Split(field, ":")
				foundFields = append(foundFields, keyValPair[0])
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Passport Count: %v", validPassportCount)
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
