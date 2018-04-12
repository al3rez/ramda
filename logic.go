package ramda

import (
	"fmt"
	"strings"
)

// Add returns true if both arguments are true; false otherwise.
func And(a, b bool) bool {
	return a && b
}

// Cond used to indicate list of [predicate, transformer] pairs.
type Cond []func() interface{}

// NewCond returns a function, fn, which encapsulates if/else, if/else, ...
// logic.  ramda.NewCond takes a Cond{[predicate, transformer], ...} struct. All
// of the arguments to fn are applied to each of the predicates in turn until
// one returns a "truthy" value, at which point fn returns the result of
// applying its arguments to the corresponding transformer. If none of the
// predicates matches, fn returns undefined.
func NewCond(pairs Cond) func(interface{}) interface{} {
	return func(i interface{}) interface{} {
		for j := 0; j < len(pairs); j++ {
			predicate, transformer := pairs[j], pairs[j+1]
			if predicate() == i || predicate() == true {
				switch v := transformer().(type) {
				case string:
					if strings.Contains(v, "%") {
						return fmt.Sprintf(v, i)
					}

					return v
				default:
					return v
				}
			}

		}
		return nil
	}
}

// T a function that always returns true. Any passed in parameters are ignored.
func T() func() interface{} {
	return func() interface{} {
		return true
	}
}
