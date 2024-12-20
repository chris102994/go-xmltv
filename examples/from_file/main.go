package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/chris102994/xmltv/pkg/go-xmltv/models"
	"golang.org/x/net/html/charset"
	"os"
)

func main() {
	var received models.TV

	xmlData, err := os.ReadFile("test_xmltv.xml")
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	reader := bytes.NewReader(xmlData)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	if err := decoder.Decode(&received); err != nil {
		fmt.Printf("error: %v", err)
	}

	output, err := xml.MarshalIndent(received, "", "    ")

	if err := os.WriteFile("out_xmltv.xml", output, 0644); err != nil {
		fmt.Printf("error: %v", err)
	}
}
