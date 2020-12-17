package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
)

type Bus struct {
	id, offset *big.Int
}

func main() {
	file, err := os.Open("./day_13_1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	scanner.Scan()

	a := scanner.Text()
	busTimes := strings.Split(a, ",")

	busList := make([]Bus, 0)
	for i, bus := range busTimes {
		busID, err := strconv.Atoi(bus)
		if err != nil {
			continue
		}

		busTime := Bus{
			id:     big.NewInt(int64(busID)),
			offset: big.NewInt(int64(busID - i)),
		}

		busList = append(busList, busTime)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	answer, _ := solveChineseRemainderTheorem(busList)
	log.Printf("Answer: %v", *answer)
}

//https://en.wikipedia.org/wiki/Chinese_remainder_theorem#Using_the_existence_construction
func solveChineseRemainderTheorem(input []Bus) (*big.Int, error) {
	var one = big.NewInt(1)

	p := new(big.Int).Set(input[0].id)
	for _, n1 := range input[1:] {
		p.Mul(p, n1.id)
	}
	var x, q, s, z big.Int
	for _, n1 := range input {
		q.Div(p, n1.id)
		z.GCD(nil, &s, n1.id, &q)
		if z.Cmp(one) != 0 {
			return nil, fmt.Errorf("%d not coprime", n1)
		}
		x.Add(&x, s.Mul(n1.offset, s.Mul(&s, &q)))
	}
	return x.Mod(&x, p), nil
}
