package main

import (
	"run_docker/cgroups"
	log"github.com/sirupsen/logrus"
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
		cli.StringFlag{
			Name:  "mem",
			Usage: "memory limit",
		},
		cli.StringFlag{
			Name:  "cpu",
			Usage: "CPU usage",
		},
		cli.StringFlag{
			Name:  "cpuset",
			Usage: "CPU number",
		},
		cli.StringFlag{
			Name: "v",
			Usage: "-v <宿主机目录>:<容器内目录>",

		},
		cli.BoolFlag{  //有-d即为真
			Name: "d",
			Usage: "-d 后台运行某个容器",

		},
	},
	Action: func(context *cli.Context) error {
		
		args := context.Args()
		
		var params string
		for i := 0; i < len(args); i++ {
    	param := args.Get(i)+" "
		params =params+param
	}	
		log.Info(args)
		log.Info(context.Bool("it"))
		log.Info(context.String("mem"))
		log.Info(context.String("cpu"))
		log.Info(context.String("cpuset"))
		log.Info(context.String("v"))
		log.Info(context.Bool("d"))

		it:=context.Bool("it")
		detach := context.Bool("d")
		volume := context.String("v")

		resource_config:=&cgroups.Resource{
			Mem:context.String("mem"),
			Cpu:context.String("cpu"),
			Cpuset:context.String("cpuset"),
		}
		log.Info(resource_config)  //相关资源配置
		if it&&detach{
			log.Fatal("it 和 d 不能同时使用 error ！！")
		}
		contain.Contain_run(params,it,resource_config,volume)
		
		return nil
	},
}
var initcommand = cli.Command{
	Name:  "init",
	Usage: "init a container ",
	Action: func(context *cli.Context)error {
		contain.Contain_init()
		return nil
	},
}

var commitcommand = cli.Command{
	Name:  "commit",
	Usage: "commit  image",
	Action: func(context *cli.Context)error {
		args := context.Args()//commit后参数的作为这里
		if len(args)<1 {
			log.Fatal("missing the image name you want to save ")
		}

		log.Info(args)
		contain.Contain_commit(args)
		return nil
	},
}