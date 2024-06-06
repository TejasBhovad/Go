package main

import (
	"fmt"
	"strconv"
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
	cV1 := 1.7
	cV2 := 2
	cV3 := int(cV1) + cV2
	pl("Sum:", cV3)
	// string conversion
	cV4 := "3"
	cV5, err := strconv.Atoi(cV4)
	if err != nil {
		pl("Error converting string to int")
	}
	pl("Int", cV5)
	//  number to string
	cV6 := strconv.Itoa(cV5)
	pl("String", cV6)

	cV7 := "3.14"
	if cV8, err := strconv.ParseFloat(cV7, 64); err == nil {
		pl("Float", cV8)
	}
	cV9 := fmt.Sprintf("%f", 3.14)
	pl("String", cV9)
}
