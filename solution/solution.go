package solution

import (
	"errors"
	"sort"
	"time"
)

type Transaction struct {
	Value     int
	Timestamp time.Time
}

func FormatTransactions(transactions []*Transaction, interval string) ([]Transaction, error) {
	sort.Slice(transactions, func(i, j int) bool {
		return transactions[i].Timestamp.Before(transactions[j].Timestamp)
	})

	if interval == "month" {
		return formatToMonth(transactions), nil
	} else if interval == "week" {
		return formatToWeek(transactions), nil
	} else if interval == "day" {
		return formatToDay(transactions), nil
	} else if interval == "hour" {
		return formatToHour(transactions), nil
	}

	return nil, errors.New("Incorrect interval name")
}
func formatToMonth(transactions []*Transaction) []Transaction {
	ans := make([]Transaction, 0)
	for i := 0; i < len(transactions)-1; i++ {
		if transactions[i].Timestamp.Month() != transactions[i+1].Timestamp.Month() {
			ans = append(ans, *transactions[i])
		}
	}

	ans = append(ans, *transactions[len(transactions)-1])

	result := make([]Transaction, 0)
	for _, val := range ans {
		formated := Transaction{val.Value, time.Date(val.Timestamp.Year(), val.Timestamp.Month(), 1, 0, 0, 0, 0, time.UTC)}
		result = append(result, formated)
	}

	return result
}
func formatToWeek(transactions []*Transaction) []Transaction {
	ans := make([]Transaction, 0)
	for i := 0; i < len(transactions)-1; i++ {

		week1 := transactions[i].Timestamp.Unix() / 604800
		week2 := transactions[i+1].Timestamp.Unix() / 604800

		if week1 != week2 {
			ans = append(ans, *transactions[i])
		}
	}

	ans = append(ans, *transactions[len(transactions)-1])

	result := make([]Transaction, 0)
	for _, val := range ans {
		weekUNIX := (val.Timestamp.Unix() / 604800) * 604800
		t := time.Unix(weekUNIX, 0)
		formated := Transaction{val.Value, t}
		result = append(result, formated)
	}

	return ans
}
func formatToDay(transactions []*Transaction) []Transaction {
	ans := make([]Transaction, 0)
	for i := 0; i < len(transactions)-1; i++ {
		if transactions[i].Timestamp.Day() != transactions[i+1].Timestamp.Day() {
			ans = append(ans, *transactions[i])
		}
	}

	ans = append(ans, *transactions[len(transactions)-1])

	result := make([]Transaction, 0)
	for _, val := range ans {
		formated := Transaction{val.Value, time.Date(val.Timestamp.Year(), val.Timestamp.Month(), val.Timestamp.Day(), 0, 0, 0, 0, time.UTC)}
		result = append(result, formated)
	}

	return result
}
func formatToHour(transactions []*Transaction) []Transaction {
	ans := make([]Transaction, 0)
	for i := 0; i < len(transactions)-1; i++ {
		if transactions[i].Timestamp.Hour() != transactions[i+1].Timestamp.Hour() {
			ans = append(ans, *transactions[i])
		}
	}

	ans = append(ans, *transactions[len(transactions)-1])

	result := make([]Transaction, 0)
	for _, val := range ans {
		formated := Transaction{val.Value, time.Date(val.Timestamp.Year(), val.Timestamp.Month(), val.Timestamp.Day(), val.Timestamp.Hour(), 0, 0, 0, time.UTC)}
		result = append(result, formated)
	}

	return result
}
