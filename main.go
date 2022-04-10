package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
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
	cmd := args[0]
	port, err := strconv.Atoi(args[1])
	if err != nil {
		panic(err)
	}

	switch cmd {
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
		switch cmd {
		case "ps":
			ports := ""
			for _, p := range c.Ports {
				ports += fmt.Sprintf("%s:%d->%d/%s, ", p.IP, p.PrivatePort, p.PrivatePort, p.Type)
			}
			names := ""
			for _, n := range c.Names {
				names += fmt.Sprintf("%s, ", n)
			}
			fmt.Fprintf(w, "%v\t%v\t\"%v\"\t%v\t%v\t%v\n", c.ID[0:12], c.Image, c.Command, c.Status, strings.TrimRight(ports, ", "), strings.TrimRight(names, ", "))
		case "stop":
			if err := cli.ContainerStop(ctx, c.ID, nil); err != nil {
				panic(err)
			}
			fmt.Println(c.ID[0:12])
		}
	}
	w.Flush()
}
