package cmd

import (
	"fmt"
	"strings"
	"text/tabwriter"

	"github.com/docker/docker/api/types"
)

func Ps(w *tabwriter.Writer, container types.Container) {
	ports := ""
	for _, p := range container.Ports {
		ports += fmt.Sprintf("%s:%d->%d/%s, ", p.IP, p.PrivatePort, p.PrivatePort, p.Type)
	}
	names := ""
	for _, n := range container.Names {
		names += fmt.Sprintf("%s, ", n)
	}
	fmt.Fprintf(w, "%v\t%v\t\"%v\"\t%v\t%v\t%v\n", container.ID[0:12], container.Image, container.Command, container.Status, strings.TrimRight(ports, ", "), strings.TrimRight(names, ", "))
}
