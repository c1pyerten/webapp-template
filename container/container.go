package container

import (
	"c1pherten/yet-webapp2/config"
	"c1pherten/yet-webapp2/log"
	// "c1pherten/yet-webapp2/repository"
)


// Container global context
type Container interface{
	Logger() log.Logger
	Config() *config.Config
	Env() string
	// Repository() repository.Repository
}

type container struct{
	logger log.Logger
	config *config.Config
	env string
	// repo repository.Repository
}

func (c *container) Logger() log.Logger {
	return c.logger
}
func (c *container) Env() string {
	return c.env
}
// func (c *container) Repository() repository.Repository {
// 	return c.repo
// }
func (c *container) Config() *config.Config {
	return c.config
}

// NewContainer TODO: add parameters
func NewContainer(l log.Logger, config *config.Config, env string) Container {
	return &container{
		logger: l,
		config: config,
		env:    env,
	}
}