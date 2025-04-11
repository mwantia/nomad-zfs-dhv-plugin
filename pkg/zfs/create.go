package zfs

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/mwantia/nomad-zfs-dhv-plugin/pkg/config"
	"github.com/mwantia/nomad-zfs-dhv-plugin/pkg/system"
)

func CreateVolume(mount, path, quota string, params config.DynamicHostVolumeParameters) error {
	args := []string{
		"create",
		"-o", fmt.Sprintf("mountpoint=%s", mount),
	}

	if quota != "" {
		args = append(args, "-o", fmt.Sprintf("quota=%s", quota))
	}

	if params.RecordSize != "" {
		args = append(args, "-o", fmt.Sprintf("recordsize=%s", params.RecordSize))
	}
	if params.Atime != "" {
		args = append(args, "-o", fmt.Sprintf("atime=%s", params.Atime))
	}
	if params.Compression != "" {
		args = append(args, "-o", fmt.Sprintf("compression=%s", params.Compression))
	}

	args = append(args, path)

	zfs, err := system.FindPath("zfs")
	if err != nil {
		return fmt.Errorf("zfs not found: %w", err)
	}

	cmd := exec.Command(zfs, args...)
	cmd.Stderr = os.Stderr

	log.Printf("Creating zfs dataset with command: zfs %s", strings.Join(args, " "))

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to create zfs dataset: %w", err)
	}

	return nil
}
