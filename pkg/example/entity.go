package example

import (
	"fmt"
	"time"

	"github.com/neoxelox/kit/util"
)

const (
	EXAMPLE_MINIMUM_AGE = 18
)

const (
	ExampleRoleAdmin  = "ADMIN"
	ExampleRoleMember = "MEMBER"
)

func IsExampleRole(value string) bool {
	return value == ExampleRoleAdmin ||
		value == ExampleRoleMember
}

type ExampleSettings struct {
	WantsNewsletter bool
}

type Example struct {
	ID        string
	Name      string
	Age       int
	Country   string
	Role      string
	Settings  ExampleSettings
	DeletedAt *time.Time
}

func NewExample() *Example {
	return &Example{}
}

func (self Example) String() string {
	return fmt.Sprintf("<Example: %s (%s)>", self.Name, self.ID)
}

func (self Example) Equals(other Example) bool {
	return util.Equals(self, other)
}

func (self Example) Copy() *Example {
	return util.Copy(self)
}
