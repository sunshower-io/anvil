package maps

import (
    "github.com/sunshower-io/anvil/maps/hash"
    "github.com/sunshower-io/anvil/collections"
)

func DefaultLinearProbingMap(
        function hash.HashFunction,
) collections.Map {
    return hash.NewLinearProbeHashMap(function)
} 
