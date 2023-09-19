package cache

import (
	cacheV1 "github.com/MukanMasud1992/GoApp/internal"
)

func Version() string {
	return cacheV1.Version
}

func NewCache() *cacheV1.Cache {
	return cacheV1.NewCache()
}
