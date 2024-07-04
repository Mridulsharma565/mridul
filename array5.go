package main
import("fmt")
func main(){
var x= [4]int{1,2,3,4}
y:=[...]int{5,6,7,8,9,10}

fmt.Println(len(x))
fmt.Println(len(y))
fmt.Println(cap(x))
}