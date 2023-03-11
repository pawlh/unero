package main

import "github.com/pawlh/unero/internal/server"

func main() {
	server.NewServer(8080).Start()
}
