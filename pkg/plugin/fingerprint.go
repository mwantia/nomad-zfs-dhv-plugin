package plugin

import (
	"encoding/json"
	"fmt"

	"github.com/mwantia/nomad-zfs-dhv-plugin/pkg/config"
)

func Fingerprint(cfg config.DynamicHostVolumeConfig) error {
	resp := FingerprintResponse{
		Version: config.Version,
	}

	json, err := json.Marshal(resp)
	if err != nil {
		return fmt.Errorf("failed to marshal response: %v", err)
	}

	fmt.Print(string(json))
	return nil
}
