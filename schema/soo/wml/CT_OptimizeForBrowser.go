// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package wml

import (
	"encoding/xml"
	"fmt"

	"github.com/techone577/unioffice/schema/soo/ofc/sharedTypes"
)

type CT_OptimizeForBrowser struct {
	TargetAttr *string
	// On/Off Value
	ValAttr *sharedTypes.ST_OnOff
}

func NewCT_OptimizeForBrowser() *CT_OptimizeForBrowser {
	ret := &CT_OptimizeForBrowser{}
	return ret
}

func (m *CT_OptimizeForBrowser) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.TargetAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:target"},
			Value: fmt.Sprintf("%v", *m.TargetAttr)})
	}
	if m.ValAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"},
			Value: fmt.Sprintf("%v", *m.ValAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_OptimizeForBrowser) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "target" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TargetAttr = &parsed
			continue
		}
		if attr.Name.Local == "val" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.ValAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_OptimizeForBrowser: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_OptimizeForBrowser and its children
func (m *CT_OptimizeForBrowser) Validate() error {
	return m.ValidateWithPath("CT_OptimizeForBrowser")
}

// ValidateWithPath validates the CT_OptimizeForBrowser and its children, prefixing error messages with path
func (m *CT_OptimizeForBrowser) ValidateWithPath(path string) error {
	if m.ValAttr != nil {
		if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
			return err
		}
	}
	return nil
}
