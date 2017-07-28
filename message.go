package main

import (
	"fmt"
	"time"
)

type Message struct {
	Err       bool   `json:"err"`
	TimeStamp int64  `json:"time"`
	Message   string `json:"msg"`
}

func (m Message) String() string {
	return fmt.Sprintf("Error: %v Message: %v, timestamp: %v", m.Err, m.Message, m.TimeStamp)
}

// MakeMsg returns a Message with all defaults set except Message.Message
func MakeMsg(msg string) Message {
	return Message{
		false,
		time.Now().Unix(),
		msg,
	}
}

// MakeErr is the same as MakeMsg except Message.Err = true
func MakeErr(msg string) Message {
	return Message{
		true,
		time.Now().Unix(),
		msg,
	}
}
