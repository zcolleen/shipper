package internal

import (
	"context"
	"testing"
	"time"

	"go.uber.org/mock/gomock"
)

func TestTaskManager_Run_Gomock(t *testing.T) {
	ctl := gomock.NewController(t)
	ctx := context.Background()

	task1 := NewMockTaskGomock(ctl)
	call1 := task1.EXPECT().Apply().Return()

	task2 := NewMockTaskGomock(ctl)
	call2 := task2.EXPECT().Apply().Return()

	task3 := NewMockTaskGomock(ctl)
	call3 := task3.EXPECT().Apply().Return()

	gomock.InOrder(call1, call2, call3)

	tm := NewTaskManager(time.Nanosecond)
	tm.AddTasks(task1, task2, task3)

	tm.Run(ctx)
}

func TestTaskManager_Run_Mockery(t *testing.T) {
	ctx := context.Background()

	task1 := NewMockTask(t)
	call1 := task1.EXPECT().Apply().Return()

	task2 := NewMockTask(t)
	call2 := task2.EXPECT().Apply().Return()
	call2.NotBefore(call1.Call)

	task3 := NewMockTask(t)
	call3 := task3.EXPECT().Apply().Return()
	call3.NotBefore(call2.Call)

	tm := NewTaskManager(time.Nanosecond)
	tm.AddTasks(task1, task2, task3)

	tm.Run(ctx)
}
