package dataset

type Nullable[T any] struct {
	IsValid bool
	Value   T
}

type String struct {
	Name string
	Data []Nullable[string]
}

type Integer struct {
	Name string
	Data []Nullable[int]
}

type Float struct {
	Name string
	Data []Nullable[float32]
}
