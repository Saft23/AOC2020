package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"regexp"
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

func simplify(line string, addPrio bool)string{
	regexPattern := regexp.MustCompile(`\([\*\+\d ]*\)`)
	matches := regexPattern.FindAllString(line, -1)
	for _, val := range matches{
		line = strings.ReplaceAll(line, val, (simplify(val[1:len(val)-1], addPrio)))
	}
	if addPrio && len(matches) > 0{
		line = prioritizeAdditions(line)
	}

	if len(matches) == 0{
		splits := strings.Split(line, " ")
		if len(splits) == 1 {
			return line
		}
		var result = 0
		for i:=0; i < len(splits)-1; i=i+2{
			first, _ := strconv.Atoi(splits[i])
			second, _ := strconv.Atoi(splits[i+2])
			op := splits[i+1]
			if i > 0{
				first = result
			}
			if op == "*"{
				result = first * second
			}else if op == "+"{
				result = first + second
			}
		}

		return strconv.Itoa(result)
	}else{
		return simplify(line, addPrio)
	}
	fmt.Println(line)
	return "Fuck"
}

func Part1(data []string)int64{
	var result int64 = 0
	for _, line := range data{
		tmpResult, _  := strconv.Atoi(simplify(line, false))
		result = result + int64(tmpResult)
	}
	return result
}


func Part2(data []string)int64{
var result int64 = 0
	for _, line := range data{

		line = prioritizeAdditions(line)
		tmpResult, _ := strconv.Atoi(simplify(line, true))
		result = result + int64(tmpResult)

	}
	return result
}

func prioritizeAdditions(line string)string{
	regexPatternPart := regexp.MustCompile(`\d+ \+ \d+`)
	matches2 := regexPatternPart.FindAllString(line, -1)
	for _, val := range matches2 {
	line = strings.ReplaceAll(line, val, "(" + val + ")")
	}
	return line
}

func main(){
	data := ReadFile(input)
	part1 := Part1(data)
	part2 := Part2(data)
	fmt.Printf("Part 1: %v\nPart 2: %v\n", part1, part2)
}
