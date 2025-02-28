# overflow
Check for integer overflow in Golang arithmetic and type conversion.

Other language README: [中文版](./README.zh_CN.md)
### Install

*This is still a maintenance fork until it is merged upstream.*

This repository forks from `johncgriffin/overflow` and adds overflow detection for unsigned integer arithmetic and type conversions between integers.

It has been well tested and benchmarked, and passed the code security scan provided by Github Workflow.

[![CodeQL](https://github.com/rwxe/overflow/actions/workflows/codeql.yml/badge.svg)](https://github.com/rwxe/overflow/actions/workflows/codeql.yml)

```sh
go get github.com/rwxe/overflow
```

In order to be compatible with the old code and keep the code simple and readable, the new code still does not use generics, but uses templates to generate code. So the majority of repetitive code is generated by `overflow_template.sh`. 

If you have to change an algorithm, change it there and regenerate the Go code via: 
```sh
go generate
```
### Synopsis

#### Arithmetic overflow detection
```go
import "github.com/rwxe/overflow"

func main() {
    addend := math.MaxInt64 - 5
    for i := 0; i < 10; i++ {
        sum, ok := overflow.Add(addend, i)
        fmt.Printf("%v+%v -> (%v,%v)\n", addend, i, sum, ok)
    }
}
```
yields the output
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
For (u)int types, provide (U)Add, (U)Sub, (U)Mul, (U)Div, (U)Quotient, etc.

#### Type conversion overflow detection
```go
func main() {
    var i uint
    for i = math.MaxInt - 5; i <= math.MaxInt+5; i++ {
        ret, ok := overflow.UintToInt(i)
        fmt.Printf("%v -> (%v,%v)\n", i, ret, ok)
    }
}
```
yields the output
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
Provide UintToInt, IntToUint, Uint64ToInt32, Int32ToUint64, etc.

#### Get absolute value
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
yields the output
```go
9007199254740992 // Mantissa overflow, precision lost
9007199254740993
9007199254740993 true
-9223372036854775808
-9223372036854775808
-9223372036854775808 false // Overflow detected
```

For int, provides an absolute value including overflow detection.

### Stay calm and panic

There's a good case to be made that a panic is an unidiomatic but proper response. If you believe that there's no valid way to continue your program after math goes wayward, you can use the easier Addp, Mulp, Subp, Divp, IntToUintp, Absp etc, which return the normal result or panic.

### Performance considerations

Compared to integer safe libraries in other languages (e.g. C++), this library uses some seemingly slow operations, such as division. But that doesn't mean these methods will be slow. In fact, because of Go's automatic inlining strategy for short functions and compiler optimizations, these operations are actually very fast in benchmark tests, lightning fast. 

In addition, Go defines the overflow behavior of signed integers, making the detection of signed integer overflow simple and efficient.

Note that using `//go:noinline` in your business function will not affect the inlining of the library function. Only disabling global inlining through `-gcflags="-l"` will affect the inlining of this library function.

### Basis and dependencies

The library is developed based on Go's official compiler implementation (gc) and language specifications.

### License

[MIT LICENSE](./LICENSE.md)

