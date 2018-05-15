package main

import (
	"flag"
	"log"

	"github.com/c-bata/rtmp"
)

func main() {
	var addr string
	flag.StringVar(&addr, "addr", ":1935", `proxy local address`)
	flag.Parse()

	log.Printf("Serving RTMP on %s", addr)
	err := rtmp.ListenAndServe(addr)
	if err != nil {
		log.Fatalf("Catch Error: %s", err)
	}
}