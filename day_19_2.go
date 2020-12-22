package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var ruleSets map[string]string

func main() {
	file, err := os.Open("./day_19_1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ruleSets = make(map[string]string, 0)

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
		ruleSets[res1[0]] = res1[1]
	}

	//ruleZeroRegex := "^" + buildRegex("0") + "$"

	//log.Printf("%v", ruleZeroRegex)

	// Replace with loops
	ruleSets["8"] = `"` + buildRegex("42") + `+"` // 42 | 42 8 - one or more rule 42
	// Unloop Rule 11 several times
	ruleSets["11"] = ""
	for i := 1; i <= 10; i++ {
		ruleSets["11"] += fmt.Sprintf("|%s{%d}%s{%d}", buildRegex("42"), i, buildRegex("31"), i)
	}
	ruleSets["11"] = `"(?:` + ruleSets["11"][1:] + `)"`

	newRuleZeroRegex := `^` + buildRegex("0") + `$`

	log.Printf("New %v", buildRegex("0"))

	answer := 0
	for scanner.Scan() {
		line := scanner.Text()
		matches := len(regexp.MustCompile(newRuleZeroRegex).FindAllString(line, -1))
		if matches > 0 {
			answer += 1
			log.Printf("Match %v", line)
		}
	}

	log.Printf("Answer: %v", answer)

}

func buildRegex(ruleNum string) string {

	rule := ruleSets[ruleNum]
	re := ""
	// Rule is a character or custom replacement
	if ruleSets[ruleNum][0] == '"' {
		return ruleSets[ruleNum][1 : len(ruleSets[ruleNum])-1]
	}

	for _, s := range strings.Split(rule, " | ") {
		re += "|"
		for _, s := range strings.Fields(s) {
			re += buildRegex(s)
		}
	}

	return "(?:" + re[1:] + ")"
}
