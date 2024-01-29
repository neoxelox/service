package example

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/neoxelox/errors"
	"github.com/neoxelox/kit"

	"service/pkg/config"
)

var (
	ErrExampleEndpointsInvalidRole = errors.New("invalid example role")
)

type ExampleEndpoints struct {
	config         config.Config
	observer       *kit.Observer
	localizer      *kit.Localizer
	exampleGetter  *ExampleGetter
	exampleCreator *ExampleCreator
	exampleDeleter *ExampleDeleter
}

func NewExampleEndpoints(observer *kit.Observer, localizer *kit.Localizer, exampleGetter *ExampleGetter,
	exampleCreator *ExampleCreator, exampleDeleter *ExampleDeleter, config config.Config) *ExampleEndpoints {
	return &ExampleEndpoints{
		config:         config,
		observer:       observer,
		localizer:      localizer,
		exampleGetter:  exampleGetter,
		exampleCreator: exampleCreator,
		exampleDeleter: exampleDeleter,
	}
}

func (self *ExampleEndpoints) GetExample(ctx echo.Context) error {
	exampleCopy := self.localizer.Localize(ctx.Request().Context(), "EXAMPLE_COPY", "🚀")

	return ctx.Render(http.StatusOK, "example.html", struct{ EXAMPLE_COPY string }{exampleCopy})
}

type ExampleEndpointsGetExampleByIDRequest struct {
	ID string `param:"id"`
}

type ExampleEndpointsGetExampleByIDResponse struct {
	ExamplePayload
}

func (self *ExampleEndpoints) GetExampleByID(ctx echo.Context) error {
	request := ExampleEndpointsGetExampleByIDRequest{}

	err := ctx.Bind(&request)
	if err != nil {
		return kit.HTTPErrInvalidRequest.Cause(err)
	}

	example, err := self.exampleGetter.Get(ctx.Request().Context(), ExampleGetterGetParams{
		ID: request.ID,
	})
	if err != nil {
		return kit.HTTPErrServerGeneric.Cause(err)
	}

	if example == nil {
		return kit.HTTPErrNotFound
	}

	response := ExampleEndpointsGetExampleByIDResponse{}
	response.ExamplePayload = *NewExamplePayload(*example)

	return ctx.JSON(http.StatusOK, &response)
}

type ExampleEndpointsPostExampleRequest struct {
	Name            string `json:"name"`
	Age             int    `json:"age"`
	Role            string `json:"role"`
	WantsNewsletter *bool  `json:"wants_newsletter"`
}

type ExampleEndpointsPostExampleResponse struct {
	ExamplePayload
}

func (self *ExampleEndpoints) PostExample(ctx echo.Context) error {
	request := ExampleEndpointsPostExampleRequest{}

	err := ctx.Bind(&request)
	if err != nil {
		return kit.HTTPErrInvalidRequest.Cause(err)
	}

	// NOTE: Ideally this would be inside of an ExampleValidator
	if !IsExampleRole(request.Role) {
		return kit.HTTPErrInvalidRequest.Cause(ErrExampleEndpointsInvalidRole.Raise())
	}

	var settings *ExampleSettings
	if request.WantsNewsletter != nil {
		settings = &ExampleSettings{
			WantsNewsletter: *request.WantsNewsletter,
		}
	}

	example, err := self.exampleCreator.Create(ctx.Request().Context(), ExampleCreatorCreateParams{
		Name:     request.Name,
		Age:      request.Age,
		Role:     request.Role,
		Settings: settings,
	})
	if err != nil {
		if ErrExampleCreatorMinor.Is(err) {
			return kit.HTTPErrInvalidRequest.Cause(err)
		}

		return kit.HTTPErrServerGeneric.Cause(err)
	}

	response := ExampleEndpointsPostExampleResponse{}
	response.ExamplePayload = *NewExamplePayload(*example)

	return ctx.JSON(http.StatusOK, &response)
}

type ExampleEndpointsDeleteExampleRequest struct {
	ID string `param:"id"`
}

type ExampleEndpointsDeleteExampleResponse struct {
}

func (self *ExampleEndpoints) DeleteExample(ctx echo.Context) error {
	request := ExampleEndpointsDeleteExampleRequest{}

	err := ctx.Bind(&request)
	if err != nil {
		return kit.HTTPErrInvalidRequest.Cause(err)
	}

	err = self.exampleDeleter.Delete(ctx.Request().Context(), ExampleDeleterDeleteParams{
		ID: request.ID,
	})
	if err != nil {
		return kit.HTTPErrServerGeneric.Cause(err)
	}

	response := ExampleEndpointsDeleteExampleResponse{}

	return ctx.JSON(http.StatusOK, &response)
}
