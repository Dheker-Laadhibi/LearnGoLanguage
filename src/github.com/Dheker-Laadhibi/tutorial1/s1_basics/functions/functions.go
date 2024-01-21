package main

import  f "fmt"
	func add(x,y int) int {
		return x + y
	}
	//return multiple values 
	//sum is the sum of the values
	func domath(x,y int) (sum int , diff int) {
		return x+y, x-y
	}
	//variadic function when you don't want to specify number of arguments
// ... = illimited number of arguments
	func sums( nums ... int)int  {
	total := 0
	for i:=0; i<len(nums); i++ {
total += nums[i]
	}
	return total
	}
	func main() {
		f.Println(add(1,2))
		f.Println(add(2,3))
		sum,diff := domath(10,5)
		f.Println(sum)
		f.Println(diff)
		f.Printf("sum = %d\t diff = %d\n", sum, diff)
		f.Println(sums(1,2,3,4,5,7,8,9,10))

		//anonymous function
		func ()  {
			f.Println("anonymous function")
		}()
	}
