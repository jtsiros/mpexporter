package mpexporter

import (
	"errors"
	"log"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const validKeepassXML = `
<Root>
	<Entry>
		<String>
			<Key>Title</Key>
			<Value>123</Value>
		</String>
		<String>
			<Key>UserName</Key>
			<Value>1</Value>
		</String>
		<String>
			<Key>Password</Key>
			<Value>a</Value>
		</String>
		<String>
			<Key>Notes</Key>
			<Value>note</Value>
		</String>
		<String>
			<Key>URL</Key>
			<Value>http://lh</Value>
		</String>
	</Entry>
	<Entry>
		<String>
			<Key>Title</Key>
			<Value>456</Value>
		</String>
		<String>
			<Key>UserName</Key>
			<Value>2</Value>
		</String>
		<String>
			<Key>Password</Key>
			<Value>b</Value>
		</String>
		<String>
			<Key>Notes</Key>
			<Value>note2</Value>
		</String>
		<String>
			<Key>URL</Key>
			<Value>http://lh-1</Value>
		</String>
	</Entry>
</Root>
`

const noEntriesKeepassXML = `<Root></Root>`
const invalidKeepassXML = `if (i > 0 && i < 2) {}`

func Test_ReadKeepassXML(t *testing.T) {
	tt := []struct {
		xmlContents     string
		expectedEntries []Entry
		expectedErr     error
	}{
		{
			validKeepassXML,
			[]Entry{
				Entry{"123", "1", "a", "http://lh", "note"},
				Entry{"456", "2", "b", "http://lh-1", "note2"},
			},
			nil,
		},
		{
			noEntriesKeepassXML,
			[]Entry{},
			nil,
		},
		{
			invalidKeepassXML,
			nil,
			errors.New("invalid"),
		},
	}

	for _, tc := range tt {
		entries, err := ReadKeepassXML(strings.NewReader(tc.xmlContents))
		log.Println(entries, err)
		if tc.expectedErr != nil {
			assert.NotNil(t, err)
		} else {
			assert.True(t, reflect.DeepEqual(entries, tc.expectedEntries))
		}
	}
}
