package main

import (
	"fmt"
	"strings"
)

type IStringBuilder struct {
	append     func(string) string
	remove     func(string, int) string
	replace    func(string, string) string
	replaceAll func(string, string) string
	format     func([]string) string
	toString   func() string
}

func StringBuilder(original string) *IStringBuilder {
	var string_val string = original

	value := &IStringBuilder{
		append: func(value string) string {
			string_val = original + value

			return string_val
		},
		remove: func(to_replace string, amount int) string {
			string_val = strings.Replace(string_val, to_replace, "", amount)

			return string_val
		},
		replace: func(to_replace string, replaceble string) string {
			string_val = strings.Replace(string_val, to_replace, replaceble, -1)

			return string_val
		},
		replaceAll: func(to_replace string, replaceble string) string {
			string_val = strings.ReplaceAll(string_val, to_replace, replaceble)

			return string_val
		},
		format: func(values []string) string {
			str := FormatString(string_val, values)

			return str
		},
		toString: func() string {
			return string_val
		},
	}

	return value
}

func FormatString(to_format string, values []string) string {
	var builder = StringBuilder(to_format)

	for i := 0; i < len(values); i++ {
		var val string = fmt.Sprint(i)

		builder.replace("{"+val+"}", values[i])
	}

	return builder.toString()
}
