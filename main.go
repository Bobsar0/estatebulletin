package main

import (
	"fmt"
	"log"

	"os"
	"os/signal"
	"syscall"

	"github.com/bobsar0/estatebulletin/http"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	port := os.Getenv("PORT")
	appHandler := http.NewAppHandler()
	httpHandler := http.NewHttpHandler(appHandler)
	server := http.NewServer(port, httpHandler)
	if err := server.Open(done, sigs); err != nil {
		log.Fatalf("Unable to Open Server for listen and serve: %v", err)
	}
	fmt.Println("Listening on: ", server.Port())
	<-done
	fmt.Println("Good Bye")
}
