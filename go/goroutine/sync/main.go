package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
)

var s *semaphore.Weighted = semaphore.NewWeighted(10)

func longProcess(ctx context.Context, i int) error {
	if i > 80 {
		return fmt.Errorf("could not be processed: %d", i)
	}
	if err := s.Acquire(ctx, 1); err != nil {
		return err
	}
	defer s.Release(1)
	fmt.Printf("Wait... %d\n", i)
	time.Sleep(1 * time.Second)
	fmt.Printf("Done %d\n", i)
	return nil
}

func main() {
	var eg errgroup.Group
	ctx := context.Background()
	for i := 0; i < 100; i++ {
		i := i
		eg.Go(func() error {
			return longProcess(ctx, i)
		})
	}
	if err := eg.Wait(); err != nil {
		fmt.Println(err)
	}
}
