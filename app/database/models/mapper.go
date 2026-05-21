package models

// Mappable is implemented by every model that can be returned over the wire.
// Call ToResponse() before passing to utils.OK or utils.Created.
type Mappable interface {
	ToResponse() any
}

// MapSlice converts a typed slice of Mappable models into []any for wire serialization.
func MapSlice[T Mappable](items []T) []any {
	out := make([]any, len(items))
	for i, item := range items {
		out[i] = item.ToResponse()
	}
	return out
}
