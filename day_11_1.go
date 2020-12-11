package main

import (
	"bufio"
	"log"
	"os"
	"reflect"
	"strings"
)

func main() {
	file, err := os.Open("./day_11_1.txt")
	//file, err := os.Open("./test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	grid := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]string, 0)
		for _, c := range strings.Split(line, "") {
			row = append(row, c)
		}
		grid = append(grid, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	isStable := false

	gridCopy := make([][]string, len(grid))
	for i := range grid {
		log.Printf("%v", grid[i])
		gridCopy[i] = make([]string, len(grid[i]))
		copy(gridCopy[i], grid[i])
	}
	log.Printf("----------------------------")

	for !isStable {
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[0]); j++ {
				seat := grid[i][j]
				empty := 0
				occupied := 0
				for k := -1; k <= 1; k++ {
					for l := -1; l <= 1; l++ {
						if k == 0 && l == 0 {
							continue
						}
						if i+k >= 0 && i+k < len(grid) && j+l >= 0 && j+l < len(grid[0]) {
							if grid[i+k][j+l] == "L" {
								empty += 1
							}
							if grid[i+k][j+l] == "#" {
								occupied += 1
							}
						}
					}
				}
				if seat == "L" && occupied == 0 {
					gridCopy[i][j] = "#"
				}
				if seat == "#" && occupied >= 4 {
					gridCopy[i][j] = "L"
				}
			}
		}

		if reflect.DeepEqual(gridCopy, grid) {
			isStable = true
		}
		for i := range gridCopy {
			grid[i] = make([]string, len(gridCopy[i]))
			copy(grid[i], gridCopy[i])
			log.Printf("%v", grid[i])
		}
		log.Printf("----------------------------")
	}
	answer := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == "#" {
				answer += 1
			}
		}
	}

	log.Printf("Answer: %v", answer)

}
