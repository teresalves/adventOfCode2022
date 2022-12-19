package main  

import (
    "fmt"
    "os"
	"bufio"
    "strings"
    "strconv"
) 
 
func main() {  

	readFile, err := os.Open("input.txt")
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
    var fileLines []string

	var resultMap = make(map[string]int) // key is filepath
  
    for fileScanner.Scan() {
        fileLines = append(fileLines, fileScanner.Text())
    }
	childDirectories := goThroughDirectories(fileLines, resultMap)
    recalculateSizes(resultMap, childDirectories, "//")
	result := getSmallestDeletableDir(resultMap)

    fmt.Println(result)
}

func goThroughDirectories(fileLines []string, resultMap map[string]int) map[string][]string {
    var childDirectories = make(map[string][]string);
    currPath := ""
    for _,line := range fileLines {
        input := strings.Split(line, " ")
        if input[0] == "$" && input[1] == "cd" && input[2] != ".." {
            currPath = currPath + "/" + input[2]
            childDirectories[currPath] = []string{}
        } else if input[0] == "$" && input[1] == "cd" && input[2] == ".." {
            dirs := strings.Split(currPath,"/")
            currPath = currPath[:len(currPath) - len(dirs[len(dirs)-1]) - 1] // remove last path
        } else if input[0] == "dir" {
            child := currPath + "/" + input[1]
            childDirectories[currPath] = append(childDirectories[currPath], child)
        } else if input[0] != "$" && input[0] != "dir" {
            value, _ := strconv.Atoi(input[0])
            if _, ok := resultMap[currPath]; ok {
                resultMap[currPath] += value
            } else {
                resultMap[currPath] = value
            }
        }
    }
    return childDirectories
}


func recalculateSizes(resultMap map[string]int, childDirectories map[string][]string, key string) {
    children := childDirectories[key]
    if(len(children) != 0 ) {
        for _,child := range children {
            recalculateSizes(resultMap, childDirectories, child)
            resultMap[key] += resultMap[child]
        }
    }
}

func sumResult(resultMap map[string]int) int{
    result := 0
    for _, value := range resultMap {
        if value < 100000 {
            result += value
        }
    }
    return result
}

func getSmallestDeletableDir(resultMap map[string]int) int{
    currOccupiedSpace := resultMap["//"]
    totalSpace := 70000000
    spaceNeeded := 30000000
    spaceRequired := spaceNeeded - (totalSpace - currOccupiedSpace)
    currMinValue := totalSpace
    for _, value := range resultMap {
        if value < currMinValue && value > spaceRequired {
            currMinValue = value
        }
    }
    return currMinValue
}
