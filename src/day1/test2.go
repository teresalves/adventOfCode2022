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

	var fattestElfSizes [3]int
	var currElfSize int = 0;
	for scanner.Scan() {
		var input string = scanner.Text()
		if input == "" {
			min, minPos := getMinValue(fattestElfSizes)
			if currElfSize > min {
				fattestElfSizes[minPos] = currElfSize
			}
			currElfSize = 0
		} else {
			value, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println(err)
			}
			currElfSize += value
		}
    }

	fmt.Println(fattestElfSizes[0] + fattestElfSizes[1] + fattestElfSizes[2])
}

func getMinValue(array [3]int) (int,int) {
	min := array[0]
	minPos := 0
	for i:= 0; i < len(array); i++ {
		if array[i] < min {
			min = array[i]
			minPos = i
		}
	}
	return min, minPos
}
