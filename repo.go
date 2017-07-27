package main

import (
	"encoding/base64"
	"strings"
	"time"
)

// var datastore *diskv.Diskv

var datastore map[string]Redirect
var reverseLookup map[string]*Redirect

func init() {
	// flatTransform := func(s string) []string { return []string{} }

	// datastore = diskv.New(diskv.Options{
	// 	BasePath:     "data",
	// 	Transform:    flatTransform,
	// 	CacheSizeMax: 1024 * 1024,
	// })
	datastore = make(map[string]Redirect)
	reverseLookup = make(map[string]*Redirect)
}

func AddRedirect(r Redirect) {
	r.Created = time.Now()
	datastore[r.Src] = r
	reverseLookup[r.Dest] = &r
}

func GetRedirect(src string) (Redirect, bool) {
	redirect, ok := datastore[src]
	return redirect, !ok
}

func SrcAvailable(src string) bool {
	_, ok := datastore[src]
	return !ok
}

func GetAvailableSrc(dest string) string {
	rev, exists := reverseLookup[dest]
	if exists {
		return rev.Src
	}

	src := base64.StdEncoding.EncodeToString([]byte(dest))
	src = strings.Replace(src, "=", "", -1)[:3]
	for i := 0; !SrcAvailable(src) || i > 10; i++ {
		src = base64.StdEncoding.EncodeToString([]byte(src))
		src = strings.Replace(src, "=", "", -1)[:3]
	}
	return src
}
