package strings

import (
    "math"
    "github.com/sunshower-io/anvil/collections"
)

func DefaultHash(value collections.Value) int {
    j := 0
    v := value.(string)
    l := len(v)
    for i := 0; i < l; i++ {
        j += int(v[i]) * int(math.Pow(31, float64(l - 1 - i)))
    }
    return j
}
