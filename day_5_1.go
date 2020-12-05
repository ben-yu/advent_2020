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
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Max Seat ID: %v", maxSeatID)
}
