package main

import (
	"github.com/docker/machine/libmachine/drivers/plugin"
	"github.com/goldfix/docker-machine-vmwareworkstation"
)

func main() {
	plugin.RegisterDriver(vmwareworkstation.NewDriver("", ""))
}
