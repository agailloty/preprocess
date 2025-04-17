package dataset

import "fmt"

func DisplayColumn(column DataSetColumn, n int) {
	switch v := column.(type) {
	case Float:
		fmt.Printf("%s (Float) \n", v.Name)
		for i := range n {
			fmt.Printf("%.2f\n", v.Data[i])
		}
	case String:
		fmt.Printf("%s (Float)", v.Name)
		for i := range n {
			fmt.Printf("%s", v.Data[i])
		}
	case Integer:
		fmt.Printf("%s (Float)", v.Name)
		for i := range n {
			fmt.Printf("%d", v.Data[i])
		}
	}
}
