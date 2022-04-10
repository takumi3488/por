package cmd

import (
	"context"
	"fmt"

	"github.com/docker/docker/client"
)

func Stop(cli *client.Client, ctx context.Context, container string) {
	if err := cli.ContainerStop(ctx, container, nil); err != nil {
		panic(err)
	}
	fmt.Println(container[0:12])
}
