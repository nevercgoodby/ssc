package main

import (
	"flag"
	"fmt"
)

var cqssc CQSSC

func main() {
	flag.Parse()

	conf, err := ConfigInit("./config.toml")
	if err != nil {
		panic(err)
	}
	cqssc.BuyHost = conf.BuyHost
	cqssc.TimeHost = conf.TimeHost
	fmt.Println(conf)
	fmt.Println(cqssc)
	cqssc.Run()
}
