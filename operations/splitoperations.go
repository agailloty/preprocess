package operations

import (
	"math/rand/v2"

	"github.com/agailloty/preprocess/dataset"
)

func splitUsingRandom(df *dataset.DataFrame, seed uint64, ratio float64, train string, test string) []dataset.SplitSpec {
	if train == "" {
		train = df.Name + "_train"
	}
	if test == "" {
		test = df.Name + "_test"
	}

	s1, s2 := generateRandomSlices(df.RowsCount, seed, ratio)
	result := make([]dataset.SplitSpec, 2)

	result[0] = dataset.SplitSpec{Name: train, Rows: s1}
	result[1] = dataset.SplitSpec{Name: test, Rows: s2}

	return result

}

func generateRandomSlices(N int, seed uint64, ratio float64) ([]int, []int) {
	r := rand.New(rand.NewPCG(seed, seed+1))

	numbers := make([]int, N)
	for i := 0; i < N; i++ {
		numbers[i] = i + 1
	}

	r.Shuffle(N, func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})

	countS1 := int(float64(N) * ratio)
	S1 := numbers[:countS1]
	S2 := numbers[countS1:]
	return S1, S2
}
