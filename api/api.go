package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Runner() {
	log.Println("api runner called")
	router := gin.New()

	setupRouter(router)

	serveHttp(router)
}

func serveHttp(router *gin.Engine) {
	log.Println("server run's at http://127.0.0.1:8080")

	srv := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	killSignal := <-interrupt
	switch killSignal {
	case os.Interrupt:
		log.Println("Got SIGINT...")
	case syscall.SIGTERM:
		log.Println("Got SIGTERM...")
	}

	log.Println("The service is shutting down...")
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Println(err.Error())
	}
	log.Println("Done")
}
