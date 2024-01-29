package example

import (
	"context"
	"encoding/json"

	"github.com/hibiken/asynq"
	"github.com/neoxelox/errors"
	"github.com/neoxelox/kit"

	"service/pkg/config"
)

const (
	ExampleTasksMakeOnboarding = "example:onboarding"
	ExampleTasksReconcile      = "example:reconcile"
)

var (
	ErrExampleTasksGeneric = errors.New("example task failed")
)

type ExampleTasks struct {
	config            config.Config
	observer          *kit.Observer
	exampleRepository *ExampleRepository
}

func NewExampleTasks(observer *kit.Observer, exampleRepository *ExampleRepository, config config.Config) *ExampleTasks {
	return &ExampleTasks{
		config:            config,
		observer:          observer,
		exampleRepository: exampleRepository,
	}
}

type ExampleTasksMakeOnboardingParams struct {
	ID string
}

func (self *ExampleTasks) MakeOnboarding(ctx context.Context, task *asynq.Task) error {
	params := ExampleTasksMakeOnboardingParams{}

	err := json.Unmarshal(task.Payload(), &params)
	if err != nil {
		// Fail silently to not retry the task as the params were wrongly serialized
		self.observer.Error(ctx, kit.ErrWorkerGeneric.Raise().Cause(err))
		return nil
	}

	example, err := self.exampleRepository.Get(ctx, params.ID)
	if err != nil {
		return ErrExampleTasksGeneric.Raise().Cause(err)
	}

	if example == nil {
		// Fail silently to not retry the task as the example ID was not found
		self.observer.Error(ctx, kit.ErrWorkerGeneric.Raise().
			With("example id not found").Extra(map[string]any{"id": params.ID}))
		return nil
	}

	self.observer.Infof(ctx, "onboarding %s", example.Name)

	return nil
}

func (self *ExampleTasks) Reconcile(ctx context.Context, task *asynq.Task) error {
	self.observer.Info(ctx, "reconciling examples")

	return nil
}
