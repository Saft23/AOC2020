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

func simplify(line string)string{
	regexPattern := regexp.MustCompile(`\([\*\+\d ]*\)`)
	matches := regexPattern.FindAllString(line, -1)
	for _, val := range matches{
		line = strings.ReplaceAll(line, val, (simplify(val[1:len(val)-1])))
	}

	if len(matches) == 0{
		splits := strings.Split(line, " ")
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
		return simplify(line)
	}
	fmt.Println(line)
	return "Fuck"
}

func Part1(data []string)int64{
	var result int64 = 0
	for _, line := range data{
		tmpResult, _  := strconv.Atoi(simplify(line))
		result = result + int64(tmpResult)
	}
	return result
}



func main(){
	data := ReadFile(input)
	part1 := Part1(data)
	fmt.Printf("Part 1: %v", part1)
}
