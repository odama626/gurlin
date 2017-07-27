package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Bounce(w http.ResponseWriter, r *http.Request) {
	src := mux.Vars(r)["src"]
	redirect, err := GetRedirect(src)
	if err {
		fmt.Fprintf(w, "Not found")
	} else {
		http.Redirect(w, r, redirect.Dest, 301)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func Available(w http.ResponseWriter, r *http.Request) {
	src := mux.Vars(r)["src"]
	err := SrcAvailable(src)
	msg := "available"
	if !err {
		msg = "taken"
	}

	if err := SendResponse(w, MakeMsg(fmt.Sprintf("%v %v", src, msg))); err != nil {
		panic(err)
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	form := r.Form
	to := form.Get("to")
	from := form.Get("from")
	var msg Message

	switch {
	case err != nil:
		msg = MakeErr("Unknown error")
	case to == "":
		msg = MakeErr("Destination address required, did you forget to add 'to:'?")
	case from != "":
		if !SrcAvailable(from) {
			msg = MakeErr(fmt.Sprintf("%v taken", from))
		} else {
			to = FixDestination(to)
			defer AddRedirect(Redirect{Src: from, Dest: to})
			msg = MakeMsg(fmt.Sprintf("%v -> %v", from, to))
		}
	default:
		from = GetAvailableSrc(to)
		defer AddRedirect(Redirect{Src: from, Dest: to})
		msg = MakeMsg(fmt.Sprintf("%v -> %v", from, to))
	}
	if err = SendResponse(w, msg); err != nil {
		panic(err)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
}

func Disable(w http.ResponseWriter, r *http.Request) {
}

func SendResponse(w http.ResponseWriter, msg interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(msg)
}
