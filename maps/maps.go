package maps

import (
    "github.com/sunshower-io/anvil/maps/hash"
    "github.com/sunshower-io/anvil/maps/tree"
    "github.com/sunshower-io/anvil/collections"
)

func NewDefaultLinearProbing(
        function hash.HashFunction,
) collections.Map {
    return hash.NewLinearProbeHashMap(function)
}

func NewSortedMap(
        comparator func(
                collections.Value,
                collections.Value,
        ) int,
) collections.SortedMap {
    return tree.NewTreeMap(comparator)
}
