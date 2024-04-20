package contain
import (
	log "github.com/sirupsen/logrus"
	"os"
	"syscall"
)


func Contain_init(cmd string){
	log.Infof("init %s",cmd)
	argv:=[]string{}
	mountflags:=syscall.MS_NOEXEC|syscall.MS_NOSUID|syscall.MS_NODEV
	syscall.Mount("proc","/proc","proc",uintptr(mountflags),"")
	if err:=syscall.Exec(cmd,argv,os.Environ());err!=nil{
		log.Error(err)
	}
	log.Infof("-it %s finish",cmd)
}