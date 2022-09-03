package main

import (
	"os"

	"github.com/tcc-uniftec-5s/internal/app"
)

func main() {
	rootdir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	app.Init(rootdir)
	ch := make(chan bool, 1)
	<-ch
}
