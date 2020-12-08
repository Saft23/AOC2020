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
	PrevInstructions []string
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
	var instruction = strings.Split(program.Instructions[program.Pointer], " ")
	var op = instruction[0]
	var argument, _ = strconv.Atoi(instruction[1])
	switch op{
		case "acc":
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
	fmt.Println(program)
}
