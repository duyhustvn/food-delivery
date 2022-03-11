package main

import (
	"context"
	"errors"
	"food-delivery/component/asyncjob"
	"log"
	"time"
)

func main() {
	job1 := asyncjob.NewJob(func(ctx context.Context) error {
		time.Sleep(1 * time.Second)
		log.Println("Hello from the job1")
		//return nil
		return errors.New("Error from job1")
	})

	job2 := asyncjob.NewJob(func(ctx context.Context) error {
		time.Sleep(1 * time.Second)
		log.Println("Hello from the job2")
		return nil
	})

	job3 := asyncjob.NewJob(func(ctx context.Context) error {
		time.Sleep(1 * time.Second)
		log.Println("Hello from the job3")
		return nil
	})

	group := asyncjob.NewGroup(true, job1, job2, job3)
	if err := group.Run(context.Background()); err != nil {
		log.Printf("Error in Run: %+v", err)
	}
}
