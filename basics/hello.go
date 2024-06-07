package main

import (
	"fmt"
	"math"
	"math/rand"
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
	// timeFunctions()

	// mathOps()
	// loops()
	arrays()
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

func mathOps() {
	pl(5 + 5)
	pl(5 - 3)
	mInt := 1
	mInt += 1
	pl(mInt)
	mInt++
	pl(mInt)

	// random number logic
	seedSecs := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seedSecs))
	ranNum := r.Intn(50) + 1
	pl("Random: ", ranNum)

	// operators
	pl("Abs(-12):", math.Abs(-12))
	pl("Power 2 to 3:", math.Pow(2, 3))
	pl("Sqrt(9):", math.Sqrt(9))

	// Format specifiers
	fmt.Printf("%s %d %c %f %t %o %x\n",
		"stuff", 1, 'A', 3.14, true, 1, 1)

	sp1 := fmt.Sprintf("%9.f\n", 3.141592)
	pl(sp1)
}

func loops() {
	for x := 1; x <= 5; x++ {
		pl(x)
	}

	fX := 0
	for fX < 5 {
		pl(fX)
		fX++
	}
	pl("Final value of fX:", fX)

	pl("While loop")
	// while loop
	fY := 0
	for { // this equals to `for true {`
		pl(fY)
		if fY == 4 {
			break
		}
		fY++
	}

	pl("iterating through an array")
	aNums := []int{1, 2, 3, 4}
	for _, num := range aNums {
		pl(num)
	}

}

func arrays() {
	var arr1 [5]int
	arr1[0] = 3
	arr1[2] = 4
	arr2 := [5]int{1, 2, 3, 4, 5}
	pl("index 0:", arr2[0])
	pl("length arr2:", len(arr2))

	for i := 0; i < len(arr2); i++ {
		pl(arr2[i])
	}
	for i, v := range arr2 {
		fmt.Printf("%d : %d\n", i, v)
	}

	pl("multi dimensional array")
	arr3 := [2][3]int{
		{1, 2, 4},
		{2, 3, 3},
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			pl(arr3[i][j])
		}
	}
}

// https://youtu.be/YzLrWHZa-Kc?t=3545
