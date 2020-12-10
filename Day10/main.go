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
	memory := make(map[int][][]int)

	var Perm func(val int)int
	Perm = func(val int)int{
		if val > 1{
			return val * Perm(val-1)
		}
		return 1
	}

	//var Comb func(n int, r int)int
	//Comb = func(n int, r int)int{
	//return Perm(n)/(Perm(r) * Perm(n-r))
	//}

	var GetCombinationsFromCurrentState2 func(list []int)uint64

	//NOW!
	GetCombinationsFromCurrentState2 = func(list []int)uint64{
		//fmt.Println("HELLO")
		var result uint64 = 0
		//var superResult = 1
		var combinations [][]int
		//listClone := make([]int, len(list))
		//copy(listClone, list)

		//var latestElement = listClone[len(listClone)-1:len(listClone)][0]
		var latestElement = list[len(list)-1]
		var latestIndex = len(list)-1
		var skipped = false
		var checkMemory func(latestIndex int)([][]int, bool)
		checkMemory = func(latestIndex int)([][]int, bool){
			res,ok := memory[latestIndex]
			if ok{
				for _, val := range res{
					checkMemory(val[len(val)-1])
				}
			}else{
				for _, val := range res{
					result = result + GetCombinationsFromCurrentState2(val)
				}
			}
			return res,ok
		}
		if _, ok := checkMemory(latestIndex); ok {
			skipped = true
			//fmt.Println("SKIPPED")
		}else{

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
				//var a = intData[len(intData)-1:len(intData)][0]
				//var b = loopListClone[len(loopListClone)-1:len(loopListClone)][0]
				//fmt.Printf("Latest THINGY %v, %v\n", a, b)
				if loopListClone[len(loopListClone)-1] == intData[len(intData)-1]{
					fmt.Println(loopListClone)
					result = result + 1
					//loopListClone = loopListClone[:0]
					combinations = append(combinations, loopListClone)
				}else{
					fmt.Printf("Bad: %v\n",loopListClone)
					combinations = append(combinations, loopListClone)
				}
				counter = counter + 1
			}
			//superResult = len(combinations)
			//fmt.Println(len(combinations))
			//fmt.Println(superResult)
		}
		}
		if !skipped{
		  memory[latestIndex] = combinations
			for _, val := range combinations{
				//fmt.Println(val)
					result = result + GetCombinationsFromCurrentState2(val)
			}
		}else{
			combinations = memory[latestIndex] 

			for _, val := range combinations{
				//fmt.Println(val)
				//fmt.Println(val)
				if val[len(val)-1] == intData[len(intData)-1]{
					result = result + 1
					if result % 10 == 0{
						fmt.Println(result)
					}
				}
				result = result + GetCombinationsFromCurrentState2(val)

			}
		}
		//for _, val := range combinations{
		////fmt.Printf("Val: %v\n",val)
		//result = result + GetCombinationsFromCurrentState2(val)
		////superResult = superResult + GetCombinationsFromCurrentState2(val)
		////fmt.Println(superResult)
		//}
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
	//
	//intData = append(intData, 0)
	sort.Ints(intData)
	//intData = append(intData, intData[len(intData)-1]+3)
	//sort.Ints(intData)
	//
	//fmt.Println(intData)
	//
	//var superResult = 1
	//var counter = 0
	//for i := 0; i < len(intData)-2; i++{
	//fmt.Println(superResult)
	//var val = intData[i]
	//var currentNum = val
	//var nextNum = intData[i+2]
	//var diff = nextNum - currentNum
	//
	//if diff <= 3{
	//if diff == 2 {
	//counter = counter + 1
	//}else{
	//fmt.Println("FUCK")
	//superResult = superResult * 2
	//}
	//}else if counter > 0{
	//var sum = 1
	//for i := 0; i <= counter; i++{
	//if i % 2 == 1{
	//sum = sum - (counter-2)
	//}
	//fmt.Print(counter)
	//sum = sum + Comb(counter, i)
	//fmt.Printf("Sum: %v\n",sum)
	//}
	//counter = 0
	//superResult = superResult * sum
	//}
	//}
	//fmt.Printf("Megaultra result: %v",superResult)
	adapterList = append(adapterList, 0)
	//adapterList = append(adapterList, 1)
	//adapterList = append(adapterList, 4)

	cpy := make([]int, len(adapterList))
	copy(cpy, adapterList)
	//var result = 0
	//var allCombinations [][]int
	//fmt.Printf("Sending in: %v\n",cpy)
	var part2 = GetCombinationsFromCurrentState2(cpy)
	fmt.Printf("Part 2: %v\n",part2)
	//fmt.Printf("Part 2: %v", part2)
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
	fmt.Printf("Part 1: %v\n", adapterList)

	AllCombinationsOfAdapters(data)
}
