package domain

/*
need for:
- field not provided
- field provided: value
- field provided: null
*/
type Nullable[T any] struct {
	Value *T
	Set   bool
}
