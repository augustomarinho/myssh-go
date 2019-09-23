package console

import (
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

func (command *Command) Ssh(args ...string) {

	cmd := exec.Command("ssh", args...)
	cmd.Stdout = os.Stdout
	err := cmd.Run()

	if err != nil {
		panic(err)
	}
}
