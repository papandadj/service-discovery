package main

import "time"

type requestBody struct {
	Name string `json:"name"`
	Host string `json:"host"`
	Port int    `json:"port"`
}

var deadline = 5 * time.Second
