package dataset

type DataSetColumn interface {
	GetName() string
	GetType() string
	Length() int
	ValueAt(i int) string
	SetName(newName string)
}
