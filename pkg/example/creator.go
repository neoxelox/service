package example

import (
	"context"

	"github.com/neoxelox/errors"
	"github.com/neoxelox/kit"
	"github.com/rs/xid"

	"service/pkg/config"
)

var (
	ErrExampleCreatorMinor = errors.New("example age (%d) does not reach minimum age (%d)")
)

type ExampleCreator struct {
	config            config.Config
	observer          *kit.Observer
	exampleService    *ExampleService
	exampleRepository *ExampleRepository
	enqueuer          *kit.Enqueuer
}

func NewExampleCreator(observer *kit.Observer, exampleService *ExampleService, exampleRepository *ExampleRepository,
	enqueuer *kit.Enqueuer, config config.Config) *ExampleCreator {
	return &ExampleCreator{
		config:            config,
		observer:          observer,
		exampleService:    exampleService,
		exampleRepository: exampleRepository,
		enqueuer:          enqueuer,
	}
}

type ExampleCreatorCreateParams struct {
	Name     string
	Age      int
	Role     string
	Settings *ExampleSettings
}

func (self *ExampleCreator) Create(ctx context.Context, params ExampleCreatorCreateParams) (*Example, error) {
	if params.Age < EXAMPLE_MINIMUM_AGE {
		return nil, ErrExampleCreatorMinor.Raise(params.Age, EXAMPLE_MINIMUM_AGE)
	}

	result, err := self.exampleService.GetCountry(ctx, ExampleServiceGetCountryParams{
		Name: params.Name,
	})
	if err != nil {
		return nil, err
	}

	settings := ExampleSettings{
		WantsNewsletter: false,
	}
	if params.Settings != nil {
		settings = *params.Settings
	}

	example := NewExample()
	example.ID = xid.New().String()
	example.Name = params.Name
	example.Age = params.Age
	example.Country = result.Country
	example.Role = params.Role
	example.Settings = settings
	example.DeletedAt = nil

	example, err = self.exampleRepository.Create(ctx, *example)
	if err != nil {
		return nil, err
	}

	err = self.enqueuer.Enqueue(ctx, ExampleTasksMakeOnboarding, ExampleTasksMakeOnboardingParams{
		ID: example.ID,
	})
	if err != nil {
		self.observer.Error(ctx, err)
	}

	return example, nil
}
