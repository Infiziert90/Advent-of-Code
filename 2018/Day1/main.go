package main

import (
	"strconv"
	"fmt"
	"os"
	"strings"
	"io/ioutil"
	"time"
)

var (
	finalInput 	int
	finData		[]int
	foundInput = make(map[int]bool)
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}

func readFile(filename string) string {
	data, err := ioutil.ReadFile(filename)
	check(err)
	strData := string(data)

	return strData
}

func convertArray(arrData []string) []int {
	var finData []int

	for _, val := range arrData {
		i, err := strconv.Atoi(val)
		check(err)
		finData = append(finData, i)
	}

	return finData
}

func part1() {
	finalInput = 0

	for _, val := range finData {
		finalInput += val
	}

	fmt.Printf("Final Input is: %d\n", finalInput)
}

func part2() {
	t := time.Now()
	finalInput = 0

	Loop:
		for {
			for _, val := range finData {
				finalInput += val
				if foundInput[finalInput] {
					break Loop
				}
				foundInput[finalInput] = true
			}
		}

	fmt.Printf("Found final input twice: %d\n", finalInput)
	fmt.Println(time.Since(t))
}

func main() {
	strData := readFile("input.txt")
	arrData := strings.Split(strData, "\n")
	finData = convertArray(arrData)

	part1()
	part2()
}
