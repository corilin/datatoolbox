package main
import  "fmt"
func main(){

	for x:=1; x<1000000; x++ {
		fmt.Print(x)
	}
	fmt.Println()
	fmt.Println("-----------")
	for x:=9; x>0; x-- {
		fmt.Print(x)
	}
	fmt.Println()
	fmt.Println("-----------")
	for x:=6; x<10; x++ {
		fmt.Print(x)
	}
	fmt.Println()
	fmt.Println("-----------")
	for x:=2; x<11; x=x+2 {
		fmt.Print(x)
	}
	fmt.Println()
	fmt.Println("-----------")
	for x:=2; x<11; x+=2{
		fmt.Print(x)
	}
	fmt.Println()
	fmt.Println("-----------")
	for x:=1; x<10; x+=2{
		fmt.Print(x)
	}
	fmt.Println()
	fmt.Println("-----------")
	for x:=1; x<11; x++{
		if x%2 != 0 {
		  fmt.Print(x)
		}
	}
}