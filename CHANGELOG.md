## Unreleased

DOCUMENTATION

- Fix README header links and the Slack badge
- Add install and usage examples
- Document Lookup input handling, Valid, CNH and withdrawn codes
- Add a package comment and Example functions
- Record the required Go version (1.20 or newer) and the `make check` lint script

## v0.4.0 (Released 2026-07-20)

IMPROVEMENTS

- Add SLE, VED, XCG and XAD to the currency table
- For numeric code 532, return XCG (the current assignment). Keep ANG available by its alphabetic code for historical lookups.

## v0.3.2 (Released 2024-09-10)

Improvements

- Add new Yuan Fen currency code (CNH)
- Update dependencies

## v0.3.1 (Released 2024-09-10)

IMPROVEMENTS

- Add new Zimbabwe Gold currency code (ZWG)

## v0.3.0 (Released 2023-07-26)

IMPROVEMENTS

- feat: Add support for numeric codes in CurrencyCode lookup

## v0.2.1 (Released 2023-06-06)

UPDATES

- update `DecimalPlaces` type to be `uint8` instead of `int`

## v0.2.0 (Released 2023-06-05)

IMPROVEMENTS

- feat: Add ISO 4217 MicroUnit as `DecimalPlaces` on CurrencyCode

## v0.1.0 (Released 2023-06-02)

This is the initial release of `iso4217` which was extracted from moov-io/ach.

The complete release history is also on [GitHub Releases](https://github.com/moov-io/iso4217/releases).
