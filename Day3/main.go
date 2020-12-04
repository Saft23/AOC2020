package main

import (
  "fmt"
  "bufio"
  "os"
)
var input = "input"

func check(e error){
  if e != nil {
    panic(e)
  }
}

func calculateLinearCollitions(data []string, stepsRight int, stepsDown int) int{
  var posX = 0
  var posY = 0
  var collitions = 0
  for len(data)-1 > posY {
    posX = posX + stepsRight
    posY = posY + stepsDown
    if(posX > len(data[posY])-1){
      posX = posX % len(data[posY])
    }
    if((data[posY][posX]) == '#') {
      collitions = collitions + 1
    }
  }
  return collitions
}

func readFile(input string) (text []string){
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

func main(){
  var data = readFile(input)
  var resultPart1 = calculateLinearCollitions(data, 3, 1)

  var path1 = calculateLinearCollitions(data, 1, 1)
  var path2 = resultPart1
  var path3 = calculateLinearCollitions(data, 5, 1)
  var path4 = calculateLinearCollitions(data, 7, 1)
  var path5 = calculateLinearCollitions(data, 1, 2)
  var resultPart2 = path1 * path2 * path3 * path4 * path5

  fmt.Printf("Part 1: %v, Part 2: %v\n", resultPart1, resultPart2)
}
