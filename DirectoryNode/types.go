package main

import "time"

type connection struct {
	From string    `json:"from" binding:"required"`
	To   string    `json:"to" binding:"required"`
	Time time.Time `json:"time"`
}

type register struct {
	Address string `json:"address"`
}
