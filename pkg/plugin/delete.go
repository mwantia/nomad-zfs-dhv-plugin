package plugin

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/mwantia/nomad-zfs-dhv-plugin/pkg/config"
	"github.com/mwantia/nomad-zfs-dhv-plugin/pkg/zfs"
)

func Delete(cfg config.DynamicHostVolumeConfig) error {
	if cfg.VolumesDir == "" {
		return fmt.Errorf("variable 'DHV_VOLUMES_DIR' must not be empty when 'DHV_CREATED_PATH' is not provided")
	}
	if cfg.VolumeID == "" {
		return fmt.Errorf("variable 'DHV_VOLUME_ID' must not be empty when 'DHV_CREATED_PATH' is not provided")
	}

	params, err := cfg.GetParams()
	if err != nil {
		log.Printf("Warning: Unable to parse parameters, using defaults: %v", err)
	}

	dataset := filepath.Join(params.Pool, "nomad", cfg.Namespace, cfg.VolumeID)
	mount := filepath.Join(cfg.VolumesDir, cfg.VolumeID)

	if err := zfs.Destroy(dataset); err != nil {
		return fmt.Errorf("failed to destroy dataset: %w", err)
	}

	if err := os.RemoveAll(mount); err != nil {
		log.Printf("Warning: Failed to remove '%s': %v", mount, err)
	}

	return nil
}
