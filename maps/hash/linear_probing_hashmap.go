package hash

import (
    "math"
    "github.com/sunshower-io/anvil/collections"
)

const (
    INITIAL_CAPACITY            = 10
    
    
    DEFAULT_LOAD_FACTOR         =  float32(0.90)
)

type entry struct {
    key         collections.Key
    value       collections.Value
}

func (e *entry) Key() collections.Key {
    return e.key
}

func (e *entry) Value() collections.Value {
    return e.value
}

type HashFunction func(collections.Value) int

func NewLinearProbeHashMap(f HashFunction) *linearProbeHashMap {
    return &linearProbeHashMap{
        len                 : 0,
        HashFunction        : f,
        values              : make([]*entry, INITIAL_CAPACITY),
        loadfactor          : DEFAULT_LOAD_FACTOR,
    }
}

type linearProbeHashMap struct {
    collections.Map
    len                 int
    values              []*entry
    HashFunction        HashFunction
    loadfactor          float32
}

func (h *linearProbeHashMap) Iterator() collections.Iterator {
    return &linearProbeIterator{
        m:h,
        vlen: len(h.values),
    }
}


func(h *linearProbeHashMap) Size() int {
    return h.len
}




func (h *linearProbeHashMap) Put(
        k collections.Key, 
        v collections.Value,
) collections.Value {
    
    values      := h.values
    i, e := locate(k, h.values, h.HashFunction, h.len, true)
    var (
        r collections.Value
    )
    
    if e != nil {
        e.value = v 
    } else {
        l           := h.len
        lv          := len(values)
        lf          := h.loadfactor
        hf          := h.HashFunction
        load        := int(lf * float32(lv))
        
        
        if l >= load {
            nv := make([]*entry, lv * 2)
            for _, ev := range values {
                if ev != nil {
                    set(ev.key, ev.value, nv, hf, l)
                }
            }
            i, e = locate(k, nv, hf, l, true)
            values = nv
        }
        
        values[i] = &entry{k, v}
    }
    
    h.len += 1
    h.values = values
    return r
    
}



const (
    mask = math.MaxInt32
)


func (h *linearProbeHashMap) Remove(
        key collections.Key,
) collections.Value {
   
    length     := h.len
    values     := h.values
    hf         := h.HashFunction
    vlength    := len(values)
   
    var value collections.Value
    i, e := locate(key, values, hf, length, false)
    
    if e != nil {
        value = e.value
    }
    
    
    
    
    j := i
    for {
        j = ((j + 1) & mask) % vlength 
        v := values[j]
        if v == nil {
            break
        } 
        k := (hf(v.key) & mask) % vlength 
        if (j > i && (k <= i || k > j)) || 
                (j < i && (k <= i && k > j)) {
            values[i] = values[j]
            i = j
            values[j] = nil
        }
    }
    values[i] = nil
    h.len -= 1
    return value
}

func (h *linearProbeHashMap) ContainsKey(
        key collections.Key,
) bool {
    values      := h.values
    l           := len(values)
    hf          := h.HashFunction
    
    i           := (hf(key) & mask) % l
    for {
        v := values[i]
        if v == nil {
            return false
        }
        
        if v.key == key {
            return true 
        }
        i = ((i + 1) & mask) % l
    }
    
}



func (h *linearProbeHashMap) Get(
        k collections.Key,
) collections.Value {
    
    v := h.values
    if _, e := locate(k, v, h.HashFunction, h.len, false); e != nil {
        return e.value
    }
    return nil
}



func set(
        k collections.Key, 
        v collections.Value, 
        values[]*entry,
        h HashFunction,
        hl int,
) *entry {
    i, e := locate(k, values, h, hl, false)
    var r *entry
    if e != nil {
        r = e
        e.value = v
    } else {
        r = &entry{k, v}
        values[i] = r
    }
    return r 
}


func locate(
        key collections.Key,
        values []*entry,
        h HashFunction,
        hl int,
        insert bool,
) (int, *entry) {
    l           := len(values)
    if insert && hl >= l {
        return -1, nil
    }
    i           := (h(key) & mask) % l
    for {
        v := values[i]
        if v == nil || values[i].key == key {
            return i, values[i]
        }
        i = ((i + 1) & mask) % l 
    }
}




type linearProbeIterator struct {
    collections.Iterator
    current         int
    cidx            int
    m               *linearProbeHashMap
    vlen            int
}

func (l *linearProbeIterator) HasNext() bool {
    return l.current < l.m.len
}

func (l *linearProbeIterator) Next() (collections.Value, error) {
    for i := l.cidx; i < l.vlen; i++ {
        v := l.m.values[i]
        if v != nil {
            l.cidx = i
            l.current++
            return v, nil
        }
    }
    return nil, collections.IteratorOverflow
}
