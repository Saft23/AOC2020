package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"strconv"
)

var input = "input"

type Program struct {
	Pointer int
	Accumilator int
	Instructions []string
	PrevInstructions []int
	RepeatedInstruction bool
	Counter int
	Infinite bool
	OutOfBounds bool
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

func check(e error)bool{
	if e != nil {
		panic(e)
	}
	return true
}

func LoadProgram(data []string) *Program {
	program := Program{}
	program.Pointer = 0
	program.Accumilator = 0
	program.Counter = 0
	program.RepeatedInstruction = false
	program.Infinite = false
	program.OutOfBounds = false

	for _, line := range data {
		program.Instructions = append(program.Instructions, line)
	}
	return &program
}

func RunProgramUntilRepeatedInstruction(program *Program)int{
	for{
		ParseAndExecuteInstruction(program)
		if(program.RepeatedInstruction){
			return program.Accumilator
		}
	}
}

func RunUntilProgramOutOfBounds(programs *[]Program)int{
	for _, program := range *programs{
		for{
			ParseAndExecuteInstruction(&program)
			if program.OutOfBounds{
				return program.Accumilator
			}else if program.Infinite || program.RepeatedInstruction{
				break
			}
		}
	}
	return 0
}

func ParseAndExecuteInstruction(program *Program){
	var StepProgram func(steps int)
	var CheckIfRepeatedInstruction func()
	var CheckIfInfinite func()
	var CheckIfOutOfBounds func()

	CheckIfOutOfBounds = func(){
		if program.Pointer > len(program.Instructions)-1{
			program.OutOfBounds = true
		}
	}

	CheckIfInfinite = func(){
		if(program.Counter > 100000){
			program.Infinite = true
		}
	}

	CheckIfRepeatedInstruction = func(){
		for _, val := range program.PrevInstructions{
			if val == program.Pointer{
				program.RepeatedInstruction = true
			}
		}
	}

	StepProgram = func(steps int){
		program.Counter = program.Counter + 1
		program.PrevInstructions = append(program.PrevInstructions, program.Pointer)
		program.Pointer = program.Pointer + steps

		CheckIfInfinite()
		CheckIfRepeatedInstruction()
		CheckIfOutOfBounds()
	}

	var instruction = strings.Split(program.Instructions[program.Pointer], " ")
	var op = instruction[0]
	var argument, _ = strconv.Atoi(instruction[1])
	switch op{
		case "acc":
		program.Accumilator = program.Accumilator + argument
		StepProgram(1)
		break

		case "jmp":
		StepProgram(argument)
		break

		case "nop":
		StepProgram(1)
		break
	}
}

func BuildProgramPermutations(program *Program)*[]Program{
	//Deep copy
	var CopyProgram = func(program Program)Program{
		newProgram := program
		newProgram.Instructions = make([]string, len(program.Instructions))
		newProgram.PrevInstructions = make([]int, len(program.PrevInstructions))
		copy(newProgram.Instructions, program.Instructions)
		copy(newProgram.PrevInstructions, program.PrevInstructions)
		return newProgram
	}

	programList := []Program{}
	var numberOfNops = 0
	var numberOfJmps = 0
	var nopIndex []int
	var jmpIndex []int

	//Build 2 lists of jumps and nops
	for index, val := range program.Instructions{
		if strings.Contains(val, "nop"){
			numberOfNops = numberOfNops + 1
			nopIndex = append(nopIndex, index)
		}else if strings.Contains(val, "jmp"){
			numberOfJmps = numberOfJmps + 1
			jmpIndex = append(jmpIndex, index)
		}
	}

	for _, index := range nopIndex {
		newProgram := CopyProgram(*program)
		newProgram.Instructions[index] = strings.Replace(newProgram.Instructions[index], "nop", "jmp", 1)
		programList = append(programList, newProgram)
	}
	for _, index := range jmpIndex {
		newProgram := CopyProgram(*program)
		newProgram.Instructions[index] = strings.Replace(newProgram.Instructions[index], "jmp", "nop", 1)
		programList = append(programList, newProgram)
	}
	return &programList
}

func main(){
	var data = ReadFile(input)
	var program = LoadProgram(data)
	var part1 = RunProgramUntilRepeatedInstruction(program)
	fmt.Printf("Part 1: %v\n", part1)

	program = LoadProgram(data)
	var allThePrograms = BuildProgramPermutations(program)
	var part2 = RunUntilProgramOutOfBounds(allThePrograms)
	fmt.Printf("Part 2: %v", part2)
}
