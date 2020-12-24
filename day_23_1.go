package main

import (
	"container/ring"
	"log"
	"strconv"
	"strings"
)

func main() {

	input := "583976241"

	labels := strings.Split(input, "")
	ncups := len(labels)

	cups := ring.New(ncups)
	prevSoln := map[int]*ring.Ring{}

	for i := 0; i < ncups; i++ {
		labelVal, _ := strconv.Atoi(labels[i])
		cups.Value = labelVal
		prevSoln[labelVal] = cups
		cups = cups.Next()
	}

	steps := 100

	for i := 0; i < steps; i++ {
		// Pick next 3
		pickup := cups.Unlink(3)

		// Find Destination
		destinationLabel := (ncups+cups.Value.(int)-2)%ncups + 1
		unavailable := map[int]bool{}
		for i := 0; i < 3; i++ {
			unavailable[pickup.Value.(int)] = true
			pickup = pickup.Next()
		}

		for unavailable[destinationLabel] {
			destinationLabel = (ncups+destinationLabel-2)%ncups + 1
		}

		//log.Printf("Destination: %v", destinationLabel)
		prevSoln[destinationLabel].Link(pickup)
		cups = cups.Next()
	}

	answer := ""
	prevSoln[1].Unlink(ncups - 1).Do(func(p interface{}) {
		answer += strconv.Itoa(p.(int))
	})

	log.Printf("Answer: %v", answer)
}
