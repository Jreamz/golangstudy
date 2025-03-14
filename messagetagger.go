package main

import "strings"

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	for index, message := range messages {
		messages[index].tags = tagger(message)
	}
	return messages
}

func tagger(msg sms) []string {
	tags := []string{}
	message := strings.ToLower(msg.content)

	if strings.Contains(message, "urgent") {
		tags = append(tags, "Urgent")
	}
	if strings.Contains(message, "sale") {
		tags = append(tags, "Promo")
	}

	return tags
}

/*
Message Tagger

Textio needs a way to tag messages based on specific criteria.
Assignment

    Complete the tagMessages function. It should take a slice of sms messages, and a function (that takes a sms as input
	and returns a slice of strings) as inputs. And it should return a slice of sms messages.

    It should loop through each message and set the tags to the result of the passed in function.
    Be sure to modify the messages of the original slice.
    Ensure the tags field contains an initialized slice. No nil slices.

    Complete the tagger function. It should take an sms message and return a slice of strings.

    For any message that contains "urgent" (regardless of casing) in the content, the Urgent tag should be applied first.
    For any message that contains "sale" (regardless of casing), the Promo tag should be applied second.

Note: regardless of casing just means that even "uRGent" or "SALE" should trigger the tag.
*/
