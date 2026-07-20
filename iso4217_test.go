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
)

// Currencies added to the official ISO 4217 List One (active) after the table was
// last topped up: SLE (Sierra Leone, 2022), VED (Venezuela, 2021), XCG (Curaçao/
// Sint Maarten, 2025) and XAD (Arab Monetary Fund, 2025). All carry a 2-digit minor
// unit per List One.
func TestActiveListOneAdditions(t *testing.T) {
	cases := []struct {
		alpha, numeric, name string
		decimals             uint8
	}{
		{"SLE", "925", "Leone", 2},
		{"VED", "926", "Bolívar Soberano", 2},
		{"XCG", "532", "Caribbean Guilder", 2},
		{"XAD", "396", "Arab Accounting Dinar", 2},
	}
	for _, tc := range cases {
		cc, exists := iso4217.Lookup(tc.alpha)
		assert.Truef(t, exists, "Lookup(%q)", tc.alpha)
		assert.Equal(t, tc.alpha, cc.Code)
		assert.Equal(t, tc.numeric, cc.NumericCode)
		assert.Equal(t, tc.name, cc.Name)
		assert.Equal(t, tc.decimals, cc.DecimalPlaces)
	}
}

// The three non-colliding numeric codes resolve directly to the new currency.
func TestActiveListOneNumericLookups(t *testing.T) {
	for numeric, alpha := range map[string]string{"925": "SLE", "926": "VED", "396": "XAD"} {
		cc, exists := iso4217.Lookup(numeric)
		assert.Truef(t, exists, "Lookup(%q)", numeric)
		assert.Equal(t, alpha, cc.Code)
	}
}

// XCG replaced ANG and inherited ISO numeric code 532. Numeric 532 resolves to the
// currently-assigned currency (XCG); the withdrawn ANG stays reachable by its alpha
// code so historical records still validate.
func TestNumericCode532IsXCG(t *testing.T) {
	cc, exists := iso4217.Lookup("532")
	assert.True(t, exists)
	assert.Equal(t, "XCG", cc.Code)

	cc, exists = iso4217.Lookup("ANG")
	assert.True(t, exists)
	assert.Equal(t, "ANG", cc.Code)
	assert.Equal(t, "532", cc.NumericCode)
}

// Withdrawn predecessors are deliberately retained for historical lookups.
func TestWithdrawnPredecessorsRetained(t *testing.T) {
	for _, code := range []string{"SLL", "VES", "ANG"} {
		_, exists := iso4217.Lookup(code)
		assert.Truef(t, exists, "Lookup(%q)", code)
	}
}

func TestCurrencyCodeString(t *testing.T) {
	assert.Equal(t, "USD", iso4217.USD.String())
	assert.Equal(t, "EUR", iso4217.EUR.String())

	cc, exists := iso4217.Lookup("AUD")
	assert.True(t, exists)
	assert.Equal(t, "AUD", cc.String())
}

func TestCurrencyCodeValid(t *testing.T) {
	assert.True(t, iso4217.USD.Valid())
	assert.True(t, iso4217.EUR.Valid())

	// Zero-value and unknown codes are invalid.
	assert.False(t, iso4217.CurrencyCode{}.Valid())
	assert.False(t, iso4217.CurrencyCode{Code: "QZA"}.Valid())
}
