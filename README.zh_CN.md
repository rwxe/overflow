# overflow 溢出
检查 Golang 算术和类型转换中的整数溢出。
### 安装

*在这个派生分叉被上游合并前，这仍然是一个维护的派生分叉。*

该库从`johncgriffin/overflow`派生而来并添加了无符号整数的算术和整数之间的溢出检测。

它经过了良好的测试和基准测试，并通过了由GitHub提供的Workflow代码安全扫描。

[![CodeQL](https://github.com/rwxe/overflow/actions/workflows/codeql.yml/badge.svg)](https://github.com/rwxe/overflow/actions/workflows/codeql.yml)

```sh
go get github.com/rwxe/overflow
```

为了兼容旧代码并保持代码简单易读，新的代码仍然不使用泛型，而是使用模板来生成代码。 所以大多数重复代码由`overflow_template.sh`生成。

如果您必须更改算法，请在那里更改并通过以下方式重新生成 Go 代码：
```sh
go generate
```
### 概要

#### 算术溢出检测
```go
package main

import "fmt"
import "math"
import "github.com/rwxe/overflow"

func main() {
	addend := math.MaxInt64 - 5
	for i := 0; i < 10; i++ {
		sum, ok := overflow.Add(addend, i)
		fmt.Printf("%v+%v -> (%v,%v)\n",
			addend, i, sum, ok)
	}
}
```
输出
```go
9223372036854775802+0 -> (9223372036854775802,true)
9223372036854775802+1 -> (9223372036854775803,true)
9223372036854775802+2 -> (9223372036854775804,true)
9223372036854775802+3 -> (9223372036854775805,true)
9223372036854775802+4 -> (9223372036854775806,true)
9223372036854775802+5 -> (9223372036854775807,true)
9223372036854775802+6 -> (0,false)
9223372036854775802+7 -> (0,false)
9223372036854775802+8 -> (0,false)
9223372036854775802+9 -> (0,false)
```
对于 (u)int 类型，提供 (U)Add、(U)Sub、(U)Mul、(U)Div、(U)Quotient 等操作。


#### 类型转换溢出检测
```go
func main() {
	var i uint
	for i = math.MaxInt - 5; i <= math.MaxInt+5; i++ {
		ret, ok := overflow.UintToInt(i)
		fmt.Printf("%v -> (%v,%v)\n",
			i, ret, ok)
	}
}
```
输出
```go
9223372036854775802 -> (9223372036854775802,true)
9223372036854775803 -> (9223372036854775803,true)
9223372036854775804 -> (9223372036854775804,true)
9223372036854775805 -> (9223372036854775805,true)
9223372036854775806 -> (9223372036854775806,true)
9223372036854775807 -> (9223372036854775807,true)
9223372036854775808 -> (-9223372036854775808,false)
9223372036854775809 -> (-9223372036854775807,false)
9223372036854775810 -> (-9223372036854775806,false)
9223372036854775811 -> (-9223372036854775805,false)
9223372036854775812 -> (-9223372036854775804,false)
```
提供UintToInt、IntToUint、Uint64ToInt32、Int32ToUint64等操作。

#### 取绝对值
```go
func main() {
    normalAbs := func(x int64) int64 {
        if x < 0 {
            x = -x
        }
        return x
    }
    var i1, j1, k1 int64 = -9007199254740993, -9007199254740993, -9007199254740993
    fmt.Println(int64(math.Abs(float64(i1))))
    fmt.Println(normalAbs(j1))
    fmt.Println(overflow.Abs64(k1))

    var i2, j2, k2 int64 = math.MinInt64, math.MinInt64, math.MinInt64
    fmt.Println(int64(math.Abs(float64(i2))))
    fmt.Println(normalAbs(j2))
    fmt.Println(overflow.Abs64(k2))
}
```
输出
```go
9007199254740992 // Mantissa overflow, precision lost
9007199254740993
9007199254740993 true
-9223372036854775808
-9223372036854775808
-9223372036854775808 false // Overflow detected
```

对于整数，提供包含溢出检测的取绝对值操作。

### 保持冷静并恐慌

有充分的证据表明，恐慌是一种不惯用但正确的反应。 如果你相信在算术和转换出现问题后，没有有效的方法可以继续你的程序，你可以使用更简单的 Addp、Mulp、Subp 、Divp、UintToIntp、Absp等版本，它们返回正常结果或恐慌。

### 性能考虑

与其他语言（例如C++）的整数类型安全库相比，该库使用了一些看似缓慢的操作，例如除法。 但这并不意味着这些方法会很慢，实际上，因为Go为短函数自动内联的策略以及编译器优化，这些操作实际上在基准测试里非常快，快如闪电。

此外，Go定义了有符号整数的溢出行为，使得对于有符号整数溢出的检测简单而高效。

请注意，在业务函数中使用 `//go:noinline` 不会影响本库函数的内联。 只有通过 `-gcflags="-l"` 禁用全局内联才会影响该本库函数的内联。

### 基于和依赖

该库基于Go的官方编译器实现（gc）和语言规范开发。

### 许可

[MIT LICENSE](./LICENSE.md)

