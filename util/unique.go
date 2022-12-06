package util

func Unique[T comparable](s []T) []T {
	// I stole this :D
	// https://gosamples.dev/generics-remove-duplicates-slice/
	inResult := make(map[T]bool)
	var result []T
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}
