package example

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/neoxelox/errors"
	"github.com/neoxelox/kit"
	"github.com/neoxelox/kit/util"

	"service/pkg/config"
)

const (
	EXAMPLE_SERVICE_TIMEOUT = 2 * time.Second
)

var (
	ErrExampleServiceGeneric  = errors.New("example service failed")
	ErrExampleServiceTimedOut = errors.New("example service timed out")
)

type ExampleService struct {
	config   config.Config
	observer *kit.Observer
	client   *kit.HTTPClient
}

func NewExampleService(observer *kit.Observer, config config.Config) *ExampleService {
	client := kit.NewHTTPClient(observer, kit.HTTPClientConfig{
		Timeout:          EXAMPLE_SERVICE_TIMEOUT,
		AllowedRedirects: util.Pointer(0),
		DefaultRetry: util.Pointer(kit.RetryConfig{
			Attempts:     3,
			InitialDelay: 1 * time.Second,
			LimitDelay:   5 * time.Second,
			Retriables: []error{
				kit.ErrHTTPClientTimedOut,
				kit.ErrHTTPClientBadStatus,
			},
		}),
	})

	return &ExampleService{
		config:   config,
		observer: observer,
		client:   client,
	}
}

type ExampleServiceGetCountryParams struct {
	Name string
}

type ExampleServiceGetCountryResult struct {
	Country string
}

func (self *ExampleService) GetCountry(ctx context.Context,
	params ExampleServiceGetCountryParams) (*ExampleServiceGetCountryResult, error) {
	response, err := self.client.Request(
		ctx, "GET", fmt.Sprintf("%s/?name=%s", self.config.ExampleService.BaseURL, params.Name), nil)
	if err != nil {
		if kit.ErrHTTPClientTimedOut.Is(err) {
			return nil, ErrExampleServiceTimedOut.Raise().Cause(err)
		}

		return nil, ErrExampleServiceGeneric.Raise().Cause(err)
	}
	defer response.Body.Close()

	var body map[string]any

	err = json.NewDecoder(response.Body).Decode(&body)
	if err != nil {
		return nil, ErrExampleServiceGeneric.Raise().Cause(err)
	}

	country, ok := body["country"].([]any)[0].(map[string]any)["country_id"].(string)
	if !ok {
		return nil, ErrExampleServiceGeneric.Raise().With("country not present")
	}

	result := ExampleServiceGetCountryResult{}
	result.Country = country

	return &result, nil
}

func (self *ExampleService) Close(ctx context.Context) error {
	err := util.Deadline(ctx, func(exceeded <-chan struct{}) error {
		self.observer.Info(ctx, "Closing Example service")

		err := self.client.Close(ctx)
		if err != nil {
			return ErrExampleServiceGeneric.Raise().Cause(err)
		}

		self.observer.Info(ctx, "Closed Example service")

		return nil
	})
	if err != nil {
		if util.ErrDeadlineExceeded.Is(err) {
			return ErrExampleServiceTimedOut.Raise().Cause(err)
		}

		return err
	}

	return nil
}
