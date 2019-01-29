package exporter

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ExporterTypes(t *testing.T) {

	tt := []struct {
		expType      string
		expectedType reflect.Type
		expectedErr  error
	}{
		{"1password", reflect.TypeOf(OnePasswordExporter{}), nil},
		{"unkown", nil, errors.New("Unsupported format: unkown")},
	}

	for _, tc := range tt {
		expType, err := Get(tc.expType)
		if tc.expectedErr != nil {
			assert.Equal(t, tc.expectedErr.Error(), err.Error())
		} else {
			assert.Equal(t, reflect.TypeOf(expType), tc.expectedType)
		}
	}
}
