package main  

import (
    "fmt"
    "os"
	"bufio"
	"math"
	"strings"
	"strconv"
	) 

type Key struct {
    X, Y int
}

type Move struct {
	X, Y int
}

func main() {  

	readFile, _ := os.Open("input.txt")
  
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
    var fileLines []string
  
    for fileScanner.Scan() {
        fileLines = append(fileLines, fileScanner.Text())
    }
	var nKnots = 10

	knots := make([]Key, nKnots)
	initKnots(knots, nKnots)

	var items = make(map[Key]int)

	for _, line := range fileLines {
		values := strings.Split(line, " ")
		direction := string(values[0])
		moves, _ := strconv.Atoi(values[1])
		switch direction{
			case "U":
				knots = cases(knots, nKnots, moves, items, Move{-1,0})
			case "D":
				knots = cases(knots, nKnots, moves, items, Move{1,0})
			case "L":
				knots = cases(knots, nKnots, moves, items, Move{0,-1})
			case "R":
				knots = cases(knots, nKnots, moves, items, Move{0,1})
		}
	}

    fmt.Println(len(items))
}

func cases(knots []Key, nKnots int, moves int, items map[Key]int, move Move) []Key{
	for i:= 0; i < moves; i++ {
		knots[0] = Key{knots[0].X + move.X, knots[0].Y + move.Y}
		for j:= 1; j < nKnots; j++ {
			knots[j] = newKnotPosition(knots[j-1], knots[j])
			
			if(j == nKnots - 1) {
				items[knots[j]] = 1
			}
		}
	}
	return knots
}

func newKnotPosition(posH Key, posT Key) Key {
	dx := float64(posH.X - posT.X)
	dy := float64(posH.Y - posT.Y)
	if math.Abs(dx) <= 1 && math.Abs(dy) <= 1 {
		return posT
	}

	if(dx > 0) {
		posT.X++
	} else if(dx < 0) {
		posT.X--
	} 
	if(dy > 0) {
		posT.Y++
	} else if (dy < 0){
		posT.Y--
	}
	return posT
}

func initKnots(knots []Key, nKnots int) {
	for i:= 0; i < nKnots; i++ {
		knots[i] = Key{0,0}
	}
}
