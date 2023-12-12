package main

import (
    "fmt"
    "os"
    "regexp"
    "strconv"
    "strings"
)


func main() {
    file, err := os.ReadFile("input.txt")
    if err != nil {
	fmt.Printf("File failed to read: %v", err) 
    }
    inputs := []int{}
    inputsPartTwo := []int{}
    sections := []*Section{}
    for i, inputSection := range strings.Split(string(file), "---") {
	if i == 0 {
	    inputs = getNumbers(strings.Split(inputSection, ":")[1])
	    continue
	}
	lines := strings.Split(inputSection, "\n")
	section := newSection()
	for _, line := range lines {
	    numbers := getNumbers(line)
	    if len(numbers) == 0 {
		continue
	    }
	    section.lookups = append(section.lookups, newLookup(numbers[1], numbers[0], numbers[2]))
	}
	sections = append(sections, section)
    }

    for i, input := range inputs {
	if i % 2 == 0 {
	    for k := input; k < input + inputs[i + 1]; k++ {
		inputsPartTwo = append(inputsPartTwo, k)
	    }  
	}
    }
    
    fmt.Printf("Size of part 2: %d\n", len(inputsPartTwo))
    
    fmt.Printf("Part One: %d\n", process(inputs, sections))
    fmt.Printf("Part Two: %d\n", process(inputsPartTwo, sections))
    
}

func process(inputs []int, sections []*Section) int {
    for i, _ := range inputs {
	for _, section := range sections {
	    inputs[i] = mapSection(section, inputs[i])
	}
	
    }
   
    min := inputs[0]
    for _, input := range inputs {
	if input < min {
	    min = input
	}
    } 
    return min
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

type Section struct {
    lookups []*Lookup
}

type Lookup struct {
    start int
    dest int 
    offset int
}

func newSection() *Section {
    return &Section {
	lookups: []*Lookup{},
    }
}

func newLookup(start, dest, offset int) *Lookup {
    return &Lookup {
	start: start,
	dest: dest,
	offset: offset,
    }
}

func shouldMap(lookup *Lookup, input int) bool {
    return input >= lookup.start && input < lookup.start + lookup.offset
}

func mapLookup(lookup *Lookup, input int) int {
    return lookup.dest + (input - lookup.start)
}

func mapSection(section *Section, input int) int {
    for _, lookup := range section.lookups {
	if (shouldMap(lookup, input)) {
	    return mapLookup(lookup, input)
	}
    }
    return input
}
