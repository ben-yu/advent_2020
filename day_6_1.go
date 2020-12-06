package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./day_6_1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	answerCount := 0
	currentAnswers := make(map[string]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			answerCount += len(currentAnswers)
			currentAnswers = make(map[string]int, 0)
		} else {
			for _, answer := range line {
				currentAnswers[string(answer)] = 1
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Answer Count: %v", answerCount)
}
