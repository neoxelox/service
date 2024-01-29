package example

import (
	"context"
	"time"

	"github.com/leporo/sqlf"
	"github.com/neoxelox/kit"

	"service/pkg/config"
)

const (
	EXAMPLE_REPOSITORY_TABLE = "\"example\""
)

type ExampleRepository struct {
	config   config.Config
	observer *kit.Observer
	database *kit.Database
}

func NewExampleRepository(observer *kit.Observer, database *kit.Database, config config.Config) *ExampleRepository {
	return &ExampleRepository{
		config:   config,
		observer: observer,
		database: database,
	}
}

func (self *ExampleRepository) Create(ctx context.Context, example Example) (*Example, error) {
	e := NewExampleModel(example)

	stmt := sqlf.
		InsertInto(EXAMPLE_REPOSITORY_TABLE).
		Set("id", e.ID).
		Set("name", e.Name).
		Set("age", e.Age).
		Set("country", e.Country).
		Set("role", e.Role).
		Set("settings", e.Settings).
		Set("deleted_at", e.DeletedAt).
		Returning("*").To(&e)

	err := self.database.Query(ctx, stmt)
	if err != nil {
		return nil, err
	}

	return e.ToEntity(), nil
}

func (self *ExampleRepository) Get(ctx context.Context, id string) (*Example, error) {
	var e ExampleModel

	stmt := sqlf.
		Select("*").To(&e).
		From(EXAMPLE_REPOSITORY_TABLE).
		Where("id = ?", id)

	err := self.database.Query(ctx, stmt)
	if err != nil {
		if kit.ErrDatabaseNoRows.Is(err) {
			return nil, nil
		}

		return nil, err
	}

	return e.ToEntity(), nil
}

func (self *ExampleRepository) Delete(ctx context.Context, id string) error {
	stmt := sqlf.
		Update(EXAMPLE_REPOSITORY_TABLE).
		Set("deleted_at", time.Now()).
		Where("id = ?", id)

	affected, err := self.database.Exec(ctx, stmt)
	if err != nil {
		return err
	}

	if affected != 1 {
		return kit.ErrDatabaseUnexpectedEffect.Raise(affected, 1)
	}

	return nil
}
