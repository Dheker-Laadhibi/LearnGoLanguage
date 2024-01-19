package main

import "fmt"

var name string
var favNumb int
const pi = 3.14
var score float32 = 5.0

//iota 
/*
The iota keyword in Go is particularly useful when you want to create
 a sequence of related constant values with a pattern. Here are some common use cases
 for iota:
 */
 // to start with 1 change  the iota keyword  to iota+1
 const (
    Red = iota
    Green
    Blue
)
func main() {
    otherFunction()
}

func otherFunction() {
    myage := 21
    fmt.Println("this is my declared variables ")
    fmt.Println(pi)
    fmt.Println(name)
    fmt.Println(favNumb)
    fmt.Println(score)
    fmt.Println("declared variable inside function", myage)
    fmt.Println(Red , Green,Blue)
}