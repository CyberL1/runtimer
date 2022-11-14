package cache

import (
	"encoding/json"
	"os"
	"runtimer/constants"
)

func Get() (map[string]bool, error) {
	file, err := os.ReadFile(constants.CachedFile)
	cache := make(map[string]bool)
	if err != nil {
		_, dirErr := os.Stat(constants.CacheDir)
		if dirErr != nil {
			os.Mkdir(constants.CacheDir, 0755)
		}

		_, fileErr := os.Stat(constants.CachedFile)
		if fileErr != nil {
			os.WriteFile(constants.CachedFile, []byte("{}"), 0644)
		}
	}
	err = json.Unmarshal(file, &cache)
	return cache, err
}