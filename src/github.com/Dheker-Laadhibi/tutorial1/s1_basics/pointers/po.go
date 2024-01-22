package main

import  f "fmt"




/*
a pointer is a programming lang object that stores a memory location
or adress . Go supports pointer types allowing u to pass references to value 


*/

func WithoutPointer(num int){
	num = 20
}

func WithPointer(num *int){
*num=30
}







func main() {
	f.Println("map course")
x:=10
f.Println(" x value  : ",x)

WithoutPointer(x)
f.Println("x(without pointer)" ,x)
f.Print(" with pointer  :---- - - - ---\n")
// here we send the memory ADDRESS not the value of x
WithPointer(&x)

f.Println("x(with pointer)",x)


a:=7
b:=&a
c:=a
f.Print("a value :" , a)
// adress of a 
f.Print("b pointer:\n",b)
//value of a in b 
f.Println("b  value :", *b)
//value in c
f.Print("c :" , c)

*b=20
f.Print(a)
c = 14
f.Println("a value :",a)









}