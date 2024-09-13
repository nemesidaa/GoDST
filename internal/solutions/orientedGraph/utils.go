package orientedgraph

import "errors"

type State byte

const (
	Ready State = iota
	NotReady
)

var (
	errStateNotReady = errors.New("state not ready")
)

// После этого капасити скажет что не хочет жить...
func removeElement[T int | *GraphNode](slice []T, values ...T) []T {
	for _, value := range values {
		for idx, v := range slice {
			if v == value {
				slice = append(slice[:idx], slice[idx+1:]...)
			}
		}
	}
	return slice
}

func max[T int8 | int16 | int32 | int64 | int](values ...T) T {
	var max T

	for _, value := range values {
		if value > max {
			max = value
		}
	}
	return max
}
