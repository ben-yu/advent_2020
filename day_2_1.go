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

		count := strings.Count(password, letter)
		if lower <= count && count <= upper {
			validPasswords += 1
			log.Printf("Valid Password: %v-%v, %v, %v", lower, upper, letter, password)
		}
	}
	log.Printf("Valid Password Count: %v", validPasswords)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
