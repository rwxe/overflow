/*
Package overflow offers overflow-checked integer arithmetic operations
for int, int32, and int64. Each of the operations returns a
result,bool combination.  This was prompted by the need to know when
to flow into higher precision types from the math.big library.

For instance, assuing a 64 bit machine:

10 + 20 -> 30
int(math.MaxInt64) + 1 -> -9223372036854775808

whereas

overflow.Add(10,20) -> (30, true)
overflow.Add(math.MaxInt64,1) -> (0, false)

Add, Sub, Mul, Div are for int.  Add64, Add32, etc. are specifically sized.

If anybody wishes an unsigned version, submit a pull request for code
and new tests.
*/
package overflow

//go:generate ./overflow_template.sh

import "math"

func _is64Bit() bool {
	maxU32 := uint(math.MaxUint32)
	return ((maxU32 << 1) >> 1) == maxU32
}

/********** PARTIAL TEST COVERAGE FROM HERE DOWN *************

The only way that I could see to do this is a combination of
my normal 64 bit system and a GopherJS running on Node.  My
understanding is that its ints are 32 bit.

So, FEEL FREE to carefully review the code visually.

*************************************************************/

// Unspecified size, i.e. normal signed int

// Add sums two ints, returning the result and a boolean status.
func Add(a, b int) (int, bool) {
	if _is64Bit() {
		r64, ok := Add64(int64(a), int64(b))
		return int(r64), ok
	}
	r32, ok := Add32(int32(a), int32(b))
	return int(r32), ok
}

// UAdd sums two uints, returning the result and a boolean status.
func UAdd(a, b uint) (uint, bool) {
	if _is64Bit() {
		r64, ok := UAdd64(uint64(a), uint64(b))
		return uint(r64), ok
	}
	r32, ok := UAdd32(uint32(a), uint32(b))
	return uint(r32), ok
}

// Sub returns the difference of two ints and a boolean status.
func Sub(a, b int) (int, bool) {
	if _is64Bit() {
		r64, ok := Sub64(int64(a), int64(b))
		return int(r64), ok
	}
	r32, ok := Sub32(int32(a), int32(b))
	return int(r32), ok
}

// USub returns the difference of two uints and a boolean status.
func USub(a, b uint) (uint, bool) {
	if _is64Bit() {
		r64, ok := USub64(uint64(a), uint64(b))
		return uint(r64), ok
	}
	r32, ok := USub32(uint32(a), uint32(b))
	return uint(r32), ok
}

// Mul returns the product of two ints and a boolean status.
func Mul(a, b int) (int, bool) {
	if _is64Bit() {
		r64, ok := Mul64(int64(a), int64(b))
		return int(r64), ok
	}
	r32, ok := Mul32(int32(a), int32(b))
	return int(r32), ok
}

// UMul returns the product of two uints and a boolean status.
func UMul(a, b uint) (uint, bool) {
	if _is64Bit() {
		r64, ok := UMul64(uint64(a), uint64(b))
		return uint(r64), ok
	}
	r32, ok := UMul32(uint32(a), uint32(b))
	return uint(r32), ok
}

// Div returns the quotient of two ints and a boolean status
func Div(a, b int) (int, bool) {
	if _is64Bit() {
		r64, ok := Div64(int64(a), int64(b))
		return int(r64), ok
	}
	r32, ok := Div32(int32(a), int32(b))
	return int(r32), ok
}

// UDiv returns the quotient of two uints and a boolean status
func UDiv(a, b uint) (uint, bool) {
	if _is64Bit() {
		r64, ok := UDiv64(uint64(a), uint64(b))
		return uint(r64), ok
	}
	r32, ok := UDiv32(uint32(a), uint32(b))
	return uint(r32), ok
}

// Quotient returns the quotient, remainder and status of two ints
func Quotient(a, b int) (int, int, bool) {
	if _is64Bit() {
		q64, r64, ok := Quotient64(int64(a), int64(b))
		return int(q64), int(r64), ok
	}
	q32, r32, ok := Quotient32(int32(a), int32(b))
	return int(q32), int(r32), ok
}

// UQuotient returns the quotient, remainder and status of two uints
func UQuotient(a, b uint) (uint, uint, bool) {
	if _is64Bit() {
		uq64, ur64, ok := UQuotient64(uint64(a), uint64(b))
		return uint(uq64), uint(ur64), ok
	}
	uq32, ur32, ok := UQuotient32(uint32(a), uint32(b))
	return uint(uq32), uint(ur32), ok
}

/************* Panic versions for int ****************/

// Addp returns the sum of two ints, panicking on overflow
func Addp(a, b int) int {
	r, ok := Add(a, b)
	if !ok {
		panic("addition overflow")
	}
	return r
}

// UAddp returns the sum of two uints, panicking on overflow
func UAddp(a, b uint) uint {
	r, ok := UAdd(a, b)
	if !ok {
		panic("addition overflow")
	}
	return r
}

// Subp returns the difference of two ints, panicking on overflow.
func Subp(a, b int) int {
	r, ok := Sub(a, b)
	if !ok {
		panic("subtraction overflow")
	}
	return r
}

// USubp returns the difference of two uints, panicking on overflow.
func USubp(a, b uint) uint {
	r, ok := USub(a, b)
	if !ok {
		panic("subtraction overflow")
	}
	return r
}

// Mulp returns the product of two ints, panicking on overflow.
func Mulp(a, b int) int {
	r, ok := Mul(a, b)
	if !ok {
		panic("multiplication overflow")
	}
	return r
}

// UMulp returns the product of two uints, panicking on overflow.
func UMulp(a, b uint) uint {
	r, ok := UMul(a, b)
	if !ok {
		panic("multiplication overflow")
	}
	return r
}

// Divp returns the quotient of two ints, panicking on overflow.
func Divp(a, b int) int {
	r, ok := Div(a, b)
	if !ok {
		panic("division failure")
	}
	return r
}

// UDivp returns the quotient of two uints, panicking on overflow.
func UDivp(a, b uint) uint {
	r, ok := UDiv(a, b)
	if !ok {
		panic("division failure")
	}
	return r
}
