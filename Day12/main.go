package main

import (
	"fmt"
	"bufio"
	"os"
	"math"
	"strconv"
)

var input = "input"

type Ship struct{
	North int
	East int
	Heading int
	OrigNorth int
	OrigEast int

	Relative bool
	NorthWaypoint int
	EastWaypoint int
}

func (s Ship) Move(dir string, value int)Ship{
	if s.Relative{
		switch dir{
			case "N":
			s.NorthWaypoint = s.NorthWaypoint + value
			break

			case "S":
			s.NorthWaypoint = s.NorthWaypoint - value
			break

			case "E":
			s.EastWaypoint = s.EastWaypoint + value
			break

			case "W":
			s.EastWaypoint = s.EastWaypoint - value
			break

			case "F":
			s.North = s.North + value * s.NorthWaypoint
			s.East = s.East + value * s.EastWaypoint
			break
		}

	}else{
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
	}
	return s
}

func (s Ship) Turn(dir string, value int)Ship{
	if s.Relative {
		var radians float64 = float64(value) * 3.14/180

		switch dir{
			case "R":
			radians = radians
			break

			case "L":
			radians = -radians
			break
		}
		newEast := float64(s.EastWaypoint) * math.Cos(radians) + float64(s.NorthWaypoint) * math.Sin(radians)
		newNorth := -float64(s.EastWaypoint) * math.Sin(radians) + float64(s.NorthWaypoint) * math.Cos(radians)
		s.EastWaypoint = int(math.Round(newEast))
		s.NorthWaypoint = int(math.Round(newNorth))
	}else{
		switch dir{
			case "L":
			s.Heading = (s.Heading - value) % 360
			break

			case "R":
			s.Heading = (s.Heading + value) % 360
			break
		}
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
	shipPart1 := Ship{}
	shipPart1.Heading = 90
	shipPart1 = shipPart1.RunInstructions(data)
	part1 := shipPart1.ManhattanDistainceFromOrig()


	shipPart2 := Ship{}
	shipPart2.EastWaypoint = 10
	shipPart2.NorthWaypoint = 1
	shipPart2.Relative = true
	shipPart2 = shipPart2.RunInstructions(data)
	part2 := shipPart2.ManhattanDistainceFromOrig()

	fmt.Printf("Part 1: %v\nPart 2: %v\n", part1, part2)
}
