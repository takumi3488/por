package cmd

import (
	"context"
	"fmt"

	"github.com/docker/docker/client"
)

func Unpause(cli *client.Client, ctx context.Context, container string) {
	if err := cli.ContainerUnpause(ctx, container); err != nil {
		panic(err)
	}
	fmt.Println(container[0:12])
}
