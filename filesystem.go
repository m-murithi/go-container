package main

import (
    "os"
    "os/exec"
    "syscall"
    "fmt"
)

func main() {
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
    syscall.Mount("", "/", "", syscall.MS_PRIVATE|syscall.MS_REC, "")
    syscall.Mount("rootfs", "rootfs", "", syscall.MS_BIND, "")
    os.Mkdir("rootfs/oldrootfs", 0700)
    syscall.PivotRoot("rootfs", "rootfs/oldrootfs")
    os.Chdir("/")

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
