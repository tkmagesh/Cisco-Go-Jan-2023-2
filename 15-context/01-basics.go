package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	rootCtx := context.Background()
	ctx, cancel := context.WithCancel(rootCtx)
	defer cancel()

	go func() {
		fmt.Scanln()
		cancel()
	}()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go fn(ctx, wg)
	wg.Wait()
}

func fn(ctx context.Context, wg *sync.WaitGroup) {
	i := 1
	wg.Add(1)
	go f1(ctx, wg)
LOOP:
	for {
		select {
		case <-ctx.Done():
			break LOOP
		default:
			fmt.Println("[fn] -", i*10)
			time.Sleep(500 * time.Millisecond)
			i++
		}

	}
	fmt.Println("fn shutting down")
	wg.Done()
}

func f1(ctx context.Context, wg *sync.WaitGroup) {
	i := 1
LOOP:
	for {
		select {
		case <-ctx.Done():
			break LOOP
		default:
			fmt.Println("[f1] -", i*3)
			time.Sleep(300 * time.Millisecond)
			i++
		}

	}
	fmt.Println("f1 shutting down")
	wg.Done()
}
