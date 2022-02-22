package asyncjob

import (
	"context"
	"errors"
	"time"
)

// JOB description
// 1. Job can do something (handler)
// 2. Job can retry
//   2.1 Config retry times and duration
// 3. Should be stateful
// 4. We should have job manager to manage jobs

type JobState int

const (
	defaultMaxTimeout = time.Second * 10
)

var (
	defaultRetryTime = []time.Duration{time.Second, time.Second * 2, time.Second * 4}
)

const (
	StateInit JobState = iota
	StateRunning
	StateFailed
	StateTimeout
	StateCompleted
	StateRetryFailed
)

func (js JobState) ToString(jsState int) string {
	return [6]string{"Init", "Running", "Failed", "Timeout", "Completed", "RetryFailed"}[jsState]
}

type Job interface {
	Execute(ctx context.Context) error
	Retry(ctx context.Context) error
	State() JobState
	SetRetryDurations(times []time.Duration)
}

type JobHandler func(ctx context.Context) error

type jobConfig struct {
	MaxTimeout time.Duration
	Retries    []time.Duration
}

type job struct {
	config     jobConfig
	handler    JobHandler
	state      JobState
	retryIndex int       // how many retry it is
	stopChan   chan bool // using to stop job when it is executed concurrency/parallel
}

func NewJob(handler JobHandler) *job {
	return &job{
		config: jobConfig{
			MaxTimeout: defaultMaxTimeout,
			Retries:    defaultRetryTime,
		},
		handler:    handler,
		state:      StateInit,
		retryIndex: -1,
		stopChan:   make(chan bool),
	}
}

func (j *job) Execute(ctx context.Context) error {
	j.state = StateRunning

	var err error
	if err = j.handler(ctx); err != nil {
		j.state = StateFailed
		return err
	}

	j.state = StateCompleted
	return nil
}

func (j *job) Retry(ctx context.Context) error {
	j.retryIndex++
	time.Sleep(j.config.Retries[j.retryIndex])
	j.state = StateRunning

	if err := j.handler(ctx); err != nil {
		if j.retryIndex == len(j.config.Retries)-1 {
			j.state = StateRetryFailed
			return errors.New("Execced max retry")
		}
		j.state = StateFailed
		return err
	}

	j.state = StateCompleted
	return nil
}

func (j *job) State() JobState {
	return j.state
}

func (j *job) RetryIndex() int {
	return j.retryIndex
}

func (j *job) SetRetry(times []time.Duration) {
	j.config.Retries = times
}
