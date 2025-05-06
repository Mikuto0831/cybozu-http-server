package server

import (
	"net/http"

	"github.com/mikuto0831/cybozu-http-server/internal/infra"
	"github.com/mikuto0831/cybozu-http-server/internal/pkg/logger"
)

func StartServer() {
	l := logger.NewLogger("server")

	mux := http.NewServeMux()
	router := infra.NewRouter()
	router.RegisterRoutes(mux)

	server := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	l.Info("Starting server on :8000")

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}