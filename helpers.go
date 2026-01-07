package helpers

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GetAbsolutePath(path string) string {
	return filepath.Abs(path)
}

func GetRelativePath(path, base string) string {
	return filepath.Rel(base, path)
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	return err == nil && s.IsDir()
}

func IsFile(path string) bool {
	s, err := os.Stat(path)
	return err == nil &&!s.IsDir()
}

func GetFileExtension(path string) string {
	ext := filepath.Ext(path)
	if ext == "" {
		return ""
	}
	return ext[1:] // Remove the dot from the extension
}

func GetFileNameWithoutExtension(path string) string {
	return filepath.Base(path)
}

func GetFileExtensionWithoutDot(path string) string {
	return GetFileExtension(path)
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if strings.Contains(strings.ToLower(b), strings.ToLower(a)) {
			return true
		}
	}
	return false
}

func LogError(format string, a...interface{}) {
	log.Printf(fmt.Sprintf(format, a...))
}