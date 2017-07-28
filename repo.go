package main

import (
	"time"
)

// var datastore *diskv.Diskv

var datastore map[string]Redirect
var reverseLookup map[string]*Redirect
var curIttr int

func init() {
	// flatTransform := func(s string) []string { return []string{} }

	// datastore = diskv.New(diskv.Options{
	// 	BasePath:     "data",
	// 	Transform:    flatTransform,
	// 	CacheSizeMax: 1024 * 1024,
	// })
	datastore = make(map[string]Redirect)
	reverseLookup = make(map[string]*Redirect)
	curIttr = 0
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
	avail := ItoS(curIttr)
	curIttr++
	return avail
}
