package main

import (
	"github.com/e421083458/golang_common/lib"
	"github.com/noovertime7/go-gateway/router"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := lib.InitModule("./conf/dev/", []string{"base", "mysql", "redis"})
	if err != nil {
		log.Fatal(err)
		return
	}
	defer lib.Destroy()
	router.HttpServerRun()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	router.HttpServerStop()
}
