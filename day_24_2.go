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
	log.Printf("Answer #1 : %v", blackTileCount)

	dayCount := 100

	for day := 0; day < dayCount; day++ {

		newState := make(map[CubeCoord]bool, 0)
		// Determine area bound to check
		maxX, maxY, maxZ := -99999, -99999, -99999
		minX, minY, minZ := 99999, 99999, 99999
		for k, _ := range tileState {
			if k.x > maxX {
				maxX = k.x
			}
			if k.y > maxY {
				maxY = k.y
			}
			if k.z > maxZ {
				maxZ = k.z
			}
			if k.x < minX {
				minX = k.x
			}
			if k.y < minY {
				minY = k.y
			}
			if k.z < minZ {
				minZ = k.z
			}
		}

		for x := minX - 1; x <= maxX+1; x++ {
			for y := minY - 1; y <= maxY+1; y++ {
				for z := minZ - 1; z <= maxZ+1; z++ {

					pos := CubeCoord{
						x: x,
						y: y,
						z: z,
					}
					currentState, ok := tileState[pos]
					if !ok {
						currentState = false
					}

					blackCount := adjBlackCount(pos, tileState)

					if currentState && (blackCount == 0 || blackCount > 2) {
						newState[pos] = false
					} else if !currentState && (blackCount == 2) {
						newState[pos] = true
					}
				}
			}
		}

		// Set new tile states
		for k, v := range newState {
			tileState[k] = v
		}
	}
	blackTileCount = 0
	for _, v := range tileState {
		if v {
			blackTileCount += 1
		}
	}
	log.Printf("Answer #2 : %v", blackTileCount)

}

func move(pos CubeCoord, dir string) CubeCoord {
	coord := CubeCoord{
		x: pos.x,
		y: pos.y,
		z: pos.z,
	}

	switch dir {
	case "e":
		coord.x += 1
		coord.y -= 1
	case "se":
		coord.z += 1
		coord.y -= 1
	case "sw":
		coord.z += 1
		coord.x -= 1
	case "w":
		coord.y += 1
		coord.x -= 1
	case "nw":
		coord.y += 1
		coord.z -= 1
	case "ne":
		coord.x += 1
		coord.z -= 1
	}

	return coord
}

var Directions = []string{"e", "se", "sw", "w", "nw", "ne"}

func adjBlackCount(pos CubeCoord, state map[CubeCoord]bool) int {
	count := 0
	for _, dir := range Directions {
		coord := move(pos, dir)
		searchState, ok := state[coord]
		if !ok {
			searchState = false
		}

		if searchState {
			count += 1
		}

	}

	return count
}
