package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os/exec"
	"time"

	"github.com/gorilla/mux"
)

type Proxy struct {
	srv *http.Server
	cmd *exec.Cmd
}

func (p *Proxy) Start() {

	target, err := url.Parse("http://localhost:8080")
	p.cmd = exec.Command("java", "-jar", "./runner/jitexecutor-runner-2.0.0-SNAPSHOT-runner.jar")

	go startRunner(p.cmd)

	if err != nil {
		log.Fatal(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	r := mux.NewRouter()
	r.HandleFunc("/", proxyHandler(proxy, p.cmd))
	r.HandleFunc("/ping", pingHandler)

	p.srv = &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go p.srv.ListenAndServe()
}

func (p *Proxy) Stop() {
	log.Println("Shutting down")

	stopRunner(p.cmd)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	if err := p.srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Println("Shutdown complete")
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
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func stopRunner(cmd *exec.Cmd) {
	cmd.Process.Kill()
}
