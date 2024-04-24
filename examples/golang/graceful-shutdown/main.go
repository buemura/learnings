package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    time.Sleep(3 * time.Second)
    w.Write([]byte("Hello World!"))
  })
  server := &http.Server{Addr: ":8080"}

  go func ()  {
    if err := server.ListenAndServe(); err != nil && http.ErrServerClosed != err {
      panic(err)
    } 
  }()

  stop := make(chan os.Signal)
  signal.Notify(stop, syscall.SIGTERM, os.Interrupt, syscall.SIGINT)
  <-stop

  ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
  defer cancel()
  fmt.Println("Stopping...")

  if err := server.Shutdown(ctx); err != nil {
    panic(err)
  }
  fmt.Println("Server stopped")
}

