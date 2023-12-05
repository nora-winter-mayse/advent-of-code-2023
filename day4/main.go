package main

import (
    "bufio"
    "fmt"
    "os"
    "slices"
    "sort"
    "strings"
    "strconv"
    "regexp"

)


func main() {
    file, err := os.OpenFile("input.txt", os.O_RDONLY, os.ModePerm)
    if err != nil {
        fmt.Print("open file error: %v", err)
        return
    }
    defer file.Close()

    input := bufio.NewScanner(file)
    partOneSum := 0
    partTwoSum := 0
    cards := []*Card{}
    for input.Scan() {
	line := input.Text()
	cards = append(cards, parseLine(line))
    }

    for _, card := range cards {
	partOneSum += getPoints(card)
	partTwoSum += getCopies(card, cards)
    } 

    fmt.Printf("Part One: %d\n", partOneSum)
    fmt.Printf("Part Two: %d\n", partTwoSum)
}

func parseLine(line string) *Card {
    card := newCard()
    colonSplit := strings.Split(line, ":")
     
    re := regexp.MustCompile(`[0-9]+`)
    cardNumber, _ := strconv.Atoi(string(re.Find([]byte(colonSplit[0]))))
    card.cardNumber = cardNumber 
    
    pipeSplit := strings.Split(colonSplit[1], "|")
    card.winningNumbers = getNumbers(pipeSplit[0])
    card.numbers = getNumbers(pipeSplit[1])
    
    sort.Ints(card.winningNumbers)
    sort.Ints(card.numbers)
    return card 
}

func getNumbers(input string) []int {
    output := []int{}
    re := regexp.MustCompile(`[0-9]+`)
    numberBytes := re.FindAll([]byte(input), -1)
    for _, element := range numberBytes {
	elementInt, _ := strconv.Atoi(string(element))
	output = append(output, elementInt)
    }
    return output
}

type Card struct {
    winningNumbers []int
    numbers []int
    cardNumber int
    copies int
}

func newCard() *Card {
    card := Card{
	winningNumbers: []int{},
	numbers: []int{},
	copies: 1,
    }
    return &card
}

func getPoints(card *Card) int {
    points := 0
    for _, number := range card.numbers {
	if _, found := slices.BinarySearch(card.winningNumbers, number); found {
	    if points == 0 {
		points = 1
	    } else {
		points *= 2
	    }
	}
    }    
    return points
}

func getCopies(card *Card, cards []*Card) int {
    matches := 0
    for _, number := range card.numbers {
	if _, found := slices.BinarySearch(card.winningNumbers, number); found {
	    matches += 1
	}
    }
    baseIndex, _ := slices.BinarySearchFunc(cards, card, compareCardNumbers)
    baseIndex++ 
    for i := baseIndex; i < baseIndex + matches; i++ {
	cards[i].copies += card.copies
    }
    return card.copies
}

func compareCardNumbers(a, b *Card) int {
    return a.cardNumber - b.cardNumber    
}
