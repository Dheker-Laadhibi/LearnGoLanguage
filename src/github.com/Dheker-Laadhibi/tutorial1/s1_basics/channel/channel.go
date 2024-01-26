package main
import  f "fmt"

func WriteMsgTochannel(c chan string, msg string){
// send msg to channel
    c <- msg
	


}













func main() {
	f.Println("Hello channel!")
// declaring channel 
// printing type and value 
// nil channel is not useful you can not use to read data or pass data from a channel which is nil
/*	
var c chan int
	f.Printf("channe; value %v\n   type %T\n", c , c) 
// nil value  type chan int 
	*/

	// declaring a ready to use channel 
	//c1:= make(chan string)
	// value is an adress memory 
	// channels by default are  pointers
	//f.Printf("channel : value :  %v\n type %T\n", c1, c1)





// new channel called msg 
msg:= make(chan string)
// send msg to channel by func 
go WriteMsgTochannel(msg , "hello channel")
// read msg from channel 
f.Print("before receive message")
f.Println("message:" , <-msg)
f.Print("after receive message from channel")

/*

The actual receiving of the message 
from the channel happens when you use 
the 
<-msg syntax.
*/






}