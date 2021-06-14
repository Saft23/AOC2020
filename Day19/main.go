package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	//"regexp"
)

var input = "input"
var input2 = "input2part2"
//var input = "input2"

var rules map[string]string
type rule struct{
	rule []int
	rule2 []int
	res string
	loop int
}

type combo struct{
	list []int
	loop int
}

func (r rule) split() bool{
	return len(r.rule) > 0 && len(r.rule2) > 0
}
func (r rule) end() bool{
	return len(r.res) > 0
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

func BuildRuleList(data []string)(map[int]rule,[]string){
	rules := make(map[int]rule)
	messages :=[]string{}
	for i:=0; i<len(data); i++{
		line := data[i]
		if(strings.Contains(line,":")) {
			if(strings.Contains(line, "\"")){
				//End rules
				res := strings.Split(line, ":")
				key, _ := strconv.Atoi(res[0])
				val := strings.ReplaceAll(res[1], "\"", "")
				r := rule{res: val}
				rules[key] = r
			}else{

				res := strings.Split(line, ":")
				key, _ := strconv.Atoi(res[0])
				if(strings.Contains(res[1], "|")){
					tmprule := strings.Split(res[1], "|")
					tmprule1 := strings.Split(tmprule[0], " ")
					tmprule2 := strings.Split(tmprule[1], " ")

					var tmpruleint1 = []int{}
					var tmpruleint2 = []int{}
					for _, i := range tmprule1{
						if (i == ""){
							continue
						}
						j, _ := strconv.Atoi(i)
						tmpruleint1 = append(tmpruleint1, j)
					}

					for _, i := range tmprule2{
						if (i == ""){
							continue
						}
						j, _ := strconv.Atoi(i)
						tmpruleint2 = append(tmpruleint2, j)
					}

					r := rule{rule:tmpruleint1,rule2:tmpruleint2}
					rules[key] = r
				}else{
					tmprule := strings.Split(res[1], " ")
					var tmpruleint = []int{}
					for _, i := range tmprule{
						if (i == ""){
							continue
						}
						j, _ := strconv.Atoi(i)
						tmpruleint = append(tmpruleint, j)
					}
					r := rule{rule: tmpruleint}
					rules[key] = r
				}
			}
		}else{
			messages = append(messages, line)
		}
	}
	return rules, messages
}

func RefineRules(){
	var changes = 0
	mapping := make(map[string]string)
	mapping["4"] = "a"
	mapping["5"] = "b"
	for _, val := range rules{
		for key, replace := range mapping{
			var change = strings.Contains(val, key)
			if change{
				changes = changes + 1
			}
			val = strings.ReplaceAll(val, key, replace)
		}
	}
}

func buildAllCombinations(rules map[int]rule)[]string{
	allCombinations := [][]int{}
	allCombinations = append(allCombinations, rules[0].rule)

	stepCombinations := func(oldCombinations [][]int)([][]int, int){
		fmt.Println(len(oldCombinations))
		changes := 0
		newCombinations := [][]int{}
		for i:=0; i<len(oldCombinations); i++{
				newComb := []int{}
			for j:=0; j<len(oldCombinations[i]); j++{
				ruleId := oldCombinations[i][j]
				rule, ok := rules[ruleId]

				if !ok{
					newComb = append(newComb, ruleId)
					continue
				}else {
					//Rule exists
					if rule.end(){
						//End rule
						if rule.res == " a"{
							newComb = append(newComb, 999)
						}else {
							newComb = append(newComb, 998)
						}
						changes = changes + 1
					} else if rule.split(){
						newCombCopy := make([]int, len(newComb))
						copy(newCombCopy, newComb)
						newCombCopy = append(newCombCopy, rule.rule...)
						newComb = append(newComb, rule.rule2...)

						newCombCopy = append(newCombCopy, oldCombinations[i][j+1:]...)
						newComb = append(newComb, oldCombinations[i][j+1:]...)

						newCombinations = append(newCombinations, newCombCopy)
						changes = changes + 1
						break

					}else {
						newComb = append(newComb, rule.rule...)
						changes = changes + 1
						//Single rule
					}

				}
			}
			newCombinations = append(newCombinations, newComb)
		}

		return newCombinations, changes
	}

	res, changes := stepCombinations(allCombinations)

	for changes != 0 {
		res, changes = stepCombinations(res)
	}
	filteredResult := filterToReadable(res)
	fmt.Println("Number of possible combos: ",len(filteredResult))
	return filteredResult
}
func buildAllCombinations2(rules map[int]rule)[]string{
	allCombinations := []combo{}
	allCombinations = append(allCombinations, combo{list:rules[0].rule, loop:0})

	stepCombinations := func(oldCombinations []combo)([]combo, int){
		fmt.Println(len(oldCombinations))
		changes := 0
		newCombinations := []combo{}
		for i:=0; i<len(oldCombinations); i++{
				newComb := combo{}
			for j:=0; j<len(oldCombinations[i].list); j++{
				if oldCombinations[i].loop > 5{
					//fmt.Println("FUCKED")
					oldCombinations[i] = combo{}
					continue
				}
				ruleId := oldCombinations[i].list[j]
				rule, ok := rules[ruleId]

				//fmt.Println(ruleId)
				if ruleId == 42{
					//fmt.Println("added")
					oldCombinations[i].loop = oldCombinations[i].loop + 1
				}

				if !ok{
					newComb.list = append(newComb.list, ruleId)
					newComb.loop = oldCombinations[i].loop
					continue
				}else {
					//Rule exists
					if rule.end(){
						//End rule
						if rule.res == " a"{
							newComb.list = append(newComb.list, 999)
							newComb.loop = oldCombinations[i].loop
						}else {
							newComb.list = append(newComb.list, 998)
							newComb.loop = oldCombinations[i].loop
						}
						changes = changes + 1
					} else if rule.split(){

						tmplist := make([]int, len(newComb.list))
						newCombCopy := combo{list:tmplist, loop:newComb.loop}
						copy(newCombCopy.list, newComb.list)
						newCombCopy.list = append(newCombCopy.list, rule.rule...)
						newComb.list = append(newComb.list, rule.rule2...)

						newCombCopy.list = append(newCombCopy.list, oldCombinations[i].list[j+1:]...)
						newCombCopy.loop = oldCombinations[i].loop
						newComb.list = append(newComb.list, oldCombinations[i].list[j+1:]...)
						newComb.loop = oldCombinations[i].loop

						newCombinations = append(newCombinations, newCombCopy)
						changes = changes + 1
						break

					}else {
						newComb.list = append(newComb.list, rule.rule...)
						newComb.loop = oldCombinations[i].loop
						changes = changes + 1
						//Single rule
					}

				}
			}
			newCombinations = append(newCombinations, newComb)
		}

		return newCombinations, changes
	}

	res, changes := stepCombinations(allCombinations)

	for changes != 0 {
		res, changes = stepCombinations(res)
	}
	filteredResult := filterToReadable2(res)
	fmt.Println("Number of possible combos: ",len(filteredResult))
	return filteredResult
}

func filterToReadable(possibleCombos [][]int)[]string{
	filteredList := []string{}
	for i:=0; i<len(possibleCombos); i++{
		strRow := ""
		for j:=0; j<len(possibleCombos[i]); j++{
			if possibleCombos[i][j] == 999{
				strRow = strRow + "a"
			}else if possibleCombos[i][j] == 998{
				strRow = strRow + "b"
			}else{
				fmt.Println("We fucked")
			}
		}
		filteredList = append(filteredList, strRow)
	}
	return filteredList
}

func filterToReadable2(possibleCombos []combo)[]string{
	filteredList := []string{}
	for i:=0; i<len(possibleCombos); i++{
		strRow := ""
		for j:=0; j<len(possibleCombos[i].list); j++{
			if possibleCombos[i].list[j] == 999{
				strRow = strRow + "a"
			}else if possibleCombos[i].list[j] == 998{
				strRow = strRow + "b"
			}else{
				fmt.Println("We fucked")
			}
		}
		filteredList = append(filteredList, strRow)
	}
	return filteredList
}

func calculateHowManyIsTrue(filteredList []string, messages []string)int{
	solutions := 0
	for _, m := range messages{
		for _, f := range filteredList{
			if m == f{
				solutions = solutions + 1
				break
			}
		}
	}
	return solutions
}
func check(e error)bool{
	if e != nil {
		panic(e)
	}
	return true
}






func main(){
	//data := ReadFile(input)
	data := ReadFile(input2)
	rules, messages := BuildRuleList(data)

	//filteredResult := buildAllCombinations(rules)
	//part1 := calculateHowManyIsTrue(filteredResult, messages)
	//fmt.Println("Part 1: ",part1)

	filteredResult := buildAllCombinations2(rules)
	part2 := calculateHowManyIsTrue(filteredResult, messages)
	fmt.Println("Part 2: ",part2)
}
