package main

import (
	"github.com/urfave/cli"
)

var commonFlags = []cli.Flag{
	cli.StringFlag{
		Name:   "bind-grpc",
		Usage:  "bind address for gRPC",
		EnvVar: "BIND_GRPC",
		Value:  ":2338",
	},
}