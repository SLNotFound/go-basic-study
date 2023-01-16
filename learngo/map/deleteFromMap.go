package main

import "fmt"

func main() {
	ageMap := make(map[string]int)
	ageMap["alice"] = 21
	ageMap["bob"] = 22
	ageMap["cindy"] = 23
	fmt.Println(ageMap)
	delete(ageMap, "alice")
	fmt.Println(ageMap)
}
