package contain

import (
	"io"
	"os"
	"os/exec"
	"path/filepath"
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
	/*
	path,err:=exec.LookPath(cmd[0]) //挂载前得到相关命令的路径会出现找不到
	if err!=nil{
		log.Fatal(err)
	}
	argv:=[]string{}
    for i:=1;i<len(cmd);i++{
		argv=append(argv,cmd[i-1])
	}*/

	log.Infof("prepare for mount ")

	mount()
	path,err:=exec.LookPath(cmd[0]) //挂载后再得到相关命令的路径等
	if err!=nil{
		log.Fatal(err)
	}
	argv:=[]string{}
    for i:=1;i<len(cmd);i++{
		argv=append(argv,cmd[i-1])
	}


	log.Infof("prepare for exec cmd ")
	log.Info(path)
	log.Info(argv)
	log.Info(os.Environ())

	if err:=syscall.Exec(path,argv,os.Environ());err!=nil{
		log.Infof("exec error")
		log.Info(err)
	}
	
}

func mount(){
	pwd,err:=os.Getwd()
	log.Info("pwd ",pwd)
	if err !=nil {
		log.Fatal(err)
	}
	log.Info("pwd ",pwd)
	
	err=syscall.Mount("","/","",syscall.MS_PRIVATE|syscall.MS_REC,"")
	log.Info("finish mount /")
	if err!=nil{
		log.Fatal(err)
	}
	
	err=pivoroot(pwd)
	log.Info("finish pivoroot /")
	if err!=nil {
		log.Fatal(err)
	}
	
	mountflags:=syscall.MS_NOEXEC|syscall.MS_NOSUID|syscall.MS_NODEV
	if err=syscall.Mount("proc","/proc","proc",uintptr(mountflags),"");err!=nil{
		log.Fatal(err)
	}
	if err=syscall.Mount("tmpfs", "/dev", "tmpfs", syscall.MS_NOSUID|syscall.MS_STRICTATIME, "mode=755");err!=nil{
		log.Fatal("mount /dev error")
	}
		//切换了根文件系统，此时busybox比较简陋，还需要自己挂载一些当前namespace的文件系统来辅助功能
}

func pivoroot(root string) error{
	log.Info("start pivoroot /")
	if err := syscall.Mount(root, root, "bind", syscall.MS_BIND|syscall.MS_REC, ""); err != nil {
		log.Info("bind error")
		log.Fatal(err)
	}
	
	old_root := filepath.Join(root, "old_root")
	log.Info(old_root)
	if err := os.Mkdir(old_root, 0777); err != nil {
		log.Info("mkdir error")
		return err
	}
	
	if err := syscall.PivotRoot(root, old_root); err != nil {
		log.Info("pivot root error")
		log.Fatal(err)
	}
	if err := syscall.Chdir("/"); err != nil {
		log.Info("chdir root error")
		log.Fatal(err)
	}
	
	old_root = filepath.Join("/", "old_root")
	log.Info(old_root)
	if err := syscall.Unmount(old_root, syscall.MNT_DETACH); err != nil {
		log.Info("unmount error")
		log.Fatal(err)
	}
	
	/*不unmount依然也能正常使用，但为了严谨还是unmount因为不需要了*/
	return os.Remove(old_root)
	/*不unmount也删除不了*/
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


