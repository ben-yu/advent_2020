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
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(`([\w]+) ([+|-]+)([0-9]+)?`)
	jmpIndexes := make([]int, 0)
	nopIndexes := make([]int, 0)
	for i, line := range program {
		matchStr := re.FindStringSubmatch(line)
		instruction := matchStr[1]

		switch instruction {
		case "jmp":
			jmpIndexes = append(jmpIndexes, i)
		case "nop":
			nopIndexes = append(nopIndexes, i)
		}
	}

	for _, index := range jmpIndexes {
		programCopy := make([]string, len(program))
		copy(programCopy, program)
		programCopy[index] = strings.ReplaceAll(programCopy[index], "jmp", "nop")
		halts, acc := programHalts(programCopy)
		if halts {
			log.Printf("Program: %v", programCopy)
			log.Printf("Answer: %v %v", index, acc)
		}
	}

	for _, index := range nopIndexes {
		programCopy := make([]string, len(program))
		copy(programCopy, program)
		programCopy[index] = strings.ReplaceAll(programCopy[index], "nop", "jmp")
		halts, acc := programHalts(programCopy)
		if halts {
			log.Printf("Answer: %v", acc)
		}
	}

}

func programHalts(program []string) (bool, int) {
	visited := make(map[int]bool, 0)
	halt := false
	currentLineIndex := 0
	acc := 0
	re := regexp.MustCompile(`([\w]+) ([+|-]+)([0-9]+)?`)

	for !halt {
		if currentLineIndex >= len(program) {
			return true, acc
		}
		if visited[currentLineIndex] {
			return false, acc
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

	return true, acc

}
