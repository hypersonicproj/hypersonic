package filesystem

import (
	"path/filepath"
	"slices"
)

func isSupportedAudioFile(path string) bool {
	return slices.Contains([]string{".mp3", ".m4a", ".wav"}, filepath.Ext(path))
}
