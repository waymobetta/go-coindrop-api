package reddit

// removeDuplicates removes duplicate subreddits from slice
func removeDuplicates(slice []string) []string {
	// initialize map to store unique elements
	nonUniqueMap := make(map[string]bool)

	// create map of all unique elements
	for i := range slice {
		nonUniqueMap[slice[i]] = true
	}

	// initialize slice to store unique elements
	var uniqueSlice []string

	// iterate over mapping and place all keys in slice
	for key := range nonUniqueMap {
		uniqueSlice = append(uniqueSlice, key)
	}

	return uniqueSlice
}