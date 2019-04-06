package main

import (
  "fmt"
  "os"
  "os/exec"
)

func main() {
  var (
    err error
    cmdOut []byte
  )

  cmdOut,err = exec.Command("which", "psql").Output()

  fmt.Println(string(cmdOut))

  err = exec.Command("sudo","-u","postgres","psql").Run()

  if(err == nil) {
    cmdOut,err = exec.Command("")
  }
}
