package main  

import (
    "fmt"
    "os"
	"bufio"
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

	nCycles := 0
	result := 0
	addition := 1
	for _, line := range fileLines {
		values := strings.Split(line, " ")
		value := 0
		newCycles := 0
		if values[0] == "noop" {
			newCycles = 1
		} else {
			newCycles = 2
			value, _ = strconv.Atoi(values[1])
		}

		if(nCycles + 20)%40 > 2 && (nCycles + newCycles + 20)%40 < 2 {
			var cycles int = (nCycles + 20)/40
			fmt.Println(line, nCycles, cycles, addition, result)
			result += addition*(20 + cycles*40)
			fmt.Println("RESULT ", addition, cycles, result)
		}
		fmt.Println("ADD ", addition)
		addition += value
		nCycles += newCycles
	}

    fmt.Println(result)
}
