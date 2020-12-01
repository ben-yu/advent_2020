package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./day_1_1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	m := make(map[int64]int64)
	for scanner.Scan() {
		n, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		m[n] = n
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for k := range m {
		for l := range m {
			res := 2020 - k - l
			v, found := m[res]
			if found {
				log.Printf("Result is: %v * %v * %v  = %v", k, l, v, k*l*v)
				os.Exit(0)
			}
		}
	}

	log.Printf("Not Found")
}
