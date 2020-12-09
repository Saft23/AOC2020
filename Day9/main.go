package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
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

func part1(data []string)int{
	var maxSize = 25
	var AddToQueue = func(queue []int, element int)[]int{
		if len(queue) == maxSize {
			queue = queue[1:]
		}
		queue = append(queue, element)
		return queue
	}

	var CheckIfValid = func(queue []int, element int)bool{

		for _, val := range queue{
			for _, val2 := range queue{
				if val + val2 == element{
					return true
				}
			}
		}
		return false
	}
	var AddPremable = func(queue []int, data []string)[]int{
		for i := 0; i < maxSize; i++ {
			element, _ := strconv.Atoi(data[i])
			queue = append(queue, element)
		}
		return queue
	}
	var queue []int
	queue = AddPremable(queue, data)
	for i:= maxSize; i < len(data); i++{
		nextElement, _ := strconv.Atoi(data[i])
		if CheckIfValid(queue, nextElement) {
			queue = AddToQueue(queue, nextElement)
		}else{
			return nextElement
		}
	}
	return 0
}

func GetContigousList(data []string, goal int)[]int{
	var SumIsGoal = func(list []int)bool{
		var sum =  0
		for _, val := range list{
			sum = sum + val
		}
		return sum == goal
	}

	var upperBound = 0
	for index, val := range data {
		intData, _ := strconv.Atoi(val)
		if intData == goal{
			upperBound = index
		}
	}

	for i := 0; i < upperBound; i++{
		var list[]int
		for j := 0; j+i < upperBound; j++{
			nextElement, _ := strconv.Atoi(data[j+i])
			list = append(list, nextElement)
			//fmt.Println(list)
			if SumIsGoal(list){
				return list
			}
		}
	}
	var list []int
	return list
}

func SumOfLargestAndSmallest(data []int)int{
	var smallest = data[0]
	var largest = data[0]
	for _, val := range data{
		if val < smallest{
			smallest = val
		}else if val > largest {
			largest = val
		}
	}
	return smallest + largest
}

func main(){
	var data = ReadFile(input)
	var part1 = part1(data)
	fmt.Printf("Part 1: %v\n", part1)

	var contiguousList = GetContigousList(data, part1)
	var part2 = SumOfLargestAndSmallest(contiguousList)
	fmt.Printf("Part 2: %v\n", part2)
}
