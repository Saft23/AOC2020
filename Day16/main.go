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
func contains(s []int, e int) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func FindFieldsAndGetPart2Result(data []string)int{
	var ranges = make(map[string]map[int]bool)
	var newLines = 0
	var result = 0
	var myTicket string
	var validTickets []string
	var RowsNotFittingCols map[string][]int
	var AssignedRowsToCols map[string]int

	var SkippedCols []int

	var CheckIfValidTicketAndAddToList = func(line string){
		numbers := strings.Split(line, ",")
		for _, val := range numbers{
			valInt, _ := strconv.Atoi(val)
			var flag = false
			for _, submap := range ranges{
				if _, ok := submap[valInt]; ok{
					flag = true
				}
			}
			if !flag{
				return
			}
		}
		validTickets = append(validTickets, line)
	}
	var AddRangeToMap = func(line string){
		var myRegex = regexp.MustCompile(`(?P<name>[\w\s].*): (?P<first>\d+-\d+) or (?P<second>\d+-\d+)`)

		match := myRegex.FindStringSubmatch(line)
		result := make(map[string]string)
		for i, name := range myRegex.SubexpNames(){
			if i != 0 && name != ""{
				result[name] = match[i]
			}
		}

		first := result["first"]
		second := result["second"]
		name := result["name"]

		firstRanges := strings.Split(first, "-")
		secondRanges := strings.Split(second, "-")

		firstLowerLimit, _ := strconv.Atoi(firstRanges[0])
		firstUpperLimit, _ := strconv.Atoi(firstRanges[1])
		secondLowerLimit, _ := strconv.Atoi(secondRanges[0])
		secondUpperLimit, _ := strconv.Atoi(secondRanges[1])

		ranges[name] = make(map[int]bool)
		for i := firstLowerLimit; i <= firstUpperLimit; i++{
			ranges[name][i] = true
		}
		for i := secondLowerLimit; i <= secondUpperLimit; i++{
			ranges[name][i] = true
		}
	}
	var sortColsAndRows = func(){

	RowsNotFittingCols = make(map[string][]int)
		for row, submap := range ranges{
			for _, ticket := range validTickets{
				numbers := strings.Split(ticket, ",")
				for i:=0; i<len(numbers);i++{
					tmpNumber, _ := strconv.Atoi(numbers[i])
					if _, ok := submap[tmpNumber];!ok{
						RowsNotFittingCols[row] = append(RowsNotFittingCols[row], i)
					}
				}
			}

			if len(RowsNotFittingCols[row]) == 0 {
				RowsNotFittingCols[row] = nil
			}
		}
	}
	var AssignRowsToCols = func(){

		AssignedRowsToCols = make(map[string]int)
		for j := 0; j <= 25; j++{

		maxSizeOfTickets := 20 //3, 20

		var maxSize  = -1
		var RowWithMostNotFitting = ""
		for row, val := range RowsNotFittingCols{
			if len(val) > maxSize{
				maxSize = len(val)
				RowWithMostNotFitting = row
			}
		}
		tmpRow := RowsNotFittingCols[RowWithMostNotFitting]
		for i := 0; i <= maxSizeOfTickets; i++{
			if contains(SkippedCols, i){
				continue
			}
			if !contains(tmpRow, i){
				AssignedRowsToCols[RowWithMostNotFitting] = i
				SkippedCols = append(SkippedCols, i)
			}
		}
		delete(RowsNotFittingCols, RowWithMostNotFitting)
		}

	}
	var CalculateFromMyTicket = func()int{
		myTicketNumbers := strings.Split(myTicket, ",")
		var myTicketResult = 1
		for key, val := range AssignedRowsToCols{
			if strings.HasPrefix(key, "departure"){
				multi, _ := strconv.Atoi(myTicketNumbers[val])
				myTicketResult = myTicketResult * multi
			}
		}
		return myTicketResult
	}

	for _, val := range data{
		if val == ""{
			newLines = newLines + 1
		}

		if val == "nearby tickets:" {
			continue
		}
		if val == "your ticket:" {
			continue
		}

		switch newLines{
			case 0:
			//Ranges
			AddRangeToMap(val)
			break

			case 1:
			myTicket = val
			break

			case 2:
			CheckIfValidTicketAndAddToList(val)
			break
		}
	}
	sortColsAndRows()
	AssignRowsToCols()
	result = CalculateFromMyTicket()
	return result
	
}

func main(){
	data := ReadFile(input)
	part1 := GetInvalidTicketsSum(data)
	part2 := FindFieldsAndGetPart2Result(data)

	fmt.Printf("Part 1: %v, part 2: %v", part1, part2)
}
