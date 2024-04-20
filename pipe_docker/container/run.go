package contain

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)
func Contain_run(cmd string,it bool) error{
	command,writepipe:=new_contain_process(it)
	log.Infof("cmd %s it %d",cmd,it)
	
	
	if err:=command.Start();err!=nil{
		log.Error(err)
	}
	log.Infof("cmd %s",cmd)
	cmdstring:=strings.Split(cmd, " ")
	log.Infof("split cmd %s",cmdstring)
	write_to_pipe(cmdstring,writepipe)
	
	command.Wait()
	// 从容器内的命令行中的退出才会wait结束
	log.Infof("exit -1 !!!")
	os.Exit(-1) 
	return nil
}

func write_to_pipe(cmd []string,pipe *os.File){
	command:=strings.Join(cmd, " ")
	log.Infof("write to pipe command %s",command)
	pipe.WriteString(command)
	pipe.Close()
}