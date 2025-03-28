package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
)

const LISTEN_ADDRESS = ":9999"

// How often should we update the info? this doesn't tell you how long it's
// going to take to update the information; that depends on the system calls. It
// is, however, the minimum time that passes between updates.
const UPDATE_INTERVAL = time.Millisecond * 1000
const BUFFER_SIZE = 1024

var upgrader = websocket.Upgrader{
	ReadBufferSize:  BUFFER_SIZE,
	WriteBufferSize: BUFFER_SIZE,
}

type cpuStats struct {
	Overall float64   `json:"overall"`
	PerCore []float64 `json:"per_core"`
	Freq    uint      `json:"freq"`
}

func getCpuStats() (cpuStats, error) {
	perCore, err := cpu.Percent(0, false)
	if err != nil {
		return cpuStats{}, err
	}
	overall, err := cpu.Percent(0, true)
	if err != nil {
		return cpuStats{}, err
	}

	freq := cpu.ClocksPerSec / 1_000_000

	return cpuStats{
		overall[0],
		perCore,
		uint(freq),
	}, nil
}

// All of these are in bytes
type memStats struct {
	Used  uint64 `json:"used"`
	Total uint64 `json:"total"`
}

func getMemStats() (memStats, error) {
	vmem, err := mem.VirtualMemory()
	if err != nil {
		return memStats{}, err
	}

	return memStats{
		vmem.Used,
		vmem.Total,
	}, nil
}

type Stats struct {
	Cpu cpuStats `json:"cpu"`
	Mem memStats `json:"mem"`
}

func getStats() (Stats, error) {
	cpuStats, err := getCpuStats()
	if err != nil {
		return Stats{}, err
	}

	memStats, err := getMemStats()
	if err != nil {
		return Stats{}, err
	}

	return Stats{
		cpuStats,
		memStats,
	}, nil
}

func websocketEndpointHandler(w http.ResponseWriter, r *http.Request) {
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

		time.Sleep(UPDATE_INTERVAL)
	}
}

func main() {
	http.HandleFunc("/", websocketEndpointHandler)

	log.Fatal(http.ListenAndServe(LISTEN_ADDRESS, nil))
}
