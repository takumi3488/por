package cmd

import (
	"io"
	"os"
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func Logs(cli *client.Client, ctx context.Context, container string) {
	options := types.ContainerLogsOptions{ShowStdout: true}
	out, err := cli.ContainerLogs(ctx, container, options)
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, out)
}
