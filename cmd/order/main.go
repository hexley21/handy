package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hexley21/handy/pkg/config"
	"github.com/hexley21/handy/pkg/logger"
	"github.com/hexley21/handy/pkg/logger/zap"
)

func main() {
	cfg, err := config.LoadConfig("./config/config.yml")
	if err != nil {
		log.Fatalf("could not load config: %v\n", err)
	}

	var zapLogger logger.Logger = zap.InitLogger(cfg.Logging, cfg.Server.IsProd)


	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
		return
	})

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Server.HTTP.Port),
		Handler:      mux,
		IdleTimeout:  cfg.Server.HTTP.IdleTimeout,
		ReadTimeout:  cfg.Server.HTTP.ReadTimeout,
		WriteTimeout: cfg.Server.HTTP.WriteTimeout,
	}
	
	srv.ListenAndServe()

	zapLogger.Debug("Server Started")
}
