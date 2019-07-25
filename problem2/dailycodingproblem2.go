package main

import "fmt"

func mult(x, y int) int {
	return x * y
}

func main(){
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{0,0,0,0,0}
	temp := 1
	for k := 0; k < len(b); k++{
		b[k] = temp/a[k]
		temp = 1
		for j := 0; j < len(a); j++{
			temp *= a[j]
		}
	}
	fmt.Println(b)
}
