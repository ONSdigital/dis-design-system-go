package helper

import "html/template"

func SafeHTML(s string) template.HTML {
	//nolint:gosec // known risk, input is expected to be safe
	return template.HTML(s)
}
