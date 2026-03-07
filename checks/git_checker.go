package checks

import (
	"context"
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func Checkgit(ctx context.Context) (string, error) {
	path, err := exec.LookPath("git")
	if err != nil {
		return "", errors.New("git not found in PATH")
	}
	cmd, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	cctx := exec.CommandContext(cmd, path, "--version")
	fullCmd, err := cctx.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("git found at %s but `git --version` failed: %w | output: %s", path, err, strings.TrimSpace(string(fullCmd)))
	}
	return strings.TrimSpace(string(fullCmd)), nil
}

func Fixgit(ctx context.Context, dryRun bool) (string, error) {
	cmdargs := []string{"dnf", "install", "-y", "git"}
	if dryRun {
		return fmt.Sprintf("DRY-RUN: %v", cmdargs), nil
	}
	cctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	cmd := exec.CommandContext(cctx, "sudo", cmdargs...)
	fullcmd, err := cmd.CombinedOutput()
	if err != nil {
		return string(fullcmd), fmt.Errorf("failed to install git: %w | output: %s", err, strings.TrimSpace(string(fullcmd)))
	}
	return strings.TrimSpace(string(fullcmd)), nil
}
