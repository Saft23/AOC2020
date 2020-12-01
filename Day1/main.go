package main

import (
  "fmt"
  "bufio"
  "os"
  "strconv"
)
var input = "input"

func check(e error){
  if e != nil {
    panic(e)
  }
}

func part1(text []string) (int, string, int, string){
  for firstIndex, firstValue := range text {
    for secondIndex, secondValue := range text {
      a, _ := strconv.Atoi(firstValue)
      b, _ := strconv.Atoi(secondValue)
      if(a + b == 2020){
        fmt.Printf("Part 1: %v\n",a*b)
        return firstIndex, firstValue, secondIndex, secondValue
      }
    }
  }
  return 0,"",0,""
}

func part2(text []string) (int, string, int, string, int, string){
  for firstIndex, firstValue := range text {
    for secondIndex, secondValue := range text {
      for thirdIndex, thirdValue := range text {
        a, _ := strconv.Atoi(firstValue)
        b, _ := strconv.Atoi(secondValue)
        c, _ := strconv.Atoi(thirdValue)
        if(a + b + c == 2020){
          fmt.Printf("Part 2: %v\n",a*b*c)
          return firstIndex, firstValue, secondIndex, secondValue, thirdIndex, thirdValue
        }
      }
    }
  }
  return 0,"",0,"",0,""
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
  part1(data)
  part2(data)
}
