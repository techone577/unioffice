// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/techone577/unioffice"
)

type CT_QueryTableFields struct {
	// Column Count
	CountAttr *uint32
	// QueryTable Field
	QueryTableField []*CT_QueryTableField
}

func NewCT_QueryTableFields() *CT_QueryTableFields {
	ret := &CT_QueryTableFields{}
	return ret
}

func (m *CT_QueryTableFields) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	if m.QueryTableField != nil {
		sequeryTableField := xml.StartElement{Name: xml.Name{Local: "ma:queryTableField"}}
		for _, c := range m.QueryTableField {
			e.EncodeElement(c, sequeryTableField)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_QueryTableFields) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "count" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CountAttr = &pt
			continue
		}
	}
lCT_QueryTableFields:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "queryTableField"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "queryTableField"}:
				tmp := NewCT_QueryTableField()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.QueryTableField = append(m.QueryTableField, tmp)
			default:
				unioffice.Log("skipping unsupported element on CT_QueryTableFields %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_QueryTableFields
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_QueryTableFields and its children
func (m *CT_QueryTableFields) Validate() error {
	return m.ValidateWithPath("CT_QueryTableFields")
}

// ValidateWithPath validates the CT_QueryTableFields and its children, prefixing error messages with path
func (m *CT_QueryTableFields) ValidateWithPath(path string) error {
	for i, v := range m.QueryTableField {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/QueryTableField[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
