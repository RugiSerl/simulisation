package assets

import (
	"path"
	"runtime"
)

func AssetPath(assetPath string) string {
	if runtime.GOOS == "android" {
		return assetPath
	} else {
		return path.Join("android/assets/", assetPath)
	}
}
