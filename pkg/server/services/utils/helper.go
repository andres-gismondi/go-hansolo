package utils

func ConcatenateSlices(messages [][]string) ([]string, bool) {
	var cleanSlice []string

	anotherSlice := InitializeSlice(len(messages[0]))
	for _, str := range messages {
		cleanSlice = append(cleanSlice, str...)
		for i := range(str) {
			if str[i] != "" {
				anotherSlice[i] = str[i]
			}
		}
	}
	return anotherSlice, FindEmptyItem(anotherSlice)
}

func InitializeSlice(length int)  []string{
	conditions := make([]string, length)
	for i:=0; i<length; i++ {
		conditions[i] = ""
	}
	return conditions
}

func FindEmptyItem(slice []string) bool {
	for _, item := range slice {
		if item == "" {
			return false
		}
	}
	return true
}
