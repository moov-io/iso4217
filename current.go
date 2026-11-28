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

import "sort"

// withdrawnCodes is the set of alphabetic codes in this table that ISO 4217
// List Three (historic denominations, published 2026-01-01) lists as withdrawn,
// and that List One (published 2026-09-17) no longer carries. The package keeps
// the codes so historical records can still be looked up.
var withdrawnCodes = map[string]struct{}{
	"ANG": {}, // Netherlands Antillean Guilder, withdrawn 2025-03; succeeded by XCG
	"BGN": {}, // Bulgarian Lev, withdrawn 2026-01; succeeded by EUR
	"CUC": {}, // Peso Convertible, withdrawn 2021-06
	"HRK": {}, // Kuna, withdrawn 2023-01; succeeded by EUR
	"SLL": {}, // Leone, withdrawn 2023-12; succeeded by SLE
	"ZWL": {}, // Zimbabwe Dollar, withdrawn 2024-09; succeeded by ZWG
}

// All returns every currency in the table, one entry per alphabetic code,
// sorted by alphabetic code. The result includes withdrawn codes and CNH.
func All() []CurrencyCode {
	out := make([]CurrencyCode, 0, len(lookupTable)/2)
	for key, cc := range lookupTable {
		if key == cc.Code {
			out = append(out, cc)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Code < out[j].Code })
	return out
}

// Withdrawn reports whether the currency code is on the official ISO 4217
// List Three (historic denominations). The package keeps these codes so
// historical records can still be looked up. CNH is a non-ISO market code
// and Withdrawn returns false for it.
//
// The method uses the same input rules as Lookup, so CurrencyCode{Code: "ang"}
// and CurrencyCode{Code: "HRK"} both report as withdrawn.
func (cc CurrencyCode) Withdrawn() bool {
	found, ok := Lookup(cc.Code)
	if !ok {
		return false
	}
	_, withdrawn := withdrawnCodes[found.Code]
	return withdrawn
}

// Current reports whether the currency code is on the official ISO 4217
// List One (current currencies and funds). It is false for withdrawn codes
// and for the non-ISO market code CNH.
//
// The method uses the same input rules as Lookup.
func (cc CurrencyCode) Current() bool {
	found, ok := Lookup(cc.Code)
	if !ok {
		return false
	}
	if found.Code == "CNH" {
		return false
	}
	_, withdrawn := withdrawnCodes[found.Code]
	return !withdrawn
}
