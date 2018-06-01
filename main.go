package main

import (
	"log"

	"git.shopex.cn/luban/lib"
)

func main() {
	cmd_runner := &lib.Commands{
		cmdServer,
	}
	log.SetFlags(log.LstdFlags | log.Llongfile)

	cmd_runner.Run()
}
