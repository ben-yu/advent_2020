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
	maskStr := ""
	re := regexp.MustCompile(`mem\[([0-9]+)\]`)

	for scanner.Scan() {
		line := scanner.Text()
		results := strings.Split(line, " = ")
		if results[0] == "mask" {
			mask = make(map[int]int, 0)
			maskStr = results[1]
			for i, v := range strings.Split(maskStr, "") {
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

			addressMasks := []string{maskStr}
			for i, v := range strings.Split(maskStr, "") {
				if v == "X" {
					for j, w := range addressMasks {
						addressMasks[j] = replaceAtIndex(w, 'X', i)
						addressMasks = append(addressMasks, replaceAtIndex(w, '1', i))
					}
				}
			}

			memVal, _ := strconv.Atoi(results[1])
			//log.Printf("%v", memVal)
			for _, addrMask := range addressMasks {
				newMemPos := memPos
				//log.Printf("AddrMask %v", addrMask)
				for k, v := range addrMask {
					if v == '1' {
						newMemPos = newMemPos | (1 << (35 - k))
					} else if v == 'X' {
						clearMask := ^(1 << (35 - k))
						newMemPos = newMemPos & clearMask
					}

				}
				log.Printf("NewMemPos %v", newMemPos)
				memory[newMemPos] = memVal
			}
		}
		//log.Printf("Memory %v", memory)
	}

	sum := 0
	for _, v := range memory {
		sum += v
	}

	log.Printf("Answer: %v", sum)
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}
