package main

func getMessageCosts(messages []string) []float64 {
	messagesTotal := make([]float64, len(messages))

	for index, message := range messages {
		messageCost := (float64(len(message))) * 0.01
		messagesTotal[index] += messageCost
	}
	return messagesTotal
}

/*
Complete the getMessageCosts() function. It takes a slice of messages and returns a slice of message costs.
    Preallocate a slice for the message costs of the same length as the messages slice.
    Fill the costs slice with costs for each message. The cost in the cost slice should correspond to the message
in the messages slice at the same index. The cost of a message is the length of the message multiplied by 0.01.
*/
