package util

import (
	"fmt"
	"strings"
	"time"
)

func DateSqlFormat() string {
	t := time.Now()
	return fmt.Sprintf(
		"%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(),
	)
}

func ScapeString(s string) string {
	res := strings.ReplaceAll(s, "'", "")
	res = strings.ReplaceAll(res, "\"", "")
	return res
}
