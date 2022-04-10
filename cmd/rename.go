package cmd

import (
	"context"
	"fmt"

	"github.com/docker/docker/client"
)

func Rename(cli *client.Client, ctx context.Context, container string, name string) {
	if err := cli.ContainerRename(ctx, container, name); err != nil {
		panic(err)
	}
	fmt.Println(container[0:12])
}
