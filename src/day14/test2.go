package main  

import (
    "fmt"
    "os"
	"bufio"
	"strconv"
	"strings"
	"math"
) 

type Position struct {
	Y,X int
}

func main() {  

	readFile, _ := os.Open("input.txt")
  
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
    var fileLines []string


    for fileScanner.Scan() {
        fileLines = append(fileLines, fileScanner.Text())
    }

	minX := 600
	minY := 600

	maxX := 0
	maxY := 0

	for i:= 0; i < len(fileLines) ; i++ {
		values := strings.Split(fileLines[i], " -> ")
		for _, value := range values {
			xyArray := strings.Split(value,",")
			x,_ := strconv.Atoi(xyArray[0])
			y,_ := strconv.Atoi(xyArray[1])
			minX = int(math.Min(float64(minX), float64(x)))
			minY = int(math.Min(float64(minY), float64(y)))
			maxX = int(math.Max(float64(maxX), float64(x)))
			maxY = int(math.Max(float64(maxY), float64(y)))
		}
	}

	
	lenY := maxY + 1 + 2
	offset := 2*lenY
	lenX := 4*lenY

	startXPos := 500 - minX + offset
	startPos := Position{0, startXPos}

	pitArray := make([][]string, lenY)
	for i := 0; i < lenY; i++ {
		pitArray[i] = make([]string, lenX)
		for j:= 0; j < lenX; j++ {
			pitArray[i][j] = "."
			if( i == lenY - 1) {
				pitArray[i][j] = "#"
			}
		}
	}
	pitArray[0][startXPos] = "+"


	for i:= 0; i < len(fileLines); i++ {
		values := strings.Split(fileLines[i], " -> ")
		for j := 0; j < len(values)-1; j++ {
			x1,y1 := getXY(values[j], minX, offset)
			x2, y2 := getXY(values[j+1], minX, offset)
			addPath(pitArray,x1,y1,x2,y2,minY)
		}
	}

	result := addSand(pitArray, startPos, lenY, lenX)
   
	//printBoard(pitArray)
	fmt.Println(result)
	
}

func getXY(value string, minX int, offset int) (int,int){
	xyArray := strings.Split(value,",")
	x,_ := strconv.Atoi(xyArray[0])
	y,_ := strconv.Atoi(xyArray[1])
	x=x-minX+offset
	y=y

	return x,y
}

func addPath(pitArray [][]string, x1,y1,x2,y2,minY int) {
	if y1 == y2 {
		for i:= 0; i + x1 != x2; i += signal(x2-x1) {
			pitArray[y1][x1 + i] = "#"
		}
	} else if x1 == x2 {
		for i:= 0; i + y1 != y2; i += signal(y2-y1) {
			pitArray[y1+i][x1] = "#"
		}
	}
	pitArray[y2][x2] = "#"
}

func signal(a int) int {
	if a > 0 {
		return 1
	} else if a < 0 {
		return -1
	}
	return 0
}

func printBoard(pitArray[][]string) {
	for i := 0;i < len(pitArray); i++ {
		fmt.Println(pitArray[i])
	}
}

func addSand(pitArray [][]string, startPos Position, lenY, lenX int) int {
	total := 0
	
	for true {
		//add one grain
		grain := addGrain(pitArray, startPos, lenY, lenX)
		if grain == true {
			total++
		} else {
			return ++total
		}
	}
	return 0
}


func addGrain(pitArray [][]string, firstPosition Position, lenY, lenX int) bool {
	currPos := firstPosition
	for true { // go through grain's path
		//go down
		newPos := Position{currPos.Y + 1, currPos.X}
		if newPos.Y  >= lenY { // abyss
			return false
		}
		if pitArray[newPos.Y][newPos.X] == "." {
			currPos = newPos
			continue
		}

		// go left
		newPos = Position{currPos.Y + 1, currPos.X - 1}
		if currPos.X  - 1 < 0 { //abyss
			return false
		}
		if pitArray[newPos.Y][newPos.X] == "." {
			currPos = newPos
			continue
		}

		// go right
		newPos = Position{currPos.Y + 1, currPos.X + 1}
		if currPos.X  + 1 >= lenX { // abyss
			return false
		}
		if pitArray[newPos.Y][newPos.X] == "." {
			currPos = newPos
			continue
		}

		// found its place
		pitArray[currPos.Y][currPos.X] = "o"
		
		if currPos != firstPosition {
			return true
		} else {
			return false
		}
	}
	return false
}
