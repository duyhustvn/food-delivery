package asyncjob

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
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

func (g *group) Run1(ctx context.Context) error {
	errChan := make(chan error, len(g.jobs))

	for _, job := range g.jobs {
		g.wg.Add(1)
		go g.runJobChan(ctx, job, errChan)
		time.Sleep(time.Second * 1)
	}

	g.wg.Wait()

	select {
	case err := <-errChan:
		if err != nil {
			return err
		}
		log.Printf("err: %+v", err)
		g.wg.Done()
	}

	return nil
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

func runGroup(ctx context.Context) {

}

func (g *group) runJobChan(ctx context.Context, job Job, c chan error) {
	if err := job.Execute(ctx); err != nil {
		for {
			if job.State() == StateRetryFailed {
				c <- err
			}

			if err := job.Retry(ctx); err == nil {
				c <- nil
			}
		}
	}
}
