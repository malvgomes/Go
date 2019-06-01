package mathematics

import (
	"fmt"
	"strconv"
)

// Avg ...
func Avg(nums ...float64) float64 {
	total := 0.0
	for _, num := range nums {
		total += num
	}
	avg := total / float64(len(nums))
	rndAvg, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", avg), 64)
	return rndAvg
}
