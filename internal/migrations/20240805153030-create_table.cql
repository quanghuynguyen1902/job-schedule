-- Create jobs table
CREATE TABLE IF NOT EXISTS jobs (
    user_id text,
    job_id text,
    retry_times int,
    created timestamp,
    interval text,
    PRIMARY KEY ((user_id), job_id)
) WITH CLUSTERING ORDER BY (job_id ASC);

-- Create job_executions table (history)
CREATE TABLE IF NOT EXISTS job_histories (
     job_id text,
     execution_time timestamp,
     status text,
     retry_count int,
     last_updated_time timestamp,
     PRIMARY KEY ((job_id), execution_time)
) WITH CLUSTERING ORDER BY (execution_time DESC);

-- Create job_schedules table
CREATE TABLE IF NOT EXISTS job_schedules (
     next_execution_time int,
     job_id text,
     PRIMARY KEY ((next_execution_time), job_id)
) WITH CLUSTERING ORDER BY (job_id ASC);

-- Create indexes for efficient querying
CREATE INDEX IF NOT EXISTS idx_jobs_interval ON jobs (interval);
CREATE INDEX IF NOT EXISTS idx_history_status ON job_histories (status);