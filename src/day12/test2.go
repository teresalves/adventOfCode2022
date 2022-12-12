package main  

import (
    "fmt"
    "os"
	"bufio"
	"strings"
) 

type Key struct {
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

	start := Key{0,0}
	end := Key{0,0}

	var items = make(map[Key]Key) //value is parent
	lengthHorizontal := len(fileLines[0])
	lengthVertical := len(fileLines)

	for i, line := range fileLines {
		for j,value := range line {
			if string(value) == "S" {
				start = Key{i,j}
				fileLines[i] = strings.Replace(fileLines[i], "S", "a", 1)
			} else if string(value) == "E" {
				end = Key{i,j}
				fileLines[i] = strings.Replace(fileLines[i], "E", "z", 1)
			}
		}
	}
	var queue []Key
	queue = append(queue, end)
	

	prevPosition := Key{-1,-1}
	items[end] = prevPosition

	for (len(queue) != 0) {	
		currPosition := queue[0]

		queue = queue[1:] // remove current element
		if(fileLines[currPosition.X][currPosition.Y] == 'a') {
			start = currPosition
			break
		} 
		queue = addAdjacents(currPosition, queue, lengthHorizontal, lengthVertical, items, fileLines)
	}

	fmt.Println(calcDistance(items, end, start))
}

func addAdjacents(parent Key, queue []Key, lenH int, lenV int, items map[Key]Key, lines []string) []Key{
	x := parent.X
	y := parent.Y 

	element := lines[x][y]

	queue = validate(x-1 >= 0, Key{x-1,y}, parent, element, lines, queue, items)
	queue = validate(y-1 >= 0, Key{x,y-1}, parent, element, lines, queue, items)
	queue = validate(x+1 < lenV, Key{x+1,y}, parent, element, lines, queue, items)
	queue = validate(y+1 < lenH, Key{x,y+1}, parent, element, lines, queue, items)

	return queue
}

func validate(validation bool, pos Key, parent Key, element byte, lines []string, queue []Key, items map[Key]Key) []Key{
	result := validation && !keyExists(items, pos) && lines[pos.X][pos.Y] >= element - 1
	if result {
		items[pos] = parent
		queue = append(queue, pos)
	}
	return queue
}

func calcDistance(items map[Key]Key, start Key, end Key) int{
	element := end 
	total := 0
	for true {
		total++
		element = items[element]
		if element == start {
			break
		}
	}
	return total
}

func keyExists(items map[Key]Key, key Key) bool{
	if _, ok := items[key]; ok {
		return true
	} 
	return false
}