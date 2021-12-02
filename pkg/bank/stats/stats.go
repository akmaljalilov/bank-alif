package stats

import "github.com/akmaljalilov/alif-academy/pkg/bank/types"

func Avg(payments []types.Payment) types.Money {
	ser := types.Money(0)
	for _, payment := range payments {
		ser += payment.Amount
	}
	return types.Money(int(ser) / len(payments))
}
