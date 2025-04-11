package dataset

type String struct {
	Name string
	Data []string
}

type Integer struct {
	Name string
	Data []int
}

type Float struct {
	Name string
	Data []float32
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
