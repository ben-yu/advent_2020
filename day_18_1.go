package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./day_18_1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	answer := 0
	for scanner.Scan() {
		line := scanner.Text()
		// Fix Spacing
		line = strings.ReplaceAll(line, "(", "( ")
		line = strings.ReplaceAll(line, ")", " )")
		log.Printf(line)
		postFixExpr := parseInfix(line)
		log.Printf(postFixExpr)
		result := evalPostfix(postFixExpr)
		log.Printf(result)
		res, _ := strconv.Atoi(result)
		answer += res
	}
	log.Printf("Answer %v", answer)

}

var opa = map[string]struct {
	prec   int
	rAssoc bool
}{
	// Swap priority for Part 1/2
	"*": {1, false},
	"+": {2, false},
}

func evalPostfix(rpn string) string {
	var stack []string // hold operands
	for _, tok := range strings.Fields(rpn) {
		if _, isOp := opa[tok]; isOp {
			// Evaluate operators
			// pop second operand
			n2, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-1]
			// pop first operand
			n1, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-1]

			switch tok {
			case "+":
				res := strconv.Itoa(n1 + n2)
				stack = append(stack, res)
			case "*":
				res := strconv.Itoa(n1 * n2)
				stack = append(stack, res)
			}
		} else {
			// Push operands
			stack = append(stack, tok)
		}
	}
	return stack[0]
}

func parseInfix(e string) (rpn string) {
	var stack []string // holds operators and left parenthesis
	for _, tok := range strings.Fields(e) {
		switch tok {
		case "(":
			stack = append(stack, tok) // push "(" to stack
		case ")":
			var op string
			for {
				// pop item ("(" or operator) from stack
				op, stack = stack[len(stack)-1], stack[:len(stack)-1]
				if op == "(" {
					break // discard "("
				}
				rpn += " " + op // add operator to result
			}
		default:
			if o1, isOp := opa[tok]; isOp {
				// token is an operator
				for len(stack) > 0 {
					// consider top item on stack
					op := stack[len(stack)-1]
					if o2, isOp := opa[op]; !isOp || o1.prec > o2.prec ||
						o1.prec == o2.prec && o1.rAssoc {
						break
					}
					// top item is an operator that needs to come off
					stack = stack[:len(stack)-1] // pop it
					rpn += " " + op              // add it to result
				}
				// push operator (the new one) to stack
				stack = append(stack, tok)
			} else { // token is an operand
				if rpn > "" {
					rpn += " "
				}
				rpn += tok // add operand to result
			}
		}
	}
	// drain stack to result
	for len(stack) > 0 {
		rpn += " " + stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	return
}
