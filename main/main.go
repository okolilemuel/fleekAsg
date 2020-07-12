package main

import (
	"os"

	"github.com/okolilemuel/fleekAsg/filemanager"
	"github.com/okolilemuel/fleekAsg/server"
)

func main() {
	source := os.Args[1]
	destination := os.Args[2]
	go filemanager.Watch(source, destination)
	server.Run(destination)
}
