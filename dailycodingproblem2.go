package main

import "fmt"

// this is a comment

func mult(x, y int) int {
	return x * y
}
func main(){
	a := [5]int{1, 2, 3, 4, 5}
	for j := 0; j < len(a)-2; j++ {
		temp := mult(a[j+1],a[j+2])
		temp *= a[j+3]
		fmt.Println(temp)
    }
}
