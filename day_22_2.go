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

	log.Printf("%v %v", deckOne, deckTwo)
	player, answer := playGame(deckOne, deckTwo)
	log.Printf("answer %v, player %v", answer, player)
}

func deckCheckSum(deck []int) int {
	sum := 0
	for i := 0; i < len(deck); i++ {
		sum += deck[i] * (len(deck) - i)
	}
	return sum
}

func playGame(deckOne, deckTwo []int) (int, int) {
	if len(deckOne) == 0 {
		return 2, deckCheckSum(deckTwo)
	} else if len(deckTwo) == 0 {
		return 1, deckCheckSum(deckOne)
	}

	deckOneCopy := make([]int, len(deckOne))
	copy(deckOneCopy, deckOne)
	deckTwoCopy := make([]int, len(deckTwo))
	copy(deckTwoCopy, deckTwo)

	deckOneChecks := make(map[int]bool, 0)
	deckTwoChecks := make(map[int]bool, 0)

	for len(deckOneCopy) > 0 && len(deckTwoCopy) > 0 {

		oneCheck := deckCheckSum(deckOneCopy)
		_, ok1 := deckOneChecks[oneCheck]
		twoCheck := deckCheckSum(deckTwoCopy)
		_, ok2 := deckTwoChecks[twoCheck]
		//log.Printf("Checks: %v %v", oneCheck, twoCheck)

		if ok1 && ok2 {
			return 1, -1
		}

		a := deckOneCopy[0]
		deckOneCopy = append([]int{}, deckOneCopy[1:]...)
		b := deckTwoCopy[0]
		deckTwoCopy = append([]int{}, deckTwoCopy[1:]...)

		if len(deckOneCopy) >= a && len(deckTwoCopy) >= b {
			winner, _ := playGame(deckOneCopy[:a], deckTwoCopy[:b])
			if winner == 1 {
				deckOneCopy = append(deckOneCopy, a, b)
				//log.Printf("%v  %v %v", winner, deckOneCopy, deckTwoCopy)
			} else {
				deckTwoCopy = append(deckTwoCopy, b, a)
				//log.Printf("%v %v %v", winner, deckOneCopy, deckTwoCopy)
			}
		} else if a > b {
			deckOneCopy = append(deckOneCopy, a, b)
			//log.Printf("%v %v %v %v", a, b, deckOneCopy, deckTwoCopy)
		} else {
			deckTwoCopy = append(deckTwoCopy, b, a)
			//log.Printf("%v %v %v %v", a, b, deckOneCopy, deckTwoCopy)
		}
		deckOneChecks[oneCheck] = true
		deckTwoChecks[twoCheck] = true
	}
	if len(deckOneCopy) == 0 {
		//log.Printf("%v %v", 2, deckTwoCopy)
		return 2, deckCheckSum(deckTwoCopy)
	} else if len(deckTwoCopy) == 0 {
		//log.Printf("%v %v", 1, deckOneCopy)
		return 1, deckCheckSum(deckOneCopy)
	}

	return 0, -1
}
