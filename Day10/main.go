package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"sort"
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

func BuildSingleAdapterList(data []string)int{
	var intData []int
	var adapterList []int

	//Cast to int
	for index, _ := range data{
		intVal, _ := strconv.Atoi(data[index])
		intData = append(intData, intVal)
	}

	sort.Ints(intData)
	adapterList = append(adapterList, 0)

	var singleJoltDifferenceCounter = 0
	var trippleJoltDifferenceCounter = 0
	for _, val := range intData{
		var latestElement = adapterList[len(adapterList)-1:len(adapterList)][0]
		var difference = val - int(latestElement)
		if difference > 0 && difference < 4{
			adapterList = append(adapterList, val)
			if difference == 3{
				trippleJoltDifferenceCounter = trippleJoltDifferenceCounter + 1
			}else if difference == 1{
				singleJoltDifferenceCounter = singleJoltDifferenceCounter + 1
			}
		}
	}
	//Adding a tripple to final adapter
	trippleJoltDifferenceCounter = trippleJoltDifferenceCounter + 1
	return singleJoltDifferenceCounter * trippleJoltDifferenceCounter
}

func AllCombinationsOfAdapters(data []string)int{
	var intData []int
	for index, _ := range data{
	intVal, _ := strconv.Atoi(data[index])
	intData = append(intData, intVal)
	}
	sort.Ints(intData)
	lastElement := intData[len(intData)-1]

	intData = append(intData, 0)
	intData = append(intData, lastElement+3)
	sort.Ints(intData)

	var CheckIfValueExists func(a int)bool
	CheckIfValueExists = func(a int) bool {
	  for _, b := range intData {
		  if b == a {
			  return true
		  }
	  }
	  return false
  }

	paths := make([]int,intData[len(intData)-1]+1)

	paths[0] = 1

	for i := 1; i < intData[len(intData)-1]+1; i++{
		for j := 1; j < 4; j++{
			if CheckIfValueExists(i-j){
				paths[i] = paths[i] + paths[i-j]
			}
		}
	}
	return paths[len(paths)-1]
}
func main(){

	var data = ReadFile(input)
	var adapterList = BuildSingleAdapterList(data)
	fmt.Printf("Part 1: %v\n", adapterList)

	var part2 = AllCombinationsOfAdapters(data)
	fmt.Printf("Part 2: %v\n", part2)
}
