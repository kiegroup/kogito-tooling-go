package kogito

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os/exec"
	"strconv"
	"time"

	"github.com/adrielparedes/kogito-local-server/pkg/config"
	"github.com/adrielparedes/kogito-local-server/pkg/utils"
	"github.com/gorilla/mux"
	"github.com/phayes/freeport"
)

type Proxy struct {
	view       *KogitoSystray
	srv        *http.Server
	cmd        *exec.Cmd
	Started    bool
	URL        string
	Port       int
	RunnerPort int
}

func (self *Proxy) New() {
	self.Started = false
}

func (self *Proxy) Start() {

	var config config.Config
	conf := config.GetConfig()

	self.RunnerPort = getFreePort()
	runnerPort := strconv.Itoa(self.RunnerPort)
	self.URL = "http://127.0.0.1:" + runnerPort
	target, err := url.Parse(self.URL)

	self.cmd = exec.Command("java", "-Dquarkus.http.port="+runnerPort, "-jar", utils.GetBaseDir()+"/"+conf.Runner.Location)
	stdout, _ := self.cmd.StdoutPipe()

	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			msg := scanner.Text()
			fmt.Printf("msg: %s\n", msg)
		}
	}()

	go startRunner(self.cmd)

	if err != nil {
		log.Fatal(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	r := mux.NewRouter()
	r.HandleFunc("/ping", pingHandler)
	r.PathPrefix("/").HandlerFunc(proxyHandler(proxy, self.cmd))

	addr := conf.Proxy.IP + ":" + strconv.Itoa(self.Port)

	self.srv = &http.Server{
		Handler:      r,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Server started: %s \n", addr)

	go self.srv.ListenAndServe()

	self.Refresh()
}

func (self *Proxy) Stop() {
	log.Println("Shutting down")

	stopRunner(self.cmd)

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*15)
	defer cancel()

	if err := self.srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Println("Shutdown complete")
	self.RunnerPort = 0
	self.Refresh()
}

func (self *Proxy) Refresh() {
	started := false
	countDown := 5
	retry := true

	for countDown > 0 && retry {
		resp, err := http.Get(self.URL)
		if err != nil {
			fmt.Println(err.Error())
			retry = true
			countDown--
		} else {
			fmt.Println(strconv.Itoa(resp.StatusCode) + " -> " + resp.Status)
			if resp.StatusCode == 200 {
				started = true
			}
			retry = false
		}
		time.Sleep(1 * time.Second)
	}

	self.Started = started
	self.view.Refresh()
}

func proxyHandler(proxy *httputil.ReverseProxy, cmd *exec.Cmd) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Host = r.URL.Host
		proxy.ServeHTTP(w, r)
	}
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	result := map[string]string{"status": "ok"}
	json, _ := json.Marshal(result)
	w.Write(json)
}

func startRunner(cmd *exec.Cmd) {
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
}

func stopRunner(cmd *exec.Cmd) {
	cmd.Process.Kill()
}

func getFreePort() int {
	port, err := freeport.GetFreePort()
	if err != nil {
		log.Fatal(err)
	}
	return port
}
