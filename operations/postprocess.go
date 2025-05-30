package operations

import (
	"log"
	"time"

	"github.com/agailloty/preprocess/config"
	"github.com/agailloty/preprocess/dataset"
)

func operateSplit(df *dataset.DataFrame, split *config.DataSetSplit) {
	if split.Method == METHOD_TRAIN_TEST_SPLIT {
		operateTrainTestSplit(df, split)
	}
}

func operateTrainTestSplit(df *dataset.DataFrame, split *config.DataSetSplit) {
	seed := uint64(time.Now().UnixMilli())
	if split == nil {
		return
	}

	if split.RandomSeed != nil {
		seed = *split.RandomSeed
	}

	ratio := float64(0.7)

	if split.TrainTestSplitRatio != nil {
		ratio = *split.TrainTestSplitRatio
	}

	train, test := "", ""

	if split.SplitNames != nil {
		if len(*split.SplitNames) != 2 {
			log.Fatal("[Train Test Split Operation] You must specify exactly 2 split names")
		}
		train, test = (*split.SplitNames)[0], (*split.SplitNames)[1]
	}

	splitFunc := func(df *dataset.DataFrame, args ...any) []dataset.SplitSpec {
		return splitUsingRandom(df, seed, ratio, train, test)
	}

	df.SaveSplittedDataframeToCSV(splitFunc)
}
