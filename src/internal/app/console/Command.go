package console

import (
	"fmt"
	"os"
	"os/exec"
)

type Command struct {
}

func NewCommand() *Command {
	command := new(Command)
	return command
}

func (command *Command) Clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (command *Command) Ssh(host string) {
	fmt.Println("Implements")
}
