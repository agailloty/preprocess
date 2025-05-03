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

// DataSetColumn.GetName()

func (c String) GetName() string {
	return c.Name
}

func (c Float) GetName() string {
	return c.Name
}

func (c Integer) GetName() string {
	return c.Name
}

// DataSetColumn.GetType()

func (c Float) GetType() string {
	return "float"
}

func (c String) GetType() string {
	return "string"
}

func (c Integer) GetType() string {
	return "int"
}

// DataSetColumn.Length()

func (c Float) Length() int {
	return len(c.Data)
}

func (c String) Length() int {
	return len(c.Data)
}
func (c Integer) Length() int {
	return len(c.Data)
}
