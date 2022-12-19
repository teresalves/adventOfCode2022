package main  

import (
    "fmt"
    "os"
	"log"
	"bufio"
	"strconv"
) 
 
func main() {  

	f, err := os.Open("input.txt")
	
	if err != nil {
		log.Fatal(err)
   	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var fattestElfSize int = 0;
	var currElfSize int = 0;
	for scanner.Scan() {
		var input string = scanner.Text()
		if input == "" {
			currElfSize = 0
		} else {
			value, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println(err)
			}
			currElfSize += value
			if currElfSize > fattestElfSize {
				fattestElfSize = currElfSize
			}
		}
    }

	fmt.Println(fattestElfSize)
}