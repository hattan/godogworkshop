package main

import (
	"fmt"
	"time"

	api "github.com/hattan/learngo/pkg/api"
	context "github.com/hattan/learngo/pkg/context"
)

func main() {
	ctx := context.NewContext()

	start := time.Now()
	defer finalize(&start, ctx)

	api.CallApi(ctx, 4)
	api.CallApi(ctx, 6)

	ctx.Wait(channelHandler)
}

func finalize(start *time.Time, ctx context.MyContext) {
	fmt.Printf("\nTotal time taken: %.2f\n", time.Since(*start).Seconds())
	ctx.Cleanup()
}

func channelHandler(ctx context.MyContext) {
	ch, done, _ := ctx.GetItems()

	for {
		select {
		case <-done:
			return
		case result := <-ch:
			fmt.Printf("Result: %s\n", result)
		}
	}
}
