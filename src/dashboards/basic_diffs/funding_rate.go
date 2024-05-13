package basic_diffs

import (
	"bytes"
	"text/template"
)

func GetFundingQuery() string {
	var rawTmpl bytes.Buffer

	tmpl, err := template.New("price_diff.sql").ParseFiles("price_diff.sql")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(&rawTmpl, nil)
	if err != nil {
		panic(err)
	}

	return rawTmpl.String()
}
