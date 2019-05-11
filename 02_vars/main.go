package main // to run the main function like java

import "fmt" // to output

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int  int8  int16  int32  int64
	// uint uint8 uint16 uint32 uint64 uintptr for 0+ (no negative nos)
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	//Using var
	var name string = "Sahil" // can ommit type string
	// if declared, it should always be used -- othersie it shows error
	var age int = 22 // can ommit type int

	fmt.Println(name, age)

	// %T	a Go-syntax representation of the type of the value
	fmt.Printf("%T\n", age)
}
