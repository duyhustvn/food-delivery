package asyncjob

import (
	"context"
	"fmt"
	"sync"
)

type group struct {
	jobs         []Job
	isConcurrent bool
	wg           *sync.WaitGroup
}

func NewGroup(isConcurrent bool, jobs ...Job) *group {
	return &group{jobs: jobs, isConcurrent: isConcurrent, wg: new(sync.WaitGroup)}
}

func (g *group) Run(ctx context.Context) error {
	errChan := make(chan error, len(g.jobs))

	for i := range g.jobs {
		g.wg.Add(1)
		if g.isConcurrent {
			go func(job Job) {
				err := g.runJob(ctx, job)
				if err != nil {
					errChan <- err
				}
				g.wg.Done()
			}(g.jobs[i])
			continue
		}

		if err := g.runJob(ctx, g.jobs[i]); err != nil {
			fmt.Println(err)
			return err
		}
		g.wg.Done()
	}

	g.wg.Wait()

	var err error
	if g.isConcurrent {
		for i := 0; i < len(errChan); i++ {
			err = <-errChan
		}
	}

	return err
}

func (g *group) runJob(ctx context.Context, job Job) error {
	if err := job.Execute(ctx); err != nil {
		for {
			if job.State() == StateRetryFailed {
				return err
			}

			if err := job.Retry(ctx); err == nil {
				return nil
			}
		}
	}
	return nil
}
