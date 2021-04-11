package cron

import (
	"testing"
	"time"
)

func TestScheduler(t *testing.T) {

	t.Run("only one job to run", func(t *testing.T) {
		scheduler := ScheduleJobs()

		got := len(scheduler.Jobs())
		want := 1

		if got != want {
			t.Errorf("got '%d' but want '%d'", got, want)
		}
	})

	t.Run("job should run ones a day", func(t *testing.T) {
		scheduler := ScheduleJobs()

		got := scheduler.Jobs()[0].ScheduledTime().Day()
		want := time.Now().Day() + 1

		if got != want {
			t.Errorf("got '%d' but want '%d'", got, want)
		}
	})

	t.Run("job should be schedule at mid night", func(t *testing.T) {
		scheduler := ScheduleJobs()

		got := scheduler.Jobs()[0].ScheduledAtTime()
		want := "0:1"

		if got != want {
			t.Errorf("got '%s' but want '%s'", got, want)
		}
	})
}
