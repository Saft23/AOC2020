package main

import (
  "fmt"
  "strings"
  "bufio"
  "os"
  "strconv"
  "regexp"
)
type Passport struct {
  Byr int     //Birth year
  Iyr int     //Issue year
  Eyr int     //Expiration year
  Hgt string     //Height
  Hcl string  //Hair color
  Ecl string  //Eye color
  Pid string     //Passport ID
  Cid string     //Optional Country ID
}
var input = "input"

func check(e error){
  if e != nil {
    panic(e)
  }
}

func Split(r rune) bool{
  return r == ':' || r == ' '
}

func sortPassports(data []string)(passports []Passport){
  var listOfPassports []Passport
  var passport Passport
  var dataValue = ""
  for _, line := range data{
    if(line == ""){
      listOfPassports = append(listOfPassports, passport)
      passport = Passport{}
    } else{
      var segments = strings.FieldsFunc(line, Split)
      var err error
      for i, block := range segments{
        if (i % 2) == 0 {
          dataValue = block
        }else{
          switch dataValue {
          case "byr":
            passport.Byr, err = strconv.Atoi(block)
            check(err)
          case "iyr":
            passport.Iyr, err = strconv.Atoi(block)
            check(err)
          case "eyr":
            passport.Eyr, err = strconv.Atoi(block)
            check(err)
          case "hgt":
            passport.Hgt = block
          case "hcl":
            passport.Hcl = block
          case "ecl":
            passport.Ecl = block
          case "pid":
            passport.Pid = block
          case "cid":
            passport.Cid = block
          }
        }
        dataValue = block
      }
    }
  }
  listOfPassports = append(listOfPassports, passport)
  return listOfPassports
}

func calculateNumberOfValidPassportsPart1(passports []Passport)(int){
  var validPassports = 0
  for _, passport := range passports{
    if(passport.Byr != 0 &&
       passport.Iyr != 0 &&
       passport.Eyr != 0 &&
       passport.Hgt != "" &&
       passport.Hcl != "" &&
       passport.Ecl != "" &&
       passport.Pid != ""){
      validPassports = validPassports + 1
    }
  }
  return validPassports
}

func calculateNumberOfValidPassportsPart2(passports []Passport)(int){
  var validPassports = 0
  var hgtValue = 0
  for _, passport := range passports{
    hgtCmValid, err := regexp.Match(`^\d+cm$`, []byte(passport.Hgt))
    check(err)
    hgtInValid, err := regexp.Match(`^\d+in$`, []byte(passport.Hgt))
    check(err)
    hclValid, err := regexp.Match(`^#[0-9, a-f]{6}$`, []byte(passport.Hcl))
    check(err)
    eclValid, err := regexp.Match(`(^amb$|^blu$|^brn$|^gry$|^grn$|^hzl$|^oth$)`, []byte(passport.Ecl))
    check(err)
    pidValid, err := regexp.Match(`^0*[0-9]{9}$`, []byte(passport.Pid))
    check(err)
    if hgtCmValid || hgtInValid{
      hgtValue, err = strconv.Atoi(passport.Hgt[0:len(passport.Hgt)-2])
      check(err)
    }
    if(passport.Byr >= 1920 && passport.Byr <= 2002 &&
       passport.Iyr >= 2010 && passport.Iyr <= 2020 &&
       passport.Eyr >= 2020 && passport.Eyr <= 2030 &&
       ( (hgtCmValid && hgtValue >= 150 && hgtValue <= 193) ||
       (hgtInValid && hgtValue >= 59 && hgtValue <= 76) )&&
       hclValid &&
       eclValid &&
       pidValid){
        validPassports = validPassports + 1
       }
  }
  return validPassports
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
  var listOfPassports =  sortPassports(data)
  var validPassportsPart1 = calculateNumberOfValidPassportsPart1(listOfPassports)
  var validPassportsPart2 = calculateNumberOfValidPassportsPart2(listOfPassports)
  fmt.Printf("Part 1: %v\nPart 2: %v\n", validPassportsPart1, validPassportsPart2)
}
