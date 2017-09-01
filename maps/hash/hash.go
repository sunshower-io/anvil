package hash

import "github.com/sunshower-io/anvil/collections"

const (
    INITIAL_CAPACITY            = 10
    
    
    DEFAULT_LOAD_FACTOR         =  0.75
)

type entry struct {
    key         collections.Key
    value       collections.Value
}

type hashFunction func(collections.Value) int

func NewHashMap(f hashFunction) *HashMap {
    return &HashMap{
        len: 0,
        hashFunction: f,
        values :make([]*entry, INITIAL_CAPACITY),
    }
}

type HashMap struct {
    len                 int
    values              []*entry
    hashFunction        hashFunction
}


func (m *HashMap) Get(
        key collections.Key,
) collections.Value {
   
    h       := m.hashFunction
    values  := m.values
    j       := 0
    len     := len(values)
    k       := h(key)
    
    for {
        code := (k + j) % len 
        if j >= len {
            return nil
        }
        if result := values[code]; result != nil {
            return result.value
        }
    }
}

func (m *HashMap) Remove(
        key collections.Key,
)  collections.Value {
   
    h           := m.hashFunction
    values      := m.values
    j           := 0
    k           := h(key)
    idx         := 0
    len         := len(m.values)
    
    var result *entry
    
    
    for {
        code    := (k + j) % len
        result  = values[code]
        
        if result == nil {
            return nil
        }
        
        if result.key == key {
            idx = code
            break
        }
        if j == len {
            return nil
        }
    }
    
    j = idx
    v := result
    
    for {
        j = (j+ 1) % len
        
        result = values[j]
        
        if result == nil {
            return v.value
        }
        
        k = h(values[j].key) % len
        
        if j > idx && (k <= idx || k > j) || (j < idx && (k <= idx && k > j)) {
            values[idx] = values[j]
            idx = j
        }
    }
    result = values[idx]
    values[idx] = nil
    return result.value
}



func (m *HashMap) Put(
        key collections.Key, 
        value collections.Value,
) collections.Value {
   
    
    h           := m.hashFunction
    values      := m.values
    j           := 0
    k           := h(key)
    len         := len(m.values) 
    
    
    var result *entry 
    values = checkBounds(values, m.len)
    
    
    for {
        code := (k + j) % len
        result := values[code]
        
        if result == nil {
            values[code] = &entry{
                key, 
                value,
            }
            break
        }
        
        
        if result.key == key {
            result.value = value
        } 
        j++
    }
    m.len += 1
    m.values = values
    return result
}



func checkBounds(
        values []*entry, 
        l int,
) []*entry {
    
    vl := len(values)
    
    if l >= vl {
        resize := make([]*entry, len(values)*2)
        return append(values, resize...)
    } 
    return values 
}



