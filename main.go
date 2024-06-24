package main

import (
	"fmt"
	"time"

	"github.com/subannn/testTask/solution"
)

func main() {
	input := []*solution.Transaction{
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

	// answer := []solution.Transaction{
	// 	{
	// 		4456,
	// 		time.Unix(1616025600, 0).UTC(),
	// 	},
	// 	{
	// 		4231,
	// 		time.Unix(1615939200, 0).UTC(),
	// 	},
	// 	{
	// 		4321,
	// 		time.Unix(1615852800, 0).UTC(),
	// 	},
	// }
	res1, _ := solution.FormatTransactions(input, "month")
	res2, _ := solution.FormatTransactions(input, "week")
	res3, _ := solution.FormatTransactions(input, "day")
	res4, _ := solution.FormatTransactions(input, "hosur")


	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)
	fmt.Println(res4)

}
