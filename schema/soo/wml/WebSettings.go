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

	"github.com/techone577/unioffice"
)

type WebSettings struct {
	CT_WebSettings
}

func NewWebSettings() *WebSettings {
	ret := &WebSettings{}
	ret.CT_WebSettings = *NewCT_WebSettings()
	return ret
}

func (m *WebSettings) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/wordprocessingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:m"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/math"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:ma"}, Value: "http://schemas.openxmlformats.org/schemaLibrary/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:pic"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/picture"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:s"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:w"}, Value: "http://schemas.openxmlformats.org/wordprocessingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:wp"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "w:webSettings"
	return m.CT_WebSettings.MarshalXML(e, start)
}

func (m *WebSettings) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_WebSettings = *NewCT_WebSettings()
lWebSettings:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "frameset"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "frameset"}:
				m.Frameset = NewCT_Frameset()
				if err := d.DecodeElement(m.Frameset, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "divs"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "divs"}:
				m.Divs = NewCT_Divs()
				if err := d.DecodeElement(m.Divs, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "encoding"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "encoding"}:
				m.Encoding = NewCT_String()
				if err := d.DecodeElement(m.Encoding, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "optimizeForBrowser"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "optimizeForBrowser"}:
				m.OptimizeForBrowser = NewCT_OptimizeForBrowser()
				if err := d.DecodeElement(m.OptimizeForBrowser, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "relyOnVML"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "relyOnVML"}:
				m.RelyOnVML = NewCT_OnOff()
				if err := d.DecodeElement(m.RelyOnVML, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "allowPNG"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "allowPNG"}:
				m.AllowPNG = NewCT_OnOff()
				if err := d.DecodeElement(m.AllowPNG, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "doNotRelyOnCSS"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "doNotRelyOnCSS"}:
				m.DoNotRelyOnCSS = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotRelyOnCSS, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "doNotSaveAsSingleFile"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "doNotSaveAsSingleFile"}:
				m.DoNotSaveAsSingleFile = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotSaveAsSingleFile, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "doNotOrganizeInFolder"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "doNotOrganizeInFolder"}:
				m.DoNotOrganizeInFolder = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotOrganizeInFolder, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "doNotUseLongFileNames"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "doNotUseLongFileNames"}:
				m.DoNotUseLongFileNames = NewCT_OnOff()
				if err := d.DecodeElement(m.DoNotUseLongFileNames, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "pixelsPerInch"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "pixelsPerInch"}:
				m.PixelsPerInch = NewCT_DecimalNumber()
				if err := d.DecodeElement(m.PixelsPerInch, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "targetScreenSz"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "targetScreenSz"}:
				m.TargetScreenSz = NewCT_TargetScreenSz()
				if err := d.DecodeElement(m.TargetScreenSz, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "saveSmartTagsAsXml"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "saveSmartTagsAsXml"}:
				m.SaveSmartTagsAsXml = NewCT_OnOff()
				if err := d.DecodeElement(m.SaveSmartTagsAsXml, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on WebSettings %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lWebSettings
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the WebSettings and its children
func (m *WebSettings) Validate() error {
	return m.ValidateWithPath("WebSettings")
}

// ValidateWithPath validates the WebSettings and its children, prefixing error messages with path
func (m *WebSettings) ValidateWithPath(path string) error {
	if err := m.CT_WebSettings.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
