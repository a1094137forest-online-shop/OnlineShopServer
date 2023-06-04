package main

import (
	"OnlineShopServer/config"
	"OnlineShopServer/constant"
	"OnlineShopServer/router"
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

func init() {
	constant.ReadConfig(".env")
}

func main() {
	gin.SetMode(config.RunMode)

	port := config.Port

	routerInit := router.InitRouter()

	server := &http.Server{
		Addr:           port,
		Handler:        routerInit,
		ReadTimeout:    time.Duration(config.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(config.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGHUP, syscall.SIGINT,
		syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2, syscall.SIGILL, syscall.SIGFPE)
	defer stop()
	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		log.Println("[info] start http server listening", port)
		return server.ListenAndServe()
	})

	g.Go(func() error {
		<-gCtx.Done()

		c, cancel := context.WithTimeout(context.Background(), time.Duration(config.ShutdownTimeout)*time.Second)
		defer cancel()

		err := server.Shutdown(c)
		if err != nil {
			log.Println("[error] http server shutdown:", err)
		}
		return err
	})

	if err := g.Wait(); err != nil {
		log.Println("[error]", err)
	}
	log.Println("[info] Server Exit")
}
