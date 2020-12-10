package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("./day_10_1_input.txt")
	//file, err := os.Open("./test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	adapters := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		voltage, _ := strconv.Atoi(line)
		adapters = append(adapters, voltage)
	}

	sort.Ints(adapters)

	differences := map[int]int{1: 0, 2: 0, 3: 0}

	for i, _ := range adapters {
		if i > 0 && adapters[i]-adapters[i-1] <= 3 {
			differences[adapters[i]-adapters[i-1]] += 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Answer: %v", (differences[3]+1)*(differences[1]+1))
}
