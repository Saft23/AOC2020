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
type RealIntSet map[string]int

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

func createSetListOfDataPart2(data []string)[]RealIntSet{
  listOfSets := []RealIntSet{}
  set := RealIntSet{}
  var numberOfPeople = 0
  for _, line := range data {
    if line == "" {
      set["size"] = numberOfPeople
      numberOfPeople = 0
      listOfSets = append(listOfSets, set)
      set = RealIntSet{}
    }else {
      numberOfPeople = numberOfPeople + 1
      for _, char := range line {
        set[string(char)] = set[string(char)] + 1
      }
    }
  }
  set["size"] = numberOfPeople
  listOfSets = append(listOfSets, set)
  return listOfSets
}

func calculateSumOfVotesInSetsPart2(setData []RealIntSet) int {
  var amountOfVotes = 0
  for _, set := range setData{
    for key, setAmount := range set {
      if(key != "size"){
        if set["size"] == setAmount {
          amountOfVotes = amountOfVotes + 1
        }
      }
    }
  }
  return amountOfVotes
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
  var dataSetPart1 = createSetListOfData(data)
  var dataSetPart2 = createSetListOfDataPart2(data)
  var part1Result = calculateSumOfVotesInSets(dataSetPart1)
  var part2Result = calculateSumOfVotesInSetsPart2(dataSetPart2)
  fmt.Printf("Part 1: %v\n", part1Result)
  fmt.Printf("Part 2: %v\n", part2Result)
}
