package main

type ScaleRequest struct {
	Action    string `json:"action"`
	ServiceID string `json:"service_id"`
	Replicas  int    `json:"replicas"`
}

type ServiceStatus struct {
	ServiceID string `json:"service_id"`
	Status    string `json:"status"`
	Uptime    string `json:"uptime"`
}

type ClusterStatus struct {
	Services []ServiceStatus `json:"services"`
}
