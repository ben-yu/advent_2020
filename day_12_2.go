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
	wX := 10
	wY := 1
	re := regexp.MustCompile(`(\w)([0-9]+)`)
	for scanner.Scan() {
		line := scanner.Text()
		result := re.FindAllStringSubmatch(line, -1)

		action := result[0][1]
		amt, _ := strconv.Atoi(result[0][2])
		switch action {
		case "N":
			wY += amt
		case "S":
			wY -= amt
		case "E":
			wX += amt
		case "W":
			wX -= amt
		case "R":
			angle := amt / 90
			for i := 0; i < angle; i++ {
				wX, wY = wY, -wX
			}
		case "L":
			angle := amt / 90
			for i := 0; i < angle; i++ {
				wX, wY = -wY, wX
			}
		case "F":
			x += amt * wX
			y += amt * wY
		}
		log.Printf("Action %v, %v", action, amt)
		log.Printf("Ship (%v, %v, %v)", x, y, dir)
		log.Printf("Way (%v, %v, %v)", wX, wY, dir)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Answer: %v", math.Abs(float64(x))+math.Abs(float64(y)))
}
