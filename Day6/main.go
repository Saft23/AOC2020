package main

import (
  "fmt"
  //"strings"
  "bufio"
  "os"
  //"strconv"
  //"sort"
)

var input = "input"
type IntSet map[string]bool

func check(e error)bool{
  if e != nil {
    panic(e)
  }
  return true
}

func createSetListOfData(data []string)[]IntSet{
  listOfSets := []IntSet{}
  set := IntSet{}

  for _, line := range data {
    if line == "" {
      listOfSets = append(listOfSets, set)
      set = IntSet{}
    }else {
      for _, char := range line {
        //:)
        set[string(char)] = true
      }
    }
  }
  listOfSets = append(listOfSets, set)
  return listOfSets
}

func calculateSumOfVotesInSets(setData []IntSet) int {
  var amountOfVotes = 0
  for _, set := range setData {
    amountOfVotes = amountOfVotes + len(set)
  }
  return amountOfVotes
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
  var dataSet = createSetListOfData(data)
  var part1Result = calculateSumOfVotesInSets(dataSet)
  fmt.Printf("Part 1: %v\n", part1Result)
}
