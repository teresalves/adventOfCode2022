package main  

import (
    "fmt"
    "os"
	"log"
	"bufio"
	"strings"
) 
 
func main() {  

	f, err := os.Open("input.txt")
	
	if err != nil {
		log.Fatal(err)
   	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var opponentMoves = make(map[string]int)
	var myMoves = make(map[string]int)

	opponentOptions := []string{"A", "B", "C"}
	myOptions := []string{"X", "Y", "Z"}

	initMoves(opponentMoves, opponentOptions)
	initMoves(myMoves, myOptions)

	matrix := [3][3]int {
		{3, 6, 0},
		{0, 3, 6},
		{6, 0, 3},
	}

	score := 0

	for scanner.Scan() {
		var input string = scanner.Text()
		moves := strings.Split(input, " ")
		opponentMoveIndex := opponentMoves[moves[0]]
		myMoveIndex := myMoves[moves[1]]
		score += matrix[opponentMoveIndex][myMoveIndex]
		score += myMoveIndex + 1
    }

	fmt.Println(score)
}

func initMoves(moves map[string]int, values []string) {
	for i:= 0; i < len(values); i++ {
		moves[values[i]] = i
	}
}