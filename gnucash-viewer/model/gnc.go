package model

import (
	"compress/gzip"
	"encoding/xml"
	"os"
)

// GnuCash type
type GnuCash struct {
	XMLName xml.Name `xml:"gnc-v2"`
	Books   []Book   `xml:"book"`
}

// LoadFromXMLFile function
func LoadFromXMLFile(path string) (*GnuCash, error) {
	// open gnucash file
	gnucashFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer gnucashFile.Close()

	// decompress gnucash file
	reader, err := gzip.NewReader(gnucashFile)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	// unmarshall XML
	gnc := GnuCash{}
	err = xml.NewDecoder(reader).Decode(&gnc)
	if err != nil {
		return nil, err
	}

	// call post Load XML for each book
	for _, b := range gnc.Books {
		err := b.postLoadXML()
		if err != nil {
			panic(err)
		}
	}

	return &gnc, nil
}
