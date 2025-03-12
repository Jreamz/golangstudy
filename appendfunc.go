package main

type cost struct {
	day   int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {
	dayCosts := make([]float64, 0)
	for _, cost := range costs {
		if cost.day == day {
			dayCosts = append(dayCosts, cost.value)
		}
	}
	return dayCosts
}

/*
Assignment
We've been asked to add a feature that extracts costs for a given day.

Complete the getDayCosts() function using the append() function. It accepts a slice of cost structs and a day int,
and it returns a float64 slice containing that day's costs. If there are no costs for that day, return an empty slice.
*/
