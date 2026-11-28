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
	"fmt"

	"github.com/moov-io/iso4217"
)

func ExampleLookup() {
	cc, ok := iso4217.Lookup("usd")
	fmt.Println(ok, cc.Code, cc.NumericCode, cc.Name, cc.DecimalPlaces)
	// Output: true USD 840 US Dollar 2
}

func ExampleLookup_numeric() {
	cc, ok := iso4217.Lookup("36")
	fmt.Println(ok, cc.Code)

	cc, ok = iso4217.Lookup("8")
	fmt.Println(ok, cc.Code)
	// Output:
	// true AUD
	// true ALL
}

func ExampleCurrencyCode_Valid() {
	fmt.Println(iso4217.USD.Valid())
	fmt.Println(iso4217.CurrencyCode{Code: "usd"}.Valid())
	fmt.Println(iso4217.CurrencyCode{Code: "840"}.Valid())
	fmt.Println(iso4217.CurrencyCode{Code: "XYZ"}.Valid())
	// Output:
	// true
	// true
	// true
	// false
}

func ExampleAll() {
	codes := iso4217.All()
	fmt.Println(codes[0].Code, len(codes) > 1)
	// Output: AED true
}

func ExampleCurrencyCode_Current() {
	fmt.Println(iso4217.USD.Current(), iso4217.USD.Withdrawn())
	fmt.Println(iso4217.ANG.Current(), iso4217.ANG.Withdrawn())
	fmt.Println(iso4217.CNH.Current(), iso4217.CNH.Withdrawn())
	// Output:
	// true false
	// false true
	// false false
}

func ExampleFormatMinor() {
	fmt.Println(iso4217.FormatMinor(123456, iso4217.USD))
	fmt.Println(iso4217.FormatMinor(123456, iso4217.JPY))
	fmt.Println(iso4217.FormatMinor(123456, iso4217.BHD))

	minor, err := iso4217.ParseMinor("1234.56", iso4217.USD)
	if err != nil {
		panic(err)
	}
	fmt.Println(minor)
	// Output:
	// 1234.56
	// 123456
	// 123.456
	// 123456
}
