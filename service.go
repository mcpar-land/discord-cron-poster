package main

import (
	"time"

	"github.com/robfig/cron/v3"
)

type Service struct {
	cron   *cron.Cron
	config Configuration
}

func NewService(config Configuration) (Service, error) {
	var service Service

	var c *cron.Cron
	if config.Timezone == "" {
		c = cron.New()
	} else {
		location, err := time.LoadLocation(config.Timezone)
		if err != nil {
			return service, err
		}
		c = cron.New(cron.WithLocation(location))
	}
	service = Service{c, config}

	for _, job := range service.config.Jobs {
		err := job.Validate()
		if err != nil {
			return service, err
		}
		_, err = service.cron.AddFunc(job.Cron, job.CronFunction(&service.config))
		if err != nil {
			return service, err
		}
	}

	return service, nil
}

func (s Service) Start() {
	s.cron.Start()
}
