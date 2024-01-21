package main

import f "fmt"

func WorkWithArrayy(arr [3]int) [3]int {

	arr[0] = 10
	arr[1] = 11
	arr[2] = 12
	return arr
}

func main() {

	f.Println("Hello World")
	var myArray [3]int
	f.Println("myArray  :\n", myArray)
	f.Println("length :\n", len(myArray))
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3
	f.Println("myArray after set values  :\t", myArray)

	//last element in myArray
	f.Println("last element in myArray", myArray[len(myArray)-1])

	//shorthand for array
	myArray2 := [3]string{"java", "angular", "dheker"}
	f.Println("myArray2 after set values :\t", myArray2)

	// call the function
	f.Println("after calling the function ", WorkWithArrayy(myArray), myArray)

	/*
	   matrix
	   multi-dimensional data structure
	*/

var myMatrix [2][3] int
for i :=0; i<2; i++{
	for j :=0; j<3; j++{
        myMatrix[i][j] = i*3 + j
    }
	
	
}

f.Println("matrix", myMatrix)
f.Print("\n")
f.Println(myMatrix[0][1], myMatrix)











}
