package captain // import "github.com/indigobio/captain"

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
)

func execute(name string, arg ...string) error {
	// Construct command for debug purposes
	var command = name
	for _, i := range arg {
		command += " " + i
	}

	debug("Executing %s", command)
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

func oneliner(name string, arg ...string) (string, error) {
	var buff bytes.Buffer
	gitCmd := exec.Command(name, arg...)
	gitCmd.Stdout = &buff
	err := gitCmd.Run()
	return strings.TrimSpace(buff.String()), err
}
