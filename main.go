package main

import (
	"fmt"
	"github.com/Jeremy-boo/go-docker-handbook/manager"
)

func main() {
	cli := new(manager.ClientManager)
	entity, err := cli.GetContainerByName("my_nginx")
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	fmt.Println(entity)
}
