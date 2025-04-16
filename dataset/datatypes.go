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

func (s String) GetName() string {
	return s.Name
}

func (s String) GetType() string {
	return "string"
}

func (f Float) GetName() string {
	return f.Name
}

func (f Float) GetType() string {
	return "float"
}

func (i Integer) GetName() string {
	return i.Name
}

func (i Integer) GetType() string {
	return "int"
}
