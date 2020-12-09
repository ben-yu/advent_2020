package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {

	searchTarget := searchNum()

	file, err := os.Open("./day_9_1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	nums := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line)

		nums = append(nums, num)
	}

	lowerIndex := 0
	upperIndex := 1
	expand := true
	shrink := false
	foundSegment := false
	for !foundSegment {
		currentTotal := sum(nums[lowerIndex : upperIndex+1])
		if currentTotal == searchTarget {
			log.Printf("Answer #2: %v", min(nums[lowerIndex:upperIndex+1])+max(nums[lowerIndex:upperIndex+1]))
			foundSegment = true
			return
		}

		if currentTotal > searchTarget {
			expand = false
			shrink = true
		} else if currentTotal < searchTarget {
			expand = true
			shrink = false
		}

		if expand {
			upperIndex += 1
		}

		if shrink {
			lowerIndex += 1
		}
	}

}

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func max(array []int) int {
	result := MinInt
	for _, v := range array {
		if v > result {
			result = v
		}
	}
	return result
}

func min(array []int) int {
	result := MaxInt
	for _, v := range array {
		if v < result {
			result = v
		}
	}
	return result
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
func searchNum() int {
	file, err := os.Open("./day_9_1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	preambleLength := 25
	preambleNums := make(map[int]bool, 0)
	preambleOrder := make([]int, 0)
	index := 0
	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line)

		if index >= preambleLength {
			sumFound := false
			for k, _ := range preambleNums {
				_, matchExists := preambleNums[num-k]
				if matchExists && num-k != k {
					sumFound = true
					log.Printf("%v + %v = %v", k, num-k, num)
					break
				}
			}
			if !sumFound {
				log.Printf("Answer #1: %v", num)
				return num
			}
		}
		preambleNums[num] = true
		preambleOrder = append(preambleOrder, num)
		if len(preambleOrder) > preambleLength {
			v := preambleOrder[0]
			preambleOrder = preambleOrder[1:]
			delete(preambleNums, v)
		}

		index += 1
	}
	return 0
}
