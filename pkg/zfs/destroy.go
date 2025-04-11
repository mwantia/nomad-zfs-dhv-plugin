package zfs

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/mwantia/nomad-zfs-dhv-plugin/pkg/system"
)

func Destroy(path string) error {
	zfs, err := system.FindPath("zfs")
	if err != nil {
		return fmt.Errorf("zfs not found: %w", err)
	}

	cmd := exec.Command(zfs, "destroy", "-f", path)
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to destroy zfs dataset: %w", err)
	}

	return nil
}
