// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package wml_test

import (
	"encoding/xml"
	"testing"

	"github.com/techone577/unioffice/schema/soo/wml"
)

func TestEG_MathContentConstructor(t *testing.T) {
	v := wml.NewEG_MathContent()
	if v == nil {
		t.Errorf("wml.NewEG_MathContent must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.EG_MathContent should validate: %s", err)
	}
}

func TestEG_MathContentMarshalUnmarshal(t *testing.T) {
	v := wml.NewEG_MathContent()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewEG_MathContent()
	xml.Unmarshal(buf, v2)
}
