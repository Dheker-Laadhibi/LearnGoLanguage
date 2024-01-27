package main

import (
	f "fmt"
	"time"
	s "sync"
)

func WriteMsgTochannel(c chan string, msg string){
// send msg to channel
    c <- msg
	


}

func squares(c chan int){
	for i:=0;i<=3;i++{
      //read msg from channel
		num:=   <-c
         f.Println(cap(c) , len(c))
		 f.Println( num * num)
	}
}

//write a msg 
func ping(pings chan<-string, msg string){
	pings <- msg
}
// rread msg 
func pong(pings <-chan string , pongs chan<- string){
msg:= <-pings
// push msg to channel
pongs <- msg
}

func server1( ch chan string) {
	// 6 seconds timeout
	time.Sleep(6*time.Second)
		// send msg to channel ch
	ch<-"from server 1"
}


func server2( ch chan string) {
	time.Sleep(3*time.Second)
	// send msg to channel ch
	ch<-"from server2 "
}

func service (wg *s.WaitGroup , instance int ){
time.Sleep(3*time.Second)
f.Println("service ", instance)
wg.Done()//decrement counter

}
 var i int // i =0
 var mutex = &s.Mutex{}
 func runrun(wg *s.WaitGroup){
i++
wg.Done()
 }



 func runrunMutex(wg *s.WaitGroup){
	mutex.Lock()
	i++
	mutex.Unlock()
	wg.Done()
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
//msg:= make(chan string)
// send msg to channel by goroutine 
//go WriteMsgTochannel(msg , "hello channel")
// read msg from channel 
//f.Print("before receive message")
//f.Println("message:" , <-msg)
//f.Print("after receive message from channel")

/*

The actual receiving of the message 
from the channel happens when you use 
the 
<-msg syntax.
*/

////////////////////////////////////////////////////////////////
// dead block handling - closing the channel

/*if there is no goroutine available  thats where
 deadlock error occurs crashing whole program*/

 //deadlock error occurs

 /* f.Println("main() starter")
m:= make(chan string)
m <-"hello world"
f.Println("message :" , <-m)
f.Println("main() end ") */
 
// closing channel
// close(channel)


//part 3 biffer size or channel capacity
// as we saw every  send operation on channel blocking current goroutine
// when the buffer size is  non             zero n , go routines  is not blocked until after buffer is full
//define buffer is full , n is the capacity of channel

// len & cap 2 built in functions to calculate length of channel buffer and capacity

// using buffred channel we can read from closed channel


f.Println("main() start ")

/* c:= make(chan int)

go squares(c)

// send data to channel 
c <-1
c <-2
c <-3 */


/* c1:= make(chan int , 3)
c2:= make(chan int )
go squares(c1) */
//current goroutine isn't blocked here
// send data to channel
/* c1<-1
c1<-2
c1<-3
c2<-1
c2<-2
c2<-3 */
// make a block
//c1<-4   

/* close(c2) */

// c:= make(<-chan) receive only channel
// c1 := make(chan<-int) send  only channel
//part 4 

/* pings:=make(chan string ,1)
pongs:=make(chan string ,1)
ping (pings, "passed message ")
pong(pings, pongs)
// read for pongs channel
f.Println(<-pongs)
 */

// 5 select 

// the select statement is used to choose send/ receive channel operations
//block main goroutine 
/* 
outp1 := make(chan string )
outp2 := make(chan string)

go server1(outp1)
go server2(outp2)

// select statement
select{
case s1 := <-outp1:
	f.Println(s1)
	case s2 := <-outp2:
		f.Println(s2)
		default:
			f.Println("default")

} */
//select statement randomly chooses one of its cases when multiple cases are ready.
// choose one of avilable go routines





// part 6 wait groups 

// waitgroups is a struct with a counter value which tracks the number of goroutines still working and finished ones

var wg s.WaitGroup //create a waitgroup emptty struct
/* for i:=0;i<=3;i++{
	wg.Add(1) // increment counter = 3  
   go service(&wg,i)    
}
 */

 //mutexes
 for i:=0;i<100;i++{
wg.Add(1) // increment counter 
//go runrun(&wg)
go runrunMutex(&wg)

 }
wg.Wait()//blocks here until counter 0


f.Println(i)

}