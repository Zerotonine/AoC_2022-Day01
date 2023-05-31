package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const INPUT_PATH = "data/input.txt"

func main() {
	fmt.Printf("Solution 1: %d\n", solutionOne())
	fmt.Printf("Solution 2: %d\n", solutionTwo())
}

func solutionOne() int64 {
	file := getFile()
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	var temp int64 = 0
	var highest int64 = 0

	for fileScanner.Scan() {
		trimmed := strings.TrimSpace(fileScanner.Text())
		if len(trimmed) == 0 {
			if temp > highest {
				highest = temp
			}
			temp = 0
			continue
		}
		temp += castToInt64(trimmed)
	}
	return highest
}

func solutionTwo() int64 {
	file := getFile()
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	var temp int64 = 0
	topThree := make([]int64, 3)

	for fileScanner.Scan() {
		trimmed := strings.TrimSpace(fileScanner.Text())
		if len(trimmed) == 0 {
			swapHighest(&topThree, &temp)
			temp = 0
			continue
		}
		temp += castToInt64(trimmed)
	}
	return sumIntSlice(&topThree)
}

func swapHighest(intSlice *[]int64, value *int64) {
	for i := 0; i < cap(*intSlice); i++ {
		if *value > (*intSlice)[i] {
			(*intSlice)[i] = *value
			return
		}
	}
}

func sumIntSlice(intSlice *[]int64) int64 {
	var temp int64 = 0
	for _, value := range *intSlice {
		temp += value
	}
	return temp
}

func getFile() *os.File {
	f, err := os.Open(INPUT_PATH)

	if err != nil {
		panic("can't open file")
	}

	return f
}

func castToInt64(s string) int64 {
	val, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatalf("can't convert %s to int\n", s)
	}
	return val
}
