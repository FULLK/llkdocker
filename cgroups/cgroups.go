package main

import (
	"os/exec"
	"os"
	"syscall"
	"io/ioutil"
	"fmt"
	"strconv"
	"path"
	"bufio"
)

const process_cgroup_root_path="/sys/fs/cgroup/"

func main(){
	if os.Args[0]=="/proc/self/exe"{
	cmd:=exec.Command("/proc/self/exe")
	fmt.Println("args[0]",os.Args[0])
	fmt.Println("/proc/self/exe",cmd)
	fmt.Println()
	cmd=exec.Command("sh","-c","stress --vm-bytes 200m --vm-keep -m 1")
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	cmd.Stdin = os.Stdin 
	cmd.Stdout = os.Stdout 
	cmd.Stderr = os.Stderr 
	bufio.NewReader(os.Stdin).ReadString('\n')
	if err:=cmd.Run(); err!= nil{ 
	 fmt.Println(err)
	  os.Exit(1)
	}
	fmt.Println("fork finished ")
	}

	cmd:= exec.Command ("/proc/self/exe") 	
	cmd.SysProcAttr = &syscall.SysProcAttr{ 
		Cloneflags: syscall.CLONE_NEWIPC|syscall.CLONE_NEWPID|syscall.CLONE_NEWUTS|syscall.CLONE_NEWNS|syscall.CLONE_NEWUSER|syscall.CLONE_NEWUSER|syscall.CLONE_NEWNET,}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout 
	cmd.Stderr = os.Stderr
	if err:= cmd.Start( ); err !=nil { 
		fmt.Println ("error", err) 
		os.Exit(1) 
	} else { 
		
		pid:=cmd.Process.Pid  
		fmt.Println("fork process pid ",pid)
		os.Mkdir(path.Join(process_cgroup_root_path,"cgroup-test"),0755)
		fmt.Println("success mkdir")
		ioutil.WriteFile(path.Join(process_cgroup_root_path, "cgroup-test","cgroup.procs"),[]byte(strconv.Itoa(cmd.Process.Pid)), 0755)
		ioutil.WriteFile(path.Join(process_cgroup_root_path,"cgroup-test","memory.max"),[]byte(strconv.Itoa(100 * 1024 * 1024)), 0755) 
		fmt.Println("change finished ")
		cmd.Process.Wait() 
		}
}

