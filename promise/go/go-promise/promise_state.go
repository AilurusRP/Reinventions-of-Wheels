package go_promise

type PromiseState int8

const (
	PENDING  PromiseState = 0
	RESOLVED PromiseState = 1
	REJECTED PromiseState = 2
)
