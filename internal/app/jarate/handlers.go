package jarate

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

const BUFFER_SIZE = 1024

var upgrader = websocket.Upgrader{
	ReadBufferSize:  BUFFER_SIZE,
	WriteBufferSize: BUFFER_SIZE,
}

type handlerContainer struct {
	updateInterval time.Duration
}

func GetHandlers(updateInterval time.Duration) handlerContainer {
	return handlerContainer{
		updateInterval,
	}
}

func (h *handlerContainer) WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("New WS client")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	log.Println("Client connected")
	defer log.Println("Client disconnected")

	for {
		s, err := getStats()
		if err != nil {
			log.Println(err)
			return
		}

		err = conn.WriteJSON(s)
		if err != nil {
			log.Println(err)
			return
		}

		time.Sleep(h.updateInterval)
	}
}

func (h *handlerContainer) OneshotHandler(w http.ResponseWriter, _ *http.Request) {
	log.Println("Got oneshot request")
	defer log.Println("Finished handling oneshot request")

	s, err := getStats()
	if err != nil {
		log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte("Unable to get usage metrics")); err != nil {
			log.Println("Unable to write HTTP error response")
		}

		return
	}

	j, err := json.Marshal(s)
	if err != nil {
		log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte("Unable to stringify metrics")); err != nil {
			log.Println("Unable to write HTTP error response")
		}

		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(j); err != nil {
		log.Println(err)

		return
	}
}
