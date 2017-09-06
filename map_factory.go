package anvil

import (
    "github.com/sunshower-io/anvil/collections"
)

type HashStrategy int 


const (
    LinearProbing   HashStrategy = 0 
    
)


type HashingStrategies struct {
    Chained             HashStrategy
    Cuckoo              HashStrategy
    LinearProbing       HashStrategy
    QuadraticProbing    HashStrategy
}



type SortingStrategy int

const (
    AVL SortingStrategy             = 0
    RedBlack SortingStrategy        = 1
)



type SortingStrategies struct {
    AVL      SortingStrategy 
    RedBlack SortingStrategy
}

type KeyStep interface {
    KeyedBy(func(collections.Value) int) collections.Map
}

type CapacityStep interface {
    InitialCapacity(int) KeyStep
    
}


type HashMapBuilder interface {
    DefaultLoadFactor() CapacityStep
}





type SortedMapBuilder interface {
    OrderBy(func(
            collections.Value,
            collections.Value,
    ) int) collections.SortedMap
}

var Trees = &SortingStrategies{
    AVL: AVL, 
    RedBlack: RedBlack,
}

var Hashing = &HashingStrategies {
    
}

type MapFactory interface {
    
    Hashes(HashStrategy) HashMapBuilder
    
    Sorted(SortingStrategy) SortedMapBuilder
    
    
}

func Maps() MapFactory {
    return defaultMapFactory{} 
}

