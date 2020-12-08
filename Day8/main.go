package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"strconv"
)

var input = "input2"

type Program struct {
	Pointer int
	Accumilator int
	Instructions []string
	PrevInstructions []int
	Terminate bool
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
	program.Terminate = false

	for _, line := range data {
		program.Instructions = append(program.Instructions, line)
	}
	return &program
}

func RunProgramUntilRepeatedInstruction(program *Program){
	for{
		ParseAndExecuteInstruction(program)
		if(program.Terminate){
			return
		}
	}
}

func ParseAndExecuteInstruction(program *Program){
	var StepProgram func(steps int)
	var CheckIfTerminate func()bool
	CheckIfTerminate = func()bool{
		for _, val := range program.PrevInstructions{
			if val == program.Pointer{
				return true
			}
		}
		return false
	}
	StepProgram = func(steps int){
		program.PrevInstructions = append(program.PrevInstructions, program.Pointer)
		program.Pointer = program.Pointer + steps
		if CheckIfTerminate(){
			return
		}
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


func main(){
	var data = ReadFile(input)
	var program = LoadProgram(data)
	RunProgramUntilRepeatedInstruction(program)
	fmt.Printf("Part 1: %v", program.Accumilator)
	fmt.Println(program)
}
