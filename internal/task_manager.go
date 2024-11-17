package internal

import (
	"context"
	"time"
)

type ShipmentSyncer struct {
}

func (s *ShipmentSyncer) Apply() {
	// some syncing work
}

type Archivator struct {
}

func (s *Archivator) Apply() {
	// some archiving work
}

type Cleaner struct {
}

func (s *Cleaner) Apply() {
	// some cleaning work
}

type TaskManager struct {
	tasks    []Task
	tickTime time.Duration
}

func NewTaskManager(tickTime time.Duration) *TaskManager {
	return &TaskManager{
		tickTime: tickTime,
	}
}

//go:generate mockgen -mock_names Task=MockTaskGomock -typed -package internal  -destination zzz_task_mock_gomock.go github.com/zcolleen/shipper/internal Task
type Task interface {
	Apply()
}

func (tm *TaskManager) AddTasks(task ...Task) {
	tm.tasks = append(tm.tasks, task...)
}

func (tm *TaskManager) Run(ctx context.Context) {
	ticker := time.NewTicker(tm.tickTime)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			for _, t := range tm.tasks {
				t.Apply()
			}
		}
	}
}

func runTasks(ctx context.Context) {
	task1 := &ShipmentSyncer{}
	task2 := &Archivator{}
	task3 := &Cleaner{}

	tm := NewTaskManager(time.Minute)
	tm.AddTasks(task1, task2, task3)

	go tm.Run(ctx)
}
