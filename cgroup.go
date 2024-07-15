package main

import (
    "os"
    "os/exec"
    "syscall"
    "io/ioutil"
    "strconv"
    "fmt"
)

func main() {
    cgroup := "/sys/fs/cgroup/mycontainer"
    os.Mkdir(cgroup, 0755)
    ioutil.WriteFile(cgroup+"/memory.limit_in_bytes", []byte("10000000"), 0700)
    pid := strconv.Itoa(os.Getpid())
    ioutil.WriteFile(cgroup+"/cgroup.procs", []byte(pid), 0700)

    cmd := exec.Command("/proc/self/exe", "child")
    cmd.SysProcAttr = &syscall.SysProcAttr{
        Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC,
    }

    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Run(); err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
}

func child() {
    fmt.Println("Running inside the container")
    cmd := exec.Command("/bin/sh")
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
}

func init() {
    if len(os.Args) > 1 && os.Args[1] == "child" {
        child()
        os.Exit(0)
    }
}
