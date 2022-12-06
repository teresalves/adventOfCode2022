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
	for i:= 0; i < len(input)-14; i++ {
		items = make(map[byte]int)
		for j := 0 ; j < 14; j++ {
			items[input[i+j]] = 1
		}
		
		if len(items) == 14 {
			result = i + 14
			break
		}
	}
	fmt.Println(result)
}
