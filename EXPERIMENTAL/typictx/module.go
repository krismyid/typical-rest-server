package typictx

import (
	"github.com/typical-go/typical-rest-server/EXPERIMENTAL/docker"
	"github.com/typical-go/typical-rest-server/EXPERIMENTAL/slice"
)

// Module of typical-go application
type Module struct {
	Config        Config
	Name          string
	OpenFunc      interface{}
	CloseFunc     interface{}
	Command       *Command
	DockerCompose *docker.Compose
	Constructors  slice.Interfaces
}

// Configure return configuration
func (m Module) Configure() Config {
	return m.Config
}
