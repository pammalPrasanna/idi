package rest

import "with-alias/internal/dtos"

type (
	FindSchedulerResponse struct {
		Scheduler []*dtos.Scheduler `json:"scheduler"`
	}
	CreateSchedulerResponse struct {
		SchedulerID int64 `json:"Scheduler_id"`
	}
	GetSchedulerResponse struct {
		Scheduler *dtos.Scheduler `json:"scheduler"`
	}
)
