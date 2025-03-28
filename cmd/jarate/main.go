package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/stanekondrej/jarate/internal/app/jarate"
)

var (
	updateInterval = flag.Duration("interval", time.Second, "Update interval (only applies to WS)")

	enableWebsocket   = flag.Bool("websocket", true, "Enable the websocket endpoint")
	websocketEndpoint = flag.String("ws-endpoint", "/", `The websocket endpoint path`)

	enableOneshot   = flag.Bool("oneshot", true, "Enable the oneshot HTTP endpoint")
	oneshotEndpoint = flag.String("oneshot-endpoint", "/oneshot", `The oneshot endpoint path`)

	listenAddress = flag.String("listen", ":9999", `Address and port to listen on`)
)

func main() {
	flag.Parse()

	log.Println("Jarate?! Noo!!")

	if !*enableOneshot && !*enableWebsocket {
		log.Fatal("No handler enabled, exiting")
	}

	h := jarate.GetHandlers(*updateInterval)
	if *enableWebsocket {
		http.HandleFunc(*websocketEndpoint, h.WebsocketHandler)
	}
	if *enableOneshot {
		http.HandleFunc(*oneshotEndpoint, h.OneshotHandler)
	}

	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}
