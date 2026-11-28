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

// Package iso4217 implements constants and a validator for ISO 4217 currency codes.
//
// Lookup finds a currency by alphabetic code (for example "USD") or numeric code
// (for example "840"). The function trims surrounding whitespace, ignores case
// and pads short numeric input with leading zeros, so Lookup("8") and Lookup("008")
// both return ALL.
//
// The table includes official ISO 4217 List One codes, some withdrawn List Three
// codes (ANG, BGN, CUC, HRK, SLL, ZWL) so historical records still validate, and
// CNH (numeric 157), a non-ISO market code for the offshore Chinese yuan.
//
// XCG replaced ANG and inherited numeric code 532. Lookup("532") returns XCG.
// Lookup("ANG") still returns the withdrawn ANG entry.
//
// CurrencyCode.Valid reports whether Lookup finds the Code field. A value can be
// valid and still have an empty Name when the caller constructs CurrencyCode
// with only Code set. Use Lookup to get the canonical table entry.
//
// All returns every currency in the table, one entry per alphabetic code.
// CurrencyCode.Current is true for official List One codes. CurrencyCode.Withdrawn
// is true for the retained List Three codes. CNH is neither current nor withdrawn.
//
// FormatMinor and ParseMinor convert between integer minor units and a decimal
// string that uses CurrencyCode.DecimalPlaces.
package iso4217
