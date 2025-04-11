package system

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func FindPath(file string) (string, error) {
	paths := []string{
		"/sbin",
		"/usr/sbin",
		"/bin",
		"/usr/bin",
		"/usr/local/bin",
		"/usr/local/sbin",
	}

	for _, path := range paths {
		full := filepath.Join(path, file)
		log.Println(full)
		if _, err := os.Stat(full); err == nil {
			if IsExecutable(full) {
				return full, nil
			}
		}
	}

	return "", fmt.Errorf("file '%s' not found", file)
}

func IsExecutable(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}

	return !info.IsDir() && (info.Mode()&0o111 != 0)
}

func FormatBytes(size int64) string {
	const (
		KB = 1024
		MB = KB * 1024
		GB = MB * 1024
		TB = GB * 1024
	)

	switch {
	case size%TB == 0:
		return fmt.Sprintf("%dT", size/TB)
	case size%GB == 0:
		return fmt.Sprintf("%dG", size/GB)
	case size%MB == 0:
		return fmt.Sprintf("%dM", size/MB)
	case size%KB == 0:
		return fmt.Sprintf("%dK", size/KB)
	default:
		return fmt.Sprintf("%dB", size)
	}
}
