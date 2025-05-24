package module1

func formatSlice(slice []string) string {
	result := ""
	for index, item := range slice {
		result += item
		if index < len(slice)-1 {
			result += ", "
		}
	}
	return result
}
