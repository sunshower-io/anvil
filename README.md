

Anvil is a curated collection of fast generic algorithms and data-structures for the Go programming language.  If it is missing
an algorithm or data-structure that you need, please feel free to submit a request-for-enhancement with the requirements.  

Usage:

# Maps


## Sorted Maps

### Red-black tree backed map

```go 

    cmp := func(lhs, rhs collections.Value) int {
        return lhs.(int) - rhs.(int)
    }
    
    
    rbtree := anvil.Maps().Sorted(anvil.Trees.RedBlack).OrderBy(cmp)

```

## Unsorted maps

```go

    hash := anvil.
        Maps().
        Hashes(anvil.Hashing.LinearProbing).
        DefaultLoadFactor().
        InitialCapacity(10).
        KeyedBy(strings.DefaultHash)


```


# Iteration

## State-based iteration

State-based iteration is fast but less idiomatic than range-based iteration

```go

for it := collection.Iterator(); it.HasNext(); {
    v, e := it.Next()
    
    if e != nil {
    ///cast v to desired type and use it
    }

}

```



## range-based iteration

Range-based iteration is the idiomatic way to iterate over collections

```go

for v := range collection {
// cast v to desired type and use it.  
}

```
# Hash-functions

Anvil provides a collection of high-quality hash-functions that can be used with its collections or in your own code

Provided functions:

- `strings.DefaultHashCode()`: product-sum hashcode for strings
- `refs.ReferenceHashCode()`: provides referential equality for values 

