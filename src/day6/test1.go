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

	var items = make(map[byte]int)
  
    for fileScanner.Scan() {
        fileLines = append(fileLines, fileScanner.Text())
    }
	var input string = fileLines[0]
	result := 0
	for i:= 0; i < len(input); i++ {
		items = make(map[byte]int)
		for j := 0 ; j < 4; j++ {
			items[input[i+j]] = 1
		}
		if len(items) == 4 {
			result = i + 4
			break
		}
	}
	fmt.Println(result)
}
