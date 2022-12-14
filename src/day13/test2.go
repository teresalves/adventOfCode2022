package main  

import (
    "fmt"
    "os"
	"bufio"
	"strconv"
) 

type Packet struct {
	val int // if it is a list this will be -1
	children []*Packet
	parent *Packet
}

func main() {  

	readFile, _ := os.Open("input.txt")
  
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
    var fileLines []string


    for fileScanner.Scan() {
        fileLines = append(fileLines, fileScanner.Text())
    }

	dividers := []*Packet{processLine("[[2]]"), processLine("[[6]]")}
	var total [2]int
	for i:= 0; i < len(fileLines) -1; i+=3 {
		for j := 0; j < 4; j++ {
			element := processLine(fileLines[i + j%2])
			if isOrdered(element.children, dividers[j/2].children) > 0 {
				total[j/2]++
			}
		}

	}
	total[0]++ //adding one for current index
	total[1]+=2 //acounting for the first divisor
   
	fmt.Println(total[0]*total[1])
	
}
   
func isOrdered(leftPacket []*Packet, rightPacket []*Packet) int{

	for i:= 0; i < len(leftPacket) && i < len(rightPacket); i++ {
		left := leftPacket[i]
		right := rightPacket[i]

		if left.val != -1 && right.val != -1 { // both numbers
			if(left.val > right.val) {
				return -1
			}
			if(left.val < right.val) {
				return 1
			}
		} else {
			leftChildren := left.children
			rightChildren := right.children
			
			if left.val != -1 {
				leftChildren = []*Packet{left}
			}
			if right.val != -1 {
				rightChildren = []*Packet{right}
			}
			aux:= isOrdered(leftChildren, rightChildren)
			if aux != 0 {
				return aux
			}
		}
	}
	if(len(leftPacket) > len(rightPacket)) {
		return -1
	} else if (len(leftPacket) < len(rightPacket)) {
		return 1
	}
	return 0
}

func processLine(line string) *Packet{
	currentPacket := &Packet{val: -1}
	for i := 0; i < len(line); i++ {
		c := line[i]
		switch c {
			case '[':
				newPacket := &Packet{val: -1, parent: currentPacket}
				currentPacket.children = append(currentPacket.children, newPacket)
				currentPacket = newPacket
			case ']':
				currentPacket = currentPacket.parent
			default:
				valueString := ""
				if c != ',' {
				 	for c != ',' &&  c != '[' && c != ']'{
				  		valueString += string(c)
				 		i++
				  		c = line[i]
				 	}
				 	value, _ := strconv.Atoi(valueString)
				 	newPacket := &Packet{val: value}
				 	currentPacket.children = append(currentPacket.children, newPacket)
				 	i--
				}
		}
	}
	return currentPacket
}

func printChildren(packets []*Packet) {
	for _, packet := range packets {
	 	if(packet.val == -1) {
	  		fmt.Printf("[")
	  		printChildren(packet.children)
	  		fmt.Printf("]")
	 	} else {
	  		fmt.Printf("%v ", packet.val)
	 	}
	}   
}