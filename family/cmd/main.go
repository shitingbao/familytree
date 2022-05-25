package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"familytree/family/conf"
	"familytree/family/server/http"
	"familytree/family/service"
)

func main() {
	flag.Parse()

	conf := conf.Init()
	svr := service.New(conf)

	http.Init(conf, svr)

	// init signal
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-sig
		log.Printf("service get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT:
			shutdown(svr)
			return
		case syscall.SIGHUP:
		//TODO reload
		default:
			return
		}
	}
}

func shutdown(sev *service.Service) {
}
