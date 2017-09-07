package tree 

import (
	"math/rand"
	"sort"
	"testing"
    "github.com/stretchr/testify/assert"
    "github.com/sunshower-io/anvil/collections"
)

func compareInt(lhs, rhs collections.Value) int {
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


func TestSimpleIterationWorks(t *testing.T) {
    
    tm := NewTreeMap(compareInt)
    tm.Put(1, 2)
    tm.Put(2, 2)
    count := 0
    for iter := tm.Iterator(); iter.HasNext(); {
        iter.Next()
        count++
    }
    assert.Equal(t, count, 2)
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
	for iter := tm.Iterator(); iter.HasNext(); {
		nextKey, _ := iter.Next()
		values = append(values, nextKey.(collections.Entry).Key().(int))
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
