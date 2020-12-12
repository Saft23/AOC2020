package main

import (
	"fmt"
	"bufio"
	"os"
	"math"
	"strconv"
	//"sort"
)

var input = "input"

type Ship struct{
	North int
	East int
	Heading int
	OrigNorth int
	OrigEast int
}

func (s Ship) Move(dir string, value int)Ship{
	switch dir{
		case "N":
		s.North = s.North + value
		break

		case "S":
		s.North = s.North - value
		break

		case "E":
		s.East = s.East + value
		break

		case "W":
		s.East = s.East - value
		break

		case "F":
		var radians float64 = float64(s.Heading) * 3.14/180
		northIncrement := math.Cos(radians) * float64(value)
		eastIncrement := math.Sin(radians) * float64(value)
		s.North = s.North + int(math.Round(northIncrement))
		s.East = s.East + int(math.Round(eastIncrement))
		break
	}
	return s
}

func (s Ship) Turn(dir string, value int)Ship{
	switch dir{
		case "L":
		s.Heading = (s.Heading - value) % 360
		break

		case "R":
		s.Heading = (s.Heading + value) % 360
		break
	}
	return s
}

func (s Ship) RunInstructions(data []string)Ship{
	for _, val := range data{
		instruction := string(val[0])
		value, _ := strconv.Atoi(val[1:])
		switch instruction{
			case "L", "R":
			s = s.Turn(instruction, value)
			break
			case "N", "E", "S", "W", "F":
			s = s.Move(instruction, value)
		}
		fmt.Println(s)
	}
	return s
}

func (s Ship) ManhattanDistainceFromOrig()float64{
	return math.Abs(float64(s.North)) + math.Abs(float64(s.East))
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




func main(){
	var data = ReadFile(input)
	ship := Ship{}
	ship.Heading = 90
	ship = ship.RunInstructions(data)
	fmt.Println(ship.ManhattanDistainceFromOrig())
}
