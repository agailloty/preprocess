package operations

type operationRunner interface {
	run() []operationResult
}
