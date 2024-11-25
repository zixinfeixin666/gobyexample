// Some command-line tools, like the `go` tool or `git`
// have many *subcommands*, each with its own set of
// flags. For example, `go build` and `go get` are two
// different subcommands of the `go` tool.
// The `flag` package lets us easily define simple
// subcommands that have their own flags.

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// We declare a subcommand using the `NewFlagSet`
	// function, and proceed to define new flags specific
	// for this subcommand.
	//fooCmd 和 barCmd 是使用 flag.NewFlagSet 创建的两个子命令，分别对应 foo 和 bar。
	//fooCmd 子命令有两个标志：enable（布尔值）和 name（字符串值）。
	//barCmd 子命令有一个标志：level（整数值）。
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	// For a different subcommand we can define different
	// supported flags.
	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	// The subcommand is expected as the first argument
	// to the program.
	//程序的主体逻辑是检查命令行的第一个参数（即子命令）并根据不同的子命令来解析对应的标志。
	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	// Check which subcommand is invoked.
	//os.Args[1] 是命令行输入的第一个参数，代表子命令（例如 foo 或 bar）。
	//如果没有给定子命令（即 os.Args 长度小于 2），程序会提示错误并退出。
	switch os.Args[1] {

	// For every subcommand, we parse its own flags and
	// have access to trailing positional arguments.
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		fmt.Println("  tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", barCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}

//例子：go run hello.go foo -enable=true -name=Example arg1 arg2
//os.Args 将是：["hello.go", "foo", "-enable=true", "-name=Example", "arg1", "arg2"]。
//程序首先检查 os.Args[1]，它是 foo，所以选择处理 foo 子命令。
//fooCmd.Parse(os.Args[2:]) 解析了从 os.Args[2:] 开始的标志和参数，即 -enable=true 和 -name=Example。
//输出：
// subcommand 'foo'
//   enable: true
//   name: Example
//   tail: [arg1 arg2]
//总结：使用 flag.NewFlagSet 创建子命令，允许为每个子命令定义独立的标志。
//子命令的参数和标志通过 Parse 方法进行解析。
//子命令可以有自己的标志和位置参数，未被标志解析的参数可以通过 Args() 获取。
//子命令模式可以让复杂的命令行工具更加清晰和结构化，例如 git、docker 等工具就使用了子命令模式。
