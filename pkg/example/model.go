package example

import (
	"encoding/json"
	"time"
)

type ExampleModel struct {
	ID        string     `db:"id"`
	Name      string     `db:"name"`
	Age       int        `db:"age"`
	Country   string     `db:"country"`
	Role      string     `db:"role"`
	Settings  []byte     `db:"settings"`
	DeletedAt *time.Time `db:"deleted_at"`
}

func NewExampleModel(example Example) *ExampleModel {
	settings, err := json.Marshal(example.Settings)
	if err != nil {
		panic(err)
	}

	return &ExampleModel{
		ID:        example.ID,
		Name:      example.Name,
		Age:       example.Age,
		Country:   example.Country,
		Role:      example.Role,
		Settings:  settings,
		DeletedAt: example.DeletedAt,
	}
}

func (self *ExampleModel) ToEntity() *Example {
	var settings ExampleSettings
	err := json.Unmarshal(self.Settings, &settings)
	if err != nil {
		panic(err)
	}

	return &Example{
		ID:        self.ID,
		Name:      self.Name,
		Age:       self.Age,
		Country:   self.Country,
		Role:      self.Role,
		Settings:  settings,
		DeletedAt: self.DeletedAt,
	}
}
