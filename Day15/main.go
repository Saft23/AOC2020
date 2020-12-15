package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"strconv"
	//"math"
)

var input = "input"


func ReadFile(input string) (text []string){
	file, err := os.Open(input)
	check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan(){
		text = append(text, scanner.Text())
	}
	file.Close()
	return text
}

func check(e error)bool{
	if e != nil {
		panic(e)
	}
	return true
}

func GetNumberSpokenFromTurn(data []string, turns int)int{
	values := strings.Split(data[0], ",")
	var sinceLastSpoken = make(map[int]int)
	for index, val := range values{
		valInt, _ := strconv.Atoi(val)
		sinceLastSpoken[valInt] = index+1
	}

	var turn int = len(values)
	lastSpokenNumber := 0

	for{
		turn = turn + 1
		lastTurn, ok := sinceLastSpoken[lastSpokenNumber]
		sinceLastSpoken[lastSpokenNumber] = turn
		if !ok{
			lastSpokenNumber = 0
		}else{
			lastSpokenNumber = turn - lastTurn
		}

		if turn == turns-1{
			return lastSpokenNumber
		}

	}
	return turn
}

func main(){
	var data = ReadFile(input)

	part1 := GetNumberSpokenFromTurn(data, 2020)
	part2 := GetNumberSpokenFromTurn(data, 30000000)
	fmt.Println(part1)
	fmt.Println(part2)
}
