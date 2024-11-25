// Go supports _constants_ of character, string, boolean,
// and numeric values.

package main

import (
	"fmt"
	"math"
)

// `const` declares a constant value.
const s string = "constant"

func main() {
	fmt.Println(s)

	// A `const` statement can appear anywhere a `var`
	// statement can.
	//'const' 语句可以出现在 'var' 语句可以出现的任何位置。
	const n = 500000000

	// Constant expressions perform arithmetic with
	// arbitrary precision.
	//常量表达式以任意精度执行算术运算。
	const d = 3e20 / n
	fmt.Println(d)

	// A numeric constant has no type until it's given
	// one, such as by an explicit conversion.
	//数值常量在给定 type 之前没有类型，例如通过显式转换。
	fmt.Println(int64(d))

	// A number can be given a type by using it in a
	// context that requires one, such as a variable
	// assignment or function call. For example, here
	// `math.Sin` expects a `float64`.
	//通过在需要 type 的上下文中使用 number，例如变量赋值或函数调用，可以为 number 指定类型。例如，这里的 'math.sin' 需要一个 'float64'。
	fmt.Println(math.Sin(n))
}
