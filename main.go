package main

import (
	"log"
	"net/http"
	"os"

	"github.com/yagmurozden/microservices/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/goodbye", gh)
	sm.Handle("/", hh)

	http.ListenAndServe(":9090", sm)
}
