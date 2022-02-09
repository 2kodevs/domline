package utils

import (
	"os"
	"os/exec"
)

func WGet(url, filepath string) error {
	// run shell `wget URL` to download a directory
	cmd := exec.Command("wget", url, "-np", "–level=0", "-erobots=off", "-O", filepath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
