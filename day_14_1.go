package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./day_14_1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	memory := make(map[int]int, 0)
	mask := make(map[int]int, 0)
	re := regexp.MustCompile(`mem\[([0-9]+)\]`)

	for scanner.Scan() {
		line := scanner.Text()
		results := strings.Split(line, " = ")
		if results[0] == "mask" {
			mask = make(map[int]int, 0)
			for i, v := range strings.Split(results[1], "") {
				bitVal, err := strconv.Atoi(v)
				if err != nil {
					continue
				}

				mask[35-i] = bitVal
			}
			log.Printf("Mask: %v", mask)
		} else {
			//			log.Printf("%v", re.FindStringSubmatch(results[0]))
			memPos, _ := strconv.Atoi(re.FindStringSubmatch(results[0])[1])
			memVal, _ := strconv.Atoi(results[1])
			for k, v := range mask {
				if v == 1 {
					memVal = memVal | (1 << k)
				} else {
					clearMask := ^(1 << k)
					memVal = memVal & clearMask
				}
				//log.Printf("v %v", v<<k)
				//log.Printf("m %v", memVal)
			}
			log.Printf("%v", memVal)
			memory[memPos] = memVal
		}
	}

	sum := 0
	for _, v := range memory {
		sum += v
	}

	log.Printf("Answer: %v", sum)
}
