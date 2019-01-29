package exporter

import (
	"testing"

	"github.com/jtsiros/mpexporter"
	"github.com/stretchr/testify/assert"
)

func Test_1Password_ToSlice(t *testing.T) {

	exp := OnePasswordExporter{}

	tt := []struct {
		entry  mpexporter.Entry
		result []string
	}{
		{mpexporter.Entry{Title: "title"}, []string{"title", "", "", "", ""}},
		{mpexporter.Entry{Title: "title", Username: "1"}, []string{"title", "", "1", "", ""}},
		{mpexporter.Entry{Title: "title", Username: "1", Password: "a", URL: "http://lh", Notes: "note"},
			[]string{"title", "http://lh", "1", "a", "note"}},
	}

	for _, tc := range tt {
		assert.Equal(t, exp.ToSlice(&tc.entry), tc.result)
	}
}
