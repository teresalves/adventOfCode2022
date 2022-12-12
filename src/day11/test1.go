package main  

import (
    "fmt"
    "os"
	"bufio"
	"strings"
	"strconv"
	) 

type ThrowingPair struct {
	ifTrue, ifFalse int
}
type Monkey struct {
    items []int
	operation string
	value int
	divisorTest int
	decision ThrowingPair
	nInspected int
}

func main() {  

	readFile, _ := os.Open("input.txt")
  
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
    var fileLines []string

	numberOfMonkeys := 8
	numberOfRounds := 20
	monkeys:= make([]Monkey, numberOfMonkeys)

    for fileScanner.Scan() {
        fileLines = append(fileLines, fileScanner.Text())
    }
	monkey := 0

	ops := map[string]func(int, int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}

	for index, line := range fileLines {
		values := strings.Split(line, " ")
		pos := index%7
		switch pos {
			case 0:
				monkey = int(values[len(values)-1][0]-48)
			case 1:
				line = strings.ReplaceAll(line, ",", "")
				values = strings.Split(line, ": ")
				values = strings.Split(values[1], " ")
				for _, value := range values {
					aux,_ := strconv.Atoi(value)
					monkeys[monkey].items = append(monkeys[monkey].items, aux)
				}
			case 2:
				aux,_ := strconv.Atoi(values[len(values)-1])
				monkeys[monkey].value = aux //will be 0 if it's a string
				monkeys[monkey].operation = values[len(values)-2]
			case 3:
				aux,_ := strconv.Atoi(values[len(values)-1])
				monkeys[monkey].divisorTest = aux
			case 4:
				aux,_ := strconv.Atoi(values[len(values)-1])
				monkeys[monkey].decision.ifTrue = aux
			case 5:
				aux,_ := strconv.Atoi(values[len(values)-1])
				monkeys[monkey].decision.ifFalse = aux
		}	
	}

	// Process everything
	for i:= 0; i < numberOfRounds; i++ {
		newMonkey := 0
		for monkey := 0; monkey < numberOfMonkeys; monkey++{
			for _,item := range monkeys[monkey].items {
				value := monkeys[monkey].value
				if value == 0 {
					value = item
				}
				value = ops[monkeys[monkey].operation](item, value)
				monkeys[monkey].nInspected++
				value = value/3

				if value % monkeys[monkey].divisorTest == 0 {
					newMonkey = monkeys[monkey].decision.ifTrue
				} else {
					newMonkey = monkeys[monkey].decision.ifFalse
				}
				monkeys[newMonkey].items = append(monkeys[newMonkey].items, value)
			}
			monkeys[monkey].items = make([]int, 0)
		}
		
	}

	result := calcResult(monkeys)
    fmt.Println(result)
}

func calcResult(monkeys []Monkey) int{
	maxMonkeyBusiness :=  make([]int, 2)
	for monkey, _ := range monkeys {
		min, index := min(maxMonkeyBusiness) 
		if monkeys[monkey].nInspected > min {
			maxMonkeyBusiness[index] = monkeys[monkey].nInspected
		}
	}
	return maxMonkeyBusiness[0]*maxMonkeyBusiness[1]
}

func min(array []int) (int, int){
	min := 500000000
	index := 0
	for i:=0; i< len(array); i++ {
		if(array[i] < min) {
			min = array[i]
			index = i
		}
	}
	return min, index
}

func printMacac(monkeys []Monkey) {
	for i, monkey := range monkeys {
		fmt.Printf("M:%v ", i)
		fmt.Println(monkey.items)
	}
}
