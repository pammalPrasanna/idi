package ports

import (
	"context"

	"with-alias/internal/dtos"
)

type ISchedulerRepository interface {
	GetScheduler(ctx context.Context, arg *dtos.GetSchedulerParams) (scheduler *dtos.Scheduler, err error)
	FindScheduler(ctx context.Context, arg *dtos.FindSchedulerParams) (scheduler []*dtos.Scheduler, err error)
	CreateScheduler(ctx context.Context, arg *dtos.CreateSchedulerParams) (id int64, err error)
	UpdateScheduler(ctx context.Context, arg *dtos.UpdateSchedulerParams) error
	DeleteScheduler(ctx context.Context, arg *dtos.DeleteSchedulerParams) error
}
