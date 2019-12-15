// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package core_properties_test

import (
	"encoding/xml"
	"testing"

	"github.com/techone577/unioffice/schema/soo/pkg/metadata/core_properties"
)

func TestCT_KeywordsConstructor(t *testing.T) {
	v := core_properties.NewCT_Keywords()
	if v == nil {
		t.Errorf("core_properties.NewCT_Keywords must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed core_properties.CT_Keywords should validate: %s", err)
	}
}

func TestCT_KeywordsMarshalUnmarshal(t *testing.T) {
	v := core_properties.NewCT_Keywords()
	buf, _ := xml.Marshal(v)
	v2 := core_properties.NewCT_Keywords()
	xml.Unmarshal(buf, v2)
}
