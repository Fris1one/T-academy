package kid_stairs

import (
	"fmt"
)

func Stairs() {
	var n, k int
	fmt.Scan(&n, &k)

	mood := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&mood[i])
	}

	const negative = -101
	mood_sum := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		mood_sum[i] = make([]int, k+1)
		for j := 0; j <= k; j++ {
			mood_sum[i][j] = negative
		}
	}
	for j := 0; j <= k; j++ {
		mood_sum[0][j] = 0
	}

	for i := 1; i <= n; i++ {
		for j := 0; j <= k; j++ {
			if j == 0 {
				if i >= 2 {
					if mood_sum[i-2][j] > mood_sum[i-1][j] {
						mood_sum[i][j] = mood_sum[i-2][j] + mood[i-1]
					} else {
						mood_sum[i][j] = mood_sum[i-1][j] + mood[i-1]
					}
				} else {
					mood_sum[i][j] = mood_sum[i-1][j] + mood[i-1]
				}
			} else {
				if mood[i-1] < 0 && i != n{
					mood_sum[i][j] = mood_sum[i][0]
				} else {
					var maxPrev = -101
					for prev_i := 0; prev_i != i; prev_i++ {
						if mood_sum[prev_i][j-1] > maxPrev {
							maxPrev = mood_sum[prev_i][j-1]
						}
					}
					if maxPrev > mood_sum[i-1][j] && maxPrev > mood_sum[i-2][j] {
						mood_sum[i][j] = maxPrev + mood[i-1]
					} else if i >= 2 {
						if mood_sum[i-2][j] > mood_sum[i-1][j] {
							mood_sum[i][j] = mood_sum[i-2][j] + mood[i-1]
						} else {
							mood_sum[i][j] = mood_sum[i-1][j] + mood[i-1]
						}
					} else {
						mood_sum[i][j] = mood_sum[i-1][j] + mood[i-1]
					}
				}
			}
			
		}
	}
	result := negative
	for j := 0; j <= k; j++ {
		if mood_sum[n][j] > result {
			result = mood_sum[n][j]
		}
	}
	fmt.Println(result)
}