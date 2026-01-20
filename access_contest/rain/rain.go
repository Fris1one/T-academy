package rain

import (
	"fmt"
)

func Rain() {
	var n, m, result int
	fmt.Scan(&n, &m)
	days_in_month := m + 14 - n
	if m+7 > days_in_month {
		result = 7 - (days_in_month - m)
	} else {
		result = m + 7
	}
	fmt.Println(result)
}