package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"math/big"
	"strconv"
)

var input = "input"


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

func GetClosestTimeOfDeparture(data []string)(int){
	time, _ := strconv.Atoi(data[0])
	busses := strings.Split(data[1],",")
	var bussesId []int

	for _, val := range busses{
		if val != "x" {
			tmp, _ := strconv.Atoi(val)
			bussesId = append(bussesId, tmp)
		}
	}
	var increasedTime = 0
	for {
		for _, buss := range bussesId{
			if (increasedTime + buss) % time == 0 {
				return (buss - (increasedTime % buss)) * buss //Fix this return
			}
		}
		increasedTime = increasedTime + 1
	}
}

//Chinese rest theorem
func GetClosestTimeOfConsecutiveDepartures(data []string) (*big.Int, error) {
	var one = big.NewInt(1)

	var crt = func(a, n []*big.Int) (*big.Int, error) {
		p := new(big.Int).Set(n[0])
		for _, n1 := range n[1:] {
			p.Mul(p, n1)
		}
		var x, q, s, z big.Int
		for i, n1 := range n {
			q.Div(p, n1)
			z.GCD(nil, &s, n1, &q)
			if z.Cmp(one) != 0 {
				return nil, fmt.Errorf("%d not coprime", n1)
			}
			x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
		}
		return x.Mod(&x, p), nil
	}
	mods := []*big.Int{}
	remainders := []*big.Int{}

	index := 0
	busses := strings.Split(data[1],",")
	for _, val := range busses{
		if val == "x"{
			index = index + 1
		}else{
			tmp, _ := strconv.Atoi(val)
			tmp64 := int64(tmp)
			mods = append(mods, big.NewInt(tmp64))
			remainders = append(remainders, big.NewInt(tmp64-int64(index)))
			index = index + 1
		}
	}
	return crt(remainders, mods)
}

func main(){
	var data = ReadFile(input)
	var part1 = GetClosestTimeOfDeparture(data)
	var part2, _ = GetClosestTimeOfConsecutiveDepartures(data)
	fmt.Printf("Part 1: %v\nPart 2: %v", part1, part2)
}
