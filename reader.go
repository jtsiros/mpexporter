package mpexporter

import (
	"io"

	"github.com/beevik/etree"
)

// ReadKeepassXML reads all Entry elements from the Keepass DB XML file.
// returns a list of Entry objects or error
func ReadKeepassXML(r io.Reader) ([]Entry, error) {
	doc := etree.NewDocument()
	_, err := doc.ReadFrom(r)
	if err != nil {
		return nil, err
	}

	entries := []Entry{}
	for _, e := range doc.FindElements("//Entry") {
		entry := Entry{}
		for _, data := range e.SelectElements("String") {
			key := data.SelectElement("Key")
			val := data.SelectElement("Value")
			switch key.Text() {
			case "Title":
				entry.Title = val.Text()
			case "UserName":
				entry.Username = val.Text()
			case "Password":
				entry.Password = val.Text()
			case "Notes":
				entry.Notes = val.Text()
			case "URL":
				entry.URL = val.Text()
			}
		}
		entries = append(entries, entry)
	}
	return entries, nil
}
