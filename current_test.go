// Licensed to The Moov Authors under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. The Moov Authors licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package iso4217_test

import (
	"testing"

	"github.com/moov-io/iso4217"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAll(t *testing.T) {
	all := iso4217.All()
	require.NotEmpty(t, all)
	assert.Equal(t, all, iso4217.All())

	seen := make(map[string]struct{}, len(all))
	for i, cc := range all {
		assert.NotEmpty(t, cc.Code)
		if i > 0 {
			assert.Truef(t, all[i-1].Code < cc.Code, "codes are not sorted: %s then %s", all[i-1].Code, cc.Code)
		}
		_, dup := seen[cc.Code]
		assert.Falsef(t, dup, "duplicate %s", cc.Code)
		seen[cc.Code] = struct{}{}

		found, ok := iso4217.Lookup(cc.Code)
		assert.True(t, ok)
		assert.Equal(t, cc, found)
	}

	// One entry per alphabetic code, including withdrawn codes and CNH.
	_, ok := seen["USD"]
	assert.True(t, ok)
	_, ok = seen["CNH"]
	assert.True(t, ok)
	_, ok = seen["ANG"]
	assert.True(t, ok)
	_, ok = seen["BGN"]
	assert.True(t, ok)
}

func TestCurrentAndWithdrawn(t *testing.T) {
	assert.True(t, iso4217.USD.Current())
	assert.False(t, iso4217.USD.Withdrawn())

	assert.True(t, iso4217.XCG.Current())
	assert.False(t, iso4217.XCG.Withdrawn())

	// VES remains on official List One alongside VED.
	assert.True(t, iso4217.VES.Current())
	assert.False(t, iso4217.VES.Withdrawn())
	assert.True(t, iso4217.VED.Current())

	withdrawn := []iso4217.CurrencyCode{
		iso4217.ANG, iso4217.BGN, iso4217.CUC, iso4217.HRK, iso4217.SLL, iso4217.ZWL,
	}
	for _, cc := range withdrawn {
		assert.Truef(t, cc.Withdrawn(), "%s.Withdrawn()", cc.Code)
		assert.Falsef(t, cc.Current(), "%s.Current()", cc.Code)
		assert.True(t, cc.Valid())
	}

	// CNH is a non-ISO market code: lookup-able, but neither current nor withdrawn.
	assert.True(t, iso4217.CNH.Valid())
	assert.False(t, iso4217.CNH.Current())
	assert.False(t, iso4217.CNH.Withdrawn())

	// Same input rules as Lookup.
	assert.True(t, iso4217.CurrencyCode{Code: "usd"}.Current())
	assert.True(t, iso4217.CurrencyCode{Code: "840"}.Current())
	assert.True(t, iso4217.CurrencyCode{Code: "ang"}.Withdrawn())
	assert.True(t, iso4217.CurrencyCode{Code: "191"}.Withdrawn())  // HRK
	assert.False(t, iso4217.CurrencyCode{Code: "532"}.Withdrawn()) // XCG
	assert.True(t, iso4217.CurrencyCode{Code: "532"}.Current())

	assert.False(t, iso4217.CurrencyCode{}.Current())
	assert.False(t, iso4217.CurrencyCode{}.Withdrawn())
	assert.False(t, iso4217.CurrencyCode{Code: "XYZ"}.Current())
	assert.False(t, iso4217.CurrencyCode{Code: "XYZ"}.Withdrawn())
}

func TestAllCurrentWithdrawnPartition(t *testing.T) {
	var current, withdrawn, extension int
	for _, cc := range iso4217.All() {
		switch {
		case cc.Current():
			current++
			assert.False(t, cc.Withdrawn())
		case cc.Withdrawn():
			withdrawn++
			assert.False(t, cc.Current())
		default:
			extension++
			assert.Equal(t, "CNH", cc.Code)
		}
	}
	assert.Equal(t, 178, current)
	assert.Equal(t, 6, withdrawn)
	assert.Equal(t, 1, extension)
	assert.Equal(t, current+withdrawn+extension, len(iso4217.All()))
}
