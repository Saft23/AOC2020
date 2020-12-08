package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"strconv"
)

var input = "input"

type Rule map[string]Packing
type Packing map[string]int

func check(e error)bool{
	if e != nil {
		panic(e)
	}
	return true
}

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

func CountAmountOfBagsWithShinyGold(rules Rule)int{
	var result = 0
	var recursivePacking func(packing Packing) bool

	recursivePacking = func(packing Packing) bool {
		var result = false
		if _, ok := packing["shinygoldbag"]; ok {
			result = true
		}else if _, ok := packing["ootherbag"]; !ok {
			for bag, _ := range packing {
				var tmpResult = recursivePacking(rules[bag])
				if !result {
					result = tmpResult
				}
			}
		}
		return result
	}
	for _, rule := range rules {
		if recursivePacking(rule){
			result = result + 1
		}
	}
	return result 
}

func CountAmountOfBagsWithinShinyGoldBag(rules Rule)int{
	var recursivePacking func(packing Packing) int
	recursivePacking = func(packing Packing) int {
		var result = 1
		if _, ok := packing["ootherbag"]; ok {
			return result
		}else{
			for bag, _ := range packing {
				result = result + recursivePacking(rules[bag]) * packing[bag]
			}
		}
		return result
	}
	return recursivePacking(rules["shinygoldbag"])-1
}

func BuildRuleList(data []string) Rule {
	rule := Rule{}
	for _, line := range data {
		packing := Packing{}
		line = strings.ReplaceAll(line, " bags", " bag")
		var lineList = strings.Split(line, "contain")
		var lineContents = strings.Split(lineList[1], ",")
		var origBag = strings.ReplaceAll(lineList[0], " ", "")

		for _, bag := range lineContents {
			bag = strings.ReplaceAll(bag, " ", "")
			bag = strings.ReplaceAll(bag, ".", "")
			var numberOfBags, _ = strconv.Atoi(string(bag[0]))
			var bagName = bag[1:]
			packing[bagName] = int(numberOfBags)
		}
		rule[origBag] = packing
	}
	return rule
}

func main(){
	var data = ReadFile(input)
	var ruleMap = BuildRuleList(data)
	var amountOfGoldBags = CountAmountOfBagsWithShinyGold(ruleMap)
	fmt.Printf("Part 1: %v\n", amountOfGoldBags)
	var amountOfBags = CountAmountOfBagsWithinShinyGoldBag(ruleMap)
	fmt.Printf("Part 2: %v\n", amountOfBags)
}
