package main

import (
	"blog/pkg/setting"
	"blog/routers"
	"fmt"
	"net/http"

	"blog/models"
	"blog/pkg/logging"
	"context"
	"os"
	"os/signal"
	"time"
)

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
}

func main() {
	router := routers.InitRouter()

	s := http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			logging.Info("Listen", err)
		}
	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, os.Interrupt)
	<-quit

	logging.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		logging.Info("Server Shutdown:", err)
	}

	logging.Info("Server exiting")
}
