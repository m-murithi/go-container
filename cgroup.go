package main

import (
    "io/ioutil"
    "os"
    "strconv"
)

func setupCgroup() {
    cgroup := "/sys/fs/cgroup/mycontainer"
    os.Mkdir(cgroup, 0755)
    ioutil.WriteFile(cgroup+"/memory.limit_in_bytes", []byte("10000000"), 0700)
    pid := strconv.Itoa(os.Getpid())
    ioutil.WriteFile(cgroup+"/cgroup.procs", []byte(pid), 0700)
}
