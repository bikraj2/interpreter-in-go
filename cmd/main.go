package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/bikraj2/pulse/repl"
)

func main () {
  user, err :=user.Current()
  if err!=nil {
    panic(err)
  }
  fmt.Printf("Hello %s This is the pulse programming language.\n",user.Username)
  fmt.Printf("Feel free to Type in commands\n")
  repl.Start(os.Stdin, os.Stdout)
}
