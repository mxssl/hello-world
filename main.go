package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Метрия для прометеуса - общее кол-во http запросов к приложению
// Считать будем только для `/` роута, т.к
// все остальные роуты технические/инфраструктурные
var (
	httpRequestsTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of http requests",
	})
)

func main() {
	log.Println("Loading...")

	// Определяем хендлеры для ендпоинтов
	http.HandleFunc("/", helloWorldHandler)
	http.HandleFunc("/health-check", healthCheckHandler)
	http.HandleFunc("/ready-check", readyCheckHandler)
	http.Handle("/metrics", promhttp.Handler())

	// Конфигурация веб сервера
	srv := &http.Server{
		Addr:         "0.0.0.0:80",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	go func() {
		log.Fatal(srv.ListenAndServe())
	}()

	log.Println("App is ready to accept connections!")

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	log.Println("Gracefully shutting down the app...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}

// Хендлер для `/` роута
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestsTotal.Inc()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello World!\n")
}

// Хендлер для `/health-check` роута
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// Хендлер для `/ready-check` роута
func readyCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
