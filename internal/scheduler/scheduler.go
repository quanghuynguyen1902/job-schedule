package scheduler

import (
	"encoding/json"
	"log"
	"time"

	"github.com/quanghuynguyen1902/job-schedule/internal/model"
	timeUtils "github.com/quanghuynguyen1902/job-schedule/internal/utils"
)

func (s *SchedulingService) Run() {
	ticker := time.NewTicker(s.pollInterval)
	defer ticker.Stop()

	for range ticker.C {
		if err := s.processPendingTasks(); err != nil {
			log.Printf("Error processing pending tasks: %v", err)
		}
	}
}

func (s *SchedulingService) processPendingTasks() error {
	// 3.1 - Poll the task_schedule table for pending tasks
	tasks, err := s.getPendingTasks()
	if err != nil {
		return err
	}

	for _, task := range tasks {
		// 3.2 - Produce pending tasks to task topic
		if err := s.produceTask(task); err != nil {
			log.Printf("Error producing task %s: %v", task.JobID, err)
			continue
		}

		// 3.3 - Add row to task_execution_history and update task_schedule if recurring
		if err := s.updateTaskStatus(task); err != nil {
			log.Printf("Error updating task status for %s: %v", task.JobID, err)
		}
	}

	return nil
}

func (s *SchedulingService) getPendingTasks() ([]model.JobSchedule, error) {
	return s.store.Schedule.GetSchedules(s.repo.Session(), time.Now().Unix())
}

func (s *SchedulingService) produceTask(task model.JobSchedule) error {
	value, err := json.Marshal(task)
	if err != nil {
		return err
	}

	err = s.queue.Produce("task", "", value)
	if err != nil {
		s.logger.Errorf(err, "failed to produce task %s", task.JobID)
	}

	return nil
}

func (s *SchedulingService) updateTaskStatus(task model.JobSchedule) error {
	// Add row to task_execution_history
	if err := s.store.History.CreateHistoryRecord(s.repo.Session(), model.JobHistory{
		JobID:         task.JobID,
		ExecutionTime: time.Now(),
		Status:        model.JobHistoryStatusPending,
	}); err != nil {
		return err
	}

	// Check if job is recurring and update next_execution_time if needed
	job, err := s.store.Job.GetJob(s.repo.Session(), task.JobID)
	if err != nil {
		return err
	}

	if job.IsRecurring {
		nextExecutionTime, err := s.calculateNextExecutionTime(job.Interval, task.NextExecutionTime)
		if err != nil {
			return err
		}

		if err := s.store.Schedule.UpsertScheduleRecord(s.repo.Session(), model.JobSchedule{
			JobID:             task.JobID,
			NextExecutionTime: nextExecutionTime,
		}); err != nil {
			return err
		}
	}

	return nil
}

func (s *SchedulingService) calculateNextExecutionTime(interval string, lastExecutionTime int64) (int64, error) {
	duration, err := timeUtils.ParseISO8601Duration(interval)
	if err != nil {
		return 0, err
	}
	return lastExecutionTime + int64(duration.Minutes()), nil
}
