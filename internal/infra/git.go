package infra

import (
	"fmt"
	"os"
	"os/exec"
	"vex/internal/utils"
)

func run(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	// cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func GitCommitAndTag(filePath, version, message string) error {
	utils.Info(fmt.Sprintf("Committing and tagging version %s with message: %s", version, message))
	if err := run("git", "add", filePath); err != nil {
		return err
	}

	utils.Info(fmt.Sprintf("Added file %s to git staging area", filePath))
	if err := run("git", "commit", "-m", message); err != nil {
		return err
	}

	utils.Info(fmt.Sprintf("Committed changes with message: %s", message))
	if err := run("git", "tag", version); err != nil {
		return err
	}

	utils.Info(fmt.Sprintf("Tagged commit with version: %s", version))
	if err := run("git", "push"); err != nil {
		return err
	}

	utils.Info("Pushed changes to remote repository")
	if err := run("git", "push", "origin", version); err != nil {
		return err
	}
	return nil
}
