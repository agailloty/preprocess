package dataset

type String struct {
	name string
	data []string
}

type Integer struct {
	name string
	data []int
}

func (s String) Parse() error {
	return nil
}

func (i Integer) Parse() error {
	return nil
}
