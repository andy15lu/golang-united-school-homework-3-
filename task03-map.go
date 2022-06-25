package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	result = make([]string, len(input))
	var keys = make([]int, len(input))
	i := 0
	for key, _ := range input {
		keys[i] = key
		i++
	}
	sort.Slice(keys, func(a, b int) bool {
		return a > b
	})
	for i := 0; i < len(input); i++ {
		result[i] = input[keys[i]]
	}
	return
}
