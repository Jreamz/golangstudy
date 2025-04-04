package main

func countConnections(groupSize int) int {
	totalConnections := 0

	for num := 1; num <= groupSize; num++ {
		newConnections := num - 1
		if groupSize == 1 {
			return totalConnections
		}
		totalConnections += newConnections
	}
	return totalConnections
}

/*
Complete the countConnections function that takes an integer groupSize representing the number of people in the group
chat and returns an integer representing the number of connections between them. For each additional person in the
group, the number of new connections continues to grow. Use a for loop to accumulate the number of connections instead
of directly using a mathematical formula.
*/
