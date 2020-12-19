package main

import (
	"bufio"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type Node struct {
	id       int
	val      *string
	children []*Node
}

func main() {
	file, err := os.Open("./day_19_1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	ruleSets := make(map[string][][]string, 0)

	scanRules := true

	//answer := 0
	for scanRules {
		scanner.Scan()
		line := scanner.Text()

		if line == "" {
			scanRules = false
			break
		}

		res1 := strings.Split(line, ": ")
		res2 := strings.Split(res1[1], " | ")
		ruleSet := make([][]string, 0)
		for _, v := range res2 {
			res3 := strings.Split(v, " ")
			rule := make([]string, 0)
			for _, v2 := range res3 {
				_, err := strconv.Atoi(v2)
				if err != nil {
					rule = append(rule, string(v2[1]))
				} else {
					rule = append(rule, v2)
				}
			}
			ruleSet = append(ruleSet, rule)
		}
		ruleSets[res1[0]] = ruleSet
	}
	//log.Printf("Answer %v", ruleSets)
	//log.Printf("Answer %v", charMap)

	expandList := [][]string{ruleSets["0"][0]}
	results := make([][]string, 0)
	//log.Printf("expandList: %v", expandList)

	for len(expandList) > 0 {
		// Pop Rule
		currentRule := expandList[len(expandList)-1]
		expandList = expandList[:len(expandList)-1]
		allExpanded := true
		for i, v := range currentRule {
			if x, ok := ruleSets[v]; ok {
				allExpanded = false
				for _, replace := range x {
					newRule := make([]string, len(currentRule))
					copy(newRule, currentRule)
					newRule = append(newRule[:i], append(replace, newRule[i+1:]...)...)
					expandList = append(expandList, newRule)
					//log.Printf("newRule: %v", newRule)

				}
				break
			}
		}
		if allExpanded {
			results = append(results, currentRule)
		}
	}
	resultStrings := MapJoin(results)
	//log.Printf("results: %v", resultStrings)

	answer := 0
	for scanner.Scan() {
		line := scanner.Text()
		if containStr(resultStrings, line) {
			//log.Printf(line)
			answer += 1
		}
	}

	log.Printf("Answer: %v", answer)

}

func MapJoin(vs [][]string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = strings.Join(v, "")
	}
	return vsm
}

func containStr(s []string, str string) bool {
	for _, v := range s {
		if reflect.DeepEqual(v, str) {
			return true
		}
	}

	return false
}
func contains(s [][]string, str []string) bool {
	for _, v := range s {
		if reflect.DeepEqual(v, str) {
			return true
		}
	}

	return false
}
