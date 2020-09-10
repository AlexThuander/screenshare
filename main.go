package main

import (
	"flag"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	addr := flag.String("addr", ":3000", "Listen address")
	// tls := flag.Bool("tls", true, "Use TLS")
	// certFile := flag.String("certFile", "files/server.crt", "TLS cert file")
	// keyFile := flag.String("keyFile", "files/server.key", "TLS key file")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())
	//server := core.GetHttp()

	// if *tls {
	// 	log.Println("Listening on TLS:", *addr)
	// 	if err := http.ListenAndServeTLS(*addr, *certFile, *keyFile, server); err != nil {
	// 		log.Fatalln(err)
	// 	}
	// } else {
	log.Println("Listening:", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatalln(err)
	}
	// }
}
