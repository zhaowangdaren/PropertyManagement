package main

import (
	"../pm"
	"../server"
)

func main() {
	pm.Init()
	server.Start()
}
