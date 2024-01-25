package main 
import  f "fmt"
func main() {
	f.Println(" rang course ")
	nums:=[] int {1,2 , 3 , 4, 5, 6, 7, 8, 9}
	for i , v := range nums {
		f.Printf("value  at index %v is %v", i, v)
}

// calculate  values only. to remove index put _ instead of i

sum:=0
for _,v:=range(nums) {
sum=sum + v

}
f.Println(sum)
}