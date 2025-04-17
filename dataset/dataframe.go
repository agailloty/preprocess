package dataset

import "fmt"

func DisplayColumn(column DataSetColumn, n int) {
	switch v := column.(type) {
	case Float:
		fmt.Printf("%s (float) \n", v.Name)
		for i := range n {
			fmt.Printf("%.2f ", v.Data[i])
		}
	case String:
		fmt.Printf("%s (string) \n", v.Name)
		for i := range n {
			fmt.Printf("%s ", v.Data[i])
		}
	case Integer:
		fmt.Printf("%s (int) \n", v.Name)
		for i := range n {
			fmt.Printf("%d ", v.Data[i])
		}
	}
}
