package cron

import (
	"log"
	"time"

	"github.com/go-co-op/gocron"
)

func ScheduleJobs() *gocron.Scheduler {
	log.Println("Scheduling jobs...")
	scheduler := gocron.NewScheduler(time.UTC)

	_, err := scheduler.Every(1).Day().At("00:01").Do(func() { log.Println("every second...") })
	if err != nil {
		log.Println("Job can't be schedule")
	}
	scheduler.StartAsync()

	return scheduler
}
