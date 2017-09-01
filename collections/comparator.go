package collections



type Comparator interface {
    Compare(
            Value,
            Value,
    ) int
}

type FunctionComparator struct {
    Comparator
    
    Ord func(Value, Value) int
}

func (c *FunctionComparator) Compare(lhs, rhs Value) int {
    return c.Ord(lhs, rhs)
}
