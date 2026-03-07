package checks

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

var (
	Keypath string
)

func Checkssh(ctx context.Context) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Cannot find home directory:", err)
		return "", nil
	}

	keyNames := []string{"id_rsa", "id_ed25519", "id_ecdsa"}
	for _, v := range keyNames {
		path := filepath.Join(home, ".ssh", v)
		if _, err := os.Stat(path); err == nil {
			Keypath = path
			break
		}
	}
	if Keypath == "" {
		return "", fmt.Errorf("no SSH key found in %s", filepath.Join(home, ".ssh"))
	}
	ctxx, cancel := context.WithTimeout(ctx, 100*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctxx, "ssh", "-o", "StrictHostKeyChecking=accept-new", "-T", "git@github.com")
	output, err := cmd.CombinedOutput()
	if err != nil {
		if strings.Contains(string(output), "successfully authenticated") {
			return string(output), nil
		}
		if strings.Contains(string(output), "Permission denied") {
			pubKeyPath := Keypath + ".pub"
			pubKey, _ := os.ReadFile(pubKeyPath)
			return "", fmt.Errorf("key exists but not authorized on GitHub.\nAdd this public key:\n%s\nSee: https://github.com/settings/keys", pubKey)
		}
		return "", fmt.Errorf("SSH connection test failed: %w\n%s", err, output)
	}
	return strings.TrimSpace(string(output)), nil
}

func Fixssh(ctx context.Context, dryrun bool) (string, error) {
	cmdarg := []string{"-o", "StrictHostKeyChecking=accept-new", "-T", "git@github.com"}
	if dryrun {
		return fmt.Sprintf("Dry Run: %v", cmdarg), nil
	}
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("cannot get home directory: %w", err)
	}
	sshDir := filepath.Join(homedir, ".ssh")
	if err := os.MkdirAll(sshDir, 0700); err != nil {
		return "", fmt.Errorf("failed to create .ssh directory: %w", err)
	}
	keyPath := filepath.Join(sshDir, "id_ed25519")
	if _, err := os.Stat(keyPath); os.IsNotExist(err) {
		cmdGen := exec.Command("ssh-keygen", "-t", "ed25519", "-f", keyPath, "-N", "")
		if out, err := cmdGen.CombinedOutput(); err != nil {
			return "", fmt.Errorf("keygen failed: %w\n%s", err, out)
		}
	} else if err != nil {
		return "", fmt.Errorf("cannot check key %s: %w", keyPath, err)
	}

	ctxTest, cancel := context.WithTimeout(ctx, 100*time.Second)
	defer cancel()
	cmdTest := exec.CommandContext(ctxTest, "ssh", cmdarg...)
	output, err := cmdTest.CombinedOutput()
	if err != nil {
		if strings.Contains(string(output), "successfully authenticated") {
			return string(output), nil
		}
		if strings.Contains(string(output), "Permission denied") {
			pubKeyPath := Keypath + ".pub" // was keyPath
			pubKey, _ := os.ReadFile(pubKeyPath)
			return "", fmt.Errorf("key exists but not authorized on GitHub.\nAdd this public key:\n%s\nSee: https://github.com/settings/keys", pubKey)
		}
		return "", fmt.Errorf("SSH connection test failed: %w\n%s", err, output)
	}
	return fmt.Sprintf("SSH key generated at %s", keyPath), nil
}
