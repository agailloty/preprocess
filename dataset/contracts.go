package dataset

type DataSetColumn interface {
	GetName() string
	GetType() string
}
