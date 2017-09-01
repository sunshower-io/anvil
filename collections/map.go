package collections

type Entry interface {
    Key() Key
    Value() Value
}

type Map interface {
    Clear() 
    
    ContainsKey(Key) bool
    
    ContainsValue(Value) bool
    
    EntrySet() Set
    
    Equals(Value) bool
    
    Get(Key) Value
    
    HashCode() int
    
    IsEmpty() bool
    
    KeySet() Set
    
    Put(Key, Value) Value
    
    PutAll(Map) 
    
    Remove(Key) Value
}



type SortedMap interface {
    
    Comparator() Comparator
    
    FirstKey() Key
    
    FirstValue() Value
    
    LastKey() Key
    
    LastValue() Value
    
    Last() Entry
    
    Range(Key, Key) SortedMap
    
}