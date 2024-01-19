
///starting point for Go compiler
package main
// fmt is package that provides function for printing text 
import "fmt"
var  myINteger int32 = 0
var myINteger1 int =5
func main() {
    // print hello world
    fmt.Println("Hello, World!")


    // adding the numbers 

    x ,y,z := -5 ,"algos", true
/*
    %d: This is a format specifier used to display integers in the Go programming language.

%s: This is a format specifier used to display strings (sequences of characters) in Go.

%t: This is a format specifier used to display boolean values (true or false) in Go.
*/
    fmt.Printf("x: %d\ty: %s\tz : %t\n", x, y, z)
}
/*

multiple lines  comment
*/
//single line comment