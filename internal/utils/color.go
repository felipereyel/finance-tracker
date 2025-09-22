package utils

// deterministicly map a string to a color
// blue | green | red | yellow | purple | pink | gray
func StringToColor(value string) string {
	colors := []string{"gray", "green", "purple", "yellow", "red", "pink", "blue"}
	var hash int
	for _, char := range value {
		hash += int(char)
	}

	return colors[hash%len(colors)]
}
