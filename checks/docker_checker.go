package checks

import (
	"context"
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func Checkdocker(ctx context.Context) (string, error) {
	path, err := exec.LookPath("docker")
	if err != nil {
		return "", errors.New("Docker not found in PATH")
	}
	cctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	cmd := exec.CommandContext(cctx, path, "--version")
	fullcmd, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("git found at %s but `git --version` failed: %w | output: %s", path, err, strings.TrimSpace(string(fullcmd)))
	}
	return strings.TrimSpace(string(fullcmd)), nil
}

func Fixdocker(ctx context.Context, dryrun bool) (string, error) {
	cmdargs := []string{"dnf", "install", "docker", "-y"}
	if dryrun {
		return fmt.Sprintf("DRY-RUN: %v", cmdargs), nil
	}
	cctx, cancel := context.WithTimeout(ctx, 3*time.Second)

	defer cancel()
	cmd := exec.CommandContext(cctx, "sudo", cmdargs...)
	fullcmd, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to install docker:%w | output:%s", err, strings.TrimSpace(string(fullcmd)))
	}
	return strings.TrimSpace(string(fullcmd)), nil
}
