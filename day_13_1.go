package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./day_13_1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	time, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()

	a := scanner.Text()
	busTimes := strings.Split(a, ",")

	minBus := 0
	minTime := time
	for _, bus := range busTimes {
		busID, err := strconv.Atoi(bus)
		if err != nil {
			continue
		}

		if time%busID == 0 {
			minTime = 0
			minBus = busID
			continue
		}

		timeDiff := busID*(time/busID+1) - time
		if timeDiff < minTime {
			minTime = timeDiff
			minBus = busID
		}
	}
	//log.Printf("Action %v, %v", action, amt)
	//log.Printf("(%v, %v, %v)", x, y, dir)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Answer: %v", minBus*minTime)
}
