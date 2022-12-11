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

	posH := Key{5,0}
	posT := Key{5,0}
	var items = make(map[Key]int)
	for _, line := range fileLines {
		values := strings.Split(line, " ")
		direction := string(values[0])
		moves, _ := strconv.Atoi(values[1])
		switch direction{
			case "U":
				posH, posT = cases(posH, posT, moves, items, Move{-1,0})
			case "D":
				posH, posT = cases(posH, posT, moves, items, Move{1,0})
			case "L":
				posH, posT = cases(posH, posT, moves, items, Move{0,-1})
			case "R":
				posH, posT = cases(posH, posT, moves, items, Move{0,1})
		}
	}

    fmt.Println(len(items))
}

func cases(posH Key, posT Key, moves int, items map[Key]int, move Move) (Key, Key){
	for i:= 1; i <= moves; i++ {
		newPosH := Key{posH.X + move.X, posH.Y + move.Y}
		newPosT := newTPosition(newPosH, posT, posH)
		items[newPosT] = 1
		
		posH = newPosH
		posT = newPosT
	}
	return posH, posT
}

// check orientation
func newTPosition(posH Key, posT Key, prevPosH Key) Key {
	if math.Abs(float64(posH.X - posT.X)) <= 1 && math.Abs(float64(posH.Y - posT.Y)) <= 1 {
		return posT
	}
	return prevPosH
}
