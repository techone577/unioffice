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
	"strconv"

	"github.com/techone577/unioffice"
)

type CT_Div struct {
	// div Data ID
	IdAttr int64
	// Data for HTML blockquote Element
	BlockQuote *CT_OnOff
	// Data for HTML body Element
	BodyDiv *CT_OnOff
	// Left Margin for HTML div
	MarLeft *CT_SignedTwipsMeasure
	// Right Margin for HTML div
	MarRight *CT_SignedTwipsMeasure
	// Top Margin for HTML div
	MarTop *CT_SignedTwipsMeasure
	// Bottom Margin for HTML div
	MarBottom *CT_SignedTwipsMeasure
	// Set of Borders for HTML div
	DivBdr *CT_DivBdr
	// Child div Elements Contained within Current div
	DivsChild []*CT_Divs
}

func NewCT_Div() *CT_Div {
	ret := &CT_Div{}
	ret.MarLeft = NewCT_SignedTwipsMeasure()
	ret.MarRight = NewCT_SignedTwipsMeasure()
	ret.MarTop = NewCT_SignedTwipsMeasure()
	ret.MarBottom = NewCT_SignedTwipsMeasure()
	return ret
}

func (m *CT_Div) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	e.EncodeToken(start)
	if m.BlockQuote != nil {
		seblockQuote := xml.StartElement{Name: xml.Name{Local: "w:blockQuote"}}
		e.EncodeElement(m.BlockQuote, seblockQuote)
	}
	if m.BodyDiv != nil {
		sebodyDiv := xml.StartElement{Name: xml.Name{Local: "w:bodyDiv"}}
		e.EncodeElement(m.BodyDiv, sebodyDiv)
	}
	semarLeft := xml.StartElement{Name: xml.Name{Local: "w:marLeft"}}
	e.EncodeElement(m.MarLeft, semarLeft)
	semarRight := xml.StartElement{Name: xml.Name{Local: "w:marRight"}}
	e.EncodeElement(m.MarRight, semarRight)
	semarTop := xml.StartElement{Name: xml.Name{Local: "w:marTop"}}
	e.EncodeElement(m.MarTop, semarTop)
	semarBottom := xml.StartElement{Name: xml.Name{Local: "w:marBottom"}}
	e.EncodeElement(m.MarBottom, semarBottom)
	if m.DivBdr != nil {
		sedivBdr := xml.StartElement{Name: xml.Name{Local: "w:divBdr"}}
		e.EncodeElement(m.DivBdr, sedivBdr)
	}
	if m.DivsChild != nil {
		sedivsChild := xml.StartElement{Name: xml.Name{Local: "w:divsChild"}}
		for _, c := range m.DivsChild {
			e.EncodeElement(c, sedivsChild)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Div) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.MarLeft = NewCT_SignedTwipsMeasure()
	m.MarRight = NewCT_SignedTwipsMeasure()
	m.MarTop = NewCT_SignedTwipsMeasure()
	m.MarBottom = NewCT_SignedTwipsMeasure()
	for _, attr := range start.Attr {
		if attr.Name.Local == "id" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.IdAttr = parsed
			continue
		}
	}
lCT_Div:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "blockQuote"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "blockQuote"}:
				m.BlockQuote = NewCT_OnOff()
				if err := d.DecodeElement(m.BlockQuote, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "bodyDiv"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "bodyDiv"}:
				m.BodyDiv = NewCT_OnOff()
				if err := d.DecodeElement(m.BodyDiv, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "marLeft"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "marLeft"}:
				if err := d.DecodeElement(m.MarLeft, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "marRight"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "marRight"}:
				if err := d.DecodeElement(m.MarRight, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "marTop"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "marTop"}:
				if err := d.DecodeElement(m.MarTop, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "marBottom"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "marBottom"}:
				if err := d.DecodeElement(m.MarBottom, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "divBdr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "divBdr"}:
				m.DivBdr = NewCT_DivBdr()
				if err := d.DecodeElement(m.DivBdr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "divsChild"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "divsChild"}:
				tmp := NewCT_Divs()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.DivsChild = append(m.DivsChild, tmp)
			default:
				unioffice.Log("skipping unsupported element on CT_Div %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Div
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Div and its children
func (m *CT_Div) Validate() error {
	return m.ValidateWithPath("CT_Div")
}

// ValidateWithPath validates the CT_Div and its children, prefixing error messages with path
func (m *CT_Div) ValidateWithPath(path string) error {
	if m.BlockQuote != nil {
		if err := m.BlockQuote.ValidateWithPath(path + "/BlockQuote"); err != nil {
			return err
		}
	}
	if m.BodyDiv != nil {
		if err := m.BodyDiv.ValidateWithPath(path + "/BodyDiv"); err != nil {
			return err
		}
	}
	if err := m.MarLeft.ValidateWithPath(path + "/MarLeft"); err != nil {
		return err
	}
	if err := m.MarRight.ValidateWithPath(path + "/MarRight"); err != nil {
		return err
	}
	if err := m.MarTop.ValidateWithPath(path + "/MarTop"); err != nil {
		return err
	}
	if err := m.MarBottom.ValidateWithPath(path + "/MarBottom"); err != nil {
		return err
	}
	if m.DivBdr != nil {
		if err := m.DivBdr.ValidateWithPath(path + "/DivBdr"); err != nil {
			return err
		}
	}
	for i, v := range m.DivsChild {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/DivsChild[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
