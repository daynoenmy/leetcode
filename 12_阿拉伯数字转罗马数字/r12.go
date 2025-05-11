package main

import "fmt"

var thousands []string = []string{"", "M", "MM", "MMM"}
var hundred []string = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
var ten []string = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
var ge []string = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}

func intToRoman(num int) string {
	var str string = ""
	str += thousands[num/1000]
	num = num % 1000
	str += hundred[num/100]
	num = num % 100
	str += ten[num/10]
	num = num % 10
	str += ge[num]
	return str
}
func main() {
	fmt.Println(intToRoman(1994))
}
