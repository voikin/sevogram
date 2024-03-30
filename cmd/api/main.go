package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	httpv1 "github.com/voikin/sevogram/internal/controller/http/v1"
	"github.com/voikin/sevogram/internal/httpserver"
)

func main() {
	_, cancel := context.WithCancel(context.Background())

	ginEngine := gin.New()
	httpv1.New(ginEngine)

	httpServer := httpserver.New(ginEngine)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case <-interrupt:
		fmt.Println("Завершение работы по сингналу SIGTERM")

	case <-httpServer.Notify():
		fmt.Println("Завершение работы по желанию http-сервера")
	}
	cancel()

	_ = httpServer.Shutdown()
}
