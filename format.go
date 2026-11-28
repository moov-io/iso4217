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

package iso4217

import (
	"fmt"
	"strconv"
	"strings"
)

// FormatMinor formats an integer amount in minor units as a decimal string
// that uses cc.DecimalPlaces digits after the point. When DecimalPlaces is 0,
// the function returns the amount as a whole number.
//
//	FormatMinor(123456, USD) // "1234.56"
//	FormatMinor(123456, JPY) // "123456"
//	FormatMinor(123456, BHD) // "123.456"
//	FormatMinor(-150, USD)   // "-1.50"
func FormatMinor(minor int64, cc CurrencyCode) string {
	places := int(cc.DecimalPlaces)
	if places == 0 {
		return strconv.FormatInt(minor, 10)
	}

	s := strconv.FormatInt(minor, 10)
	negative := false
	if s[0] == '-' {
		negative = true
		s = s[1:]
	}
	for len(s) <= places {
		s = "0" + s
	}
	s = s[:len(s)-places] + "." + s[len(s)-places:]
	if negative {
		return "-" + s
	}
	return s
}

// ParseMinor parses a decimal amount into an integer number of minor units
// that uses cc.DecimalPlaces. The amount may include a leading plus or minus
// sign and surrounding whitespace. The number of fraction digits must not
// exceed DecimalPlaces. A currency with DecimalPlaces 0 does not accept a
// decimal point.
func ParseMinor(amount string, cc CurrencyCode) (int64, error) {
	s := strings.TrimSpace(amount)
	if s == "" {
		return 0, fmt.Errorf("iso4217: empty amount")
	}

	sign := int64(1)
	switch s[0] {
	case '+':
		s = s[1:]
	case '-':
		sign = -1
		s = s[1:]
	}
	if s == "" || s == "." {
		return 0, fmt.Errorf("iso4217: invalid amount %q", amount)
	}

	places := int(cc.DecimalPlaces)
	whole, frac, found := strings.Cut(s, ".")
	if found && places == 0 {
		return 0, fmt.Errorf("iso4217: %q has more than %d decimal places", amount, places)
	}
	if whole == "" {
		whole = "0"
	}
	if !allDigits(whole) || (found && !allDigits(frac)) {
		return 0, fmt.Errorf("iso4217: invalid amount %q", amount)
	}
	if found && len(frac) > places {
		return 0, fmt.Errorf("iso4217: %q has more than %d decimal places", amount, places)
	}
	if !found {
		frac = ""
	}
	for len(frac) < places {
		frac += "0"
	}

	digits := whole + frac
	if sign < 0 && digits != "0" {
		digits = "-" + digits
	}
	n, err := strconv.ParseInt(digits, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("iso4217: invalid amount %q: %w", amount, err)
	}
	return n, nil
}

func allDigits(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}
