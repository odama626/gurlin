package main

import (
	"time"
	//"github.com/garyburd/redigo/redis"
)

// var redisPool

var datastore map[string]Redirect
var reverseLookup map[string]*Redirect
var curIttr int

func init() {

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
