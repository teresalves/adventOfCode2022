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

		var position int = nCycles%40

		for i:= 0; i < newCycles; i++ {
			aux := position + i
			if aux == 0 || aux == 40 {
				fmt.Println("")
				aux = 0
			}

			if aux >= addition - 1 && aux <= addition + 1 {
				fmt.Printf("█")
			} else {
				fmt.Printf(" ")
			}
		}

		addition += value
		nCycles += newCycles
	}
	fmt.Println()
}
