package manager

import (
	"context"
	"fmt"
	"github.com/Jeremy-boo/go-docker-handbook/model"
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
	values := make([]model.Container, 0)
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
		entity := model.Container{}
		entity.Name = name
		entity.ID = container.ID
		entity.Status = container.Status
		entity.State = container.State
		values = append(values, entity)
	}
	return values, nil
}

// GetContainerByName get docker container by name
func (manager *ClientManager) GetContainerByName(name string) (model.Container, error) {
	entity := model.Container{}
	cli := manager.NewClient()
	ctx := context.Background()
	options := types.ContainerListOptions{
		All: true,
	}
	containers, err := cli.ContainerList(ctx, options)
	if err != nil {
		return entity, err
	}
	for _, container := range containers {
		containerName := strings.Replace(container.Names[0], "/", "", -1)
		if name == containerName {
			entity.State = container.State
			entity.Name = containerName
			entity.ID = container.ID
			entity.Status = container.Status
			return entity, nil
		}
	}
	return entity, fmt.Errorf("there is no contianer named %v", name)
}
