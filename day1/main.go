package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
)


var NUMBER_MAPPINGS = map[int]string {
    1: "one",
    2: "two",
    3: "three",
    4: "four",
    5: "five",
    6: "six",
    7: "seven",
    8: "eight",
    9: "nine",
} 

func main() {
    file, err := os.OpenFile("input.txt", os.O_RDONLY, os.ModePerm)
    if err != nil {
        fmt.Print("open file error: %v", err)
        return
    }
    defer file.Close()

    input := bufio.NewScanner(file)
    basicSum := 0
    complexSum := 0
    for input.Scan() {
	line := input.Text()
	basicSum += basicExtractNumber(line)
	complexSum += complexExtractNumber(line)
    }
    fmt.Printf("Part One: %d\n", basicSum)
    fmt.Printf("Part Two: %d\n", complexSum)
}

func basicExtractNumber(input string) int {
    first, last := 0, 0
    for _, char := range input {
	output := int(char) - '0'
	if output >= 0 && output <= 9 {
	    if first == 0 {
		first = output;
	    }
	    last = output;
	}
    }

    return (first * 10) + last
}

func complexExtractNumber(input string) int {
    indexMap := map[int]int{}
    for i := 1; i < 10; i++ {
	tempInput := input
	for strings.Contains(tempInput, NUMBER_MAPPINGS[i]) {
	    index := strings.Index(tempInput, NUMBER_MAPPINGS[i])
	    indexMap[index] = i
	    expunge := getExpungeString(NUMBER_MAPPINGS[i])
	    tempInput = strings.Replace(tempInput, NUMBER_MAPPINGS[i], expunge, 1)
	}
	for strings.Contains(tempInput, strconv.Itoa(i)) {
	    index := strings.Index(tempInput, strconv.Itoa(i))
	    indexMap[index] = i
	    tempInput = strings.Replace(tempInput, strconv.Itoa(i), "-", 1)
	}
    }
    minIndex := -1
    maxIndex := -1
    for k := range indexMap {
	if minIndex == -1 {
	    minIndex = k
	}
	if k < minIndex {
	    minIndex = k
	}
	if k > maxIndex {
	    maxIndex = k
	}
    }
    return (10 * indexMap[minIndex]) + indexMap[maxIndex]
}

func getExpungeString(input string) string {
    output := ""
    for i := 0; i < len(input); i++ {
	output += "-"
    }
    return output
}

