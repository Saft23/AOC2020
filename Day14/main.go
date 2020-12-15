package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"strconv"
	"math"
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

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func LoadMemoryFromDataWithFloatingMemoryAndReturnSum(data []string)int64{
	memory := make(map[int]string)
	mask := "00000000000000000000000000000000000"
	var endResult int64 = 0

	var AddPadding = func(valueBinary string)string{
		var padding = strings.Repeat("0", 36-len(valueBinary))
		var paddedBinary = padding + valueBinary
		return paddedBinary
	}
	var AddPaddingCustom = func(valueBinary string, length int)string{
		var padding = strings.Repeat("0", length-len(valueBinary))
		var paddedBinary = padding + valueBinary
		return paddedBinary
	}
	var CombineMaskAndValue = func(valueBinary string)[]int{
		indexes := []int{}
		newAddresses := []int{}
		for index, _ := range valueBinary {
			if mask[index] == 'X'{
				indexes = append(indexes, index)
			}else if mask[index] == '0'{
				continue
				//valueBinary = replaceAtIndex(valueBinary, '0', index)
			}else if mask[index] == '1'{
				valueBinary = replaceAtIndex(valueBinary, '1', index)
			}else{
				panic("BAD HERE")
			}
		}

		if len(indexes) > 0{
			var combinations = []int{}
			var combinationsBinary = []string{}
			var nbrOfCombinations = int(math.Pow(2,float64(len(indexes))))
			for i:=0; i < nbrOfCombinations;i++{
				combinations = append(combinations, i)
			}
			for _, comb := range combinations{
				combinationBinary := strconv.FormatInt(int64(comb), 2)
				combinationBinaryPadded := AddPaddingCustom(combinationBinary, len(indexes))
				combinationsBinary = append(combinationsBinary, combinationBinaryPadded)
			}
			for i := 0; i < nbrOfCombinations; i++{
				tmpAddress := valueBinary
				for j := 0; j < len(indexes); j++{
					tmpAddress = replaceAtIndex(tmpAddress, rune(combinationsBinary[i][j]), indexes[j])
				}
				valInt, _ := strconv.ParseInt(tmpAddress, 2, 64)
				newAddresses = append(newAddresses, int(valInt))
			}
			return newAddresses
			//Do weird combination here, permutations of X:s

		}else{
			valInt, _ := strconv.ParseInt(valueBinary, 2, 64)
			newAddresses = append(newAddresses, int(valInt))
			return newAddresses
		}
		panic("POop")
		return newAddresses

	}

	for _, val := range data{
		if strings.HasPrefix(val, "mask"){
			maskValue := strings.Split(val, " = ")[1]
			mask = maskValue

		}else if strings.HasPrefix(val, "mem"){
			memValues := strings.Split(val, " = ")
			value, _ := strconv.Atoi(memValues[1])
			valueBinary := strconv.FormatInt(int64(value), 2)
			address := strings.Trim(memValues[0], "mem[]")
			addressInt, _ := strconv.Atoi(address)
			addressBinary := strconv.FormatInt(int64(addressInt), 2)
			addressBinaryWithPadding := AddPadding(addressBinary)

			valueBinaryWithPadding := AddPadding(valueBinary)
			var addresses = CombineMaskAndValue(addressBinaryWithPadding)
			for _, val := range addresses{
				memory[val] = valueBinaryWithPadding
			}

		}else{
			panic("Input data is really bad")
		}
	}

	for _, val := range memory{
		valInt, _ := strconv.ParseInt(val, 2, 64)
		endResult = endResult + valInt
	}

	return endResult
}
func LoadMemoryFromDataAndReturnSum(data []string)int64{
	memory := make(map[int]string)
	mask := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"

	var CombineMaskAndValue = func(valueBinary string)string{
		for index, _ := range valueBinary {
			//binaryPointer := len(valueBinary)
			if mask[index] == 'X'{
				continue
			}else if mask[index] == '0'{
				valueBinary = replaceAtIndex(valueBinary, '0', index)
			}else if mask[index] == '1'{
				valueBinary = replaceAtIndex(valueBinary, '1', index)
			}else{
				panic("BAD HERE")
			}
		}
		return valueBinary
	}

	var AddPadding = func(valueBinary string)string{
		var padding = strings.Repeat("0", 36-len(valueBinary))
		var paddedBinary = padding + valueBinary
		return paddedBinary
	}

	for _, val := range data{
		if strings.HasPrefix(val, "mask"){
			maskValue := strings.Split(val, " = ")[1]
			mask = maskValue

		}else if strings.HasPrefix(val, "mem"){
			memValues := strings.Split(val, " = ")
			value, _ := strconv.Atoi(memValues[1])
			valueBinary := strconv.FormatInt(int64(value), 2)
			address := strings.Trim(memValues[0], "mem[]")
			addressInt, _ := strconv.Atoi(address)

			valueBinaryWithPadding := AddPadding(valueBinary)
			var result = CombineMaskAndValue(valueBinaryWithPadding)
			//fmt.Printf("Added value %v, to address %v, with mask %v and result %v\n", valueBinaryWithPadding, addressInt, mask, result)
			memory[addressInt] = result

		}else{
			panic("Input data is really bad")
		}
	}

	var result int64 = 0
	for _, val := range memory{
		valInt, _ := strconv.ParseInt(val, 2, 64)
		result = result + valInt
	}

	return result
}



func main(){
	var data = ReadFile(input)

	part1 := LoadMemoryFromDataAndReturnSum(data)
	fmt.Printf("Part 1: %v\n", part1)
	part2 := LoadMemoryFromDataWithFloatingMemoryAndReturnSum(data)
	fmt.Printf("Part 2: %v\n", part2)
}
