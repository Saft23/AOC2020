package main

import (
	"fmt"
	"strings"
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

func replaceAtIndex(in string, r rune, i int) string {
    out := []rune(in)
    out[i] = r
    return string(out)
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
}
