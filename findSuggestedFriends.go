package main

import "fmt"

func findSuggestedFriends(username string, friendships map[string][]string) []string {
	suggestedFriends := make([]string, 0)
	usernameFriends := friendships[username]
	for _, friend := range usernameFriends {
		for _, ff := range friendships[friend] {
			if ff != username && !contains(usernameFriends, ff) && !contains(suggestedFriends, ff) {
				suggestedFriends = append(suggestedFriends, ff)
			}
		}
	}
	if len(suggestedFriends) == 0 {
		return nil
	}
	return suggestedFriends
}

func contains(slice []string, word string) bool {
	for _, value := range slice {
		if value == word {
			return true
		}
	}
	return false
}

func main() {
	friendships := map[string][]string{
		"Alice":   {"Bob", "Charlie"},
		"Bob":     {"Alice", "Charlie", "David"},
		"Charlie": {"Alice", "Bob", "David", "Eve"},
		"David":   {"Bob", "Charlie"},
		"Eve":     {"Charlie"},
	}
	suggestedFriends := findSuggestedFriends("Alice", friendships)
	fmt.Println(suggestedFriends)
}
