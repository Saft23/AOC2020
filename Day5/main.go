package main

import (
  "fmt"
  "strings"
  "bufio"
  "os"
  "strconv"
  "sort"
)

var input = "input"

func check(e error)bool{
  if e != nil {
    panic(e)
  }
  return true
}

func calculateHighestIdPart1(data []string)(int){
  var highestId = 0
  for _, line := range data{
    line = strings.ReplaceAll(line, "F","0")
    line = strings.ReplaceAll(line, "B", "1")
    line = strings.ReplaceAll(line, "R", "1")
    line = strings.ReplaceAll(line, "L", "0")
    row := line[0:7]
    col := line[7:10]
    if rowValue, err := strconv.ParseInt(row, 2, 10); check(err) {
      if colValue, err := strconv.ParseInt(col, 2, 10); check(err) {
        if(int(rowValue) * 8 + int(colValue) > highestId){
          highestId = int(rowValue) * 8 + int(colValue)
        }
      }
    }
  }
  fmt.Printf("Part 1: %v\n", highestId)
  return highestId
}

func findMissingSeatPart2(data []string)(int){
  var s []int
  for _, line := range data{
    line = strings.ReplaceAll(line, "F","0")
    line = strings.ReplaceAll(line, "B", "1")
    line = strings.ReplaceAll(line, "R", "1")
    line = strings.ReplaceAll(line, "L", "0")
    row := line[0:7]
    col := line[7:10]
    if rowValue, err := strconv.ParseInt(row, 2, 10); check(err) {
      if colValue, err := strconv.ParseInt(col, 2, 10); check(err) {
        var value = int(rowValue) * 8 + int(colValue)
        s = append(s, int(value))
      }
    }
  }
  sort.Ints(s)
  var result = 0
  for i := 0; i < len(s)-1; i++{
    if(s[i+1] - s[i] != 1){
      result = s[i]+1
    }
  }
  fmt.Printf("Part 2: %v\n", result)
  return result
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
  calculateHighestIdPart1(data)
  findMissingSeatPart2(data)
}
