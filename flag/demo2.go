package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	//name := flag.String("name", "张三", "姓名")
	//age := flag.Int("age", 18, "年龄")
	//married := flag.Bool("married", false, "婚否")
	//delay := flag.Duration("d", 0, "时间间隔")

	var name string
	var age int
	var married bool
	var delay time.Duration
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "时间间隔")

	flag.Parse()
	fmt.Println(name, age, married, delay)
	fmt.Println(flag.Args())
	fmt.Println(flag.NArg())
	fmt.Println(flag.NFlag())
}
