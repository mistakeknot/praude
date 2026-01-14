package agents

import (
	"fmt"
	"os/exec"
)

type Profile struct {
	Command string
	Args    []string
}

func Resolve(cfg map[string]Profile, name string) (Profile, error) {
	p, ok := cfg[name]
	if !ok {
		return Profile{}, fmt.Errorf("unknown agent %s", name)
	}
	return p, nil
}

func Launch(p Profile, briefPath string) error {
	if _, err := exec.LookPath(p.Command); err != nil {
		return err
	}
	args := append([]string{}, p.Args...)
	args = append(args, briefPath)
	cmd := exec.Command(p.Command, args...)
	return cmd.Start()
}
