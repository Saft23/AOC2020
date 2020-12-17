package main

import (
	"fmt"
	"bufio"
	"os"
)

var input = "input"

type Cord struct {
	x int
	y int
	z int
}
type Cube map[Cord]bool

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


func LoadDataIntoCube(data []string, cube *Cube){
	for indexRow, row := range data{
		for indexCol, col := range row{
			if  col == '#'{
				cord := Cord{x:indexRow, y:indexCol, z:0}
				(*cube)[cord] = true
			}else if col == '.'{
				cord := Cord{x: indexRow, y:indexCol, z:0}
				(*cube)[cord] = false
			}
		}
	}
}

func (cube Cube) GetAllActiveCubes() int{
	var sum = 0
	for _, val := range cube{
		if val{
			sum = sum + 1
		}
	}
	return sum
}

func (cube Cube) Init(size int){
	for x:=-(size/2); x<size/2; x++{
		for y:=-(size/2); y<size/2; y++{
			for z:=-(size/2); z<size/2; z++{
				cord := Cord{x:x, y:y, z:z}
				cube[cord] = false
			}
		}
	}
}

func (cube Cube) GetAmountOfActiveNeighbors(cord Cord)int{
	var result = 0
	for x := cord.x-1; x <= cord.x+1; x++{
		for y := cord.y-1; y <= cord.y+1; y++{
			for z := cord.z-1; z <= cord.z+1; z++{
				newCord := Cord{x:x, y:y, z:z}
				if newCord != cord{
					if cube[newCord]{
						result = result + 1
					}
				}
			}
		}
	}
	return result
}

func (cube Cube) Cycle(){
	tmpCube := make(Cube)
	for k, v := range cube{
		tmpCube[k] = v
	}

	for cord, _ := range cube{
		if cube[cord]{
			if cube.GetAmountOfActiveNeighbors(cord) == 2 || cube.GetAmountOfActiveNeighbors(cord) == 3{
				tmpCube[cord] = true
			}else{
				tmpCube[cord] = false
			}
		}else{
			if cube.GetAmountOfActiveNeighbors(cord) == 3{
				tmpCube[cord] = true
			}
		}
	}

	for k, v := range tmpCube{
		cube[k] = v
	}
}

func main(){
	data := ReadFile(input)
	cube := make(Cube)
	cube.Init(30)
	LoadDataIntoCube(data, &cube)
	for i := 0; i < 6; i++{
		cube.Cycle()
	}
	part1 := cube.GetAllActiveCubes()
	fmt.Printf("Part 1: %v", part1)
}
