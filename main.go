package main

import (
	"fmt"
	"github.com/Jeremy-boo/go-docker-handbook/manager"
)

func main() {
	cli := new(manager.ClientManager)
	values, err := cli.ListContainer()
	if err != nil {
		fmt.Printf("error:%v", err.Error())
		return
	}
	fmt.Println(values)
}
