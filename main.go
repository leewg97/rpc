package main

import (
	"flag"
	"log"
)

var configFlag = flag.String("config", "./config.toml", "config path")

func main() {
	flag.Parse()

	log.Println("env 경로는 다음과 같다,", *configFlag)
	//cfg := config.NewConfig(*configFlag)
	//
	//if err := server.NewGRPCServer(cfg); err != nil {
	//	panic(err)
	//} else {
	//	time.Sleep(1e9)
	//	cmd.NewApp(cfg)
	//}
}
