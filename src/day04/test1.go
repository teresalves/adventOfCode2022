package main  

import (
    "fmt"
    "os"
	"log"
	"bufio"
	"strings"
	"strconv"
	"math"
) 
 
func main() {  

	f, err := os.Open("input.txt")
	
	if err != nil {
		log.Fatal(err)
   	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var sections1 = make(map[int]int)
	var sections2 = make(map[int]int)
	totalPairs := 0
	for scanner.Scan() {
		var input string = scanner.Text()
		isPair1 := 1
		isPair2 := 1
		sections1 = make(map[int]int)
		sections2 = make(map[int]int)
		zones := strings.Split(input, ",")

		firstElfZones := strings.Split(zones[0], "-")
		secondElfZones := strings.Split(zones[1], "-")

		val1, _ := strconv.Atoi(firstElfZones[0])
		val2, _ := strconv.Atoi(firstElfZones[1])

		val21, _ := strconv.Atoi(secondElfZones[0])
		val22, _ := strconv.Atoi(secondElfZones[1])

		for i:= val1; i <= val2; i++ {
			sections1[i] = 1
		}
		for i:= val21; i <= val22; i++ {
			if(sections1[i] != 1) {
				isPair1 = 0
			}
			sections2[i] = 1
		}

		if isPair1 == 0 {
			for i:= val1; i <= val2; i++ {
				if(sections2[i] != 1) {
					isPair2 = 0
					break
				}
			}
		}
				
		totalPairs += int(math.Max(float64(isPair1), float64(isPair2)))
    }

	fmt.Println(totalPairs)
}