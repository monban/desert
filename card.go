package main

import (
	"fmt"
)

type card struct {
	CardType string     `json:"cardType"`
	Storm    *stormCard `json:"storm,omitempty"`
}

type stormCard struct {
	Direction string `json:"direction"`
	Distance  int    `json:"distance"`
}

func (c card) String() string {
	s := fmt.Sprintf("[%v", c.CardType)
	if c.Storm != nil {
		s = s + fmt.Sprintf(" %v %v", c.Storm.Direction, c.Storm.Distance)
	}
	s = s + "]"
	return s
}
