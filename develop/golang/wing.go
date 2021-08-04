package main

import "fmt"

func main() {
	for x := 1; x < 10; x++ {
		for y := 1; y < 10; y++ {
			content := fmt.Sprintf("%d * %d = %d", x, y, (x * y))
			fmt.Println(content)
		}

		fmt.Println("-------------")
	}
}
