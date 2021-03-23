package main

import (
	"flag"

	"github.com/adrielparedes/kogito-local-server/pkg/systray"
)

func main() {

	port := flag.Int("p", 0, "DMN Runner Port")
	flag.Parse()
	systray.Systray(*port)
}
