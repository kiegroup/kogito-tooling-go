package main

import (
	"flag"
	"fmt"

	"github.com/adrielparedes/kogito-local-server/pkg/config"
	"github.com/adrielparedes/kogito-local-server/pkg/kogito"
)

func main() {

	var config config.Config
	conf := config.GetConfig()
	fmt.Println(conf.Proxy.Port)
	port := flag.Int("p", conf.Proxy.Port, "DMN Runner Port")
	flag.Parse()
	kogito.Systray(*port)
}
