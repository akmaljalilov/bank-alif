package stats

import (
	"fmt"
	"github.com/akmaljalilov/alif-academy/pkg/bank/types"
)

func ExampleAvg() {
	fmt.Println(Avg([]types.Payment{
		{Amount: 10},
		{Amount: 20},
		{Amount: 30},
		{Amount: 60},
	}))
	// Output:
	// 30

}
