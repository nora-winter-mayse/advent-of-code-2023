package main

import (
    "bufio"
    "fmt"
    "os"
    "slices"
    "strconv"
    "strings"
)

var powerMappings = map[string]int {
    "1": 1,
    "2": 2,
    "3": 3,
    "4": 4,
    "5": 5,
    "6": 6,
    "7": 7,
    "8": 8,
    "9": 9,
    "T": 10,
    "J": 0,
    "Q": 12,
    "K": 13,
    "A": 14,
}

func main() {
    file, err := os.OpenFile("input.txt", os.O_RDONLY, os.ModePerm)
    if err != nil {
        fmt.Print("open file error: %v", err)
        return
    }
    defer file.Close()

    input := bufio.NewScanner(file)

    hands := []*Hand{}
    for input.Scan() {
	text := strings.Split(input.Text(), " ")
	cardsText := text[0]
	bid, _ := strconv.Atoi(text[1])
	hands = append(hands, newHand(cardsText, bid))
    }
    slices.SortFunc(hands, compareHands)
    partOneSum := 0
    for i, hand := range hands {
	rank := i + 1
	fmt.Printf("Bid %d has rank %d\n", hand.bid, rank)
	partOneSum += rank * hand.bid
    }
    fmt.Printf("Part One: %d\n", partOneSum)
}

type Hand struct {
    cards []*Card
    bid int
    power int
}

func newHand(input string, bid int) *Hand {
    cards := []*Card{}
    characterOccurance := map[string]int{}
    for key, _ := range powerMappings {
	characterOccurance[key] = 0
    }
    jaysContained := 0
    for _, character := range input {
	cards = append(cards, newCard(string(character)))
	if (string(character) == "J") {
	    jaysContained++
	} else {
	    characterOccurance[string(character)] += 1
	}
    }
    max := 0
    twosFound := 0
    for _, power := range characterOccurance {
	if power > max {
	    max = power
	}
	if power == 2 {
	    twosFound++
	}
    }
   
    if jaysContained == 1 {
	if twosFound > 0 {
	    twosFound--
	}
    }
    
    max += jaysContained
    if max > 3 || (max == 3 && twosFound > 0) {
	max+=2
    } else if max == 3 || twosFound == 2 {
	max++
    }
    
    
    return &Hand{
	cards: cards,
	bid: bid,
	power: max,
    }
}

func compareHands(a, b *Hand) int {
    if a.power != b.power {
	return a.power - b.power
    }
    for i := 0; i < len(a.cards); i++ {
	power := compareCards(a.cards[i], b.cards[i])
	if power != 0 {
	    return power
	}
    }
    return 0
}

type Card struct {
    letter string
    power int
}

func newCard(letter string) *Card {
    return &Card{
	letter: letter,
	power: powerMappings[letter],
    }
}

func compareCards(a, b *Card) int {
    return a.power - b.power
}
