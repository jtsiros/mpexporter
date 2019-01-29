package exporter

import (
	"fmt"

	"github.com/jtsiros/mpexporter"
)

func Get(t string) (mpexporter.Exporter, error) {
	switch t {
	case "1password":
		return OnePasswordExporter{}, nil
	default:
		return nil, fmt.Errorf("Unsupported format: %s", t)
	}
}
