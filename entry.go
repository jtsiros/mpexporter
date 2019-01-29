package mpexporter

type Exporter interface {
	ToSlice(e *Entry) []string
}

// An Entry represents an XML Entry element in the Keepass DB.
type Entry struct {
	Title    string
	Username string
	Password string
	URL      string
	Notes    string
}
