package main

import (
	
	log "github.com/sirupsen/logrus"
	"run_docker/container"
	"github.com/urfave/cli"
)

var runcommand = cli.Command{
	Name:  "run",
	Usage: "start a container ",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "it",
			Usage: "start in some type",
		},
	},
	Action: func(context *cli.Context) error {
		log.Info("run contain")
		cmd:=context.Args().Get(0)
		it:=context.Bool("it")
		contain.Contain_run(cmd,it)
		log.Info("finish run contain")
		return nil
	},
}
var initcommand = cli.Command{
	Name:  "init",
	Usage: "init a container ",
	Action: func(context *cli.Context)error {
		log.Info("init contain")
		cmd:=context.Args().Get(0)
		contain.Contain_init(cmd)
		log.Info("finish init contain")
		return nil
	},
}
