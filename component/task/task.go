package task

import "github.com/robfig/cron"

func New() *cron.Cron {
	return cron.New()
}
