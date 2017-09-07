package hash

import (
    "bytes"
    "testing"
    "math/rand"
    "github.com/stretchr/testify/assert"
    "github.com/sunshower-io/anvil/collections"
    "github.com/sunshower-io/anvil/f/strings"
)

func intFunc(value collections.Value) int {
    return value.(int)
}



func TestInsertionWithManyCollisionsWorks(t *testing.T) {
    
    a := func (v collections.Value)  int {
        return 0
    }
    
    m := NewLinearProbeHashMap(a)
    
    m.Put("hello", 1)
    m.Put("world", 2)
    
    assert.Equal(t, m.Remove("hello"), 1)
    assert.Equal(t, m.Remove("world"), 2)
}

func TestInsertionWithCollisionWorks(t *testing.T) {
    
    a := func (v collections.Value)  int {
        return 0
    }
    
    m := NewLinearProbeHashMap(a)
    
    m.Put("hello", 1)
    m.Put("world", 2)
    
    assert.Equal(t, m.Remove("hello"), 1)
    
}

func TestInsertionAndRemovalOfVeryManyStringsWorks(t *testing.T) {
    
    m := NewLinearProbeHashMap(strings.DefaultHash)
    k := make(map[string]int)
    
    for i := 0; i < 100; i++ {
        r := randomString(20)
        m.Put(r, i)
        k[r] = i
    }
    
    c := 0
    for key := range k {
        r := m.Remove(key)
        if r == nil {
            c++
        }
    }
    assert.Equal(t, 0, c)
    
}

func TestInsertionAndRemovalOfManyStringsWorks(t *testing.T) {
    
    m := NewLinearProbeHashMap(strings.DefaultHash)
    k := make(map[string]int)
    
    for i := 0; i < 10; i++ {
        r := randomString(20)
        m.Put(r, i)
        k[r] = i
    }
   
    c := 0
    for key := range k {
        r := m.Remove(key)
        if r == nil {
            c++
        }
    }
    assert.Equal(t, 0, c)
    assert.Equal(t, m.Size(), 0)
    
    checkSize(t, m.values, 0)
    
}
func checkSize(t *testing.T, entries []*entry, expected int) {
    c := 0
    for _, v := range entries {
        if v != nil {
            c++
        }
    }
    assert.Equal(t, c, expected)
}


func TestInsertionResultsInValueBeingInsertedAndRetrievableByKey(t *testing.T) {
    
    m := NewLinearProbeHashMap(strings.DefaultHash)
    m.Put("hello", "world")
    v := m.Get("hello")
    assert.Equal(t, "world", v)
}


func TestInsertionAndRemovalOfValuesWorks(t *testing.T) {
    m := NewLinearProbeHashMap(strings.DefaultHash)
    m.Put("hello", "world")
    v := m.Get("hello")
    assert.Equal(t, "world", v)
    u := m.Remove("hello")
    assert.Nil(t, m.Get("hello"))
    assert.Equal(t, "world", u)
}


func TestIteratingViaRangeWorks(t *testing.T) {
    
    m := NewLinearProbeHashMap(strings.DefaultHash)
    
    for i := 0; i < 1000; i++ {
        m.Put(randomString(40), i)
    }
   
    count := 0
    for value := range m.Range() {
        count++
        assert.True(t, m.ContainsKey(value.(collections.Entry).Key()))
    }
    assert.Equal(t, count, 1000)
}


func TestIteratingOverManyStringsWorks(t *testing.T) {
    
    m := NewLinearProbeHashMap(strings.DefaultHash) 
    
    for i := 0; i < 1000; i++ {
        m.Put(randomString(40), i)
    }
    
    count := 0
    for iter := m.Iterator(); iter.HasNext(); {
        count++
        value, _ := iter.Next()
        assert.True(t, m.ContainsKey(value.(collections.Entry).Key()))
    }
    assert.Equal(t, 1000, count)
    
}

func TestInsertionOfManyStringsWorks(t *testing.T) {
    m := NewLinearProbeHashMap(strings.DefaultHash) 
    hm := make(map[string]int)
    
    for i := 0; i < 10000; i++ {
        s := randomString(20)
        hm[s] = i
        m.Put(s, i)
    }
    
    for k, v := range hm {
        assert.Equal(t, v, m.Get(k))
    }
}

func TestInsertionOfManyElementsWorks(t *testing.T) {
    
    m := NewLinearProbeHashMap(intFunc)
    
    cmp := make(map[int]int)
    
    rand := rand.NewSource(100)
    for i := 0; i < 100000; i++ {
        v := int(rand.Int63())
        //cmp[i] = v
        m.Put(i, v)
    }
    for i, k := range cmp {
        assert.Equal(t, k, m.Get(i))
    }
}

func TestInsertionsAndDeletionsWork(t *testing.T) {
    m := NewLinearProbeHashMap(strings.DefaultHash)
    m.Put("hello", "world")
    v := m.Get("hello")
    assert.Equal(t, "world", v)
}


func TestIteratingWorks(t *testing.T) {
    
    m := NewLinearProbeHashMap(strings.DefaultHash)
    m.Put("hello", "world")
    count := 0 
    for iter := m.Iterator(); iter.HasNext(); {
        iter.Next()
        count++
    }
    assert.Equal(t, count, 1)
}

const alphabet = "abcdefghijklmnopqrstuvwxwy"

func randomString(l int) string {
    
    
    var r bytes.Buffer
    for i := 0; i < l; i++ {
        ch := alphabet[int(rand.Int()) % len(alphabet)]
        r.WriteRune(rune(ch))
    }
    c := r.String()
    return c
}

