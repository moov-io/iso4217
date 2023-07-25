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

//go:build ignore
// +build ignore

// Generates iso4217.go.
//
// This file grabs the ISO 4217 currency codes and writes them
// into source code so we don't rely on any external files (zip,
// json, etc).
//
// The data is pulled from datahub.io as the ISO.org site only offers
// XML.
//
// https://datahub.io/core/currency-codes
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"log"
	"net/http"
	"os"
	"os/user"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var (
	downloadUrl    = "https://datahub.io/core/currency-codes/r/codes-all.json"
	outputFilename = "iso4217.go"

	charCleaner = strings.NewReplacer(`"`, ``, `’`, `'`, `’`, `'`, `“`, `"`, `”`, `"`, `\`, ``)
)

// {"AlphabeticCode": "AFN", "NumericCode": 971.0, "Currency": "Afghani", ... }
type currency struct {
	Code        string  `json:"AlphabeticCode"`
	NumericCode float64 `json:"NumericCode"`
	Name        string  `json:"Currency"`
	MinorUnit   string  `json:"MinorUnit"`
}

func main() {
	when := time.Now().Format("2006-01-02T03:04:05Z")
	who, err := user.Current()
	if err != nil {
		log.Fatalf("Unable to get user on %s", runtime.GOOS)
	}

	// Write copyright header
	var buf bytes.Buffer
	fmt.Fprintf(&buf, `// Licensed to The Moov Authors under one or more contributor
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

// Generated on %s by %s, any modifications will be overwritten
package iso4217

type CurrencyCode struct {
    Code, NumericCode, Name string

    // DecimalPlaces represents the unsigned integer value of a currency's minor unit.
	// DecimalPlaces is 0 if the currency doesn't have a minor unit.
    DecimalPlaces uint8
}

func (cc CurrencyCode) String() string {
    return cc.Code
}

func (cc CurrencyCode) Valid() bool {
    _, exists := Lookup(cc.Code)
    return exists
}
`, when, who.Username)

	// Download certs
	resp, err := http.Get(downloadUrl)
	if err != nil {
		log.Fatalf("error while downloading %s: %v", downloadUrl, err)
	}
	defer resp.Body.Close()

	var currencies []currency
	if err := json.NewDecoder(resp.Body).Decode(&currencies); err != nil {
		log.Fatalf("error while parsing currency response: %v", err)
	}

	// The JSON file contains duplicates so we need to dedup them..
	cs := make(map[string]bool, 150)

	// Write countries to source code
	var varBuffer bytes.Buffer
	fmt.Fprintln(&varBuffer, "var (")

	var lookupBuffer bytes.Buffer
	fmt.Fprintln(&lookupBuffer, "var lookupTable = map[string]CurrencyCode{")

	var skipCount int
	for i := range currencies {
		code, name, minorunit := currencies[i].Code, currencies[i].Name, currencies[i].MinorUnit
		if code == "" || name == "" || minorunit == "" {
			fmt.Printf("SKIPPING: code=%q currency=%q minorunit=%q\n", code, name, minorunit)
			skipCount++
			continue
		}

		name = charCleaner.Replace(name)
		minorunit = charCleaner.Replace(minorunit)
		if _, exists := cs[code]; !exists {
			cs[code] = true // mark as seen

			var decimalPlaces uint8 = 0
			if minorunit != "-" {
				d, err := strconv.Atoi(minorunit)
				if err != nil {
					log.Fatalf("error while parsing currency minor unit: %v", err)
				}
				decimalPlaces = uint8(d)
			}

			numericCode := fmt.Sprintf("%03d", int(currencies[i].NumericCode))
			fmt.Fprintf(&varBuffer, fmt.Sprintf(`  %s = CurrencyCode{Code: "%s", NumericCode: "%s", Name: "%s", DecimalPlaces: %d}`+"\n", code, code, numericCode, name, decimalPlaces))
			fmt.Fprintf(&lookupBuffer, fmt.Sprintf(`"%s": %s, // %s`+"\n", code, code, name))
			fmt.Fprintf(&lookupBuffer, fmt.Sprintf(`"%s": %s, // %s`+"\n", numericCode, code, name))
		}
	}
	fmt.Printf("Skipped %d currencies\n", skipCount)

	fmt.Fprintln(&varBuffer, ")")
	fmt.Fprintln(&lookupBuffer, "}")

	// Add code to file
	buf.Write(varBuffer.Bytes())
	buf.WriteString("\n")
	buf.Write(lookupBuffer.Bytes())

	// format source code and write file
	out, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Println(buf.String())
		log.Fatalf("error formatting output code, err=%v", err)
	}

	err = os.WriteFile(outputFilename, out, 0644)
	if err != nil {
		log.Fatalf("error writing file, err=%v", err)
	}
}
