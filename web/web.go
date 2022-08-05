package web

import (
	"embed"
	"io/fs"
	"log"
)

//go:embed all:dist
var f embed.FS

func BuildHTTPS() fs.FS {
	build, err := fs.Sub(f, "dist")
	if err != nil {
		log.Fatal(err)
	}
	return build
}
