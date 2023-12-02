package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
)

const MAX_RED = 12
const MAX_GREEN = 13
const MAX_BLUE = 14
    
func main() {
    file, err := os.OpenFile("input.txt", os.O_RDONLY, os.ModePerm)
    if err != nil {
        fmt.Print("open file error: %v", err)
        return
    }
    defer file.Close()

    input := bufio.NewScanner(file)
    gamesSum := 0
    powerSetSum := 0
    for input.Scan() {
	line := input.Text()
	gamesSum += extractNumber(line)
	powerSetSum += getPowerSet(line)
    }
    fmt.Printf("Part One: %d\n", gamesSum)
    fmt.Printf("Part Two: %d\n", powerSetSum)
}

func extractNumber(line string) int {
    inputs := strings.Split(line, ":")
    gameNumber, _ := strconv.Atoi(strings.Split(inputs[0], " ")[1])
    if isGamePossible(inputs[1]) {
	return gameNumber
    }
    return 0
}

func isGamePossible(input string) bool {
    red, green, blue := getMaximumsForGame(input) 
    return red <= MAX_RED && green <= MAX_GREEN && blue <= MAX_BLUE
    
}

func getPowerSet(input string) int {
    red, green, blue := getMaximumsForGame(strings.Split(input, ":")[1])
    return red * green * blue
}

func getMaximumsForGame(input string) (int, int, int) {
    reveals := strings.Split(input, ";")
    red, green, blue := 0, 0, 0
    for _, reveal := range reveals {
	currentRed, currentGreen, currentBlue := processReveal(reveal)
	if currentRed > red {
	    red = currentRed
	}
	if currentGreen > green {
	    green = currentGreen
	}
	if currentBlue > blue {
	    blue = currentBlue
	}
    }
    return red, green, blue
}

func processReveal(input string) (int, int, int) {
    colors := strings.Split(input, ",")
    red, green, blue := 0, 0, 0
    for _, color := range colors {
	colorReveal := strings.Split(color, " ")
	colorName := colorReveal[2]
	if colorName == "red" {
	    red, _ = strconv.Atoi(colorReveal[1])
	} else if colorName == "green" {
	    green, _ = strconv.Atoi(colorReveal[1])
	} else if colorName == "blue" {
	    blue, _ = strconv.Atoi(colorReveal[1])
	}
    }
    return red, green, blue
}
