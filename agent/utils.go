package main

import (
	"fmt"
	"log"
	"os/exec"
	"sync"
)

// Определение глобальной переменной services
var services = make(map[string]ServiceInstance)
var servicesMutex = &sync.Mutex{}

func handleManageRequest(req ManageRequest) {
	switch req.Action {
	case "start":
		startService(req.ServiceID, req.Replicas)
	case "stop":
		stopService(req.ServiceID)
	}
}

func startService(serviceID string, replicas int) {
	for i := 0; i < replicas; i++ {
		instanceID := fmt.Sprintf("%s-%d", serviceID, i)
		cmd := exec.Command("go", "run", "load_service/main.go")
		if err := cmd.Start(); err != nil {
			log.Printf("Failed to start service: %v", err)
		} else {
			servicesMutex.Lock()
			services[instanceID] = ServiceInstance{ID: instanceID, Status: "running"}
			servicesMutex.Unlock()
			log.Printf("Started service instance %s", instanceID)
		}
	}
}

func stopService(serviceID string) {
	// Реализация остановки сервиса
}
