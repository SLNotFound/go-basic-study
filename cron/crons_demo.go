package main

import (
	"fmt"
	"github.com/robfig/cron"
	"log"
)

type TestCron struct {
}

func (this TestCron) Run() {
	fmt.Println("TestCron11111")
}

type Test2Cron struct {
}

func (this Test2Cron) Run() {
	fmt.Println("TestCron22222")
}

func main() {
	i := 0
	c := cron.New()
	// 5秒一次
	spec := "*/5 * * * * ?"
	c.AddFunc(spec, func() {
		i++
		log.Println("cron running: ", i)
	})

	c.AddJob(spec, TestCron{})
	c.AddJob(spec, Test2Cron{})

	c.Start()

	defer c.Stop()

	select {}
}
