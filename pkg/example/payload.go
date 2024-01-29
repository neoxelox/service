package example

type ExamplePayloadSettings struct {
	WantsNewsletter bool `json:"wants_newsletter"`
}

type ExamplePayload struct {
	ID       string                 `json:"id"`
	Name     string                 `json:"name"`
	Age      int                    `json:"age"`
	Country  string                 `json:"country"`
	Role     string                 `json:"role"`
	Settings ExamplePayloadSettings `json:"settings"`
}

func NewExamplePayload(example Example) *ExamplePayload {
	settings := ExamplePayloadSettings{
		WantsNewsletter: example.Settings.WantsNewsletter,
	}

	return &ExamplePayload{
		ID:       example.ID,
		Name:     example.Name,
		Age:      example.Age,
		Country:  example.Country,
		Role:     example.Role,
		Settings: settings,
	}
}
