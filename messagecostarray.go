package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	originalMessages := [3]string{primary, secondary, tertiary}

	primaryCost := len(primary)
	secondaryCost := primaryCost + len(secondary)
	tertiaryCost := secondaryCost + len(tertiary)

	totalMessageCost := [3]int{primaryCost, secondaryCost, tertiaryCost}

	return originalMessages, totalMessageCost
}

/*
When our clients don't respond to a message, they can be reminded with up to 2 additional messages.

Complete the getMessageWithRetries function. It takes three strings and returns:

    An array of 3 strings
    An array of 3 integers

The returned string array contains the original messages. The first is the primary message, the second is the first reminder, and the third is the last reminder.

The integers in the integer array represent the cost of sending each message. The cost of each message is equal to the length of the message, plus the length of any previous messages. For example:

    "hello" costs 5
    "world" costs 5, adding "hello" makes total cost 10 (5 + 5)
    "!" costs 1, adding the previous messages makes total cost 11 (5 + 5 + 1)

*/
