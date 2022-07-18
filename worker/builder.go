type IStringBuilder struct {
	append func(string) string
	remove func(string) string
	replace func(string) string
}

func StringBuilder(original string) *StringBuilder {
	var string_val string = original
}