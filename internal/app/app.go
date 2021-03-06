package app

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/indigowar/blog-site/internal/config"
	"github.com/indigowar/blog-site/internal/handlers"
	"github.com/indigowar/blog-site/internal/server"
	"log"
	"os"
	"os/signal"
	"time"
)

func Run() {
	cfg, err := config.Init("./configs", os.Getenv("APP_ENV"))

	if err != nil {
		log.Panic(err)
	}

	r := gin.Default()

	handlers.Init(r.Group("/"))

	s := server.New(cfg, r)

	go func() {
		if err := s.Run(); err != nil {
			log.Println("listen error: ", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	log.Println("Shutting down the server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Stop(ctx); err != nil {
		log.Fatal("Exit with error: ", err)
	}
	log.Println("Exit.")
}
