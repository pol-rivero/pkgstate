package common

import (
	"slices"
	"strings"
)

func Sorted(list []string) []string {
	slices.Sort(list)
	return list
}

func SplitAndTrim(s, sep string) []string {
	parts := []string{}
	for part := range strings.SplitSeq(s, sep) {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			parts = append(parts, trimmed)
		}
	}
	return parts
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

func IntersectionOfOrderedSlices(slice1, slice2 []string) []string {
	result := []string{}
	i, j := 0, 0
	for i < len(slice1) && j < len(slice2) {
		if slice1[i] < slice2[j] {
			i++
		} else if slice1[i] > slice2[j] {
			j++
		} else {
			result = append(result, slice1[i])
			i++
			j++
		}
	}
	return result
}
