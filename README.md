[![Build Status](https://travis-ci.org/JohnCGriffin/overflow.png)](https://travis-ci.org/JohnCGriffin/overflow)
# overflow
Check for integer overflow in Golang arithmetic and type conversion.
### Install
```sh
go get github.com/johncgriffin/overflow
```
Note that because Go has no template types, the majority of repetitive code is 
generated by overflow_template.sh.  If you have to change an
algorithm, change it there and regenerate the Go code via:
```sh
go generate
```
### Synopsis

#### Arithmetic overflow detection
```go
package main

import "fmt"
import "math"
import "github.com/JohnCGriffin/overflow"

func main() {

	addend := math.MaxInt64 - 5

	for i := 0; i < 10; i++ {
		sum, ok := overflow.Add(addend, i)
		fmt.Printf("%v+%v -> (%v,%v)\n",
			addend, i, sum, ok)
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
		fmt.Printf("%v -> (%v,%v)\n",
			i, ret, ok)
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

### Stay calm and panic

There's a good case to be made that a panic is an unidiomatic but proper response.  Iff you
believe that there's no valid way to continue your program after math goes wayward, you can
use the easier Addp, Mulp, Subp, and Divp versions which return the normal result or panic.

### Performance considerations

Compared with the integer type safety libraries of other languages (such as C++), this
library uses some seemingly slow operations, such as division. But this does not mean that
these methods will be slow, on the contrary, it will be faster than complex implementations
in other languages. The reason is that Go does not allow forced inlining, and any complex
functions will be abandoned for inlining, resulting in additional calling overhead. Short
functions are lightning fast due to automatic inlining. For example, for unsigned 64-bit
integer multiplication overflow detection, when inlining is disabled, division takes five
times as long as long multiplication, but after automatic inlining is allowed, division
takes 1/5 of long multiplication.

Note that using `//go:noinline` in your business function will not affect the inlining of
the library function. Only disabling global inlining through `-gcflags="-l"` will affect the
inlining of this library function.

### Basis and dependencies

This library is based on Go's official compiler implementation and language specification,
which defines the behavior when integer overflow occurs.

### License

[MIT LICENSE](./LICENSE.md)

