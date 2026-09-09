package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sankardabest/students-api/internal/config"
	"github.com/sankardabest/students-api/internal/http/handlers/student"
)

func main() {
	cfg := config.MustLoad()

	//////router setup////
	router := http.NewServeMux()

	router.HandleFunc("GET /api/v1/students", student.GetStudents())
	router.HandleFunc("POST /api/v1/students/create", student.Create())
	/////Server Setup//////////
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}
	slog.Info("Server Started", slog.String("Address:", cfg.Addr))
	// fmt.Printf("Server Started Successfully %s", cfg.HttpServer.Addr)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("Server start Failed %s:", err.Error())
		}
	}()

	<-done

	slog.Info("Shutting down the Server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Server failed to Shutdown", slog.String("Error:", err.Error()))
	}
	slog.Info("Server shutdown successfully")
}
