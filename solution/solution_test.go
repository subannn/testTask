package solution

import (
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParsePath(t *testing.T) {

	input := []*Transaction{
		{
			4456,
			time.Unix(1616026248, 0).UTC(),
		},
		{
			4231,
			time.Unix(1616022648, 0).UTC(),
		},
		{
			5212,
			time.Unix(1616019048, 0).UTC(),
		},
		{
			4321,
			time.Unix(1615889448, 0).UTC(),
		},
		{
			4567,
			time.Unix(1615871448, 0).UTC(),
		},
	}

	answer := []Transaction{
		{
			4456,
			time.Unix(1616025600, 0).UTC(),
		},
		{
			4231,
			time.Unix(1615939200, 0).UTC(),
		},
		{
			4321,
			time.Unix(1615852800, 0).UTC(),
		},
	}
	res, _ := FormatTransactions(input, "day")

	sort.Slice(answer, func(i, j int) bool {
		return answer[i].Timestamp.Before(answer[j].Timestamp)
	})
	sort.Slice(res, func(i, j int) bool {
		return res[i].Timestamp.Before(res[j].Timestamp)
	})

	assert.Equal(t, res, answer, "The two words should be the same.")
}
