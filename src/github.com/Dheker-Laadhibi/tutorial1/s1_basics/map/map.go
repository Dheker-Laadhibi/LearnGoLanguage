package main

import  f "fmt"




/*
a map is a collection of key/value pairs called dictionaries in python
var mapname map[keytype]value type = make  map[keytype]valuetye

shorthand declaration : 
mapname:= map[keytype]valuetye{key:val , key:val}



*/
func main() {
	f.Println("map course")

var m map[string]int = make(map[string]int)
m["java"]=14
m["security"]=10
m["soa"]=20
f.Println("my first map :\t ",m)
f.Println("len of m",(len(m)))
f.Println("java rate :", m["java"])













}