package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("./day_12_1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	dir := 0
	x := 0
	y := 0
	re := regexp.MustCompile(`(\w)([0-9]+)`)
	for scanner.Scan() {
		line := scanner.Text()
		result := re.FindAllStringSubmatch(line, -1)

		action := result[0][1]
		amt, _ := strconv.Atoi(result[0][2])
		switch action {
		case "N":
			y += amt
		case "S":
			y -= amt
		case "E":
			x += amt
		case "W":
			x -= amt
		case "L":
			dir += int(amt / 90)
			dir = int(dir % 4)
		case "R":
			dir -= int(amt / 90)
			dir = int(dir % 4)
		case "F":
			switch dir {
			case 0:
				x += amt
			case 1, -3:
				y += amt
			case 2, -2:
				x -= amt
			case 3, -1:
				y -= amt
			}
		}
		//log.Printf("Action %v, %v", action, amt)
		//log.Printf("(%v, %v, %v)", x, y, dir)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Answer: %v", math.Abs(float64(x))+math.Abs(float64(y)))
}
