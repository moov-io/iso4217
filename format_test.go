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
	"math"
	"strconv"
	"testing"

	"github.com/moov-io/iso4217"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFormatMinor(t *testing.T) {
	assert.Equal(t, "1234.56", iso4217.FormatMinor(123456, iso4217.USD))
	assert.Equal(t, "123456", iso4217.FormatMinor(123456, iso4217.JPY))
	assert.Equal(t, "123.456", iso4217.FormatMinor(123456, iso4217.BHD))
	assert.Equal(t, "12.3456", iso4217.FormatMinor(123456, iso4217.CLF))
	assert.Equal(t, "-1.50", iso4217.FormatMinor(-150, iso4217.USD))
	assert.Equal(t, "0.00", iso4217.FormatMinor(0, iso4217.USD))
	assert.Equal(t, "0", iso4217.FormatMinor(0, iso4217.JPY))
	assert.Equal(t, "-123456", iso4217.FormatMinor(-123456, iso4217.JPY))
	assert.Equal(t, "0.01", iso4217.FormatMinor(1, iso4217.USD))
	assert.Equal(t, "-0.01", iso4217.FormatMinor(-1, iso4217.USD))
	assert.Equal(t, "0.001", iso4217.FormatMinor(1, iso4217.BHD))

	assert.Equal(t, strconv.FormatInt(math.MaxInt64, 10), iso4217.FormatMinor(math.MaxInt64, iso4217.JPY))
	assert.Equal(t, strconv.FormatInt(math.MinInt64, 10), iso4217.FormatMinor(math.MinInt64, iso4217.JPY))
	assert.Equal(t, "92233720368547758.07", iso4217.FormatMinor(math.MaxInt64, iso4217.USD))
	assert.Equal(t, "-92233720368547758.08", iso4217.FormatMinor(math.MinInt64, iso4217.USD))
}

func TestParseMinor(t *testing.T) {
	cases := []struct {
		amount string
		cc     iso4217.CurrencyCode
		want   int64
	}{
		{"1234.56", iso4217.USD, 123456},
		{"123456", iso4217.JPY, 123456},
		{"123.456", iso4217.BHD, 123456},
		{"12.3456", iso4217.CLF, 123456},
		{"-1.50", iso4217.USD, -150},
		{"+1.5", iso4217.USD, 150},
		{"  1.5  ", iso4217.USD, 150},
		{".5", iso4217.USD, 50},
		{"1.", iso4217.USD, 100},
		{"0", iso4217.JPY, 0},
		{"-0.00", iso4217.USD, 0},
	}
	for _, tc := range cases {
		n, err := iso4217.ParseMinor(tc.amount, tc.cc)
		require.NoErrorf(t, err, "ParseMinor(%q, %s)", tc.amount, tc.cc.Code)
		assert.Equalf(t, tc.want, n, "ParseMinor(%q, %s)", tc.amount, tc.cc.Code)
	}
}

func TestParseMinorErrors(t *testing.T) {
	cases := []string{"", "   ", "+", "-", ".", "-.", "abc", "1.2.3", "1,234.56"}
	for _, amount := range cases {
		_, err := iso4217.ParseMinor(amount, iso4217.USD)
		assert.Errorf(t, err, "ParseMinor(%q)", amount)
	}

	_, err := iso4217.ParseMinor("1.234", iso4217.USD)
	assert.Error(t, err)

	_, err = iso4217.ParseMinor("123.0", iso4217.JPY)
	assert.Error(t, err)

	_, err = iso4217.ParseMinor("1.2", iso4217.JPY)
	assert.Error(t, err)

	_, err = iso4217.ParseMinor("9223372036854775808", iso4217.JPY)
	assert.Error(t, err)

	// Negative overflow (magnitude larger than -math.MinInt64).
	_, err = iso4217.ParseMinor("-9223372036854775809", iso4217.JPY)
	assert.Error(t, err)

	// Overflow uint64 during parse.
	_, err = iso4217.ParseMinor("18446744073709551616", iso4217.JPY)
	assert.Error(t, err)
}

func TestFormatParseMinorRoundTrip(t *testing.T) {
	amounts := []int64{0, 1, -1, 150, -150, 123456, math.MaxInt64, math.MinInt64}
	codes := []iso4217.CurrencyCode{iso4217.JPY, iso4217.USD, iso4217.BHD, iso4217.CLF}
	for _, cc := range codes {
		for _, amount := range amounts {
			formatted := iso4217.FormatMinor(amount, cc)
			parsed, err := iso4217.ParseMinor(formatted, cc)
			require.NoErrorf(t, err, "ParseMinor(%q, %s)", formatted, cc.Code)
			assert.Equalf(t, amount, parsed, "round trip %d %s -> %s", amount, cc.Code, formatted)
		}
	}
}
