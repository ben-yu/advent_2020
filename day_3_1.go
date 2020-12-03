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

	treeCount := 0
	for x, y := 0, 0; x < len(treeMap); y, x = ((y + 3) % len(treeMap[0])), x+1 {
		log.Printf("x,y: %v %v", x, y)

		if string(treeMap[x][y]) == "#" {
			treeCount += 1
		}
	}
	log.Printf("Tree Count: %v", treeCount)
}
