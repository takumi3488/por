package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/takumi3488/por/cmd"
)

func main() {
	help := `Available subcommnads:
  - logs <PORT>
  - pause <PORT>
  - ps <PORT>
  - rename <PORT> <NAME>
  - restart <PORT>
  - stop <PORT>
  - unpause <PORT>`
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 1, '\t', 0)
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		println(help)
		os.Exit(1)
	}
	opt := args[0]
	port, err := strconv.Atoi(args[1])
	if err != nil {
		println(help)
		os.Exit(1)
	}

	switch opt {
	case "ps":
		fmt.Fprintln(w, "CONTAINER ID\tIMAGE\tCOMMAND\tSTATUS\tPORTS\tNAMES")
	}

	var filteredContainers []types.Container
	for _, container := range containers {
		for _, p := range container.Ports {
			if uint16(port) == p.PublicPort {
				filteredContainers = append(filteredContainers, container)
				break
			}
		}
	}
	for _, c := range filteredContainers {
		switch opt {
		case "logs":
			cmd.Logs(cli, ctx, c.ID)
		case "pause":
			cmd.Pause(cli, ctx, c.ID)
		case "ps":
			cmd.Ps(w, c)
		case "rename":
			cmd.Rename(cli, ctx, c.ID, args[2])
		case "restart":
			cmd.Restart(cli, ctx, c.ID)
		case "stop":
			cmd.Stop(cli, ctx, c.ID)
		case "unpause":
			cmd.Unpause(cli, ctx, c.ID)
		}
	}
	w.Flush()
}
