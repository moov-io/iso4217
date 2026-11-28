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
	"strings"
)

// Lookup returns the ISO 4217 CurrencyCode for an alphabetic or numeric code,
// and a boolean that is true when the table contains a match.
//
// The function trims surrounding whitespace, ignores case and pads short
// numeric input with leading zeros. Examples: "USD", "usd", "840", "36", "036", "8".
func Lookup(code string) (CurrencyCode, bool) {
	cc, exists := lookupTable[normalize(code)]
	return cc, exists
}

func normalize(code string) string {
	return strings.ToUpper(fmt.Sprintf("%03s", strings.TrimSpace(code)))
}
