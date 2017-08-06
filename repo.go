package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/garyburd/redigo/redis"
)

var connection redis.Conn

func init() {
	var err error
	connection, err = redis.Dial("tcp", "10.0.0.100:6379")
	if err != nil {
		log.Fatalf("Could not connect: %v\n", err)
	}
}

func get(src string) (Redirect, error) {
	redirect := Redirect{}
	jsonRedirect, err := redis.String(connection.Do("GET", src))
	json.Unmarshal([]byte(jsonRedirect), &redirect)
	return redirect, err
}

func put(key string, redirect Redirect) error {
	js, _ := json.Marshal(redirect)
	_, err := connection.Do("SET", key, js)
	return err
}

func incr(key string) (int, error) {
	return redis.Int(connection.Do("incr", key))
}

func exists(key string) bool {
	res, _ := redis.Bool(connection.Do("EXISTS", key))
	return res
}

func AddRedirect(r Redirect) {
	r = MakeRedirect(r)
	put(r.Src, r)
	put(r.Dest, r)
}

func GetRedirect(src string) (Redirect, bool) {
	// redirect, ok := datastore[src]
	redirect, ok := get(src)
	fmt.Printf("redirect %v -> %v\n", redirect.Src, redirect.Dest)
	return redirect, ok != nil
}

func SrcAvailable(src string) bool {
	// _, ok := datastore[src]
	return !exists(src)
}

func GetAvailableSrc(dest string) string {
	// rev, exists := reverseLookup[dest]
	if exists(dest) {
		rev, _ := get(dest)
		return rev.Src
	}

	curIttr, _ := redis.Int(connection.Do("GET", "redirectRand"))
	avail := ItoS(curIttr)
	curIttr++

	for exists(avail) {
		avail = ItoS(curIttr)
		curIttr++
	}
	connection.Do("SET", "redirectRand", curIttr)
	return avail
}
