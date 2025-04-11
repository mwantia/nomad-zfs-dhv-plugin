package zfs

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/mwantia/nomad-zfs-dhv-plugin/pkg/system"
)

func GetUsedSpace(path string) (int64, error) {
	zfs, err := system.FindPath("zfs")
	if err != nil {
		return 0, fmt.Errorf("zfs not found: %w", err)
	}

	cmd := exec.Command(zfs, "get", "-Hp", "-o", "value", "used", path)
	cmd.Stderr = os.Stderr

	buffer, err := cmd.Output()
	if err != nil {
		return 0, fmt.Errorf("failed to read used storage space: %w", err)
	}

	var used int64
	_, err = fmt.Sscanf(string(buffer), "%d", &used)
	if err != nil {
		return 0, fmt.Errorf("fdailed to parse output: %w", err)
	}

	return used, nil
}
