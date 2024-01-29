package example

import (
	"context"
	"time"

	"github.com/neoxelox/kit"
	"github.com/neoxelox/kit/util"

	"service/pkg/config"
)

const (
	EXAMPLE_GETTER_TTL = 1 * time.Minute
)

type ExampleGetter struct {
	config            config.Config
	observer          *kit.Observer
	cache             *kit.Cache
	exampleRepository *ExampleRepository
}

func NewExampleGetter(observer *kit.Observer, cache *kit.Cache, exampleRepository *ExampleRepository, config config.Config) *ExampleGetter {
	return &ExampleGetter{
		config:            config,
		observer:          observer,
		cache:             cache,
		exampleRepository: exampleRepository,
	}
}

type ExampleGetterGetParams struct {
	ID string
}

func (self *ExampleGetter) Get(ctx context.Context, params ExampleGetterGetParams) (*Example, error) {
	var example Example
	err := self.cache.Get(ctx, params.ID, &example)
	if err == nil {
		return &example, nil
	} else if !kit.ErrCacheMiss.Is(err) {
		return nil, err
	}

	rExample, err := self.exampleRepository.Get(ctx, params.ID)
	if err != nil {
		return nil, err
	}

	if rExample == nil || rExample.DeletedAt != nil {
		return nil, nil
	}

	err = self.cache.Set(ctx, params.ID, rExample, util.Pointer(EXAMPLE_GETTER_TTL))
	if err != nil {
		return nil, err
	}

	return rExample, nil
}
