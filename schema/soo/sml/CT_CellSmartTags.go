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

	"github.com/techone577/unioffice"
)

type CT_CellSmartTags struct {
	// Reference
	RAttr string
	// Cell Smart Tag
	CellSmartTag []*CT_CellSmartTag
}

func NewCT_CellSmartTags() *CT_CellSmartTags {
	ret := &CT_CellSmartTags{}
	return ret
}

func (m *CT_CellSmartTags) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r"},
		Value: fmt.Sprintf("%v", m.RAttr)})
	e.EncodeToken(start)
	secellSmartTag := xml.StartElement{Name: xml.Name{Local: "ma:cellSmartTag"}}
	for _, c := range m.CellSmartTag {
		e.EncodeElement(c, secellSmartTag)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CellSmartTags) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "r" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RAttr = parsed
			continue
		}
	}
lCT_CellSmartTags:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "cellSmartTag"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "cellSmartTag"}:
				tmp := NewCT_CellSmartTag()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.CellSmartTag = append(m.CellSmartTag, tmp)
			default:
				unioffice.Log("skipping unsupported element on CT_CellSmartTags %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CellSmartTags
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_CellSmartTags and its children
func (m *CT_CellSmartTags) Validate() error {
	return m.ValidateWithPath("CT_CellSmartTags")
}

// ValidateWithPath validates the CT_CellSmartTags and its children, prefixing error messages with path
func (m *CT_CellSmartTags) ValidateWithPath(path string) error {
	for i, v := range m.CellSmartTag {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/CellSmartTag[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
