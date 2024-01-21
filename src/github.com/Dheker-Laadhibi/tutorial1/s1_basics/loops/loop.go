package main

import  f "fmt"
func main() {
i:=0
// for loop with a single condition
for i <=10{
	f.Println(i)
    i++
	// i+=i
	//i = i+1
}

f.Println("")
/*
decrement operation */
for j :=5 ; j>=0 ; j--{
	f.Println(j)

}

// continue operation
/*
if n = 1  or n mod 2 ><  0 get out of if statement only.
*/
f.Println(" for with if ")
for n :=0;n>=10 ; n++{
if n%2 != 0 {
	continue
}
f.Println(n)
}


f.Println("\n")
 
// if k =4  get out of the loop not only the if statement
for k :=0;k>=10 ; k++{
	if k == 4 {
		break
	}
	f.Println(k)
	}

// now i should be able to understand diff between break and continue


}





