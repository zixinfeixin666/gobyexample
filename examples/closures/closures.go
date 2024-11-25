// Go supports [_anonymous functions_](https://en.wikipedia.org/wiki/Anonymous_function),
// which can form <a href="https://en.wikipedia.org/wiki/Closure_(computer_science)"><em>closures</em></a>.
// Anonymous functions are useful when you want to define
// a function inline without having to name it.

package main

import "fmt"

// This function `intSeq` returns another function, which
// we define anonymously in the body of `intSeq`. The
// returned function _closes over_ the variable `i` to
// form a closure.
// intSeq 函数返回一个匿名函数。这个匿名函数没有名字，但它会被返回并赋值给调用者的变量。
// 闭包的关键点在于，返回的匿名函数“闭合”了对变量 i 的引用。即使 intSeq 函数已经执行完毕，返回的匿名函数依然能够访问并修改 i。
// 每次调用 intSeq 时，都会创建一个新的匿名函数，并且这个匿名函数拥有自己的 i 变量，因此每个调用 intSeq 返回的函数都有自己独立的 i 状态。
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	// We call `intSeq`, assigning the result (a function)
	// to `nextInt`. This function value captures its
	// own `i` value, which will be updated each time
	// we call `nextInt`.
	nextInt := intSeq()

	// See the effect of the closure by calling `nextInt`
	// a few times.
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// To confirm that the state is unique to that
	// particular function, create and test a new one.
	newInts := intSeq()
	fmt.Println(newInts())
}
