package main

import "fmt"

// 根据分数 判断获得的评级
func level(score int) {
	if score > 100 || score < 0 {
		fmt.Println("The score is out of 0 ~ 100, please re-input..")
	} else {
		switch {
		case score >= 90:
			fmt.Println("You get A!!!")
		case score > 80:
			fmt.Println("You get B!!!")
		case score > 60:
			fmt.Println("You get C!!!")
		default:
			fmt.Println("You get failed!!!")
		}
	}

}

func main() {
	level(22)
}
