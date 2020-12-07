package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type bag struct {
	color    string
	quantity int
}

func main() {
	file, err := os.Open("./day_7_1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	bagColors := make(map[string][]string, 0)
	bags := make(map[string][]bag, 0)
	re := regexp.MustCompile(` ([0-9]+) ([a-zA-Z\s]+) bags?[,|\.]?`)
	for scanner.Scan() {
		line := scanner.Text()
		res := strings.Split(line, " bags contain")
		parentColor := res[0]
		res2 := strings.Split(res[1], ",")
		//log.Printf("%v", res2)
		for _, colorStr := range res2 {
			matchStr := re.FindStringSubmatch(colorStr)
			if len(matchStr) == 0 {
				continue
			}
			quantity, _ := strconv.Atoi(matchStr[1])
			color := matchStr[2]
			val, _ := bagColors[color]
			bagColors[color] = append(val, parentColor)
			bags[parentColor] = append(bags[parentColor], bag{
				color:    color,
				quantity: quantity,
			})
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Loop through possible parent bags
	canContain := bagColors["shiny gold"]
	seen := make(map[string]bool)
	seen["shiny gold"] = true
	for len(canContain) > 0 {
		curr := canContain[0]
		canContain = canContain[1:]
		if seen[curr] {
			continue
		}
		seen[curr] = true
		canContain = append(canContain, bagColors[curr]...)
	}
	// Subtract one for shiny gold
	log.Printf("Answer Count: %v", len(seen)-1)

	log.Printf("Total Bags: %v", countContents("shiny gold", bags))
}

func countContents(target string, contains map[string][]bag) int {
	sum := int(0)
	for _, c := range contains[target] {
		sum += int(c.quantity)
		sum += int(c.quantity) * countContents(c.color, contains)
	}
	return sum
}
