package main
import  f "fmt"

func WriteMsgTochannel(c chan string, msg string){
// send msg to channel
    c <- msg
	


}













func main() {
	f.Println("Hello World")
// declaring channel 
// printing type and value 
// nil channel is not useful you can not use to read data or pass data from a channel which is nil
/*	
var c chan int
	f.Printf("channe; value %v\n   type %T\n", c , c) 
// nil value  type chan int 
	*/

	// declaring a ready to use channel 
	c1:= make(chan string)
	// value is an adress memory 
	f.Printf("channel : value :  %v\n type %T\n", c1, c1)
}