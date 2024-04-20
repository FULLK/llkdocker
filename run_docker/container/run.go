
package contain
import(
	log "github.com/sirupsen/logrus"
	"os"
)
func Contain_run(cmd string,it bool) error{
	command:=new_contain_process(cmd ,it)
	log.Infof("cmd %s it %d",cmd,it)
	
	
	if err:=command.Start();err!=nil{
		log.Error(err)
	}
	
	command.Wait()
	// 从容器内的命令行中的退出才会wait结束
	log.Infof("exit -1 !!!")
	os.Exit(-1) 
	return nil
}