package example

import (
	"context"

	"github.com/neoxelox/kit"

	"service/pkg/config"
)

type ExampleDeleter struct {
	config            config.Config
	observer          *kit.Observer
	database          *kit.Database
	exampleRepository *ExampleRepository
}

func NewExampleDeleter(observer *kit.Observer, database *kit.Database, exampleRepository *ExampleRepository,
	config config.Config) *ExampleDeleter {
	return &ExampleDeleter{
		config:            config,
		observer:          observer,
		database:          database,
		exampleRepository: exampleRepository,
	}
}

type ExampleDeleterDeleteParams struct {
	ID string
}

func (self *ExampleDeleter) Delete(ctx context.Context, params ExampleDeleterDeleteParams) error {
	err := self.database.Transaction(ctx, &kit.IsoLvlSerializable, func(ctx context.Context) error {
		example, err := self.exampleRepository.Get(ctx, params.ID)
		if err != nil {
			return err
		}

		if example != nil {
			err = self.exampleRepository.Delete(ctx, params.ID)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
