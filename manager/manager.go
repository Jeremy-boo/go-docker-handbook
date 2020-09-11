package manager

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"strings"
)

type ClientManager struct {
	Client *client.Client
}

func (manager *ClientManager) NewClient() *client.Client {
	cli, err := client.NewEnvClient()
	if err != nil {
		fmt.Printf("init docker client error:%v", err.Error())
		return nil
	}
	return cli
}

func (manager *ClientManager) ListContainer() (interface{}, error) {
	values := make(map[string]interface{}, 0)
	cli := manager.NewClient()
	ctx := context.Background()
	options := types.ContainerListOptions{
		All: true,
	}
	containers, err := cli.ContainerList(ctx, options)
	if err != nil {
		return nil, err
	}
	for _, container := range containers {
		name := strings.Replace(container.Names[0], "/", "", -1)
		values["name"] = name
		values["status"] = container.Status
		values["state"] = container.State
		values["id"] = container.ID
	}
	return values, nil
}
