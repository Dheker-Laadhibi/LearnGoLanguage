package main

import "fmt"

var name string
var favNumb int
const pi = 3.14
var score float32 = 5.0

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
}