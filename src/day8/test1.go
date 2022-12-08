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

	visibleTrees := 0
	for i := 1; i < numberOfTreesInLine - 1; i++ {
		for j := 1; j < numberOfTreesInLine - 1; j++ {
			if(checkUp(i,j,numberOfTreesInLine, fileLines) || 
				checkDown(i,j,numberOfTreesInLine, fileLines)|| 
				checkRight(i,j,numberOfTreesInLine, fileLines) ||
				checkLeft(i,j,numberOfTreesInLine, fileLines)) { 
					visibleTrees++
					// // fmt.Println(i,j, int(fileLines[i][j])-48)
			} 
		}
	}
	visibleTrees+= numberOfTreesInLine*4 - 4

    fmt.Println(visibleTrees)
}

func checkUp(i int, j int, size int, fileLines []string) bool {
	for a := 1; a + i < size; a++ {
		if(int(fileLines[i][j]) <= int(fileLines[i+a][j])) {
			return false
		}
	}
	// fmt.Println("up")
	return true
}

func checkDown(i int, j int, size int, fileLines []string) bool {
	for a := 1; i-a >= 0; a++ {
		if(int(fileLines[i][j]) <= int(fileLines[i-a][j])) {
			return false
		}
	}
	// fmt.Println("down")
	return true
}

func checkRight(i int, j int, size int, fileLines []string) bool {
	for a := 1; a + j < size; a++ {
		if(int(fileLines[i][j]) <= int(fileLines[i][j+a])) {
			return false
		}
	}
	// fmt.Println("right")
	return true
}

func checkLeft(i int, j int, size int, fileLines []string) bool {
	for a := 1; j-a >= 0; a++ {
		if(int(fileLines[i][j]) <= int(fileLines[i][j-a])) {
			return false
		}
	}
	// fmt.Println("left")
	return true
}

