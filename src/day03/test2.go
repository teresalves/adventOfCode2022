package main

import (
    "fmt"
    "os"
	"log"
	"bufio"
) 
 
func main() {  

	f, err := os.Open("input.txt")
	
	if err != nil {
		log.Fatal(err)
   	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var list []int

	numberOfLines := 0

	var items = make(map[byte]int)

	for scanner.Scan() {
		var input string = scanner.Text()

		numberOfItems := len(input)
		t := numberOfLines%3

		switch t {
			case 0:
				value := findValue(items)
				if(value != 0) {
					list = append(list, value)
				}
				
				items = make(map[byte]int) // re init map
				for i:= 0; i < numberOfItems; i++ {
					items[input[i]] = 1
				}
			case 1:
				for i:= 0; i < numberOfItems; i++ {
					if(items[input[i]] == 1) {
						items[input[i]] = 2
					}
				}
			case 2:
				for i:= 0; i < numberOfItems; i++ {
					if(items[input[i]] == 2) {
						items[input[i]] = 3
					}
				}

		}
		numberOfLines++
    }
	printSlice(list)
	fmt.Println(sumList(list))
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func sumList(values []int) int {
	sol := 0
	for _,v := range values {
		sol += v
	}
	return sol
}

func findValue(values map[byte]int) int {
	uppercaseOffset := 38
	lowecaseOffset := 96

	offset := 0
	for key, value := range values {
		if(value == 3) {
			if int(key) < 97 {
				offset = uppercaseOffset
			} else {
				offset = lowecaseOffset
			}
			return int(key) - offset
		}
	}
	return 0
}