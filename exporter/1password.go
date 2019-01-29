package exporter

import "github.com/jtsiros/mpexporter"

type OnePasswordExporter struct{}

func (ex OnePasswordExporter) ToSlice(e *mpexporter.Entry) []string {
	fields := []string{e.Title, e.URL, e.Username, e.Password, e.Notes}
	return fields
}
