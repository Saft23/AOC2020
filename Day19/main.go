package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	//"regexp"
)

var input = "input2"

var rules map[string]string
type rule struct{
	rule []int
	rule2 []int
	res string
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

func BuildRuleList(data []string)map[int]rule{
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
				//Rules
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
				fmt.Println(line)
			}
		}else{
			messages = append(messages, line)
			//Messages data
			fmt.Println(line)
		}
	}
	fmt.Println(rules)
	return rules
}


func CalculateAllPossibleCombinations(allRules map[int]rule)[]string{
	allCombinations := []string{}
	orig := allRules[0].rule

	mul := func(r rule)bool{
		return len(r.rule) > 0 && len(r.rule2) > 0
	}
	fmt.Println(mul(allRules[0]))
	end := func(r rule)bool{
		return len(r.res) > 0
	}

	for _, val := range orig{
		if end(allRules[val]){
			fmt.Println(allRules[val].res)
		}
	}
	return allCombinations
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

func check(e error)bool{
	if e != nil {
		panic(e)
	}
	return true
}

func main(){
	data := ReadFile(input)
	rules := BuildRuleList(data)
	fmt.Println(rules)
	//part1 := Part1(data)

	allCombinations := CalculateAllPossibleCombinations(rules)
	fmt.Println(allCombinations)
}
