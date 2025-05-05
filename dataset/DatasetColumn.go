package dataset

import (
	"fmt"
	"strconv"
)

type DataSetColumn interface {
	GetName() string
	GetType() string
	Length() int
	ValueAt(i int) string
	SetName(newName string)
}

// DataSetColumn.GetName()

func (c *String) GetName() string {
	return c.Name
}

func (c *Float) GetName() string {
	return c.Name
}

func (c *Integer) GetName() string {
	return c.Name
}

// DataSetColumn.GetType()

func (c *Float) GetType() string {
	return "float"
}

func (c *String) GetType() string {
	return "string"
}

func (c *Integer) GetType() string {
	return "int"
}

// DataSetColumn.Length()

func (c *Float) Length() int {
	return len(c.Data)
}

func (c *String) Length() int {
	return len(c.Data)
}
func (c *Integer) Length() int {
	return len(c.Data)
}

// DataSetColumn.ValueAt()

func (c *String) ValueAt(i int) string {

	return (c.Data)[i].Value
}

func (c *Integer) ValueAt(i int) string {
	return strconv.Itoa((c.Data)[i].Value)
}

func (c *Float) ValueAt(i int) string {
	return fmt.Sprintf("%f", (c.Data)[i].Value)
}

// DataSetColumn.SetName()

func (c *String) SetName(newName string) {
	c.Name = newName
}

func (c *Integer) SetName(newName string) {
	c.Name = newName
}

func (c *Float) SetName(newName string) {
	c.Name = newName
}
