package main
//Go supports methods defined on struct types
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

	//new struct  and his constructor
type Rectangle struct {
	width  int
	height int
}
func NewRectangle(width int, height int) *Rectangle {

    return &Rectangle{width: width, height: height}

}
// value getter receiver

func (r *Rectangle) Area() int {
    return r.width * r.height
}
//setter

func (r *Rectangle) SetWidth(width int) {
    r.width = width
}
func (r *Rectangle) SetHeight(height int) {
    r.height = height
}

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

rect:=NewRectangle(10,5)
fmt.Println("rect",*rect)
fmt.Println("area",rect.Area())

}