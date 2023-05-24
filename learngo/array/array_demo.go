package main

import "fmt"

func main() {
	// 声明一个数组，并设置零值
	var array1 [5]int
	fmt.Println(array1)

	// 使用数组字面量声明数组
	array2 := [5]int{10, 20, 30, 40, 50}
	fmt.Println(array2)

	// 容量又初始值的数量决定
	array3 := [...]int{10, 20, 30, 40, 50}
	fmt.Println(array3)

	// 声明数组指定特定元素的值
	array4 := [5]int{1: 10, 3: 30}
	fmt.Println(array4)

	// 访问数组的元素
	array5 := [5]int{10, 20, 30, 40, 50}
	fmt.Println(array5[2])

	// 修改数组元素的值
	array5[0] = 100
	fmt.Println(array5)

	// 访问指针数组的元素
	array6 := [5]*int{0: new(int), 1: new(int)}
	*array6[0] = 10
	*array6[1] = 20
	fmt.Println(*array6[0])

	// 把一个指针数组赋值给另一个
	var array7 [3]*string
	array8 := [3]*string{new(string), new(string), new(string)}

	*array8[0] = "Red"
	*array8[1] = "Blue"
	*array8[2] = "Green"

	array7 = array8
	fmt.Println(array7)
	fmt.Println(array8)
}
