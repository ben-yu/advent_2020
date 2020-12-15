package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./day_15_1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	numberList := strings.Split(scanner.Text(), ",")

	turn := 0

	maxTurn := 30000000
	lastSpoken := make(map[int]int, 0)
	lastPrevSpoken := make(map[int]int, 0)

	spokenCount := make(map[int]int, 0)
	lastNum := -1
	for _, v := range numberList {
		turn += 1
		n, _ := strconv.Atoi(v)
		lastSpoken[n] = turn
		spokenCount[n] = 1
		lastNum = n
		//log.Printf("Answer: %v", lastSpoken)
	}

	answer := 0
	for turn = turn + 1; turn <= maxTurn; turn++ {
		spokenNum := 0
		if spokenCount[lastNum] > 1 {
			spokenNum = turn - 1 - lastPrevSpoken[lastNum]
		}
		lastPrevSpoken[spokenNum] = lastSpoken[spokenNum]
		lastSpoken[spokenNum] = turn
		spokenCount[spokenNum] += 1
		lastNum = spokenNum
		//log.Printf("Answer %v: %v", turn, spokenNum)

		//log.Printf("Answer %v: %v %v", turn, lastSpoken[lastNum], lastSpoken)
		answer = spokenNum
	}
	log.Printf("Answer %v", answer)

}
