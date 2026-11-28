[![Moov Banner Logo](https://user-images.githubusercontent.com/20115216/104214617-885b3c80-53ec-11eb-8ce0-9fc745fb5bfc.png)](https://github.com/moov-io)

<p align="center">
  <a href="https://pkg.go.dev/github.com/moov-io/iso4217">Package Documentation</a>
  ·
  <a href="https://slack.moov.io/">Community</a>
  ·
  <a href="https://moov.io/blog/">Blog</a>
  <br>
  <br>
</p>

[![GoDoc](https://pkg.go.dev/badge/github.com/moov-io/iso4217.svg)](https://pkg.go.dev/github.com/moov-io/iso4217)
[![Build Status](https://github.com/moov-io/iso4217/workflows/Go/badge.svg)](https://github.com/moov-io/iso4217/actions)
[![Coverage Status](https://codecov.io/gh/moov-io/iso4217/branch/master/graph/badge.svg)](https://codecov.io/gh/moov-io/iso4217)
[![Go Report Card](https://goreportcard.com/badge/github.com/moov-io/iso4217)](https://goreportcard.com/report/github.com/moov-io/iso4217)
[![Repo Size](https://img.shields.io/github/languages/code-size/moov-io/iso4217?label=project%20size)](https://github.com/moov-io/iso4217)
[![Apache 2 License](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/moov-io/iso4217/master/LICENSE)
[![Slack Channel](https://img.shields.io/badge/Slack-Moov-e01563?logo=slack)](https://slack.moov.io/)
[![GitHub Stars](https://img.shields.io/github/stars/moov-io/iso4217)](https://github.com/moov-io/iso4217)
[![Twitter](https://img.shields.io/twitter/follow/moov?style=social)](https://twitter.com/moov?lang=en)

# moov-io/iso4217

The `iso4217` package implements constants and a validator for ISO 4217 currency codes.

## Install

This project uses [Go Modules](https://go.dev/blog/using-go-modules) and Go 1.20 or newer. See [Golang's install instructions](https://go.dev/doc/install) for help setting up Go.

```
go get github.com/moov-io/iso4217@latest
```

To clone the repository:

```
git clone https://github.com/moov-io/iso4217.git
```

Use a [tagged release](https://github.com/moov-io/iso4217/releases/latest) in production.

## Usage

Look up a currency by its alphabetic or numeric code. `Lookup` trims surrounding whitespace, ignores case and pads short numeric codes with leading zeros (`"8"` and `"008"` both return ALL).

```go
package main

import (
	"fmt"

	"github.com/moov-io/iso4217"
)

func main() {
	cc, ok := iso4217.Lookup("usd")
	if !ok {
		return
	}
	fmt.Println(cc.Code, cc.NumericCode, cc.Name, cc.DecimalPlaces)
	// USD 840 US Dollar 2

	cc, ok = iso4217.Lookup("36")
	fmt.Println(ok, cc.Code) // true AUD

	fmt.Println(iso4217.USD.Valid())           // true
	fmt.Println(iso4217.JPY.DecimalPlaces)     // 0
	fmt.Println(iso4217.BHD.DecimalPlaces)     // 3

	fmt.Println(iso4217.FormatMinor(123456, iso4217.USD)) // 1234.56
	fmt.Println(iso4217.FormatMinor(123456, iso4217.JPY)) // 123456
	fmt.Println(iso4217.FormatMinor(123456, iso4217.BHD)) // 123.456
}
```

The package also exports a `CurrencyCode` constant for each code in the table (for example `iso4217.USD`).

`CurrencyCode.Valid` reports whether `Lookup` finds the `Code` field. It does not require a canonical, fully populated struct. `CurrencyCode{Code: "usd"}.Valid()` and `CurrencyCode{Code: "840"}.Valid()` are true, but `Name` and `DecimalPlaces` stay empty on those values. Call `Lookup` when you need the table entry.

`All` returns every currency in the table, one entry per alphabetic code, sorted by code. `CurrencyCode.Current` is true for official ISO 4217 List One codes. `CurrencyCode.Withdrawn` is true for the historic List Three codes that the table retains (ANG, BGN, CUC, HRK, SLL and ZWL). CNH is neither current nor withdrawn.

`DecimalPlaces` is the ISO 4217 minor unit. `FormatMinor` and `ParseMinor` convert between an integer amount in minor units and a decimal string that uses `DecimalPlaces` (123456 with USD is `"1234.56"`; with JPY, `"123456"`; with BHD, `"123.456"`).

Runnable examples are in [example_test.go](example_test.go).

### Currency table policy

The table keeps every official ISO 4217 List One code, some withdrawn List Three codes for historical lookups, and one non-ISO market code:

- **CNH** (numeric 157, "Yuan Fen") is the offshore Chinese yuan. It is not an official ISO 4217 code. The package includes it as a convenience extension. `Lookup("CNH")` and `Lookup("157")` succeed.
- **Withdrawn codes** stay in the table so historical records still validate. The withdrawn codes in the table are ANG, BGN, CUC, HRK, SLL and ZWL (ISO 4217 List Three, published 2026-01-01). `Lookup` and `Valid` return true for them.
- **Numeric 532** is assigned to XCG (Caribbean Guilder), which replaced ANG. `Lookup("532")` returns XCG. `Lookup("ANG")` still returns ANG, whose `NumericCode` remains `"532"`. The numeric code does not round-trip for ANG.

VES remains on List One alongside VED.

## Getting help

 channel | info
 ------- | -------
Twitter [@moov](https://twitter.com/moov)	| You can follow Moov.io's Twitter feed to get updates on our project(s). You can also tweet us questions or just share blogs or stories.
[GitHub Issue](https://github.com/moov-io/iso4217/issues/new) | If you are able to reproduce a problem please open a GitHub Issue under the specific project that caused the error.
[moov-io slack](https://slack.moov.io/) | Join our slack channel to have an interactive discussion about the development of the project.

## Supported and tested platforms

- 64-bit Linux (Ubuntu, Debian), macOS, and Windows

## Contributing

Yes please! Please review our [Contributing guide](https://github.com/moov-io/ach/blob/master/CONTRIBUTING.md) and [Code of Conduct](https://github.com/moov-io/ach/blob/master/CODE_OF_CONDUCT.md) to get started.

`make check` downloads [`lint-project.sh`](https://github.com/moov-io/infra/blob/master/go/lint-project.sh) from moov-io/infra and runs it. The script requires 100% test coverage. On Windows, `make check` skips the lint script; `go test ./...` still runs on Windows in CI.

## License

Apache License 2.0 - See [LICENSE](LICENSE) for details.
