// Copyright 2019 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package custom_properties_test

import (
	"encoding/xml"
	"testing"

	"github.com/techone577/unioffice/schema/soo/ofc/custom_properties"
)

func TestCT_PropertiesConstructor(t *testing.T) {
	v := custom_properties.NewCT_Properties()
	if v == nil {
		t.Errorf("custom_properties.NewCT_Properties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed custom_properties.CT_Properties should validate: %s", err)
	}
}

func TestCT_PropertiesMarshalUnmarshal(t *testing.T) {
	v := custom_properties.NewCT_Properties()
	buf, _ := xml.Marshal(v)
	v2 := custom_properties.NewCT_Properties()
	xml.Unmarshal(buf, v2)
}
