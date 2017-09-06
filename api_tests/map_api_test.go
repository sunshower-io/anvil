package api_tests

import (
    "testing"
    "github.com/sunshower-io/anvil"
    "github.com/stretchr/testify/assert"
    "github.com/sunshower-io/anvil/f/strings"
    "github.com/sunshower-io/anvil/collections"
)

func TestApiForBuildingRedBlackTreeMakesSense(t *testing.T) {
    cmp := func(lhs, rhs collections.Value) int {
        return 0
    }
    rbtree := anvil.Maps().Sorted(anvil.Trees.RedBlack).OrderBy(cmp)
    assert.NotNil(t, rbtree)
}


func TestApiForBuildingLinearProbingHashMapMakesSense(t *testing.T) {
    
    hash := anvil.
        Maps().
        Hashes(anvil.Hashing.LinearProbing).
        DefaultLoadFactor().
        InitialCapacity(10).
        KeyedBy(strings.DefaultHash)
    
    assert.NotNil(t, hash)
}



