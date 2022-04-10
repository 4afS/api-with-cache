package migration

import (
	"strings"
)

func GenerateQueryValues(max int, query func(int) string) string {
	queries := make([]string, max)
	for i := 0; i < max; i++ {
		queries[i] = query(i + 1)
	}
	return strings.Join(queries[:], ",")
}
