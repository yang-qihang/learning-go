package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand/v2"
	"time"
)

var ErrNumberReached = errors.New("number 1234 reached")

func main() {
	var sum int64
	var iterations int

	timeoutCtx, cancelTimeout := context.WithTimeout(
		context.Background(),
		2*time.Second,
	)
	defer cancelTimeout()

	ctx, cancel := context.WithCancelCause(timeoutCtx)
	defer cancel(nil)

	for {
		if err := context.Cause(ctx); err != nil {
			break
		}

		n := rand.IntN(100_000_000)

		sum += int64(n)
		iterations++

		if n == 1234 {
			cancel(ErrNumberReached)
			break
		}
	}

	// The child context is the single point at which
	// you can observe both cancellation conditions,
	// because cancellation propagates from parent to child,
	// but never from child to parent.
	fmt.Println("Sum:", sum)
	fmt.Println("Iterations:", iterations)
	fmt.Println("Reason:", context.Cause(ctx))
}
