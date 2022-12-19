package main  

import (
    "fmt"
    "os"
	"bufio"
) 
 
func main() {  

	readFile, err := os.Open("input.txt")
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
    var fileLines []string
  
    for fileScanner.Scan() {
        fileLines = append(fileLines, fileScanner.Text())
    }

	numberOfTreesInLine := len(fileLines[0])

	highestScore := 0
	for i := 0; i < numberOfTreesInLine; i++ {
		for j := 0; j < numberOfTreesInLine; j++ {
			currScore := checkUp(i,j,numberOfTreesInLine, fileLines) *
				checkDown(i,j,numberOfTreesInLine, fileLines) *
				checkRight(i,j,numberOfTreesInLine, fileLines) *
				checkLeft(i,j,numberOfTreesInLine, fileLines) 
			fmt.Printf("%v ", currScore)	
			if highestScore < currScore {
				highestScore = currScore
			}
					
		}
		fmt.Println()
	}

    fmt.Println(highestScore)
}

func checkUp(i int, j int, size int, fileLines []string) int {
	total := 0
	for a := 1; a + i < size; a++ {
		total++
		if(int(fileLines[i][j]) <= int(fileLines[i+a][j])) {
			break
		}
	}
	return total
}

func checkDown(i int, j int, size int, fileLines []string) int {
	total := 0
	for a := 1; i-a >= 0; a++ {
		total++
		if(int(fileLines[i][j]) <= int(fileLines[i-a][j])) {
			break
		}
	}
	return total
}

func checkRight(i int, j int, size int, fileLines []string) int {
	total := 0
	for a := 1; a + j < size; a++ {
		total++
		if(int(fileLines[i][j]) <= int(fileLines[i][j+a])) {
			break
		}
	}
	return total
}

func checkLeft(i int, j int, size int, fileLines []string) int {
	total := 0
	for a := 1; j-a >= 0; a++ {
		total++
		if(int(fileLines[i][j]) <= int(fileLines[i][j-a])) {
			break
		}
	}
	return total
}
