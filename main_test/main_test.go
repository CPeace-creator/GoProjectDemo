package main

import (
	"fmt"
	"sync"
	"testing"
)

func Test(t *testing.T) {
	taskNum := 5
	wg := sync.WaitGroup{}
	wg.Add(taskNum)
	for i := 0; i < taskNum; i++ {
		go func(i int) {
			fmt.Println("info", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
