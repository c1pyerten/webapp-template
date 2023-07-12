package appctx

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

type appCtx struct{
	logger log.Logger
	config *config.Config
	env string
	// repo repository.Repository
}

func (c *appCtx) Logger() log.Logger {
	return c.logger
}
func (c *appCtx) Env() string {
	return c.env
}
// func (c *container) Repository() repository.Repository {
// 	return c.repo
// }
func (c *appCtx) Config() *config.Config {
	return c.config
}

// NewContainer TODO: add parameters
func NewContainer(l log.Logger, config *config.Config, env string) Container {
	return &appCtx{
		logger: l,
		config: config,
		env:    env,
	}
}