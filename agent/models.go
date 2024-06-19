package main

type ManageRequest struct {
	Action    string `json:"action"`
	ServiceID string `json:"service_id"`
	Replicas  int    `json:"replicas"`
}

type ServiceInstance struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}
