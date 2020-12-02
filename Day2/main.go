package main

import (
  "fmt"
  "bufio"
  "os"
  "strconv"
  "strings"
)
var input = "input"

func check(e error){
  if e != nil {
    panic(e)
  }
}

func checkValidPassword(lowerLimit int, upperLimit int, character string, password string) bool{
  count := strings.Count(password, character)
  if(count >= lowerLimit && count <= upperLimit){
    return true
  } else {
    return false
  }
}

func checkValidData(text []string) int{
  var validPasswords = 0
  for _, line := range text {
    data := strings.Split(line, " ")
    limits := strings.Split(data[0], "-")     //limits[0] = lower limit, limits[1] = upper limit
    lowerLimit, _ := strconv.Atoi(limits[0])
    upperLimit, _ := strconv.Atoi(limits[1])
    char := strings.Split(data[1], ":")[0] 
    password := data[2]                       //Password
    if (checkValidPassword(lowerLimit, upperLimit, char, password)){
      validPasswords = validPasswords + 1
    }
  }
  fmt.Printf("part 1: %v\n", validPasswords)
  return validPasswords
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
  checkValidData(data)
}
