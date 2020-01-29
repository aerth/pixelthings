package main

import (
	"flag"
	"log"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

var version string
var name = "example"

func main() {
	log.SetFlags(0)
	showversion := flag.Bool("version", false, "show version and exit")
	flag.Parse()
	if *showversion {
		log.Println(name, version)
		return
	}

	cfg := pixelgl.WindowConfig{
		Bounds:    pixel.R(0, 0, 500, 500),
		VSync:     true,
		Resizable: true,
	}
	pixelgl.Run(NewApp(cfg).run)
}
