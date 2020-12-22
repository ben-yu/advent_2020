package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./day_22_1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	deckOne := make([]int, 0)
	deckTwo := make([]int, 0)

	// Player 1
	scanner.Scan()
	doneScan := false
	for !doneScan {
		scanner.Scan()
		line := scanner.Text()
		if line == "" {
			doneScan = true
			break
		}
		val, _ := strconv.Atoi(line)
		deckOne = append(deckOne, val)
	}
	// Player 2
	scanner.Scan()
	doneScan = false
	for !doneScan {
		scanner.Scan()
		line := scanner.Text()
		if line == "" {
			doneScan = true
			break
		}
		val, _ := strconv.Atoi(line)
		deckTwo = append(deckTwo, val)
	}

	for len(deckOne) > 0 && len(deckTwo) > 0 {
		a := deckOne[0]
		deckOne = deckOne[1:]
		b := deckTwo[0]
		deckTwo = deckTwo[1:]
		if a > b {
			deckOne = append(deckOne, a, b)
		} else {
			deckTwo = append(deckTwo, b, a)
		}
	}

	answer := 0
	for i := 0; i < len(deckOne); i++ {
		answer += deckOne[i] * (len(deckOne) - i)
	}
	for i := 0; i < len(deckTwo); i++ {
		answer += deckTwo[i] * (len(deckTwo) - i)
	}

	log.Printf("answer %v", answer)
}
