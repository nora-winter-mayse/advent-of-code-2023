package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strings"
    "strconv"
    "slices"
)


var symbols = []string{"@", "/", "#", "$", "%", "&", "*", "-", "+", "="}

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
    prev := ""
    cur := ""
    next := ""
    i := 1 
    gearMappings := map[string]int{}
    encounteredNumbers := []GearEntry{}
    for input.Scan() {
	prev = cur
	cur = next
	//pad ends with periods to avoid look-behind index out of bounds shenanigans
	next = "." + input.Text() + "."
	if cur == "" {
	    cur = getPaddingRow(len(next))
	    continue
	}
	basicSum += process(prev, cur, next)
	gearMappings, encounteredNumbers = processGear(prev, cur, next, i, gearMappings, encounteredNumbers)
	i++
    }
    
    //process the final row
    prev = cur
    cur = next
    next = getPaddingRow(len(cur))
    basicSum += process(prev, cur, next)
    gearMappings, encounteredNumbers = processGear(prev, cur, next, i, gearMappings, encounteredNumbers)
    
    fmt.Printf("Part One: %d\n", basicSum)

    for x := 0; x < len(encounteredNumbers); x++ {
	for y:= x + 1; y < len(encounteredNumbers); y++ {
	    dirtyList := []string{}
	    for _, xElement := range encounteredNumbers[x].gearHash {
		for _, yElement := range encounteredNumbers[y].gearHash {
		    if !slices.Contains(dirtyList, xElement) && xElement == yElement && gearMappings[xElement] == 2 {
			complexSum += encounteredNumbers[x].value * encounteredNumbers[y].value
			dirtyList = append(dirtyList, xElement)
			fmt.Printf("Adding %d * %d\n", encounteredNumbers[x].value, encounteredNumbers[y].value)
		    }
		}
	    }
	}
    }
    fmt.Printf("Part Two: %d\n", complexSum)
}

func processGear(prev string, cur string, next string, rowIndex int, gearMappings map[string]int, encounteredNumbers []GearEntry) (map[string]int, []GearEntry) {
    
    re := regexp.MustCompile(`[0-9]+`)
    matches := re.FindAllIndex([]byte(cur), -1)
    for _, regexMatch := range matches {
	startIndex := regexMatch[0]
	endIndex := regexMatch[1]
	val, _ := strconv.Atoi(cur[startIndex:endIndex])
	gearEntry := newGearEntry(val)
    	for i := startIndex - 1; i < endIndex + 1; i++ {
	    if prev[i] == '*' {
		gearHash := hash(rowIndex - 1, i)
		gearEntry.gearHash = append(gearEntry.gearHash, gearHash)
		if _, ok := gearMappings[gearHash]; !ok {
		    gearMappings[gearHash] = 1
		    fmt.Printf("Saw first time")
		} else {
		    gearMappings[gearHash] += 1
		    fmt.Printf("Saw second time")
		}
	    }
	    if cur[i] == '*' {
		gearHash := hash(rowIndex, i)
		gearEntry.gearHash = append(gearEntry.gearHash, gearHash)
		if _, ok := gearMappings[gearHash]; !ok {
		    gearMappings[gearHash] = 1
		    fmt.Printf("Saw first time on cur")
		} else {
		    gearMappings[gearHash] += 1
		    fmt.Printf("Saw second time on cur")
		}
	    }
	    if next[i] == '*' {
		gearHash := hash(rowIndex + 1, i)
		gearEntry.gearHash = append(gearEntry.gearHash, gearHash)
		if _, ok := gearMappings[gearHash]; !ok {
		    gearMappings[gearHash] = 1
		    fmt.Printf("Saw first time on next")
		} else {
		    gearMappings[gearHash] += 1
		    fmt.Printf("Saw second time on next")
		}
	    }
	}
	for _, stuff := range gearEntry.gearHash {
	    fmt.Printf("Gear Entry Hash Value: %d\n", stuff)
	} 
	encounteredNumbers = append(encounteredNumbers, *gearEntry)
    }
    return gearMappings, encounteredNumbers
}



type GearEntry struct {
    value int
    gearHash []string
}

func newGearEntry(value int) *GearEntry {
    gearEntry := GearEntry{
	value: value,
	gearHash: []string{},
    }
    return &gearEntry
}

func hash(x int, y int) string {
    return fmt.Sprintf("%d,%d", x, y) 
}

func process(prev string, cur string, next string) int {
    re := regexp.MustCompile(`[0-9]+`)
    matches := re.FindAllIndex([]byte(cur), -1)
    sum := 0
    for _, regexMatch := range matches {
	startIndex := regexMatch[0]
	endIndex := regexMatch[1]
	if doesContainSymbol(prev[startIndex - 1:endIndex + 1]) ||
	    doesContainSymbol(cur[startIndex - 1: startIndex]) ||
	    doesContainSymbol(cur[endIndex: endIndex + 1]) ||
	    doesContainSymbol(next[startIndex - 1:endIndex + 1]) {
	    element, _ :=  strconv.Atoi(cur[startIndex:endIndex])
	    sum += element
	}	
    }
    return sum
}

func doesContainSymbol(input string) bool {
    for _, symbol := range symbols {
	if strings.Contains(input, symbol) {
	    return true
	}
    }
    return false
}

func getPaddingRow(count int) string {
    output := ""
    for i := 0; i < count; i++ {
	output += "."
    } 
    return output
}
