package imgtool

import (
	"path/filepath"
	"slices"
	"strings"
)

func IsImageFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	supportedExts := []string{".jpg", ".jpeg", ".png", ".webp", ".tiff", ".avif"}

	return slices.Contains(supportedExts, ext)
}

func GetOutputFilename(inputPath, targetExt string) string {
	base := strings.TrimSuffix(inputPath, filepath.Ext(inputPath))
	return base + targetExt
}

func ValidateExtension(ext string) bool {
	if !strings.HasPrefix(ext, ".") {
		return false
	}
	return IsImageFile("dummy" + ext)
}
