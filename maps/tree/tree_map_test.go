package tree 

import (
	"math/rand"
	"sort"
	"testing"
    "github.com/stretchr/testify/assert"
    core "github.com/sunshower-io/anvil/collections"
)

func compareInt(lhs, rhs core.Value) int {
	return lhs.(int) - rhs.(int)
}

func TestInsertingResultsInValueBeingRetrievable(t *testing.T) {

	tm := NewTreeMap(compareInt)

	tm.Put(1, 2)
	i := tm.Get(1)

	assert.Equal(t, i, 2)
}

func TestInsertingMultipleValuesResultsInEachValueBeingRetrievalbe(t *testing.T) {
	tm := NewTreeMap(compareInt)

	for i := 0; i < 100; i++ {

		tm.Put(i, i+100)
	}

	assert.Equal(t, tm.Size(), 100)

	for i := 0; i < 100; i++ {
		tm.Get(i)

		assert.Equal(t, tm.Get(i), i+100)
	}

}

func TestRemovingSingleValueWorks(t *testing.T) {

	tm := NewTreeMap(compareInt)

	tm.Put(1, 2)
	assert.Equal(t, tm.Get(1), 2)
	tm.Remove(1)
	assert.Equal(t, tm.IsEmpty(), true)
}

func TestIteratingOverMapWorks(t *testing.T) {

	tm := NewTreeMap(compareInt)

	for i := 0; i < 100; i++ {
		k := int(rand.Int()) % 100
		tm.Put(k, k+100)
	}
	min := tm.FirstKey()
	assert.Equal(t, min, 0)

	assert.Equal(t, min, 0)
	values := make([]int, tm.Size())
	for iter := tm.Iterator(); iter.HasNext(); iter = iter.Next() {
		nextKey := iter.NextKey()
		values = append(values, nextKey.(int))
	}

	assert.True(t, sort.IntsAreSorted(values))
}

func TestRemovingValueWorks(t *testing.T) {

	tm := NewTreeMap(compareInt)

	for i := 0; i < 100; i++ {
		tm.Put(i, i+100)
	}

	println(tm.String())

	assert.Equal(t, tm.Size(), 100)

	for i := 0; i < 100; i++ {
		tm.Remove(i)
	}

	assert.Equal(t, tm.IsEmpty(), true)
}
