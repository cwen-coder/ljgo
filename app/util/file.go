package util

import (
	"errors"
	"fmt"
	"io"
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

func CopyFile(source, destPath string) error {
	sourceFile, err := os.Open(source)
	defer sourceFile.Close()
	if err != nil {
		return fmt.Errorf("open source: %v", err)
	}
	destFile, err := os.Create(destPath)
	defer destFile.Close()
	if err != nil {
		return fmt.Errorf("create destfile %v: %v", destPath, err)
	}
	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return fmt.Errorf("copy %v to %v: %v", source, destPath, err)
	}
	sourceInfo, err := os.Stat(source)
	if err != nil {
		return fmt.Errorf("source %v stat: %v", source, err)
	}
	err = os.Chmod(destPath, sourceInfo.Mode())
	if err != nil {
		return fmt.Errorf("chmod %v: %v", destPath, err)
	}
	return nil
}

func CopyDir(source, destPath string) error {
	sourceInfo, err := os.Stat(source)
	if err != nil {
		return fmt.Errorf("source %v stat: %v", source, err)
	}
	err = os.MkdirAll(destPath, sourceInfo.Mode())
	if err != nil {
		return fmt.Errorf("Mkdir %v: %v", destPath, err)
	}
	directory, err := os.Open(source)
	defer directory.Close()
	if err != nil {
		return fmt.Errorf("open %v: %v", source, err)
	}
	objects, err := directory.Readdir(-1)
	if err != nil {
		return fmt.Errorf("read dir %v: %v", source, err)
	}
	for _, obj := range objects {
		sourceObjectPath := filepath.Join(source, obj.Name())
		destObjectPath := filepath.Join(destPath, obj.Name())
		if obj.IsDir() {
			err = CopyDir(sourceObjectPath, destObjectPath)
			if err != nil {
				return err
			}
		} else {
			err = CopyFile(sourceObjectPath, destObjectPath)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
