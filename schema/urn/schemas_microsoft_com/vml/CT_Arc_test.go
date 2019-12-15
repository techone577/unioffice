// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package vml_test

import (
	"encoding/xml"
	"testing"

	"github.com/techone577/unioffice/schema/urn/schemas_microsoft_com/vml"
)

func TestCT_ArcConstructor(t *testing.T) {
	v := vml.NewCT_Arc()
	if v == nil {
		t.Errorf("vml.NewCT_Arc must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.CT_Arc should validate: %s", err)
	}
}

func TestCT_ArcMarshalUnmarshal(t *testing.T) {
	v := vml.NewCT_Arc()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewCT_Arc()
	xml.Unmarshal(buf, v2)
}
