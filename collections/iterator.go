package collections


type Iterator interface {
    
    HasNext() bool 
    
    Next() (Value, error)

}

type Iterable interface {
    Iterator() Iterator 
    
    Range() <- chan Value
}