package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./day_2_1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	re := regexp.MustCompile(`(?P<lower>[0-9]+)-(?P<upper>[0-9]+) (?P<letter>\w)`)
	validPasswords := 0
	for scanner.Scan() {
		s := scanner.Text()
		temp := strings.Split(s, ":")
		password := strings.Trim(temp[1], " ")
		result := re.FindAllStringSubmatch(temp[0], -1)
		lower, _ := strconv.Atoi(result[0][1])
		upper, _ := strconv.Atoi(result[0][2])
		letter := result[0][3]

		lowerIndex := lower - 1
		upperIndex := upper - 1
		validLower := lowerIndex >= 0 && lowerIndex < len(password) && string(password[lowerIndex]) == letter
		validUpper := upperIndex >= 0 && upperIndex < len(password) && string(password[upperIndex]) == letter
		if validLower != validUpper {
			validPasswords += 1
			log.Printf("Valid Password: %v-%v, %v, %v", lowerIndex, upperIndex, letter, password)
		}
	}
	log.Printf("Valid Password Count: %v", validPasswords)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
