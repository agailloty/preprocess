package operations

type operationRunner interface {
	run() []operationResult
	isIndependant() bool
}
