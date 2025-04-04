package main

import (
	"errors"
)

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	var userMap = make(map[string]user)

	if len(names) != len(phoneNumbers) {
		err := errors.New("invalid sizes")
		if err != nil {
			return nil, err
		}
	}

	for index, name := range names {
		userMap[name] = user{
			name:        name,
			phoneNumber: phoneNumbers[index],
		}
	}

	return userMap, nil
}

type user struct {
	name        string
	phoneNumber int
}

/*

Assignment

We can speed up our contact-info lookups by using a map!

    Key-based map lookup: O(1)
    Slice brute-force search: O(n)

Complete the getUserMap function. It takes a slice of names and a slice of phone numbers, and returns a map of name ->
user structs and an error. A user struct just contains a user's name and phone number. The first name in the names slice
pairs with the first phone number, and so on.

If the length of names and phoneNumbers is not equal, return an error with the string "invalid sizes".
*/
