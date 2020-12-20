package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Tile struct {
	id    int
	val   [][]string
	edges map[string][]string
}

func main() {
	file, err := os.Open("./day_20_1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	tileRe := regexp.MustCompile(`Tile (\d+):`)
	currentTile := make([][]string, 0)
	allTiles := make(map[int]Tile, 0)
	edgeMap := make(map[string][]int, 0)
	tileID := 0

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "Tile ") {
			res := tileRe.FindStringSubmatch(line)[1]
			tileID, _ = strconv.Atoi(res)
		} else if line == "" {
			allTiles[tileID] = Tile{
				id:  tileID,
				val: currentTile,
			}
			currentTile = make([][]string, 0)
		} else {
			row := make([]string, 0)
			for _, v := range line {
				row = append(row, string(v))
			}
			currentTile = append(currentTile, row)
		}
	}

	log.Printf("%v", allTiles)

	for id, tile := range allTiles {
		topEdge := strings.Join(tile.val[0], "")
		edgeMap[topEdge] = append(edgeMap[topEdge], id)
		topEdge = Reverse(topEdge)
		edgeMap[topEdge] = append(edgeMap[topEdge], id)

		bottomEdge := strings.Join(tile.val[len(tile.val)-1], "")
		edgeMap[bottomEdge] = append(edgeMap[bottomEdge], id)
		bottomEdge = Reverse(bottomEdge)
		edgeMap[bottomEdge] = append(edgeMap[bottomEdge], id)

		leftEdge := vertJoin(tile.val, 0)
		edgeMap[leftEdge] = append(edgeMap[leftEdge], id)
		leftEdge = Reverse(leftEdge)
		edgeMap[leftEdge] = append(edgeMap[leftEdge], id)

		rightEdge := vertJoin(tile.val, len(tile.val[0])-1)
		edgeMap[rightEdge] = append(edgeMap[rightEdge], id)
		rightEdge = Reverse(rightEdge)
		edgeMap[rightEdge] = append(edgeMap[rightEdge], id)
	}

	log.Printf("%v", edgeMap)

	// Check Matches

	tileMatchCount := make(map[int]int, 0)

	for _, v := range edgeMap {
		if len(v) > 1 {
			for _, id := range v {
				if _, ok := tileMatchCount[id]; !ok {
					tileMatchCount[id] = 1
				} else {
					tileMatchCount[id] += 1
				}
			}
		}
	}

	log.Printf("%v", tileMatchCount)

	answer := 1
	for k, v := range tileMatchCount {
		if v == 4 {
			answer *= k
		}
	}
	log.Printf("Answer %v", answer)

}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func vertJoin(arr [][]string, index int) string {
	returnStr := ""
	for _, v := range arr {
		returnStr += v[index]
	}

	return returnStr
}

func rotate(matrix [][]string) {
	for i, temp, n := 0, "", len(matrix)-1; i <= n/2; i++ {
		for j := i; j < n-i; j++ {
			temp = matrix[j][n-i]
			matrix[j][n-i] = matrix[i][j]
			matrix[i][j] = matrix[n-j][i]
			matrix[n-j][i] = matrix[n-i][n-j]
			matrix[n-i][n-j] = temp
		}
	}
}
