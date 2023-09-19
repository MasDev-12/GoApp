package main

import (
	"fmt"

	cacheV1 "github.com/MukanMasud1992/GoApp/cache"
)

func main() {
	println("v1 version:", cacheV1.Version())
	v1 := cacheV1.NewCache()
	v1.Set("userId", 42)
	item, err := v1.Get("userId")
	if err != nil {
		println("error: ", err.Error())
	}
	if intValue, ok := item.(int); ok {
		println("got item: ", intValue)
	} else {
		fmt.Println("Value is not an int")
	}
}
