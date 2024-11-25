// [_Command-line arguments_](https://en.wikipedia.org/wiki/Command-line_interface#Arguments)
// are a common way to parameterize execution of programs.
// For example, `go run hello.go` uses `run` and
// `hello.go` arguments to the `go` program.

package main

import (
	"fmt"
	"os"
)

func main() {

	// `os.Args` provides access to raw command-line
	// arguments. Note that the first value in this slice
	// is the path to the program, and `os.Args[1:]`
	// holds the arguments to the program.
	//os.Args 是一个字符串切片，包含程序运行时的所有命令行参数。
	//os.Args[0] 是程序的名称或路径（即运行的可执行文件），
	//os.Args[1:] 则是实际传递给程序的命令行参数。
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	//这里尝试访问 os.Args[3]，即程序传入的第四个参数（索引从 0 开始）。
	//注意：如果运行时没有传递足够的参数，访问 os.Args[3] 将会导致 索引越界 错误，程序会 panic。因此在实际使用时需要确保传递了足够的参数。
	arg := os.Args[3]
	// You can get individual args with normal indexing.
	//使用命令行参数时，要确保传递了足够的参数，否则会出现访问越界错误（index out of range）。
	//可以使用 len(os.Args) 来检查传入的参数个数，以避免访问越界
	if len(os.Args) > 3 {
		fmt.Println(os.Args[3])
	} else {
		fmt.Println("not enough args!")
	}

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
