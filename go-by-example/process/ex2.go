package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	//c, err := exec.LookPath("ls")
	c, err := exec.LookPath("date")
	if err != nil {
		log.Println(err.Error())
	}
	//log.Println(c)

	//p, err := os.Executable()
	//log.Println(p, err)

	args := []string{""}

	//args := []string{"ls", "-lah"}

	//exec.Command(c, os.Environ()..., )
	err = syscall.Exec(c, args, os.Environ())
	if err != nil {
		log.Println(err.Error())
	}

}
