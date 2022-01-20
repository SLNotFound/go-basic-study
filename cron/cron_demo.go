package main

import (
	"github.com/robfig/cron"
	"log"
)

func main() {
	i := 0
	c := cron.New()
	spec := "0 */1 * * * ?"
	err := c.AddFunc(spec, func() {
		i++
		log.Println("cron running: ", i)
	})
	if err != nil {
		panic(err)
	}
	c.Start()
	select {}
}
