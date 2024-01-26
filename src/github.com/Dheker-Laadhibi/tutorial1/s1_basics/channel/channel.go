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
	var c chan int
	f.Printf("channe; value %v\n   type %T\n", c , c) 
}