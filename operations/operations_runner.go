package operations

type operationRunner interface {
	run() []operationResult
}

func runAllOperations(operations []operationRunner) {
	for _, op := range operations {
		op.run()
	}
}
