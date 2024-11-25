// [_Command-line flags_](https://en.wikipedia.org/wiki/Command-line_interface#Command-line_option)
// are a common way to specify options for command-line
// programs. For example, in `wc -l` the `-l` is a
// command-line flag.

package main

//Go 提供了一个内置的 flag 包，用于解析命令行标志。命令行标志（也叫选项、开关）是指传递给程序的带有特定格式的参数，
//例如 -word=foo、-numb=42、-fork=true 等。
// Go provides a `flag` package supporting basic
// command-line flag parsing. We'll use this package to
// implement our example command-line program.
import (
	"flag"
	"fmt"
)

func main() {

	// Basic flag declarations are available for string,
	// integer, and boolean options. Here we declare a
	// string flag `word` with a default value `"foo"`
	// and a short description. This `flag.String` function
	// returns a string pointer (not a string value);
	// we'll see how to use this pointer below.
	//flag.String：声明一个字符串类型的标志 -word，默认值是 "foo"，描述为 "a string"。
	wordPtr := flag.String("word", "foo", "a string")

	// This declares `numb` and `fork` flags, using a
	// similar approach to the `word` flag.
	//flag.Int：声明一个整数类型的标志 -numb，默认值是 42，描述为 "an int"
	numbPtr := flag.Int("numb", 42, "an int")
	//flag.Bool：声明一个布尔类型的标志 -fork，默认值是 false，描述为 "a bool"
	forkPtr := flag.Bool("fork", false, "a bool")

	// It's also possible to declare an option that uses an
	// existing var declared elsewhere in the program.
	// Note that we need to pass in a pointer to the flag
	// declaration function.
	var svar string
	//flag.StringVar：声明一个字符串类型的标志 -svar，但这个标志的值会直接存储在 svar 变量中，默认值是 "bar"，描述为 "a string var"。
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// Once all flags are declared, call `flag.Parse()`
	// to execute the command-line parsing.
	//flag.Parse() 调用用于解析命令行中的标志。它会解析所有传入的命令行参数，并将它们与已声明的标志进行匹配，存储相应的值。
	flag.Parse()

	// Here we'll just dump out the parsed options and
	// any trailing positional arguments. Note that we
	// need to dereference the pointers with e.g. `*wordPtr`
	// to get the actual option values.
	//解析后的标志存储在指针中，因此我们需要通过解引用（*wordPtr、*numbPtr、*forkPtr）来访问它们的实际值。
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	//对于 svar 变量，它已经被通过 flag.StringVar 函数绑定到 svar 变量，因此可以直接使用它的值。
	fmt.Println("svar:", svar)
	//flag.Args() 返回一个字符串切片，包含所有未被解析为标志的命令行参数。换句话说，它包含所有位置参数（也就是没有前缀 - 的参数）。
	fmt.Println("tail:", flag.Args())
}

//总结
//使用 flag 包，可以很方便地解析命令行参数。通过声明标志（flag），我们可以接收不同类型的输入（如字符串、整数、布尔值等）。
//标志的值存储在指针中，因此在访问时需要解引用。
//flag.Args() 用于获取未解析为标志的其他参数。
//在实际应用中，命令行标志非常常见，可以用于控制程序的行为或提供输入参数。

//例子
//go run hello.go -word=hello -numb=100 -fork=true -svar=test extraArgs1 extraArgs2
//-word=hello 会设置 word 为 "hello"。
//-numb=100 会设置 numb 为 100。
//-fork=true 会设置 fork 为 true。
//-svar=test 会设置 svar 为 "test"。
//extraArgs1 和 extraArgs2 是位置参数，它们会被 flag.Args() 捕获。
//程序运行时，输出如下：
//vbnet
//复制代码
//word: hello
//numb: 100
//fork: true
//svar: test
//tail: [extraArgs1 extraArgs2]
