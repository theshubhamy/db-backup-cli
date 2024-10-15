package utils

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

func ScheduleBackup(cronExp string, backupFunc func()) {
	c := cron.New()
	c.AddFunc(cronExp, backupFunc)
	c.Start()
	fmt.Println("Scheduled backup with cron expression:", cronExp)
}
