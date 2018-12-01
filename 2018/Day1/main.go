package main

import (
	"strconv"
	"fmt"
	"os"
	"strings"
	"io/ioutil"
	)

var (
	finalInput 	int
	finData		[]int
	foundInput 	[]int
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}

func contains(x int) bool {
	for _, v := range foundInput {
		if x == v {
			return true
		}
	}

	return false
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
	finalInput = 0

	Loop:
		for {
			for _, val := range finData {
				finalInput += val
				ok := contains(finalInput)
				foundInput = append(foundInput, finalInput)

				if ok {
					break Loop
				}
			}
		}

	fmt.Printf("Found final input twice: %d\n", finalInput)
}

func main() {
	strData := readFile("input.txt")
	arrData := strings.Split(strData, "\n")
	finData = convertArray(arrData)

	part1()
	part2()
}
