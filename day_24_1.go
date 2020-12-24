package main

import (
	"bufio"
	"log"
	"os"
)

type CubeCoord struct {
	x, y, z int
}

func main() {
	file, err := os.Open("./day_24_1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	tileState := make(map[CubeCoord]bool, 0)

	for scanner.Scan() {
		line := scanner.Text()

		currentPos := CubeCoord{
			x: 0,
			y: 0,
			z: 0,
		}

		for i := 0; i < len(line); i++ {
			dir := string(line[i])
			if line[i] == 'n' || line[i] == 's' {
				if i+1 < len(line) && (line[i+1] == 'e' || line[i+1] == 'w') {
					dir += string(line[i+1])
					i += 1
				}
			}

			switch dir {
			case "e":
				currentPos.x += 1
				currentPos.y -= 1
			case "se":
				currentPos.z += 1
				currentPos.y -= 1
			case "sw":
				currentPos.z += 1
				currentPos.x -= 1
			case "w":
				currentPos.y += 1
				currentPos.x -= 1
			case "nw":
				currentPos.y += 1
				currentPos.z -= 1
			case "ne":
				currentPos.x += 1
				currentPos.z -= 1
			}

			//log.Printf("DIR: %v", dir)
		}
		if _, ok := tileState[currentPos]; !ok {
			tileState[currentPos] = true
		} else {
			tileState[currentPos] = !tileState[currentPos]
		}
		log.Printf("%v", currentPos)
	}
	log.Printf("%v", tileState)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	blackTileCount := 0
	for _, v := range tileState {
		if v {
			blackTileCount += 1
		}
	}
	log.Printf("Answer: %v", blackTileCount)
}

func testEq(a, b []string) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
