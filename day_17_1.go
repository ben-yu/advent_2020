package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./day_17_1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	grid := make([][]bool, 0)
	// Scan in initial state at z=0
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]bool, 0)
		for _, c := range strings.Split(line, "") {
			cubeState := false
			if c == "#" {
				cubeState = true
			}
			row = append(row, cubeState)
		}
		grid = append(grid, row)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Fill in rest of grid
	stepCount := 6
	state := make(map[int]map[int]map[int]bool, 12)
	for z := -stepCount; z <= stepCount; z++ {
		slice := make(map[int]map[int]bool, 0)
		for i := -stepCount; i < len(grid)+stepCount; i++ {
			row := make(map[int]bool, 0)
			for j := -stepCount; j < len(grid[0])+stepCount; j++ {
				row[j] = false
			}
			slice[i] = row
		}
		state[z] = slice
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			state[0][i][j] = grid[i][j]
		}
	}
	for z := -stepCount; z <= stepCount; z++ {
		log.Printf("z=%v:", z)
		for i := -stepCount; i < len(grid)+stepCount; i++ {
			printStr := ""
			for j := -stepCount; j < len(grid[0])+stepCount; j++ {
				if state[z][i][j] == true {
					printStr += "#"
				} else {
					printStr += "."
				}
			}
			log.Printf("%v", printStr)
		}
	}
	log.Printf("----------------------------")

	// Make a copy with max bounds
	stateCopy := make(map[int]map[int]map[int]bool, 12)
	for z := -stepCount; z <= stepCount; z++ {
		slice := make(map[int]map[int]bool, 0)
		for i := -stepCount; i < len(grid)+stepCount; i++ {
			row := make(map[int]bool, 0)
			for j := -stepCount; j < len(grid[0])+stepCount; j++ {
				row[j] = false
			}
			slice[i] = row
		}
		stateCopy[z] = slice
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			stateCopy[0][i][j] = grid[i][j]
		}
	}
	for z := -stepCount; z <= stepCount; z++ {
		log.Printf("z=%v:", z)
		for i := -stepCount; i < len(grid)+stepCount; i++ {
			printStr := ""
			for j := -stepCount; j < len(grid[0])+stepCount; j++ {
				if stateCopy[z][i][j] == true {
					printStr += "#"
				} else {
					printStr += "."
				}
			}
			log.Printf("%v", printStr)
		}
	}
	log.Printf("----------------------------")
	for s := 0; s < stepCount; s++ {
		for z := -stepCount; z <= stepCount; z++ {
			for i := -stepCount; i < len(state[z])-stepCount; i++ {
				for j := -stepCount; j < len(state[z][0])-stepCount; j++ {
					seat := state[z][i][j]
					empty := 0
					occupied := 0
					for k := -1; k <= 1; k++ {
						for l := -1; l <= 1; l++ {
							for m := -1; m <= 1; m++ {
								if k == 0 && l == 0 && m == 0 {
									continue
								}
								if i+k >= -stepCount && i+k < len(state[z])-stepCount && j+l >= -stepCount && j+l < len(state[z][0])-stepCount && z+m >= -stepCount && z+m <= stepCount {
									if state[z+m][i+k][j+l] == false {
										empty += 1
									}
									if state[z+m][i+k][j+l] == true {
										occupied += 1
									}
								}
							}
						}
					}
					if seat == false && occupied == 3 {
						stateCopy[z][i][j] = true
					} else if seat == true && (occupied == 2 || occupied == 3) {
						stateCopy[z][i][j] = true
					} else {
						//log.Printf("%v %v %v", z, i, j)
						stateCopy[z][i][j] = false
					}
				}
			}
		}
		for z := -stepCount; z <= stepCount; z++ {
			log.Printf("t=%v:", s+1)
			log.Printf("z=%v:", z)
			for i := -stepCount; i < len(grid)+stepCount; i++ {
				printStr := ""
				for j := -stepCount; j < len(grid[0])+stepCount; j++ {
					if stateCopy[z][i][j] == true {
						printStr += "#"
					} else {
						printStr += "."
					}
					state[z][i][j] = stateCopy[z][i][j]
				}
				log.Printf("%v", printStr)
			}
		}
		log.Printf("----------------------------")
	}

	answer := 0
	for z := -stepCount; z <= stepCount; z++ {
		for i := -stepCount; i < len(grid)+stepCount; i++ {
			for j := -stepCount; j < len(grid[0])+stepCount; j++ {
				if stateCopy[z][i][j] == true {
					answer += 1
				}
			}
		}
	}
	log.Printf("Answer %v", answer)
}
