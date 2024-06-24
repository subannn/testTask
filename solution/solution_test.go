package solution

import (
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test1Day(t *testing.T) {
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

	assert.Equal(t, res, answer, "days")
}

func Test2Week(t *testing.T) {
	input := []*Transaction{
		{
			4456,
			time.Unix(0, 0).UTC(),
		},
		{
			4231,
			time.Unix(50, 0).UTC(),
		},
		{
			5212,
			time.Unix(60, 0).UTC(),
		},
		{
			4321,
			time.Unix(123123, 0).UTC(),
		},
		{
			999999,
			time.Unix(624831, 0).UTC(),
		},
	}

	answer := []Transaction{
		{
			4321,
			time.Unix(0, 0).UTC(),
		},
		{
			999999,
			time.Unix(604800, 0).UTC(),
		},
	}
	res, _ := FormatTransactions(input, "week")

	sort.Slice(answer, func(i, j int) bool {
		return answer[i].Timestamp.Before(answer[j].Timestamp)
	})
	sort.Slice(res, func(i, j int) bool {
		return res[i].Timestamp.Before(res[j].Timestamp)
	})

	assert.Equal(t, res, answer, "weeks")
}

func Test3Hour(t *testing.T) {
	input := []*Transaction{
		{
			4456,
			time.Unix(0, 0).UTC(),
		},
		{
			9999,
			time.Unix(3599, 0).UTC(),
		},
		{
			777,
			time.Unix(3999, 0).UTC(),
		},
	}

	answer := []Transaction{
		{
			9999,
			time.Unix(0, 0).UTC(),
		},
		{
			777,
			time.Unix(3600, 0).UTC(),
		},
	}
	res, _ := FormatTransactions(input, "hour")

	sort.Slice(answer, func(i, j int) bool {
		return answer[i].Timestamp.Before(answer[j].Timestamp)
	})
	sort.Slice(res, func(i, j int) bool {
		return res[i].Timestamp.Before(res[j].Timestamp)
	})

	assert.Equal(t, res, answer, "hours")
}

func Test4Month(t *testing.T) {
	input := []*Transaction{
		{
			4456,
			time.Unix(0, 0).UTC(),
		},
		{
			9999,
			time.Unix(3599, 0).UTC(),
		},
		{
			777,
			time.Unix(3999, 0).UTC(),
		},
		{
			1,
			time.Unix(2592001*13, 0).UTC(),
		},
		{
			1,
			time.Unix(2592000*45, 0).UTC(),
		},
	}

	answer := []Transaction{
		{
			777,
			time.Unix(0, 0).UTC(),
		},
		{
			1,
			time.Unix(2592000*13, 0).UTC(),
		},
		{
			1,
			time.Unix(2592000*45, 0).UTC(),
		},
	}
	res, _ := FormatTransactions(input, "month")

	sort.Slice(answer, func(i, j int) bool {
		return answer[i].Timestamp.Before(answer[j].Timestamp)
	})
	sort.Slice(res, func(i, j int) bool {
		return res[i].Timestamp.Before(res[j].Timestamp)
	})

	assert.Equal(t, res, answer, "month")
}
