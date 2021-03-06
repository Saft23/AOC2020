package main

import (
	"fmt"
	"bufio"
	"os"
	//"strconv"
	//"sort"
)

var input = "input"
//var input = "input"

const (
	EMPTY = iota
	OCCUPIED 
	FLOOR
)

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


func ConvertToSeatMap(data []string)[][]int{
	seatMap := make([][]int, len(data))
	for i := range seatMap {
		seatMap[i] = make([]int, len(data[i]))
	}
	for y, row := range data{
		for x, _ := range row{
			if data[y][x] == 'L'{
				seatMap[y][x] = EMPTY
			}else if data[y][x] == '#'{
				seatMap[y][x] = OCCUPIED
			}else{
				seatMap[y][x] = FLOOR
			}
		}
	}
	return seatMap
}

func StepOneSeatMap(seatMap [][]int, part int)([][]int, int){
	var changes = 0
	var AmountOfAdjacentOccupiedSeats func(y int,x int)int
	var AmountOfAdjacentOccupiedSeatsPart2 func(y int,x int)int
	AmountOfAdjacentOccupiedSeats = func(chairY int, chairX int)int{
		var result = 0
		
		for yCord := chairY-1; yCord <= chairY+1; yCord++{
			for xCord := chairX-1; xCord <= chairX+1; xCord++{
				if xCord == chairX && yCord == chairY{
					continue
				}else if yCord < 0 || yCord > len(seatMap)-1 || xCord < 0 || xCord > len(seatMap[0])-1{
					continue
				}else{
					if seatMap[yCord][xCord] == OCCUPIED{
						result = result + 1
					}
				}
			}
		}
		return result
	}
	AmountOfAdjacentOccupiedSeatsPart2 = func(chairY int, chairX int)int{
		var result = 0
		var maxY = len(seatMap)
		var maxX = len(seatMap[0])

		var x = 0
		var y = 0
		for i := 0; i < 8; i++{
			switch i {
			case 0:
				x = 1
				y = 1
			case 1:
				x = 1
				y = 0
				break
			case 2:
				x = 1
				y = -1
				break
			case 3:
				x = 0
				y = -1
				break
			case 4:
				x = -1
				y = -1
				break
			case 5:
				x = -1
				y = 0
				break
			case 6:
				x = -1
				y = 1
				break
			case 7:
				x = 0
				y = 1
				break
			}
			var currentX = chairX
			var currentY = chairY
			for j := 1; currentY+y*j >= 0 && currentY+y*j < maxY && currentX+x*j >= 0 && currentX+x*j < maxX; j++{
				if seatMap[currentY+y*j][currentX+x*j] == OCCUPIED{
					result = result +1
					break
				}else if seatMap[currentY+y*j][currentX+x*j] == EMPTY{
					break
				}
			}
		}
		return result
	}

	//Copy map
	seatMapCopy := make([][]int, len(seatMap))
	for i := range seatMap{
		seatMapCopy[i] = make([]int, len(seatMap[i]))
		copy(seatMapCopy[i], seatMap[i])
	}

	for rowIndex, row := range seatMap{
		for colIndex, _ := range row{
			var currentSeat = seatMap[rowIndex][colIndex]
			switch currentSeat {
			case FLOOR:
				break

			case OCCUPIED:
				if part == 1{
					if AmountOfAdjacentOccupiedSeats(rowIndex, colIndex) >= 4{ //4
						seatMapCopy[rowIndex][colIndex] = EMPTY
						changes = changes + 1
					}
				}else{
					if AmountOfAdjacentOccupiedSeatsPart2(rowIndex, colIndex) >= 5{ //4
						seatMapCopy[rowIndex][colIndex] = EMPTY
						changes = changes + 1
					}

				}
				break
			case EMPTY:
				if part == 1{
					if AmountOfAdjacentOccupiedSeats(rowIndex, colIndex) == 0{
						seatMapCopy[rowIndex][colIndex] = OCCUPIED
						changes = changes + 1
					}

				}else{
					if AmountOfAdjacentOccupiedSeatsPart2(rowIndex, colIndex) == 0{
						seatMapCopy[rowIndex][colIndex] = OCCUPIED
						changes = changes + 1
					}
				}
				break
			}
		}
	}

	return seatMapCopy, changes
}

func AmountOfOccupiedSeats(seatMap [][]int)int{
	var occupiedSeats = 0
	for rowIndex, row := range seatMap{
		for colIndex, _ := range row{
			if seatMap[rowIndex][colIndex] == 1{
				occupiedSeats = occupiedSeats + 1
			}

		}
	}
	return occupiedSeats
}

func main(){
	var data = ReadFile(input)
	seatMap := ConvertToSeatMap(data)
	changes := 1
	for changes > 0{
		seatMap, changes = StepOneSeatMap(seatMap, 1)
	}
	var part1 = AmountOfOccupiedSeats(seatMap)

	changes = 1
	seatMap = ConvertToSeatMap(data)
	for changes > 0{
		seatMap, changes = StepOneSeatMap(seatMap, 2)
	}
	var part2 = AmountOfOccupiedSeats(seatMap)
	fmt.Printf("Part 1: %v\nPart 2: %v\n", part1, part2)
}
