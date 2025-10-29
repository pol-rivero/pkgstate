package common

import (
	"slices"
)

func Sorted(list []string) []string {
	slices.Sort(list)
	return list
}

func DifferenceOfOrderedSlices(base, toRemove []string) []string {
	result := []string{}
	i, j := 0, 0
	for i < len(base) && j < len(toRemove) {
		if base[i] < toRemove[j] {
			result = append(result, base[i])
			i++
		} else if base[i] > toRemove[j] {
			j++
		} else {
			i++
			j++
		}
	}
	for i < len(base) {
		result = append(result, base[i])
		i++
	}
	return result
}
