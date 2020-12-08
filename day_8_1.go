package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("./day_8_1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	program := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		program = append(program, line)
	}

	visited := make(map[int]bool, 0)
	halt := false
	currentLineIndex := 0
	acc := 0

	re := regexp.MustCompile(`([\w]+) ([+|-]+)([0-9]+)?`)
	for !halt {
		if visited[currentLineIndex] {
			halt = true
			break
		}
		matchStr := re.FindStringSubmatch(program[currentLineIndex])

		instruction := matchStr[1]
		sign := matchStr[2]
		val, _ := strconv.Atoi(matchStr[3])
		if sign == "-" {
			val = val * -1
		}

		visited[currentLineIndex] = true

		switch instruction {
		case "acc":
			acc += val
			currentLineIndex += 1
		case "jmp":
			currentLineIndex += val
		case "nop":
			currentLineIndex += 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Answer: %v", acc)
}
