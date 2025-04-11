package dataset

type String struct {
	name string
	data []string
}

type Integer struct {
	name string
	data []int
}

type Float struct {
	name string
	data []float32
}

func (s String) Parse() error {
	return nil
}

func (i Integer) Parse() error {
	return nil
}

func (f Float) Parse() error {
	return nil
}
