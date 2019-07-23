package main

import "fmt"

// this is a comment

func add(x, y int) int {
	return x + y
}
func main(){
	a := [5]int{10, 15, 3, 7, 5}
	k := 17
	for j := 0; j < len(a); j++ {
		if add(a[0],a[j]) == k {
			fmt.Println(a[0])
			fmt.Println(a[j])
		}
    }
}