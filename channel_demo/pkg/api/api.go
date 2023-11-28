package api

import (
	"fmt"
	"sync"
	"time"

	context "github.com/hattan/learngo/pkg/context"
)

func callApi(ch chan<- string, wg *sync.WaitGroup, delay int) {
	fmt.Printf("\nStarting Api-%d\n", delay)
	// do stuff

	//fake slow api call
	time.Sleep(time.Duration(delay) * time.Second)

	result := fmt.Sprintf("Completed Api-%d", delay)
	wg.Done()
	ch <- result
}

func CallApi(ctx context.MyContext, delay int) {
	ch, _, wg := ctx.GetItems()

	go callApi(ch, wg, delay)
	wg.Add(1)
}
