package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./day_3_1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	treeMap := make([]string, 0)
	for scanner.Scan() {
		treeMap = append(treeMap, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	treeProduct := 1
	slopes := [5][2]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	for _, slope := range slopes {
		treeCount := 0
		for x, y := 0, 0; x < len(treeMap); y, x = ((y + slope[0]) % len(treeMap[0])), x+slope[1] {
			if string(treeMap[x][y]) == "#" {
				treeCount += 1
			}
		}
		treeProduct *= treeCount
	}
	log.Printf("Tree Count: %v", treeProduct)
}
