package main

import (
	"fmt"
	"strings"
)

func strSplit(str string) {
	resStr := strings.Split(str, ",")
	for _, v := range resStr {
		fmt.Println(v)
	}
}

func main() {

	str := "alice,bob,cindy"
	strSplit(str)

}
