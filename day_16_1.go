package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./day_16_1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanFields := true
	fieldRanges := make(map[string][][]int, 0)
	fieldNames := make([]string, 0)
	for scanFields {
		scanner.Scan()
		line := scanner.Text()
		if line == "" {
			scanFields = false
			break
		}
		res := strings.Split(line, ": ")
		field := res[0]
		rangeStrings := strings.Split(res[1], " or ")
		ranges := make([][]int, 0)
		for _, v := range rangeStrings {
			rangeStrs := strings.Split(v, "-")
			lower, _ := strconv.Atoi(rangeStrs[0])
			upper, _ := strconv.Atoi(rangeStrs[1])
			ranges = append(ranges, []int{lower, upper})
		}
		fieldRanges[field] = ranges
		fieldNames = append(fieldNames, field)
	}
	log.Printf("Ranges %v", fieldNames)

	scanner.Scan()
	scanner.Scan()
	myNumberList := strings.Split(scanner.Text(), ",")
	log.Printf("My Ticket %v", myNumberList)

	scanner.Scan()
	scanner.Scan()

	invalidSum := 0
	scanNearbyTickets := true

	validFieldValues := make([][]string, 0)

	for scanNearbyTickets {
		scanner.Scan()
		invalidTicket := false
		line := scanner.Text()
		if line == "" {
			scanNearbyTickets = false
			break
		}
		fieldValues := strings.Split(scanner.Text(), ",")
		for _, fieldValStr := range fieldValues {
			invalidForAll := true
			fieldVal, _ := strconv.Atoi(fieldValStr)
			for _, possibleRanges := range fieldRanges {
				for _, fieldRange := range possibleRanges {
					if fieldVal >= fieldRange[0] && fieldVal <= fieldRange[1] {
						invalidForAll = false
					}
				}
			}
			if invalidForAll {
				log.Printf("Invalid %v", fieldVal)
				invalidSum += fieldVal
				invalidTicket = true
			}
		}
		if !invalidTicket {
			validFieldValues = append(validFieldValues, fieldValues)
		}
	}

	possibleFields := make(map[int][]string, 0)
	for i, _ := range myNumberList {
		a := make([]string, len(fieldNames))
		copy(a, fieldNames)
		possibleFields[i] = a
	}
	//log.Printf("possible fields %v", possibleFields)

	for _, v := range validFieldValues {
		for i, fieldValStr := range v {
			fieldVal, _ := strconv.Atoi(fieldValStr)
			for fieldName, possibleRanges := range fieldRanges {
				possible := false
				for _, fieldRange := range possibleRanges {
					if fieldVal >= fieldRange[0] && fieldVal <= fieldRange[1] {
						possible = true
					}
				}
				if !possible {
					possibleFields[i] = removeStr(possibleFields[i], fieldName)
					//log.Printf("possible fields %v %v %v", i, fieldName, possibleFields)
				}
			}
		}
	}

	//log.Printf("Answer %v", possibleFields)
	//log.Printf("Answer %v", len(myNumberList))

	finalFieldMapping := make(map[int]string, 0)
	fieldsAdded := 0
	answer := 1
	for fieldsAdded != len(myNumberList) {
		for k, v := range possibleFields {
			if len(v) == 1 {
				finalFieldMapping[k] = v[0]
				if strings.HasPrefix(v[0], "departure") {
					r, _ := strconv.Atoi(myNumberList[k])
					answer = answer * r
				}
				log.Printf("%v -> %v -> %v", k, v[0], myNumberList[k])
				for a, _ := range possibleFields {
					possibleFields[a] = removeStr(possibleFields[a], v[0])
				}
				fieldsAdded += 1
			}
		}
	}
	log.Printf("Answer %v", answer)
}

func removeStr(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
