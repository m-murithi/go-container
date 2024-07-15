package main

import (
    "fmt"
    "os"
    "os/exec"
    "syscall"
)

func child() {
    setupCgroup()
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
