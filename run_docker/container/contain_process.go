package contain
import(
	log "github.com/sirupsen/logrus"
	"os/exec"
	"syscall"
	"os"
)

func new_contain_process(cmd string,it bool)  *exec.Cmd {	
	args:=[]string{"init",cmd}
	command:=exec.Command("/proc/self/exe",args...)
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
	return command
	
}
