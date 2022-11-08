package http_routing

import "strings"

func lengthOfNested[T any](nested [][]T) int {
	runningLength := 0
	for _, inner := range nested {
		runningLength += len(inner)
	}
	return runningLength
}

func allocateFlattened[T any](nested [][]T) []T {
	descriptionCount := lengthOfNested(nested)
	return make([]T, 0, descriptionCount)
}

func flatten[T any](nested [][]T) []T {
	out := allocateFlattened(nested)
	for _, inner := range nested {
		out = append(out, inner...)
	}
	return out
}

func flattenThenMap[T any, U any](nested [][]T, f func(T) U) []U {
	out := make([]U, 0, lengthOfNested(nested))
	for _, inner := range nested {
		for _, value := range inner {
			out = append(out, f(value))
		}
	}
	return out
}

func mapFind[T any, U any](slice []T, mapper func(item T) *U) *U {
	for _, item := range slice {
		result := mapper(item)
		if result != nil {
			return result
		}
	}
	return nil
}

func takeUntilByte(str string, target byte) (string, string) {
	index := strings.IndexByte(str, target)
	if index == -1 {
		return str, ""
	}
	return str[:index], str[index:]
}
