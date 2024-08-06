-- alter table jobs
ALTER TABLE job_scheduler.jobs ADD IF NOT EXISTS is_recurring boolean;