package cache

import (
	"encoding/json"
	"os"
	"runtimer/constants"
)

func Set(runtime string, cached bool) {
	cache, _ := Get()
	cache[runtime] = cached
	newCache, _ := json.MarshalIndent(cache, "", "\t")
	os.WriteFile(constants.CachedFile, newCache, os.FileMode(0644))
}