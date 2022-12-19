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

	uppercaseOffset := 38
	lowecaseOffset := 96

	scanner := bufio.NewScanner(f)
	var list []int

	offset := lowecaseOffset
	for scanner.Scan() {
		var input string = scanner.Text()
		var items = make(map[byte]int)
		var items2 = make(map[byte]int)
		numberOfItems := len(input)
		for i:= 0; i < numberOfItems/2; i++ {
			items[input[i]] = 1
		}
		for i:= numberOfItems/2; i< numberOfItems; i++ {
			if items[input[i]] == 1 && items2[input[i]] != 1 {
				if input[i] < 97 {
					offset = uppercaseOffset
				} else {
					offset = lowecaseOffset
				}
				value := int(input[i]) 
				list = append(list, value - offset)
			}
			items2[input[i]] = 1
		}
    }
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