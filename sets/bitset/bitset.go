package bitset

import (
	"errors"
	"github.com/Tony-Kwan/gods/sets"
	"sort"
)

const (
	memoryLimit = 512 * 1024 * 1024 // 512MB
	bytesLimit  = memoryLimit >> 3
)

var (
	outOfRangeErr = errors.New("bit offset is not an integer or out of range")
)

func assertSetImplementation() {
	var _ sets.Set = (*Set)(nil)
}

// Set holds elements in go's uint64 slice
type Set struct {
	bits []uint64
}

func New(values ...int) *Set {
	set := &Set{}
	sort.Ints(values)
	for i := len(values) - 1; i >= 0; i-- {
		set.add(values[i])
	}
	return set
}

// Add adds the items (one or more) to the set.
func (set *Set) Add(items ...interface{}) {
	for _, item := range items {
		// panic if item is not a int
		set.add(item.(int))
	}
}

func (set *Set) add(item int) {
	panic("TODO")
}

// Remove removes the items (one or more) from the set.
func (set *Set) Remove(items ...interface{}) {
	panic("TODO")
}

func (set *Set) Contains(items ...interface{}) bool {
	panic("TODO")
}

func (set *Set) Empty() bool {
	panic("TODO")
}

func (set *Set) Size() int {
	panic("TODO")
}
func (set *Set) Clear() {
	panic("TODO")
}

func (set *Set) Values() []interface{} {
	panic("TODO")
}
