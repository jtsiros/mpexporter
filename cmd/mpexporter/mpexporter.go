package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jtsiros/mpexporter"
	"github.com/jtsiros/mpexporter/exporter"
)

// takes any KeePass compatible database XML file and converts it to a provider-compatible CSV format.
func main() {

	var exportType string
	var keepassFile string
	var outputFile string
	flag.StringVar(&exportType, "export", "1password", "specifies CSV compatible password manager")
	flag.StringVar(&keepassFile, "i", "output.xml", "specifies keepass DB file in XML format")

	defaultOut := fmt.Sprintf("%s.csv", keepassFile[:strings.IndexByte(keepassFile, '.')])
	flag.StringVar(&outputFile, "o", defaultOut, "specifies output csv file")
	flag.Parse()

	ex, err := exporter.Get(exportType)
	if err != nil {
		log.Fatalf("Error with exporter: %v\n", err)
	}

	f, err := os.Open(keepassFile)
	if err != nil {
		log.Fatalf("cannot open keepass DB file: %v\n", err)
	}
	defer f.Close()

	entries, err := mpexporter.ReadKeepassXML(f)
	if err != nil {
		log.Fatalf("could not read Keepass XML file: %v\n", err)
	}

	log.Printf("read %d entries\n", len(entries))

	// write DB to csv file
	f, err = os.OpenFile(outputFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("cannot open output file: %s", outputFile)
	}

	defer f.Close()
	writer := csv.NewWriter(f)
	defer writer.Flush()

	for _, entry := range entries {
		if err = writer.Write(ex.ToSlice(&entry)); err != nil {
			log.Fatalln(err)
		}
	}
}
