package problems

import "math"

func Max[V int64 | float64](s []V) V {
	var max V = math.MinInt64
	for _, v := range s {
		if v > max {
			max = v
		}
	}
	return max
}

func Map[T, U any](s []T, f func(T) U) []U {
	// Make a new slice of Us with length of our input slice
	r := make([]U, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

func Reduce[T, U any](s []T, init U, f func(U, T) U) U {
	r := init
	for _, v := range s {
		r = f(r, v)
	}
	return r
}

func Filter[T any](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}
