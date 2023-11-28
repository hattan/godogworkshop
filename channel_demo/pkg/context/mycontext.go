package context

import (
	"context"
	"sync"
)

type MyContext struct {
	ctx context.Context
}

func (ctx MyContext) GetItems() (chan string, chan bool, *sync.WaitGroup) {
	return ctx.getChannel(), ctx.getDoneChannel(), ctx.getWeightGroup()
}

func (ctx MyContext) getChannel() chan string {
	return ctx.ctx.Value(channelKey).(chan string)
}

func (ctx MyContext) getWeightGroup() *sync.WaitGroup {
	return ctx.ctx.Value(wgKey).(*sync.WaitGroup)
}

func (ctx MyContext) getDoneChannel() chan bool {
	return ctx.ctx.Value(doneKey).(chan bool)
}

func (ctx MyContext) Cleanup() {
	ch, done, _ := ctx.GetItems()
	close(ch)
	close(done)
}

func (ctx MyContext) Wait(handler func(ctx MyContext)) {
	go func() {
		_, done, wg := ctx.GetItems()
		wg.Wait()
		done <- true
	}()
	handler(ctx)
}

func NewContext() MyContext {
	var wg sync.WaitGroup
	ch := make(chan string)
	done := make(chan bool)

	ctx := context.Background()
	ctx = context.WithValue(ctx, channelKey, ch)
	ctx = context.WithValue(ctx, wgKey, &wg)
	ctx = context.WithValue(ctx, doneKey, done)

	return MyContext{
		ctx: ctx,
	}
}

type contextKey string

const channelKey contextKey = "ch"
const doneKey contextKey = "done"
const wgKey contextKey = "wg"
