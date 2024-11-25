// We can use channels to synchronize execution
// across goroutines. Here's an example of using a
// blocking receive to wait for a goroutine to finish.
// When waiting for multiple goroutines to finish,
// you may prefer to use a [WaitGroup](waitgroups).

package main

import (
	"fmt"
	"time"
)

// This is the function we'll run in a goroutine. The
// `done` channel will be used to notify another
// goroutine that this function's work is done.
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// Send a value to notify that we're done.
	done <- true
}

func main() {

	// Start a worker goroutine, giving it the channel to
	// notify on.
	done := make(chan bool, 1)
	go worker(done)
	//
	// Block until we receive a notification from the
	// worker on the channel.
	//Go 中的 goroutine 是并发执行的，main 函数启动了 worker goroutine 后，它并不会等待 worker 完成工作，而是继续执行。
	//如果没有 <-done，main 函数可能会在 worker 完成工作之前就退出。这样就无法看到 worker 的输出（如 "working..." 和 "done"）。
	//当 worker goroutine 执行完任务并通过 done <- true 向通道发送值时，main 函数中的 <-done 会解除阻塞，程序继续执行。
	<-done
}
