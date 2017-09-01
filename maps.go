package anvil

import (
    "github.com/sunshower-io/anvil/maps/tree"
    "github.com/sunshower-io/anvil/collections"
)

func NewSortedMap(
        comparator func(
                collections.Value, 
                collections.Value,
        ) int,
) collections.SortedMap {
    return tree.NewTreeMap(comparator)
}
