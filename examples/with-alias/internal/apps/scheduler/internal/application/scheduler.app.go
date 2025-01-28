package application

import (
	"context"

	"with-alias/internal/apps/scheduler/internal/application/domain"
	"with-alias/internal/apps/scheduler/internal/application/ports"
	"with-alias/internal/dtos"
	"with-alias/internal/lib"
)

type (
	Scheduler struct {
		RootApp lib.IApp

		Repo ports.ISchedulerRepository
	}

	IScheduler interface {
		GetScheduler(ctx context.Context, arg *dtos.GetSchedulerParams) (todo *dtos.Scheduler, err error)
		FindScheduler(ctx context.Context, arg *dtos.FindSchedulerParams) (scheduler []*dtos.Scheduler, err error)
		CreateScheduler(ctx context.Context, arg *dtos.CreateSchedulerParams) (id int64, err error)
		UpdateScheduler(ctx context.Context, arg *dtos.UpdateSchedulerParams) error
		DeleteScheduler(ctx context.Context, arg *dtos.DeleteSchedulerParams) error
	}
)

var _ IScheduler = (*Scheduler)(nil)

func New(rootApp lib.IApp, db ports.ISchedulerRepository) *Scheduler {
	return &Scheduler{
		RootApp: rootApp,
		Repo:    db,
	}
}

func (t *Scheduler) GetScheduler(ctx context.Context, arg *dtos.GetSchedulerParams) (scheduler *dtos.Scheduler, err error) {
	return t.Repo.GetScheduler(ctx, arg)
}

func (t *Scheduler) FindScheduler(ctx context.Context, arg *dtos.FindSchedulerParams) (scheduler []*dtos.Scheduler, err error) {
	return t.Repo.FindScheduler(ctx, arg)
}

func (t *Scheduler) CreateScheduler(ctx context.Context, arg *dtos.CreateSchedulerParams) (id int64, err error) {
	_, err = domain.NewScheduler(arg.Task, arg.Description, arg.DueDate)
	if err != nil {
		return 0, err
	}
	return t.Repo.CreateScheduler(ctx, arg)
}

func (t *Scheduler) UpdateScheduler(ctx context.Context, arg *dtos.UpdateSchedulerParams) error {
	_, err := domain.NewScheduler(arg.Task, arg.Description, arg.DueDate)
	if err != nil {
		return err
	}
	return t.Repo.UpdateScheduler(ctx, arg)
}

func (t *Scheduler) DeleteScheduler(ctx context.Context, arg *dtos.DeleteSchedulerParams) error {
	return t.Repo.DeleteScheduler(ctx, arg)
}
