package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/yagmurozden/microservices/handlers"
)

func main() {
	l := log.New(os.Stdout, "test API - ", log.LstdFlags)

	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/goodbye", gh)
	sm.Handle("/", hh)

	s := http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	s.ListenAndServe()
	// go func() {
	// 	err := s.ListenAndServe()
	// 	if err != nil {
	// 		l.Fatal(err)
	// 	}
	// }()
	//it will wait the process and it will shot down the service, if we want to upgrade or something we may use it.
	//if the handkers still working of this time it will force shotdown

	// tc, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	// fmt.Println("cancel: ", cancel)
	// s.Shutdown(tc)
}
