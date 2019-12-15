// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package pml

import (
	"encoding/xml"

	"github.com/techone577/unioffice"
)

type CT_TLTimeNodeParallel struct {
	// Parallel TimeNode
	CTn *CT_TLCommonTimeNodeData
}

func NewCT_TLTimeNodeParallel() *CT_TLTimeNodeParallel {
	ret := &CT_TLTimeNodeParallel{}
	ret.CTn = NewCT_TLCommonTimeNodeData()
	return ret
}

func (m *CT_TLTimeNodeParallel) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	secTn := xml.StartElement{Name: xml.Name{Local: "p:cTn"}}
	e.EncodeElement(m.CTn, secTn)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TLTimeNodeParallel) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CTn = NewCT_TLCommonTimeNodeData()
lCT_TLTimeNodeParallel:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "cTn"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "cTn"}:
				if err := d.DecodeElement(m.CTn, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_TLTimeNodeParallel %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TLTimeNodeParallel
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TLTimeNodeParallel and its children
func (m *CT_TLTimeNodeParallel) Validate() error {
	return m.ValidateWithPath("CT_TLTimeNodeParallel")
}

// ValidateWithPath validates the CT_TLTimeNodeParallel and its children, prefixing error messages with path
func (m *CT_TLTimeNodeParallel) ValidateWithPath(path string) error {
	if err := m.CTn.ValidateWithPath(path + "/CTn"); err != nil {
		return err
	}
	return nil
}
