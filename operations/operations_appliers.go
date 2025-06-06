package operations

func runAllOperations(operations []operationRunner) {
	for _, op := range operations {
		op.run()
	}
}
