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
	var adapterList []int

	var GetCombinationsFromCurrentState2 func(list []int)int

	//NOW!
	GetCombinationsFromCurrentState2 = func(list []int)int{
		fmt.Println("HELLO")
		var result = 0
		//var superResult = 1
		var combinations [][]int
		//listClone := make([]int, len(list))
		//copy(listClone, list)

		//var latestElement = listClone[len(listClone)-1:len(listClone)][0]
		var latestElement = list[len(list)-1]
		var latestIndex = len(list)-1

		//fmt.Printf("latestElement: %v, latestIndex: %v", latestElement, latestIndex)

		counter := 0
		for i := latestIndex; i < len(intData); i++ {
			var diff = intData[i] - latestElement
			if diff > 3{
				break
			}else if i >= latestIndex+3{
				break
			}else {
				loopListClone := make([]int, len(list))
				copy(loopListClone, list)

				for j := 0; j < counter; j++{
					loopListClone = append(loopListClone, 0)
				}

				loopListClone = append(loopListClone, intData[i])
				combinations = append(combinations, loopListClone)
				//var a = intData[len(intData)-1:len(intData)][0]
				//var b = loopListClone[len(loopListClone)-1:len(loopListClone)][0]
				//fmt.Printf("Latest THINGY %v, %v\n", a, b)
				if loopListClone[len(loopListClone)-1] == intData[len(intData)-1]{
					result = result + 1
				}
				counter = counter + 1
			}
			//superResult = len(combinations)
			//fmt.Println(len(combinations))
			//fmt.Println(superResult)
		}
		for _, val := range combinations{
			//fmt.Printf("Val: %v\n",val)
			result = result + GetCombinationsFromCurrentState2(val)
			//superResult = superResult + GetCombinationsFromCurrentState2(val)
			//fmt.Println(superResult)
		}
		//fmt.Println(superResult)
		return result
		//if superResult != 0{
		//return superResult
		//}else{
		//return 0
		//}
	}
	//var GetCombinationsFromCurrentState2 = func(list []int)([][]int){
	//var combinations [][]int
	//listClone := make([]int, len(list))
	//copy(listClone, list)
	//
	//var latestElement = listClone[len(listClone)-1:len(listClone)][0]
	//var latestIndex = len(listClone)-1
	//
	//fmt.Printf("latestElement: %v, latestIndex: %v", latestElement, latestIndex)
	//
	//counter := 0
	//for i := latestIndex; i < latestIndex+3; latestIndex++ {
	//var diff = intData[latestIndex] - latestElement
	//if diff > 3{
	//break
	//}else {
	//loopListClone := make([]int, len(list))
	//copy(loopListClone, list)
	//
	//for i = 0; i < counter; i++{
	//loopListClone = append(loopListClone, 0)
	//}
	//
	//loopListClone = append(loopListClone, intData[latestIndex])
	//combinations = append(combinations, loopListClone)
	//counter = counter + 1
	//}
	//}
	//return combinations
	//}
	//var GetCombinationsFromCurrentState = func(list []int)([][]int, int){
	//var combinations [][]int
	//for index, _ := range intData{
	////var difference = 0
	////var possibleChoices []int
	//cpy := make([]int, len(list))
	//copy(cpy, list)
	//var latestElement = cpy[len(cpy)-1:len(cpy)][0]
	//fmt.Printf("Latest element: %v", latestElement)
	//var difference = intData[index+len(cpy)-1] - int(latestElement)
	//if(difference > 3){
	//break
	//}
	//cpy = append(cpy, intData[len(cpy)+index-1])
	//combinations = append(combinations, cpy)
	//fmt.Println("ADDED COMBINATION")
	//fmt.Printf("Compared elements: %v, difference: %v, index: %v\n",latestElement, difference, index)
	//}
	//return combinations, len(combinations)
	//}
	//Cast to int
	for index, _ := range data{
		intVal, _ := strconv.Atoi(data[index])
		intData = append(intData, intVal)
	}

	sort.Ints(intData)
	adapterList = append(adapterList, 0)
	//adapterList = append(adapterList, 1)
	//adapterList = append(adapterList, 4)

	cpy := make([]int, len(adapterList))
	copy(cpy, adapterList)
	//var result = 0
	//var allCombinations [][]int
	//fmt.Printf("Sending in: %v\n",cpy)
	var part2 = GetCombinationsFromCurrentState2(cpy)
	fmt.Printf("Part 2: %v", part2)
	//for _, val := range combinations{
	//allCombinations = append(allCombinations, val)
	//}
//fmt.Printf("Combinations: %v", allCombinations)

	//for _, val := range allCombinations {
	//combinations := GetCombinationsFromCurrentState2(val)
	//for _, val := range combinations {
	//allCombinations = append(allCombinations, val)
	//}
	//}
	//fmt.Printf("ALL COMBINATIONS: \n%v\n", allCombinations)
	return 0
}

func main(){
	var data = ReadFile(input)
	var adapterList = BuildSingleAdapterList(data)
	fmt.Printf("Part 1: %v", adapterList)

	AllCombinationsOfAdapters(data)
}
