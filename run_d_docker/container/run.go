package contain

import (
	"run_docker/cgroups"
	"os"
	"strings"
	log "github.com/sirupsen/logrus"
)
func Contain_run(cmd string,it bool,resource_config *cgroups.Resource,volume string) {
	command,writepipe,rooturl:=new_contain_process(it,volume)
	log.Infof("cmd %s it %t",cmd,it)
	
	if err:=command.Start();err!=nil{
		log.Error(err)
	}
	log.Infof("cmd %s",cmd)
	cmdstring:=strings.Split(cmd, " ")
	log.Infof("split cmd %v",cmdstring)

	llkdockercgroups:=cgroups.Cgroups{
		Cgroups_Name: "llkdockercgroups",
		Resour:resource_config,
		Sub:cgroups.Subsystemins,
	}
	log.Info(llkdockercgroups.Resour)

	cgroups_path:=cgroups.Get_cgroups_path("cgroup",llkdockercgroups.Cgroups_Name)
	log.Info("in run.go get cgroups_path")
	llkdockercgroups.Move(command.Process.Pid,cgroups_path)

	llkdockercgroups.Set(cgroups_path)
	

	write_to_pipe(cmdstring,writepipe)
	if it{ //-it实现交互，那么当前终端就不能关闭退出，还需留给子进程使用
		command.Wait()
		// 从容器内的命令行中的退出才会wait结束
		end_volume(rooturl,volume)
		end_overlays(rooturl)
		llkdockercgroups.Remove(cgroups_path) //-d后台运行的话不能删除cgroup
	}
	
	log.Infof("exit  !!!")
	
}

func write_to_pipe(cmd []string,pipe *os.File){
	command:=strings.Join(cmd, " ")
	log.Infof("write to pipe command %s",command)
	pipe.WriteString(command)
	pipe.Close()
}