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

func checkValidPasswordPart1(lowerLimit int, upperLimit int, character string, password string) bool {
  count := strings.Count(password, character)
  if(count >= lowerLimit && count <= upperLimit){
    return true
  } else {
    return false
  }
}

func checkValidPasswordPart2(lowerIndex int, upperIndex int , character string, password string) bool {
  a := string(password[lowerIndex-1])
  b := string(password[upperIndex-1])
  if((a == character) !=  (b == character)){
    return true
  }
  return false
}

func parseDataAndCheckPassword(text []string) (int, int){
  var validPasswordsPart1 = 0
  var validPasswordsPart2 = 0
  for _, line := range text {
    data := strings.Split(line, " ")
    limits := strings.Split(data[0], "-")     //limits[0] = lower limit, limits[1] = upper limit
    lowerLimit, _ := strconv.Atoi(limits[0])  //Lower limit
    upperLimit, _ := strconv.Atoi(limits[1])  //Upper limit
    char := strings.Split(data[1], ":")[0]    //Character
    password := data[2]                       //Password

    if (checkValidPasswordPart1(lowerLimit, upperLimit, char, password)){
      validPasswordsPart1 = validPasswordsPart1 + 1
    }
    if (checkValidPasswordPart2(lowerLimit, upperLimit, char, password)){
      validPasswordsPart2 = validPasswordsPart2 + 1
    }
  }
  fmt.Printf("Part 1: %v\nPart 2: %v\n", validPasswordsPart1, validPasswordsPart2)
  return validPasswordsPart1, validPasswordsPart2
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
  parseDataAndCheckPassword(data)
}
