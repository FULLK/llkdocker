package contain

import (
	"io"
	"os"
	"os/exec"
	"strings"
	"syscall"

	log "github.com/sirupsen/logrus"
)


func Contain_init(){
	
	cmd:=readpipe()
	log.Infof("init %s",cmd)
	
	log.Infof("cmd[0] %s",cmd[0]=="/bin/sh")
	log.Infof("cmd len %d",len(cmd))
	log.Infof("cmd %s",cmd[0])

	path,err:=exec.LookPath(cmd[0])
	if err!=nil{
		log.Fatal(err)
	}
	argv:=[]string{}
    for i:=1;i<len(cmd);i++{
		argv=append(argv,cmd[i-1])
	}
	
	mountflags:=syscall.MS_NOEXEC|syscall.MS_NOSUID|syscall.MS_NODEV
	syscall.Mount("proc","/proc","proc",uintptr(mountflags),"")
	defer syscall.Unmount("/proc",0)
	if err:=syscall.Exec(path,argv,os.Environ());err!=nil{
		log.Error(err)
	}
	log.Infof("-it %s finish",cmd)
}
func readpipe()[]string{
	readpipe:=os.NewFile(uintptr(3),"pipe")
	
	cmd,err:=io.ReadAll(readpipe)
	if err!=nil{
		log.Fatal(err)
	}
	cmdstr:=string(cmd)
	return strings.Split(cmdstr," ")
}