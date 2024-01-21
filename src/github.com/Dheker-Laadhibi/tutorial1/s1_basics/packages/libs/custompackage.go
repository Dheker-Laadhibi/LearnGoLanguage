package custompackage
import f "fmt"
func SayHello() {
	f.Println("Helloooo")
}

func SayHi() {
	f.Println("Hi")
}

func init() {
	SayHello()
}