package main

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
)

// shorthand for fmt.Println
var pl = fmt.Println

func main() {
	// IO
	// pl("What is your name?")
	// reader := bufio.NewReader(os.Stdin)
	// name,err := reader.ReadString('\n')
	// if err != nil {
	// 	pl("Error reading input")
	// 	log.Fatal(err)
	// 	os.Exit(1)
	// }
	// pl("Hello, " + name)

	//  Variables and Data Types
	// var name string = "John"
	// var v1,v2 = 1.7,2
	// v4:= 3

	// pl(reflect.TypeOf(name))
	// pl(reflect.TypeOf(v1))
	// pl(reflect.TypeOf(v2))
	// pl(reflect.TypeOf(v4))

	//  casting
	// cV1 := 1.7
	// cV2 := 2
	// cV3 := int(cV1) + cV2
	// pl("Sum:", cV3)
	// // string conversion
	// cV4 := "3"
	// cV5, err := strconv.Atoi(cV4)
	// if err != nil {
	// 	pl("Error converting string to int")
	// }
	// pl("Int", cV5)
	// //  number to string
	// cV6 := strconv.Itoa(cV5)
	// pl("String", cV6)

	// cV7 := "3.14"
	// if cV8, err := strconv.ParseFloat(cV7, 64); err == nil {
	// 	pl("Float", cV8)
	// }
	// cV9 := fmt.Sprintf("%f", 3.14)
	// pl("String", cV9)

	// operator()
	// stringOps()
	// runes()
	timeFunctions()
}

func operator() {
	//  Arithmetic operators
	//  +, -, *, /, %, ++, --
	//  Relational operators
	//  ==, !=, <, >, <=, >=
	//  Logical operators
	//  &&, ||, !
	//  Bitwise operators
	//  &, |, ^, <<, >>
	//  Assignment operators
	//  =, +=, -=, *=, /=, %=, <<=, >>=, &=, |=, ^=
	//  Misc operators
	//  &, *, <-, &^
	var a int = 60 // 0011 1100
	var b int = 13 // 0000 1101
	var c int = 0
	c = a & b // 0000 1100
	pl("a & b", c)
	c = a | b // 0011 1101
	pl("a | b", c)
	c = a ^ b // 0011 0001
	pl("a ^ b", c)
	c = a << 2 // 1111 0000
	pl("a << 2", c)
	c = a >> 2 // 0000 1111
	pl("a >> 2", c)

	var age int = 18
	if isAdult(age) {
		pl("You are an adult")
	} else {
		pl("You are not an adult")
	}

}

func isAdult(age int) bool {
	return age >= 18
}

func stringOps() {
	sV1 := "A word"
	replacer := strings.NewReplacer("A", "Another")
	sV2 := replacer.Replace(sV1)
	pl("After replcament:", sV2)
	pl("length:", len(sV1))
	pl("Contains Another:", strings.Contains(sV2, "Another"))
	pl("o index:", strings.Index(sV2, "o"))
	pl("Replace:", strings.Replace(sV2, "o", "0", -1))
	sV3 := "\n Some words \n"
	sV3 = strings.TrimSpace(sV3)
	pl("trimmed string:", sV3)
	pl("Split:", strings.Split("a-b-c-d", "-"))
	pl("Lower:", strings.ToLower(sV2))
	pl("Lower:", strings.ToUpper(sV2))
	pl("Prefix:", strings.HasPrefix(sV2, "Another"))
	pl("Suffix:", strings.HasSuffix("tacocat", "cat"))
}

func runes() {
	rStr := "abcdefg"
	pl("Rune count:", utf8.RuneCountInString(
		rStr))
	for i, runeVal := range rStr {
		fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal)
	}
}

func timeFunctions() {
	now := time.Now()
	pl(now.Year(), now.Month(), now.Day())
	pl(now.Hour(), now.Minute(), now.Second())

}

// https://youtu.be/YzLrWHZa-Kc?t=2084
