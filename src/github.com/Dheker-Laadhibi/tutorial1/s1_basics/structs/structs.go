package main

import "fmt"
  // declaring struct
  type Employee struct {
	name string
	age  int
}
  // new emp constructor
  func Newemployee(name string, age int) *Employee {
	return &Employee{name: name, age: age}
}
// unlike traditional object-oriented programming, go doesn't have a class object architecture rather
    /*
        we have structs
    */

func main() {
    fmt.Println("what is structs in go ")
    
  emp:= &Employee{"algo",24}
fmt.Println("first empolyee : \n" , emp)
fmt.Println("employee age  : \n" ,emp.age)
fmt.Println("employee name : \n" , emp.name)
//type of employee
fmt.Printf("Type of emp: %T\n", emp)







 emp2:=Newemployee("dhekerLaadhibi",21)
 fmt.Println("constructor of employee" ,emp2)
}