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
	state := make(map[int]map[int]map[int]map[int]bool, stepCount*2)
	for w := -stepCount; w <= stepCount; w++ {
		threeDCopy := make(map[int]map[int]map[int]bool, stepCount*2)
		for z := -stepCount; z <= stepCount; z++ {
			slice := make(map[int]map[int]bool, 0)
			for i := -stepCount; i < len(grid)+stepCount; i++ {
				row := make(map[int]bool, 0)
				for j := -stepCount; j < len(grid[0])+stepCount; j++ {
					row[j] = false
				}
				slice[i] = row
			}
			threeDCopy[z] = slice
		}
		state[w] = threeDCopy
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			state[0][0][i][j] = grid[i][j]
		}
	}
	for w := -stepCount; w <= stepCount; w++ {
		for z := -stepCount; z <= stepCount; z++ {
			//log.Printf("t=0")
			//log.Printf("z=%v, w=%v", z, w)
			for i := -stepCount; i < len(grid)+stepCount; i++ {
				printStr := ""
				for j := -stepCount; j < len(grid[0])+stepCount; j++ {
					if state[w][z][i][j] == true {
						printStr += "#"
					} else {
						printStr += "."
					}
				}
				//log.Printf("%v", printStr)
			}
		}
	}
	log.Printf("----------------------------")

	// Make a copy with max bounds
	stateCopy := make(map[int]map[int]map[int]map[int]bool, stepCount*2)
	for w := -stepCount; w <= stepCount; w++ {
		threeDCopy := make(map[int]map[int]map[int]bool, stepCount*2)
		for z := -stepCount; z <= stepCount; z++ {
			slice := make(map[int]map[int]bool, 0)
			for i := -stepCount; i < len(grid)+stepCount; i++ {
				row := make(map[int]bool, 0)
				for j := -stepCount; j < len(grid[0])+stepCount; j++ {
					row[j] = false
				}
				slice[i] = row
			}
			threeDCopy[z] = slice
		}
		stateCopy[w] = threeDCopy
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			stateCopy[0][0][i][j] = grid[i][j]
		}
	}
	for w := -stepCount; w <= stepCount; w++ {
		for z := -stepCount; z <= stepCount; z++ {
			//log.Printf("z=%v w=%v", z, w)
			for i := -stepCount; i < len(grid)+stepCount; i++ {
				printStr := ""
				for j := -stepCount; j < len(grid[0])+stepCount; j++ {
					if state[w][z][i][j] == true {
						printStr += "#"
					} else {
						printStr += "."
					}
				}
				//log.Printf("%v", printStr)
			}
		}
	}
	log.Printf("----------------------------")
	for s := 0; s < stepCount; s++ {
		for w := -stepCount; w <= stepCount; w++ {
			for z := -stepCount; z <= stepCount; z++ {
				for i := -stepCount; i < len(state[w][z])-stepCount; i++ {
					for j := -stepCount; j < len(state[w][z][0])-stepCount; j++ {
						seat := state[w][z][i][j]
						empty := 0
						occupied := 0
						for k := -1; k <= 1; k++ {
							for l := -1; l <= 1; l++ {
								for m := -1; m <= 1; m++ {
									for n := -1; n <= 1; n++ {
										if k == 0 && l == 0 && m == 0 && n == 0 {
											continue
										}
										if i+k >= -stepCount && i+k < len(state[w][z])-stepCount &&
											j+l >= -stepCount && j+l < len(state[w][z][0])-stepCount &&
											z+m >= -stepCount && z+m <= stepCount &&
											w+n >= -stepCount && w+n <= stepCount {
											if state[w+n][z+m][i+k][j+l] == false {
												empty += 1
											}
											if state[w+n][z+m][i+k][j+l] == true {
												occupied += 1
											}
										}
									}
								}
							}
						}
						if seat == false && occupied == 3 {
							stateCopy[w][z][i][j] = true
						} else if seat == true && (occupied == 2 || occupied == 3) {
							stateCopy[w][z][i][j] = true
						} else {
							stateCopy[w][z][i][j] = false
						}
					}
				}
			}
		}
		for w := -stepCount; w <= stepCount; w++ {
			for z := -stepCount; z <= stepCount; z++ {
				//log.Printf("t=%v:", s+1)
				//log.Printf("z=%v, w=%v", z, w)
				for i := -stepCount; i < len(grid)+stepCount; i++ {
					printStr := ""
					for j := -stepCount; j < len(grid[0])+stepCount; j++ {
						if stateCopy[w][z][i][j] == true {
							printStr += "#"
						} else {
							printStr += "."
						}
						state[w][z][i][j] = stateCopy[w][z][i][j]
					}
					//log.Printf("%v", printStr)
				}
			}
		}
		//log.Printf("----------------------------")
	}

	answer := 0
	for w := -stepCount; w <= stepCount; w++ {
		for z := -stepCount; z <= stepCount; z++ {
			for i := -stepCount; i < len(grid)+stepCount; i++ {
				for j := -stepCount; j < len(grid[0])+stepCount; j++ {
					if stateCopy[w][z][i][j] == true {
						answer += 1
					}
				}
			}
		}
	}
	log.Printf("Answer %v", answer)
}
