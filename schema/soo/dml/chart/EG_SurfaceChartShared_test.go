// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package chart_test

import (
	"encoding/xml"
	"testing"

	"github.com/techone577/unioffice/schema/soo/dml/chart"
)

func TestEG_SurfaceChartSharedConstructor(t *testing.T) {
	v := chart.NewEG_SurfaceChartShared()
	if v == nil {
		t.Errorf("chart.NewEG_SurfaceChartShared must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.EG_SurfaceChartShared should validate: %s", err)
	}
}

func TestEG_SurfaceChartSharedMarshalUnmarshal(t *testing.T) {
	v := chart.NewEG_SurfaceChartShared()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewEG_SurfaceChartShared()
	xml.Unmarshal(buf, v2)
}
