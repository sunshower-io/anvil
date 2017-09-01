package hash

import (
    "testing"
    "github.com/sunshower-io/anvil/collections"
    "github.com/magiconair/properties/assert"
    "math"
)

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
