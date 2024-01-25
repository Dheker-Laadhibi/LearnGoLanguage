package main
// an error is just a value  that a function can return if something unexpected happened
// Go does not provide try and catch statements 
// error is a built-in type in go and its zero value is nil 
import f "fmt"
import "errors"





type MyCustomMath struct{}


func (math MyCustomMath) Error() string {
	return "our first error"
}



// devide func  return value and error nil  or  0  with error


func devide(x, y float64)(float64 ,error){
	if y == 0 {
        return 0, errors.New("zero devision error  : y cannot be zero")
    }
    return x/y, nil
}





func main() {
	//println function is designed to call error method automatically when value is an error 
	doMath := MyCustomMath{}
	f.Println(doMath)
	//type 
f.Printf("type do math %T\n", doMath)

var x , y  float64
x=10
y=0

val , err := devide(x,y)
if err != nil {
	f.Println(err)
}else{
	f.Println(val)
}
err2:= errors.New("devide failed")
f.Println(err2)


}