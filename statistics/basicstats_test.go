package statistics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStats(t *testing.T) {

	ages := []int{15, 20, 30, 45, 25, 30, 20, 35}

	t.Run("Testing mean of integer values", func(t *testing.T) {
		ageMean := Mean(ages)
		assert.Equal(t, ageMean, 27.5)
	})

	ages = []int{}

	t.Run("Mean with empty slice returns 0", func(t *testing.T) {
		ageMean := Mean(ages)
		assert.Equal(t, ageMean, float64(0))
	})

}
