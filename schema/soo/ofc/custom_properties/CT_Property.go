// Copyright 2019 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package custom_properties

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"time"

	"github.com/techone577/unioffice/schema/soo/ofc/docPropsVTypes"
	"github.com/techone577/unioffice/schema/soo/ofc/sharedTypes"

	"github.com/techone577/unioffice"
)

type CT_Property struct {
	FmtidAttr      string
	PidAttr        int32
	NameAttr       *string
	LinkTargetAttr *string
	Vector         *docPropsVTypes.Vector
	Array          *docPropsVTypes.Array
	Blob           *string
	Oblob          *string
	Empty          *docPropsVTypes.Empty
	Null           *docPropsVTypes.Null
	I1             *int8
	I2             *int16
	I4             *int32
	I8             *int64
	Int            *int32
	Ui1            *uint8
	Ui2            *uint16
	Ui4            *uint32
	Ui8            *uint64
	Uint           *uint32
	R4             *float32
	R8             *float64
	Decimal        *float64
	Lpstr          *string
	Lpwstr         *string
	Bstr           *string
	Date           *time.Time
	Filetime       *time.Time
	Bool           *bool
	Cy             *string
	Error          *string
	Stream         *string
	Ostream        *string
	Storage        *string
	Ostorage       *string
	Vstream        *docPropsVTypes.Vstream
	Clsid          *string
}

func NewCT_Property() *CT_Property {
	ret := &CT_Property{}
	ret.FmtidAttr = "{00000000-0000-0000-0000-000000000000}"
	return ret
}

func (m *CT_Property) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fmtid"},
		Value: fmt.Sprintf("%v", m.FmtidAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "pid"},
		Value: fmt.Sprintf("%v", m.PidAttr)})
	if m.NameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
			Value: fmt.Sprintf("%v", *m.NameAttr)})
	}
	if m.LinkTargetAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "linkTarget"},
			Value: fmt.Sprintf("%v", *m.LinkTargetAttr)})
	}
	e.EncodeToken(start)
	if m.Vector != nil {
		sevector := xml.StartElement{Name: xml.Name{Local: "vt:vector"}}
		e.EncodeElement(m.Vector, sevector)
	}
	if m.Array != nil {
		searray := xml.StartElement{Name: xml.Name{Local: "vt:array"}}
		e.EncodeElement(m.Array, searray)
	}
	if m.Blob != nil {
		seblob := xml.StartElement{Name: xml.Name{Local: "vt:blob"}}
		unioffice.AddPreserveSpaceAttr(&seblob, *m.Blob)
		e.EncodeElement(m.Blob, seblob)
	}
	if m.Oblob != nil {
		seoblob := xml.StartElement{Name: xml.Name{Local: "vt:oblob"}}
		unioffice.AddPreserveSpaceAttr(&seoblob, *m.Oblob)
		e.EncodeElement(m.Oblob, seoblob)
	}
	if m.Empty != nil {
		seempty := xml.StartElement{Name: xml.Name{Local: "vt:empty"}}
		e.EncodeElement(m.Empty, seempty)
	}
	if m.Null != nil {
		senull := xml.StartElement{Name: xml.Name{Local: "vt:null"}}
		e.EncodeElement(m.Null, senull)
	}
	if m.I1 != nil {
		sei1 := xml.StartElement{Name: xml.Name{Local: "vt:i1"}}
		e.EncodeElement(m.I1, sei1)
	}
	if m.I2 != nil {
		sei2 := xml.StartElement{Name: xml.Name{Local: "vt:i2"}}
		e.EncodeElement(m.I2, sei2)
	}
	if m.I4 != nil {
		sei4 := xml.StartElement{Name: xml.Name{Local: "vt:i4"}}
		e.EncodeElement(m.I4, sei4)
	}
	if m.I8 != nil {
		sei8 := xml.StartElement{Name: xml.Name{Local: "vt:i8"}}
		e.EncodeElement(m.I8, sei8)
	}
	if m.Int != nil {
		seint := xml.StartElement{Name: xml.Name{Local: "vt:int"}}
		e.EncodeElement(m.Int, seint)
	}
	if m.Ui1 != nil {
		seui1 := xml.StartElement{Name: xml.Name{Local: "vt:ui1"}}
		e.EncodeElement(m.Ui1, seui1)
	}
	if m.Ui2 != nil {
		seui2 := xml.StartElement{Name: xml.Name{Local: "vt:ui2"}}
		e.EncodeElement(m.Ui2, seui2)
	}
	if m.Ui4 != nil {
		seui4 := xml.StartElement{Name: xml.Name{Local: "vt:ui4"}}
		e.EncodeElement(m.Ui4, seui4)
	}
	if m.Ui8 != nil {
		seui8 := xml.StartElement{Name: xml.Name{Local: "vt:ui8"}}
		e.EncodeElement(m.Ui8, seui8)
	}
	if m.Uint != nil {
		seuint := xml.StartElement{Name: xml.Name{Local: "vt:uint"}}
		e.EncodeElement(m.Uint, seuint)
	}
	if m.R4 != nil {
		ser4 := xml.StartElement{Name: xml.Name{Local: "vt:r4"}}
		e.EncodeElement(m.R4, ser4)
	}
	if m.R8 != nil {
		ser8 := xml.StartElement{Name: xml.Name{Local: "vt:r8"}}
		e.EncodeElement(m.R8, ser8)
	}
	if m.Decimal != nil {
		sedecimal := xml.StartElement{Name: xml.Name{Local: "vt:decimal"}}
		e.EncodeElement(m.Decimal, sedecimal)
	}
	if m.Lpstr != nil {
		selpstr := xml.StartElement{Name: xml.Name{Local: "vt:lpstr"}}
		unioffice.AddPreserveSpaceAttr(&selpstr, *m.Lpstr)
		e.EncodeElement(m.Lpstr, selpstr)
	}
	if m.Lpwstr != nil {
		selpwstr := xml.StartElement{Name: xml.Name{Local: "vt:lpwstr"}}
		unioffice.AddPreserveSpaceAttr(&selpwstr, *m.Lpwstr)
		e.EncodeElement(m.Lpwstr, selpwstr)
	}
	if m.Bstr != nil {
		sebstr := xml.StartElement{Name: xml.Name{Local: "vt:bstr"}}
		unioffice.AddPreserveSpaceAttr(&sebstr, *m.Bstr)
		e.EncodeElement(m.Bstr, sebstr)
	}
	if m.Date != nil {
		sedate := xml.StartElement{Name: xml.Name{Local: "vt:date"}}
		e.EncodeElement(m.Date, sedate)
	}
	if m.Filetime != nil {
		sefiletime := xml.StartElement{Name: xml.Name{Local: "vt:filetime"}}
		e.EncodeElement(m.Filetime, sefiletime)
	}
	if m.Bool != nil {
		sebool := xml.StartElement{Name: xml.Name{Local: "vt:bool"}}
		e.EncodeElement(m.Bool, sebool)
	}
	if m.Cy != nil {
		secy := xml.StartElement{Name: xml.Name{Local: "vt:cy"}}
		unioffice.AddPreserveSpaceAttr(&secy, *m.Cy)
		e.EncodeElement(m.Cy, secy)
	}
	if m.Error != nil {
		seerror := xml.StartElement{Name: xml.Name{Local: "vt:error"}}
		unioffice.AddPreserveSpaceAttr(&seerror, *m.Error)
		e.EncodeElement(m.Error, seerror)
	}
	if m.Stream != nil {
		sestream := xml.StartElement{Name: xml.Name{Local: "vt:stream"}}
		unioffice.AddPreserveSpaceAttr(&sestream, *m.Stream)
		e.EncodeElement(m.Stream, sestream)
	}
	if m.Ostream != nil {
		seostream := xml.StartElement{Name: xml.Name{Local: "vt:ostream"}}
		unioffice.AddPreserveSpaceAttr(&seostream, *m.Ostream)
		e.EncodeElement(m.Ostream, seostream)
	}
	if m.Storage != nil {
		sestorage := xml.StartElement{Name: xml.Name{Local: "vt:storage"}}
		unioffice.AddPreserveSpaceAttr(&sestorage, *m.Storage)
		e.EncodeElement(m.Storage, sestorage)
	}
	if m.Ostorage != nil {
		seostorage := xml.StartElement{Name: xml.Name{Local: "vt:ostorage"}}
		unioffice.AddPreserveSpaceAttr(&seostorage, *m.Ostorage)
		e.EncodeElement(m.Ostorage, seostorage)
	}
	if m.Vstream != nil {
		sevstream := xml.StartElement{Name: xml.Name{Local: "vt:vstream"}}
		e.EncodeElement(m.Vstream, sevstream)
	}
	if m.Clsid != nil {
		seclsid := xml.StartElement{Name: xml.Name{Local: "vt:clsid"}}
		unioffice.AddPreserveSpaceAttr(&seclsid, *m.Clsid)
		e.EncodeElement(m.Clsid, seclsid)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Property) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.FmtidAttr = "{00000000-0000-0000-0000-000000000000}"
	for _, attr := range start.Attr {
		if attr.Name.Local == "pid" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.PidAttr = int32(parsed)
			continue
		}
		if attr.Name.Local == "linkTarget" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.LinkTargetAttr = &parsed
			continue
		}
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = &parsed
			continue
		}
		if attr.Name.Local == "fmtid" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.FmtidAttr = parsed
			continue
		}
	}
lCT_Property:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "vector"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "vector"}:
				m.Vector = docPropsVTypes.NewVector()
				if err := d.DecodeElement(m.Vector, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "array"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "array"}:
				m.Array = docPropsVTypes.NewArray()
				if err := d.DecodeElement(m.Array, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "blob"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "blob"}:
				m.Blob = new(string)
				if err := d.DecodeElement(m.Blob, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "oblob"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "oblob"}:
				m.Oblob = new(string)
				if err := d.DecodeElement(m.Oblob, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "empty"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "empty"}:
				m.Empty = docPropsVTypes.NewEmpty()
				if err := d.DecodeElement(m.Empty, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "null"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "null"}:
				m.Null = docPropsVTypes.NewNull()
				if err := d.DecodeElement(m.Null, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "i1"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "i1"}:
				m.I1 = new(int8)
				if err := d.DecodeElement(m.I1, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "i2"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "i2"}:
				m.I2 = new(int16)
				if err := d.DecodeElement(m.I2, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "i4"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "i4"}:
				m.I4 = new(int32)
				if err := d.DecodeElement(m.I4, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "i8"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "i8"}:
				m.I8 = new(int64)
				if err := d.DecodeElement(m.I8, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "int"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "int"}:
				m.Int = new(int32)
				if err := d.DecodeElement(m.Int, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "ui1"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "ui1"}:
				m.Ui1 = new(uint8)
				if err := d.DecodeElement(m.Ui1, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "ui2"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "ui2"}:
				m.Ui2 = new(uint16)
				if err := d.DecodeElement(m.Ui2, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "ui4"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "ui4"}:
				m.Ui4 = new(uint32)
				if err := d.DecodeElement(m.Ui4, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "ui8"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "ui8"}:
				m.Ui8 = new(uint64)
				if err := d.DecodeElement(m.Ui8, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "uint"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "uint"}:
				m.Uint = new(uint32)
				if err := d.DecodeElement(m.Uint, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "r4"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "r4"}:
				m.R4 = new(float32)
				if err := d.DecodeElement(m.R4, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "r8"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "r8"}:
				m.R8 = new(float64)
				if err := d.DecodeElement(m.R8, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "decimal"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "decimal"}:
				m.Decimal = new(float64)
				if err := d.DecodeElement(m.Decimal, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "lpstr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "lpstr"}:
				m.Lpstr = new(string)
				if err := d.DecodeElement(m.Lpstr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "lpwstr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "lpwstr"}:
				m.Lpwstr = new(string)
				if err := d.DecodeElement(m.Lpwstr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "bstr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "bstr"}:
				m.Bstr = new(string)
				if err := d.DecodeElement(m.Bstr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "date"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "date"}:
				m.Date = new(time.Time)
				if err := d.DecodeElement(m.Date, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "filetime"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "filetime"}:
				filetimeStr := new(string)
				if err := d.DecodeElement(filetimeStr, &el); err != nil {
					return err
				}
				filetime, err := time.Parse("2006-01-02T15:04:05Z07:00", *filetimeStr)
				if err != nil {
					filetime, err = time.Parse("02-01-2006T15:04:05Z07:00", *filetimeStr)
					if err != nil {
						return err
					}
				}
				m.Filetime = &filetime
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "bool"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "bool"}:
				m.Bool = new(bool)
				if err := d.DecodeElement(m.Bool, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "cy"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "cy"}:
				m.Cy = new(string)
				if err := d.DecodeElement(m.Cy, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "error"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "error"}:
				m.Error = new(string)
				if err := d.DecodeElement(m.Error, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "stream"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "stream"}:
				m.Stream = new(string)
				if err := d.DecodeElement(m.Stream, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "ostream"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "ostream"}:
				m.Ostream = new(string)
				if err := d.DecodeElement(m.Ostream, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "storage"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "storage"}:
				m.Storage = new(string)
				if err := d.DecodeElement(m.Storage, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "ostorage"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "ostorage"}:
				m.Ostorage = new(string)
				if err := d.DecodeElement(m.Ostorage, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "vstream"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "vstream"}:
				m.Vstream = docPropsVTypes.NewVstream()
				if err := d.DecodeElement(m.Vstream, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "clsid"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "clsid"}:
				m.Clsid = new(string)
				if err := d.DecodeElement(m.Clsid, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_Property %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Property
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Property and its children
func (m *CT_Property) Validate() error {
	return m.ValidateWithPath("CT_Property")
}

// ValidateWithPath validates the CT_Property and its children, prefixing error messages with path
func (m *CT_Property) ValidateWithPath(path string) error {
	if !sharedTypes.ST_GuidPatternRe.MatchString(m.FmtidAttr) {
		return fmt.Errorf(`%s/m.FmtidAttr must match '%s' (have %v)`, path, sharedTypes.ST_GuidPatternRe, m.FmtidAttr)
	}
	if m.Vector != nil {
		if err := m.Vector.ValidateWithPath(path + "/Vector"); err != nil {
			return err
		}
	}
	if m.Array != nil {
		if err := m.Array.ValidateWithPath(path + "/Array"); err != nil {
			return err
		}
	}
	if m.Empty != nil {
		if err := m.Empty.ValidateWithPath(path + "/Empty"); err != nil {
			return err
		}
	}
	if m.Null != nil {
		if err := m.Null.ValidateWithPath(path + "/Null"); err != nil {
			return err
		}
	}
	if m.Cy != nil {
		if !docPropsVTypes.ST_CyPatternRe.MatchString(*m.Cy) {
			return fmt.Errorf(`%s/m.Cy must match '%s' (have %v)`, path, docPropsVTypes.ST_CyPatternRe, *m.Cy)
		}
	}
	if m.Error != nil {
		if !docPropsVTypes.ST_ErrorPatternRe.MatchString(*m.Error) {
			return fmt.Errorf(`%s/m.Error must match '%s' (have %v)`, path, docPropsVTypes.ST_ErrorPatternRe, *m.Error)
		}
	}
	if m.Vstream != nil {
		if err := m.Vstream.ValidateWithPath(path + "/Vstream"); err != nil {
			return err
		}
	}
	if m.Clsid != nil {
		if !sharedTypes.ST_GuidPatternRe.MatchString(*m.Clsid) {
			return fmt.Errorf(`%s/m.Clsid must match '%s' (have %v)`, path, sharedTypes.ST_GuidPatternRe, *m.Clsid)
		}
	}
	return nil
}
