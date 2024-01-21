package main

import "fmt"



func main() {
    fmt.Println("Hello, World!")

	// we can use the else if exemple nm :=3 ; num<0 print... else if num<10 print...  else ...
	if 8 %3 == 0 {
		fmt.Println("8 is divisible by 3")
	} else {
	fmt.Println("8 is not  divisible by 3")
	}
	


num:=10
switch{
	case num<0:
        fmt.Println("num is negative")
    case num==0:
        fmt.Println("num is zero")
    case num<10:
        fmt.Println("num is less than 10")
    // default is like the final else 
	default:
        fmt.Println("num is greater than 10")
}
// there is no ternary if in Go , so you will need to use full of (condition ? value if true )




}