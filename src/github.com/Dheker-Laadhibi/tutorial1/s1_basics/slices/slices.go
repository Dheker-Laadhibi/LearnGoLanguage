package main
import "fmt"
// a slice can be considred as a dynamic array
// a slice can grow as we need
// slices are refferences
// declaring slices be like  var slicename [] type = make([] type) , len)


func WorkWithArrayy(array [] int ) [] int{
	array[0] = 10
    array[1] = 11
    array[2] = 12
    return array
}



func main() {

var magicNumbers []int = make([]int, 5)
fmt.Println("magic numbers slice : " , magicNumbers)
//default length 
fmt.Println("length :" , len(magicNumbers))

fmt.Println("work with function now :\n " )

fmt.Println("slices are refferenced : " ,WorkWithArrayy(magicNumbers), magicNumbers)
fmt.Println("length of slices : " , len(magicNumbers))

// fetch  an element in a slice
//before the third element
slice1 :=magicNumbers[:3]
//after the third element
slice2 := magicNumbers[3:]

slice3 := magicNumbers[1:4]
slice33:=magicNumbers[3]
fmt.Println(slice1, slice2, slice3 ,slice33)






} 