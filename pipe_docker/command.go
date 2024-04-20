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
		args := context.Args()
		var params string
		for i := 0; i < len(args); i++ {
    	param := args.Get(i)+" "
		params =params+param
	}

		log.Infof("get args %s",params)
		it:=context.Bool("it")

		contain.Contain_run(params,it)
		log.Info("finish run contain")
		return nil
	},
}
var initcommand = cli.Command{
	Name:  "init",
	Usage: "init a container ",
	Action: func(context *cli.Context)error {
		log.Info("init contain")
		contain.Contain_init()
		log.Info("finish init contain")
		return nil
	},
}
