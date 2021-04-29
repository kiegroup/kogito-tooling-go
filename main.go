package main

import (
	_ "embed"
	"flag"
	"fmt"

	"github.com/adrielparedes/kogito-local-server/pkg/config"
	"github.com/adrielparedes/kogito-local-server/pkg/kogito"
)

// Embed the jitrunner into the runner variable, to produce a self-contained binary.
//go:embed kogito-apps/jitexecutor/jitexecutor-runner/target/jitexecutor-runner-2.0.0-SNAPSHOT-runner
var runner []byte

func main() {

	var config config.Config
	conf := config.GetConfig()
	fmt.Println(conf.Proxy.Port)
	port := flag.Int("p", conf.Proxy.Port, "DMN Runner Port")
	flag.Parse()
	kogito.Systray(*port, runner)
}
