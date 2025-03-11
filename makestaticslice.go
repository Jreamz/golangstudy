package main

func getMessageCosts(messages []string) []float64 {
	messagesTotal := make([]float64, len(messages))

	for index, message := range messages {
		messageCost := (float64(len(message))) * 0.01
		messagesTotal[index] += messageCost
	}
	return messagesTotal
}
