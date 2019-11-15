package bitset

import (
	"errors"
	"github.com/Tony-Kwan/gods/sets"
	"math/bits"
	"sort"
)

const (
	memoryLimit = 512 * 1024 * 1024 // 512MB in bytes
	offsetLimit = memoryLimit << 3
	one         = 0x01
	zero        = 0x00
)

var (
	outOfRangeErr = errors.New("bit offset is not an integer or out of range")
)

func assertSetImplementation() {
	var _ sets.Set = (*Set)(nil)
}

// Set holds elements in go's uint64 slice
type Set struct {
	bytes []byte
	size  int

	opt Option
}

type Option struct {
	AutoDownscaling bool
	AllowOverflow   bool
}

// New instantiates a new empty set and adds the passed values, if any, to the set
func New(values ...int) *Set {
	set := &Set{}
	sort.Ints(values)
	set.init(values...)
	return set
}

// New instantiates a new empty set with option and adds the passed values, if any, to the set
// TODO: impl option
func NewWithOption(opt Option, values ...int) *Set {
	set := &Set{opt: opt}
	sort.Ints(values)
	set.init(values...)
	return set
}

// Adds the passed values
func (set *Set) init(values ...int) {
	for i := len(values) - 1; i >= 0; i-- {
		set.setBit(values[i], one)
	}
	set.size = len(values)
}

// Add adds the items (one or more) to the set.
func (set *Set) Add(items ...interface{}) {
	for _, item := range items {
		// panic if item is not a int
		offset := item.(int)
		if !set.Contains(offset) {
			set.setBit(offset, one)
			set.size++
		}
	}
}

// Remove removes the items (one or more) from the set.
func (set *Set) Remove(items ...interface{}) {
	for _, item := range items {
		// panic if item is not a int
		offset := item.(int)
		if set.Contains(offset) {
			set.setBit(offset, zero)
			set.size--
		}
	}
}

// Contains check if items (one or more) are present in the set.
// All items have to be present in the set for the method to return true.
// Returns true if no arguments are passed at all, i.e. set is always superset of empty set.
func (set *Set) Contains(items ...interface{}) bool {
	for _, item := range items {
		// panic if item is not a int
		if !set.getBit(item.(int)) {
			return false
		}
	}
	return true
}

// Empty returns true if set does not contain any elements.
func (set *Set) Empty() bool {
	return set.size == 0
}

// Size returns number of elements within the set.
func (set *Set) Size() int {
	return set.size
}

// Clear clears all values in the set.
func (set *Set) Clear() {
	for i := range set.bytes {
		set.bytes[i] = 0x00
	}
	set.size = 0
}

// Values returns all items in the set.
func (set *Set) Values() []interface{} {
	values := make([]interface{}, 0, set.Size())
	for i := range set.bytes {
		byteVal := set.bytes[i]
		for byteVal != 0 {
			bit := bits.Len8(byteVal)
			values = append(values, i*8+(8-bit))
			byteVal &^= 1 << uint8(bit-1)
		}
	}
	return values
}

func (set *Set) getBit(offset int) bool {
	if offset < 0 || offset >= offsetLimit {
		panic(outOfRangeErr)
	}

	byteIdx := offset >> 3
	if byteIdx >= len(set.bytes) {
		return false
	}
	bit := uint8(7 - (offset & 0x07))
	return set.bytes[byteIdx]&(1<<bit) != 0
}

func (set *Set) setBit(offset int, on byte) {
	if offset < 0 || offset >= offsetLimit {
		panic(outOfRangeErr)
	}

	byteIdx := offset >> 3
	set.grow(byteIdx + 1)
	byteVal := set.bytes[byteIdx]
	bit := uint8(7 - (offset & 0x07))

	byteVal &= ^(1 << bit)
	byteVal |= (on & 1) << bit
	set.bytes[byteIdx] = byteVal
}

func (set *Set) grow(byteCnt int) {
	if len(set.bytes) >= byteCnt {
		return
	}
	newBytes := make([]byte, byteCnt)
	copy(newBytes, set.bytes)
	set.bytes = newBytes
}
