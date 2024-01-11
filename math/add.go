package math

import "math"

func Add(args ...any) any {
	if len(args) <= 0 {
		return nil
	}
	if len(args) == 1 {
		return args[0]
	}
	return math.MaxInt64
}
