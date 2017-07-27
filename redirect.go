package main

import (
	"fmt"
	"strings"
	"time"
)

// Valid URI chars = ALPHA / DIGIT / "-" / "." / "_" / "~"

type Redirect struct {
	Src     string    `json:"src"`
	Dest    string    `json:"dest"`
	Active  bool      `json:"active"`
	Created time.Time `json:"created"`
}

func (r Redirect) String() string {
	return fmt.Sprintf("%v -> %v active: %v created: %v", r.Src, r.Dest, r.Active, r.Created)
}

func FixDestination(dest string) string {
	dest = strings.ToLower(dest)
	if !(strings.HasPrefix(dest, "http") || strings.HasPrefix(dest, "//")) {
		dest = "//" + dest
	}
	return dest
}

type Redirects []Redirect
