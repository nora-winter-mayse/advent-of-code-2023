package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strconv"
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

    input.Scan()
    times := getNumbers(input.Text())
    input.Scan()
    dist := getNumbers(input.Text())
    
    races := []*Race{}
    for i := 0; i < len(times); i++ {
	races = append(races, newRace(times[i], dist[i]))
    }
    
    partOneSum = 1
    for _, race := range races {
	successCount := 0
	for holdDuration := 0; holdDuration < race.duration; holdDuration++ {
	    runDuration := race.duration - holdDuration
	    distance := holdDuration * runDuration
	    if distance > race.goalDistance {
		successCount++
	    }
	}
	if successCount != 0 {
	    partOneSum *= successCount
	}
    }
    
    race := newRace(makeBigNumber(times), makeBigNumber(dist))
    for holdDuration := 0; holdDuration < race.duration; holdDuration++ {
        runDuration := race.duration - holdDuration
        distance := holdDuration * runDuration
        if distance > race.goalDistance {
            partTwoSum++
        }
    }


    fmt.Printf("Part One: %d\n", partOneSum)
    fmt.Printf("Part Two: %d\n", partTwoSum)
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

func makeBigNumber(numbers []int) int {
    bigNumber := ""
    for _, number := range numbers {
	bigNumber += strconv.Itoa(number)
    }
    output, _ := strconv.Atoi(bigNumber)
    return output
}

type Race struct {
    duration int
    goalDistance int
}

func newRace(duration, goalDistance int) *Race {
    return &Race {
	duration: duration,
	goalDistance: goalDistance,
    }
}
