package infra

import (
	"os"
	"os/exec"
)

func run(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func GitCommitAndTag(filePath, version, message string) error {
	if err := run("git", "add", filePath); err != nil {
		return err
	}
	if err := run("git", "commit", "-m", message); err != nil {
		return err
	}
	if err := run("git", "tag", version, "--push"); err != nil {
		return err
	}
	return nil
}
