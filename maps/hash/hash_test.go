package hash

import (
    "math"
    "testing"
    "github.com/magiconair/properties/assert"
    "github.com/sunshower-io/anvil/collections"
    "github.com/docker/docker/pkg/random"
)

func intFunc(value collections.Value) int {
    return value.(int)
}

func stringFunc(value collections.Value) int {
    
    j := 0
    v := value.(string)
    l := len(v)
   
    for i := 0; i < l; i++ {
        j += int(v[i]) * int(math.Pow(31, float64(l - 1 - i)))
    }
    return j
}


func TestInsertionResultsInValueBeingInsertedAndRetrievableByKey(t *testing.T) {
    
    m := NewHashMap(stringFunc)
    m.Put("hello", "world")
    v := m.Get("hello")
    assert.Equal(t, "world", v)
}


func TestInsertionOfManyElementsWorks(t *testing.T) {
    
    m := NewHashMap(intFunc)
    
    cmp := make(map[int]int)
    
    rand := random.NewSource()
    for i := 0; i < 10; i++ {
        v := int(rand.Int63())
        cmp[i] = v
        m.Put(i, v)
    }
   
    for i, k := range cmp {
        assert.Equal(t, k, m.Get(i))
    }
}

func TestInsertionsAndDeletionsWork(t *testing.T) {
    m := NewHashMap(stringFunc)
    m.Put("hello", "world")
    v := m.Get("hello")
    assert.Equal(t, "world", v)
    h := m.Remove("hello")
    assert.Equal(t, h, "world")
}