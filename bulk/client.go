package bulk

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func names() map[string]map[string]string {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	ret := make(map[string]map[string]string)
	for _, container := range containers {
		fmt.Printf("%s %#v\n", container.ID[:10], container.Names)
		fmt.Printf("\t%#v\n", container.Labels)
		ret[container.Names[0]] = container.Labels
	}
	return ret
}
