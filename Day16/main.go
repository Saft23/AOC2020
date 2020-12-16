package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"strconv"
	"regexp"
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

func GetInvalidTicketsSum(data []string)int{
	var ranges = make(map[int]bool)
	var result = 0

	var AddRangeToMap = func(line string){
		var myRegex = regexp.MustCompile(`(?P<first>\d+-\d+) or (?P<second>\d+-\d+)`)

		match := myRegex.FindStringSubmatch(line)
		result := make(map[string]string)
		for i, name := range myRegex.SubexpNames(){
			if i != 0 && name != ""{
				result[name] = match[i]
			}
		}

		for _, value := range result{
			stringRanges := strings.Split(value, "-")

			lowerLimit, _ := strconv.Atoi(stringRanges[0])
			upperLimit, _ := strconv.Atoi(stringRanges[1])

			for i := lowerLimit; i <= upperLimit; i++{
				ranges[i] = true
			}
		}
	}

	var CheckIfInvalidTicket = func(line string){
		numbers := strings.Split(line, ",")

		for _, val := range numbers{
			valInt, _ := strconv.Atoi(val)
			if ranges[valInt]{
				continue
			}else{
				result = result + valInt
			}
		}
	}


	//Build maps
	var newLines = 0;
	for _, val := range data{
		if val == ""{
			newLines = newLines + 1
		}

		if val == "nearby tickets:"{
			continue
		}


		switch newLines{
			case 0:
			//Ranges
			AddRangeToMap(val)
			break

			case 1:
			//Nothing to do here
			break

			case 2:
			CheckIfInvalidTicket(val)
			break
		}
	}
	return result
}

func main(){
	data := ReadFile(input)

	part1 := GetInvalidTicketsSum(data)
	fmt.Println(part1)
	//part2 := GetNumberSpokenFromTurn(data, 30000000)
	//fmt.Println(part2)
}
