package plugin

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/mwantia/nomad-zfs-dhv-plugin/pkg/config"
	"github.com/mwantia/nomad-zfs-dhv-plugin/pkg/zfs"
)

func Create(cfg config.DynamicHostVolumeConfig) error {
	if cfg.VolumesDir == "" {
		return fmt.Errorf("variable 'DHV_VOLUMES_DIR' must not be empty")
	}
	if cfg.VolumeID == "" {
		return fmt.Errorf("variable 'DHV_VOLUME_ID' must not be empty")
	}

	params, err := cfg.GetParams()
	if err != nil {
		log.Printf("Warning: Unable to parse parameters, using defaults: %v", err)
	}

	datasetPath := filepath.Join(params.Pool, "nomad", cfg.Namespace, cfg.VolumeID)
	mountPath := filepath.Join(cfg.VolumesDir, cfg.VolumeID)

	if err := os.MkdirAll(mountPath, 0o755); err != nil {
		return fmt.Errorf("failed to create volume directory: %v", err)
	}

	// zfs create -o mountpoint=<mount> -o quota=<quota> -o recordsize=<recordsize> -o atime=<atime> -o compression=<compression> <path>
	log.Printf("Create ZFS dataset...")
	if err := zfs.CreateVolume(mountPath, datasetPath, cfg.CapacityMinBytes, *params); err != nil {
		return fmt.Errorf("failed to create volume: %w", err)
	}

	used, err := zfs.GetUsedSpace(datasetPath)
	if err != nil {
		return fmt.Errorf("failed to get used dataset storage space: %w", err)
	}

	response := VolumeCreateResponse{
		Path:  mountPath,
		Bytes: used,
	}

	out, err := json.Marshal(response)
	if err != nil {
		return fmt.Errorf("failed to marshal response: %w", err)
	}

	fmt.Print(string(out))
	return nil
}
