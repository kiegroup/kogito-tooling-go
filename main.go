package main

import (
	_ "embed"
	"flag"
	"fmt"

	"github.com/adrielparedes/kogito-local-server/pkg/config"
	"github.com/adrielparedes/kogito-local-server/pkg/kogito"
)

// Embed the jitrunner into the runner variable, to produce a self-contained binary.
//go:embed jitexecutor
var jitexecutor []byte

func main() {
	var config config.Config
	conf := config.GetConfig()
	port := flag.Int("p", conf.Proxy.Port, "DMN Runner Port")
	fmt.Printf("Running on port: %d\n", port)
	flag.Parse()
	kogito.Systray(*port, jitexecutor)
}
