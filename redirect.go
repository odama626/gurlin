package main

import (
	"fmt"
	"strings"
	"time"
)

// Redirect represents a url redirct
type Redirect struct {
	Src     string    `json:"src"`
	Dest    string    `json:"dest"`
	Active  bool      `json:"active"`
	Created time.Time `json:"created"`
}

func (r Redirect) String() string {
	return fmt.Sprintf("%v -> %v active: %v created: %v", r.Src, r.Dest, r.Active, r.Created)
}

// FixDestination prepends // when the beginning of a url isn't provided
func FixDestination(dest string) string {
	dest = strings.ToLower(dest)
	if !(strings.HasPrefix(dest, "http") || strings.HasPrefix(dest, "//")) {
		dest = "//" + dest
	}
	return dest
}

type Redirects []Redirect

func MakeRedirect(r Redirect) Redirect {
	r.Active = true
	r.Created = time.Now()
	return r
}
