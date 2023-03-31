package cmd

import (
	"fmt"
	"context"

	"github.com/docker/docker/client"
)

func Pause(cli *client.Client, ctx context.Context, container string) {
	if err := cli.ContainerPause(ctx, container); err != nil {
		panic(err)
	}
	fmt.Println(container[0:12])
}
