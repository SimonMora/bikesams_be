package util

import (
	"fmt"
	"strconv"
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

func BuildUpdateSentence(sentence string, field string, v any) string {
	var res string

	if !strings.HasSuffix(sentence, "SET ") {
		res += ", "
	}

	switch v.(type) {
	case int:
		if v.(int) > 0 {
			res += field + " = " + strconv.Itoa(v.(int))
		} else {
			res = ""
		}
	case string:
		if len(v.(string)) != 0 {
			res += field + " = '" + ScapeString(v.(string)) + "'"
		} else {
			res = ""
		}
	case float64:
		if v.(float64) > 0 {
			res += field + " = " + strconv.FormatFloat(v.(float64), 'e', -1, 64)
		} else {
			res = ""
		}
	}

	return sentence + res
}
