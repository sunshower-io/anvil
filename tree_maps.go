package anvil

import (
    "github.com/sunshower-io/anvil/maps/tree"
    "github.com/sunshower-io/anvil/collections"
    "github.com/sunshower-io/anvil/maps"
)

type defaultMapFactory struct {
    
    
}


func (f defaultMapFactory) Hashes(
        strategy HashStrategy,
) HashMapBuilder {
    
    return defaultHashMapBuilder{
        hashStrategy: strategy,
    }
    
}








type defaultHashMapBuilder struct {
    useDefaultLoadFactor    bool
    loadFactor              float64
    hashStrategy            HashStrategy
}

func (d defaultHashMapBuilder) DefaultLoadFactor() CapacityStep {
    d.useDefaultLoadFactor = true
    return d
}


func (d defaultHashMapBuilder) InitialCapacity(
        capacity int,
) KeyStep {
    return d
}

func (d defaultHashMapBuilder) KeyedBy(
        f func (collections.Value) int,
) collections.Map {
    return maps.NewDefaultLinearProbing(f)
}





// default map factory


func (f defaultMapFactory) Sorted(
        strategy SortingStrategy,
) SortedMapBuilder {
    return defaultSortedMapBuilder{strategy: strategy}
}

type defaultSortedMapBuilder struct {
    strategy SortingStrategy
}

func (d defaultSortedMapBuilder) OrderBy(cmp func(
        collections.Value,
        collections.Value,
) int) collections.SortedMap {
    
    switch d.strategy {
    case AVL:
    case RedBlack:
        return tree.NewTreeMap(cmp)
    }
    return nil
}
