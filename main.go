package main

import (
	"flag"
	"rpc-server/cmd"
	"rpc-server/config"
)

var configFlag = flag.String("config", "./config.toml", "config path")

func main() {
	flag.Parse()

	// fmt.Println(*configFlag)

	cfg := config.NewConfig(*configFlag)
	cmd.NewApp(cfg)
}
