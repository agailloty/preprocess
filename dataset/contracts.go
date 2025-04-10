package dataset

type DataSetColumn interface {
	Parse() error
}
