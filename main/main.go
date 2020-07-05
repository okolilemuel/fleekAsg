package main

import (
	"os"

	"github.com/okolilemuel/fleekAsg/encryptor"
	"github.com/okolilemuel/fleekAsg/server"
)

func main() {
	source := os.Args[1]
	destination := os.Args[2]
	go encryptor.Watcher(source, encryptor.Encryptor(destination))
	server.Run(destination)
}
