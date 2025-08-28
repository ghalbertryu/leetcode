package heetcode

/*
以下為某個國家中銀行貸款後利率及付款的計算方式，若向銀行貸款20'000'000元，月利率為0.15%，若要在30年內(360)個月內還完
每月最少需要還多少錢，最後一個月會少還一些
*/
func findMinMonthlyRepayment() float64 {
	l := 0.0
	r := 20000000.0
	for r-l > 1 {
		m := (r + l) / 2
		if isMonthlyRepaymentValid(m) {
			r = m
		} else {
			l = m
		}
	}
	return r
}

func isMonthlyRepaymentValid(monthlyRepayment float64) bool {
	remaining := 20000000.0
	monthlyRate := 1.0015
	month := 360
	for i := 0; i < month; i++ {
		remaining = remaining*monthlyRate - monthlyRepayment
		if remaining <= 0 {
			return true
		}
	}
	return false
}
