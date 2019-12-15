// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/techone577/unioffice/schema/soo/dml"
)

func TestEG_ColorChoiceConstructor(t *testing.T) {
	v := dml.NewEG_ColorChoice()
	if v == nil {
		t.Errorf("dml.NewEG_ColorChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.EG_ColorChoice should validate: %s", err)
	}
}

func TestEG_ColorChoiceMarshalUnmarshal(t *testing.T) {
	v := dml.NewEG_ColorChoice()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewEG_ColorChoice()
	xml.Unmarshal(buf, v2)
}
