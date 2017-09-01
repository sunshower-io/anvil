package collections

type Collection interface {
    
    Clear() 
    
    Add(Value) bool 
    
    AddAll(Collection) bool
    
    Contains(Value) bool
    
    ContainsAll(Value)
    
    Equals(interface{}) bool
    
    HashCode() int
    
    IsEmpty() bool
    
    Iterator() Iterator
    
    Remove(Value) bool
    
    RemoveAll(Collection) bool
    
    RetainAll(Collection) bool
    
    Size() int
    
    ToArray() []Value
    
    Into([]Value) []Value
    
}
