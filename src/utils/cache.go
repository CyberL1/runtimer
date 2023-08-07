package utils

import (
	"encoding/json"
	"os"
	"runtimer/constants"

	"golang.org/x/exp/slices"
)

func GetCache() []string {
	file, err := os.ReadFile(constants.CacheFile)
	var cache []string

	if err != nil {
		os.WriteFile(constants.CacheFile, []byte("[]"), 0644)
	}

	json.Unmarshal(file, &cache)
	return cache
}

func IsCached(name string) bool {
	cache := GetCache()
	return slices.Contains(cache, name)
}

func SetCache(name string, remove bool) {
	cache := GetCache()
	if remove {
		for i, v := range cache {
			if v == name {
				cache = append(cache[:i], cache[i+1:]...)
			}
		}
	} else {
		cache = append(cache, name)
	}

	newCache, _ := json.MarshalIndent(cache, "", "\t")
	os.WriteFile(constants.CacheFile, newCache, 0644)
}
