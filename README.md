# overflow
Check for int/int64/int32 integer overflow in Golang arithmetic.
### Install
```
go get github.com/JohnCGriffin/overflow
```
### Synopsis

```
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
```
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

For int, int64, and int32 types, provide Add, Add32, Add64, Sub, Sub32, Sub64, etc.  
Unsigned types not covered at the moment, but such additions are welcome.

