package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

type Node struct {
	Value int
	Edges []int
}

func main() {
	file, err := os.Open("./day_10_1_input.txt")
	//file, err := os.Open("./test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	adapters := make([]int, 0)
	adapterMap := make(map[int]bool, 0)

	for scanner.Scan() {
		line := scanner.Text()
		voltage, _ := strconv.Atoi(line)
		adapters = append(adapters, voltage)
		adapterMap[voltage] = false
	}

	sort.Ints(adapters)
	adapters = append([]int{0}, adapters...)
	adapters = append(adapters, adapters[len(adapters)-1]+3)

	log.Printf("%v", adapters)

	totalPaths := make([]int, len(adapters))
	totalPaths[0] = 1
	for i := 1; i < len(adapters); i += 1 {
		totalPaths[i] = 0
		for j := 1; j <= 3; j += 1 {
			if i-j >= 0 && (adapters[i]-adapters[i-j]) <= 3 {
				//log.Printf("%v - %v", adapters[i], adapters[i-j])

				totalPaths[i] += totalPaths[i-j]
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Answer: %v", totalPaths)
}

func numPaths(node Node, visited map[int]bool, paths map[int]int, nodes map[int]Node) {
	visited[node.Value] = true
	paths[node.Value] = 1

	for _, edge := range node.Edges {
		nextNode := nodes[edge]
		if !visited[nextNode.Value] {
			numPaths(nextNode, visited, paths, nodes)
		}
		paths[node.Value] += paths[nextNode.Value]
	}
	return
}
