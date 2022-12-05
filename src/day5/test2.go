package main  

import (
    "fmt"
    "os"
	"bufio"
	"strings"
	"strconv"
	"sort"
) 
 
func main() {  

	readFile, err := os.Open("input.txt")
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
    var fileLines []string

	var items = make(map[int][]byte)
  
    for fileScanner.Scan() {
        fileLines = append(fileLines, fileScanner.Text())
    }

	lineBreak := 0
	for i:= 0; i < len(fileLines); i++ {
		if fileLines[i] == "" {
			lineBreak = i
			break
		}
	}

	for i := lineBreak - 2; i >= 0 ; i-- {
		indexes := findAllLetters(fileLines[i])
		for j := 0; j < len(indexes); j++ {
			var index int = indexes[j]/4 + 1
			items[index] = append(items[index], fileLines[i][indexes[j]])
		}
	}

	for i := lineBreak + 1; i < len(fileLines) ; i++ {
		newLine := fileLines[i]
		values := strings.Split(newLine, " ")
		numberOfMoves,_ := strconv.Atoi(values[1])
		firstPile,_ := strconv.Atoi(values[3])
		secondPile,_ := strconv.Atoi(values[5])
		length := len(items[firstPile])
		
		items[secondPile] = append(items[secondPile], items[firstPile][length-numberOfMoves:]...)
		items[firstPile] = items[firstPile][:length-numberOfMoves]
	}
	printValues(items)
	
}

func findAllLetters(input string) []int{
	value := 0
	var newList []int
	newSlice := input
	previous := 0
	for true {
		value = strings.Index(newSlice, "[")
		if value == -1 {
			break
		}
		newList = append(newList, previous + value+1)
		previous += value + 3
		newSlice = input[previous:]
	}
	return newList
}

func printSlice(s []byte) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printValues(values map[int][]byte) {
	keys := make([]int, len(values))
    i := 0
    for k := range values {
        keys[i] = k
        i++
    }
	sort.Ints(keys)

	for _,key := range keys {
		length := len(values[key])
		r :=  values[key][length - 1]
		fmt.Printf("%s", string(r))
	}
	fmt.Println()
}
