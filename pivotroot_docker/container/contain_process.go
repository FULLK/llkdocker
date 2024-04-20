package contain
import(
	log "github.com/sirupsen/logrus"
	"os/exec"
	"syscall"
	"os"
)

func new_contain_process(it bool)  (*exec.Cmd,*os.File){	
	
	readpipe,writepipe,err:=os.Pipe()
	if err!=nil{
		log.Fatal(err)
	}
	command:=exec.Command("/proc/self/exe","init")
	command.SysProcAttr=&syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS |
			syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC,
	}
	if it{
		command.Stdin=os.Stdin
		command.Stdout=os.Stdout
		command.Stderr=os.Stderr
	}
	log.Info(command)
	log.Infof("sucess create a contain process")
	command.ExtraFiles=[]*os.File{readpipe}
	command.Dir = "/home/llk/Desktop/docker/src/pivotroot_docker/busybox"
	log.Info(readpipe)
	return command,writepipe
	
}
