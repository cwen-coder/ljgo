package util

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func RemoveFileByPatterns(path string, patterns []string) error {
	if strings.TrimSpace(path) == "" {
		return errors.New("RemoveFileByPatterns error: path is empty")
	}
	if len(patterns) == 0 {
		return errors.New("RemoveFileByPatterns error: patterns is empty")
	}
	for _, pattern := range patterns {
		files, err := filepath.Glob(filepath.Join(path, pattern))
		if err != nil {
			return fmt.Errorf("RemoveFileByPatterns : %v", err)
		}
		for _, path := range files {
			os.RemoveAll(path)
		}
	}

	return nil
}
