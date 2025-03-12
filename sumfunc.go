package main

func sum(nums ...int) int {
	sumTotal := 0
	for _, num := range nums {
		sumTotal += num
	}
	return sumTotal
}

/*
We need to sum up the costs of all individual messages so we can send an end-of-month bill to our customers.
Complete the sum function to return the sum of all inputs.
*/
