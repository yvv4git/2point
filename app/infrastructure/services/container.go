package services

import (
	"log"
	"os"

	"2point/app/infrastructure/components"

	"github.com/sarulabs/di"
)

// NewApp must create di container.
func NewApp() di.Container {
	builder, err := di.NewBuilder()
	if err != nil {
		panic(err)
	}

	err = builder.Add(di.Def{
		Name: "config",
		Build: func(container di.Container) (interface{}, error) {
			config := components.NewConfig()
			return config, nil
		},
	})

	err = builder.Add(di.Def{
		Name: "logger",
		Build: func(container di.Container) (interface{}, error) {
			var config *components.Config
			err = container.Fill("config", &config)
			return log.New(os.Stdout, config.GetPrefix(), 0), err
		},
	})

	err = builder.Add(di.Def{
		Name: "storage",
		Build: func(container di.Container) (interface{}, error) {
			storage, err := components.NewStorage("simple")
			return storage, err
		},
	})

	container := builder.Build()
	return container
}
