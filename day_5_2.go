package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./day_5_1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	seatMap := make(map[int]int, 0)
	maxSeatID := -1
	for scanner.Scan() {
		line := scanner.Text()
		lowerRow := 0
		upperRow := 127
		row := 0
		for i := 0; i < 8; i++ {
			if string(line[i]) == "F" {
				upperRow -= ((upperRow - lowerRow) / 2) + 1
				if i == 6 {
					row = upperRow
				}
			}
			if string(line[i]) == "B" {
				lowerRow += ((upperRow - lowerRow) / 2) + 1
				if i == 6 {
					row = lowerRow
				}
			}
		}

		lowerCol := 0
		upperCol := 7
		col := 0
		for i := 0; i < 3; i++ {
			if string(line[7+i]) == "L" {
				upperCol -= ((upperCol - lowerCol) / 2) + 1
				if i == 2 {
					col = upperCol
				}
			}
			if string(line[7+i]) == "R" {
				lowerCol += ((upperCol - lowerCol) / 2) + 1
				if i == 2 {
					col = lowerCol
				}
			}
		}

		seatID := row*8 + col
		seatMap[seatID] = 1
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}

	for k := 1; k < maxSeatID; k += 1 {
		_, ok1 := seatMap[k+1]
		_, ok2 := seatMap[k-1]
		_, ok3 := seatMap[k]

		if ok1 && ok2 && !ok3 {
			log.Printf("your seat id %v", k)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
