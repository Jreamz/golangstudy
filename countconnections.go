package main

func countConnections(groupSize int) int {
	totalConnections := 0
	for num := 0; num <= groupSize; num++ {
		if num == 1 {
			break
		}
		totalConnections += num * num / groupSize
	}
	return totalConnections
}
