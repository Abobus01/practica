package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Определяем необходимые типы, если они действительно нужны в load_service
type ManageRequest struct {
	Action    string `json:"action"`
	ServiceID string `json:"service_id"`
	Replicas  int    `json:"replicas"`
}

type ServiceInstance struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

// Определяем глобальные переменные для хранения состояния сервиса
var services = make(map[string]ServiceInstance)
var startTime time.Time

func main() {
	http.HandleFunc("/status", statusHandler)

	port := "8080"
	log.Printf("Load Service is running on port %s", port)
	startTime = time.Now()
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"status": "running", "uptime": "%s"}`, time.Since(startTime))
}
