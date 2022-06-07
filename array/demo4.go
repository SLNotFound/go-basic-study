package main

import "fmt"

func main() {
	var arr1 [3]int
	var arr2 = [3]int{1, 2}
	var arr3 = [3]string{"alice", "bob", "cindy"}
	var arr4 = [...]int{1, 2, 3, 4}
	arr5 := [...]int{1: 1, 3: 5}
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)

	fmt.Printf("type of arr4: %T\n", arr4)
	fmt.Println(arr5)
}
