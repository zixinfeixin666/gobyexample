// _Closing_ a channel indicates that no more values
// will be sent on it. This can be useful to communicate
// completion to the channel's receivers.

package main

import "fmt"

// In this example we'll use a `jobs` channel to
// communicate work to be done from the `main()` goroutine
// to a worker goroutine. When we have no more jobs for
// the worker we'll `close` the `jobs` channel.
func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// Here's the worker goroutine. It repeatedly receives
	// from `jobs` with `j, more := <-jobs`. In this
	// special 2-value form of receive, the `more` value
	// will be `false` if `jobs` has been `close`d and all
	// values in the channel have already been received.
	// We use this to notify on `done` when we've worked
	// all our jobs.
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// This sends 3 jobs to the worker over the `jobs`
	// channel, then closes it.
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	//. 关闭通道后不能再发送数据
	//一旦调用 close(jobs)，你就不能再使用 jobs <- x 向通道发送数据了。尝试发送数据到一个关闭的通道会引发运行时错误。
	//close(jobs) 仅仅是通知接收方没有更多数据会发送到通道，意味着工作已经完成或者任务已经全部传送完毕。
	//2. 仍然可以从关闭的通道接收数据
	//即使通道已关闭，你仍然可以继续从关闭的通道接收数据。此时，接收操作的行为会有不同：
	//如果通道没有数据，接收操作会立刻返回该通道类型的零值（例如，int 类型的零值是 0，string 类型的零值是空字符串）。
	//如果通道已经关闭且没有数据可接收，接收操作会返回第二个返回值 false，表示通道已经关闭且没有更多数据。
	//3. close 的作用
	//close 并不会立即释放通道的内存或销毁通道，它只是改变了通道的状态，标记这个通道已经不再用于发送数据了。
	//通道关闭的目的通常是为了告知接收方，不再有更多的数据，接收方可以基于此进行必要的清理或退出操作。
	close(jobs)
	fmt.Println("sent all jobs")

	// We await the worker using the
	// [synchronization](channel-synchronization) approach
	// we saw earlier.
	<-done

	// Reading from a closed channel succeeds immediately,
	// returning the zero value of the underlying type.
	// The optional second return value is `true` if the
	// value received was delivered by a successful send
	// operation to the channel, or `false` if it was a
	// zero value generated because the channel is closed
	// and empty.
	_, ok := <-jobs
	fmt.Println("received more jobs:", ok)
}
