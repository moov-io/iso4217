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

// Generated on 2023-07-25T04:47:20Z by daniel, any modifications will be overwritten
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

var (
	AFN = CurrencyCode{Code: "AFN", NumericCode: "971", Name: "Afghani", DecimalPlaces: 2}
	EUR = CurrencyCode{Code: "EUR", NumericCode: "978", Name: "Euro", DecimalPlaces: 2}
	ALL = CurrencyCode{Code: "ALL", NumericCode: "008", Name: "Lek", DecimalPlaces: 2}
	DZD = CurrencyCode{Code: "DZD", NumericCode: "012", Name: "Algerian Dinar", DecimalPlaces: 2}
	USD = CurrencyCode{Code: "USD", NumericCode: "840", Name: "US Dollar", DecimalPlaces: 2}
	AOA = CurrencyCode{Code: "AOA", NumericCode: "973", Name: "Kwanza", DecimalPlaces: 2}
	XCD = CurrencyCode{Code: "XCD", NumericCode: "951", Name: "East Caribbean Dollar", DecimalPlaces: 2}
	ARS = CurrencyCode{Code: "ARS", NumericCode: "032", Name: "Argentine Peso", DecimalPlaces: 2}
	AMD = CurrencyCode{Code: "AMD", NumericCode: "051", Name: "Armenian Dram", DecimalPlaces: 2}
	AWG = CurrencyCode{Code: "AWG", NumericCode: "533", Name: "Aruban Florin", DecimalPlaces: 2}
	AUD = CurrencyCode{Code: "AUD", NumericCode: "036", Name: "Australian Dollar", DecimalPlaces: 2}
	AZN = CurrencyCode{Code: "AZN", NumericCode: "944", Name: "Azerbaijan Manat", DecimalPlaces: 2}
	BSD = CurrencyCode{Code: "BSD", NumericCode: "044", Name: "Bahamian Dollar", DecimalPlaces: 2}
	BHD = CurrencyCode{Code: "BHD", NumericCode: "048", Name: "Bahraini Dinar", DecimalPlaces: 3}
	BDT = CurrencyCode{Code: "BDT", NumericCode: "050", Name: "Taka", DecimalPlaces: 2}
	BBD = CurrencyCode{Code: "BBD", NumericCode: "052", Name: "Barbados Dollar", DecimalPlaces: 2}
	BYN = CurrencyCode{Code: "BYN", NumericCode: "933", Name: "Belarusian Ruble", DecimalPlaces: 2}
	BZD = CurrencyCode{Code: "BZD", NumericCode: "084", Name: "Belize Dollar", DecimalPlaces: 2}
	XOF = CurrencyCode{Code: "XOF", NumericCode: "952", Name: "CFA Franc BCEAO", DecimalPlaces: 0}
	BMD = CurrencyCode{Code: "BMD", NumericCode: "060", Name: "Bermudian Dollar", DecimalPlaces: 2}
	INR = CurrencyCode{Code: "INR", NumericCode: "356", Name: "Indian Rupee", DecimalPlaces: 2}
	BTN = CurrencyCode{Code: "BTN", NumericCode: "064", Name: "Ngultrum", DecimalPlaces: 2}
	BOB = CurrencyCode{Code: "BOB", NumericCode: "068", Name: "Boliviano", DecimalPlaces: 2}
	BOV = CurrencyCode{Code: "BOV", NumericCode: "984", Name: "Mvdol", DecimalPlaces: 2}
	BAM = CurrencyCode{Code: "BAM", NumericCode: "977", Name: "Convertible Mark", DecimalPlaces: 2}
	BWP = CurrencyCode{Code: "BWP", NumericCode: "072", Name: "Pula", DecimalPlaces: 2}
	NOK = CurrencyCode{Code: "NOK", NumericCode: "578", Name: "Norwegian Krone", DecimalPlaces: 2}
	BRL = CurrencyCode{Code: "BRL", NumericCode: "986", Name: "Brazilian Real", DecimalPlaces: 2}
	BND = CurrencyCode{Code: "BND", NumericCode: "096", Name: "Brunei Dollar", DecimalPlaces: 2}
	BGN = CurrencyCode{Code: "BGN", NumericCode: "975", Name: "Bulgarian Lev", DecimalPlaces: 2}
	BIF = CurrencyCode{Code: "BIF", NumericCode: "108", Name: "Burundi Franc", DecimalPlaces: 0}
	CVE = CurrencyCode{Code: "CVE", NumericCode: "132", Name: "Cabo Verde Escudo", DecimalPlaces: 2}
	KHR = CurrencyCode{Code: "KHR", NumericCode: "116", Name: "Riel", DecimalPlaces: 2}
	XAF = CurrencyCode{Code: "XAF", NumericCode: "950", Name: "CFA Franc BEAC", DecimalPlaces: 0}
	CAD = CurrencyCode{Code: "CAD", NumericCode: "124", Name: "Canadian Dollar", DecimalPlaces: 2}
	KYD = CurrencyCode{Code: "KYD", NumericCode: "136", Name: "Cayman Islands Dollar", DecimalPlaces: 2}
	CLP = CurrencyCode{Code: "CLP", NumericCode: "152", Name: "Chilean Peso", DecimalPlaces: 0}
	CLF = CurrencyCode{Code: "CLF", NumericCode: "990", Name: "Unidad de Fomento", DecimalPlaces: 4}
	CNY = CurrencyCode{Code: "CNY", NumericCode: "156", Name: "Yuan Renminbi", DecimalPlaces: 2}
	COP = CurrencyCode{Code: "COP", NumericCode: "170", Name: "Colombian Peso", DecimalPlaces: 2}
	COU = CurrencyCode{Code: "COU", NumericCode: "970", Name: "Unidad de Valor Real", DecimalPlaces: 2}
	KMF = CurrencyCode{Code: "KMF", NumericCode: "174", Name: "Comorian Franc", DecimalPlaces: 0}
	CDF = CurrencyCode{Code: "CDF", NumericCode: "976", Name: "Congolese Franc", DecimalPlaces: 2}
	NZD = CurrencyCode{Code: "NZD", NumericCode: "554", Name: "New Zealand Dollar", DecimalPlaces: 2}
	CRC = CurrencyCode{Code: "CRC", NumericCode: "188", Name: "Costa Rican Colon", DecimalPlaces: 2}
	HRK = CurrencyCode{Code: "HRK", NumericCode: "191", Name: "Kuna", DecimalPlaces: 2}
	CUP = CurrencyCode{Code: "CUP", NumericCode: "192", Name: "Cuban Peso", DecimalPlaces: 2}
	CUC = CurrencyCode{Code: "CUC", NumericCode: "931", Name: "Peso Convertible", DecimalPlaces: 2}
	ANG = CurrencyCode{Code: "ANG", NumericCode: "532", Name: "Netherlands Antillean Guilder", DecimalPlaces: 2}
	CZK = CurrencyCode{Code: "CZK", NumericCode: "203", Name: "Czech Koruna", DecimalPlaces: 2}
	DKK = CurrencyCode{Code: "DKK", NumericCode: "208", Name: "Danish Krone", DecimalPlaces: 2}
	DJF = CurrencyCode{Code: "DJF", NumericCode: "262", Name: "Djibouti Franc", DecimalPlaces: 0}
	DOP = CurrencyCode{Code: "DOP", NumericCode: "214", Name: "Dominican Peso", DecimalPlaces: 2}
	EGP = CurrencyCode{Code: "EGP", NumericCode: "818", Name: "Egyptian Pound", DecimalPlaces: 2}
	SVC = CurrencyCode{Code: "SVC", NumericCode: "222", Name: "El Salvador Colon", DecimalPlaces: 2}
	ERN = CurrencyCode{Code: "ERN", NumericCode: "232", Name: "Nakfa", DecimalPlaces: 2}
	SZL = CurrencyCode{Code: "SZL", NumericCode: "748", Name: "Lilangeni", DecimalPlaces: 2}
	ETB = CurrencyCode{Code: "ETB", NumericCode: "230", Name: "Ethiopian Birr", DecimalPlaces: 2}
	FKP = CurrencyCode{Code: "FKP", NumericCode: "238", Name: "Falkland Islands Pound", DecimalPlaces: 2}
	FJD = CurrencyCode{Code: "FJD", NumericCode: "242", Name: "Fiji Dollar", DecimalPlaces: 2}
	XPF = CurrencyCode{Code: "XPF", NumericCode: "953", Name: "CFP Franc", DecimalPlaces: 0}
	GMD = CurrencyCode{Code: "GMD", NumericCode: "270", Name: "Dalasi", DecimalPlaces: 2}
	GEL = CurrencyCode{Code: "GEL", NumericCode: "981", Name: "Lari", DecimalPlaces: 2}
	GHS = CurrencyCode{Code: "GHS", NumericCode: "936", Name: "Ghana Cedi", DecimalPlaces: 2}
	GIP = CurrencyCode{Code: "GIP", NumericCode: "292", Name: "Gibraltar Pound", DecimalPlaces: 2}
	GTQ = CurrencyCode{Code: "GTQ", NumericCode: "320", Name: "Quetzal", DecimalPlaces: 2}
	GBP = CurrencyCode{Code: "GBP", NumericCode: "826", Name: "Pound Sterling", DecimalPlaces: 2}
	GNF = CurrencyCode{Code: "GNF", NumericCode: "324", Name: "Guinean Franc", DecimalPlaces: 0}
	GYD = CurrencyCode{Code: "GYD", NumericCode: "328", Name: "Guyana Dollar", DecimalPlaces: 2}
	HTG = CurrencyCode{Code: "HTG", NumericCode: "332", Name: "Gourde", DecimalPlaces: 2}
	HNL = CurrencyCode{Code: "HNL", NumericCode: "340", Name: "Lempira", DecimalPlaces: 2}
	HKD = CurrencyCode{Code: "HKD", NumericCode: "344", Name: "Hong Kong Dollar", DecimalPlaces: 2}
	HUF = CurrencyCode{Code: "HUF", NumericCode: "348", Name: "Forint", DecimalPlaces: 2}
	ISK = CurrencyCode{Code: "ISK", NumericCode: "352", Name: "Iceland Krona", DecimalPlaces: 0}
	IDR = CurrencyCode{Code: "IDR", NumericCode: "360", Name: "Rupiah", DecimalPlaces: 2}
	XDR = CurrencyCode{Code: "XDR", NumericCode: "960", Name: "SDR (Special Drawing Right)", DecimalPlaces: 0}
	IRR = CurrencyCode{Code: "IRR", NumericCode: "364", Name: "Iranian Rial", DecimalPlaces: 2}
	IQD = CurrencyCode{Code: "IQD", NumericCode: "368", Name: "Iraqi Dinar", DecimalPlaces: 3}
	ILS = CurrencyCode{Code: "ILS", NumericCode: "376", Name: "New Israeli Sheqel", DecimalPlaces: 2}
	JMD = CurrencyCode{Code: "JMD", NumericCode: "388", Name: "Jamaican Dollar", DecimalPlaces: 2}
	JPY = CurrencyCode{Code: "JPY", NumericCode: "392", Name: "Yen", DecimalPlaces: 0}
	JOD = CurrencyCode{Code: "JOD", NumericCode: "400", Name: "Jordanian Dinar", DecimalPlaces: 3}
	KZT = CurrencyCode{Code: "KZT", NumericCode: "398", Name: "Tenge", DecimalPlaces: 2}
	KES = CurrencyCode{Code: "KES", NumericCode: "404", Name: "Kenyan Shilling", DecimalPlaces: 2}
	KPW = CurrencyCode{Code: "KPW", NumericCode: "408", Name: "North Korean Won", DecimalPlaces: 2}
	KRW = CurrencyCode{Code: "KRW", NumericCode: "410", Name: "Won", DecimalPlaces: 0}
	KWD = CurrencyCode{Code: "KWD", NumericCode: "414", Name: "Kuwaiti Dinar", DecimalPlaces: 3}
	KGS = CurrencyCode{Code: "KGS", NumericCode: "417", Name: "Som", DecimalPlaces: 2}
	LAK = CurrencyCode{Code: "LAK", NumericCode: "418", Name: "Lao Kip", DecimalPlaces: 2}
	LBP = CurrencyCode{Code: "LBP", NumericCode: "422", Name: "Lebanese Pound", DecimalPlaces: 2}
	LSL = CurrencyCode{Code: "LSL", NumericCode: "426", Name: "Loti", DecimalPlaces: 2}
	ZAR = CurrencyCode{Code: "ZAR", NumericCode: "710", Name: "Rand", DecimalPlaces: 2}
	LRD = CurrencyCode{Code: "LRD", NumericCode: "430", Name: "Liberian Dollar", DecimalPlaces: 2}
	LYD = CurrencyCode{Code: "LYD", NumericCode: "434", Name: "Libyan Dinar", DecimalPlaces: 3}
	CHF = CurrencyCode{Code: "CHF", NumericCode: "756", Name: "Swiss Franc", DecimalPlaces: 2}
	MOP = CurrencyCode{Code: "MOP", NumericCode: "446", Name: "Pataca", DecimalPlaces: 2}
	MKD = CurrencyCode{Code: "MKD", NumericCode: "807", Name: "Denar", DecimalPlaces: 2}
	MGA = CurrencyCode{Code: "MGA", NumericCode: "969", Name: "Malagasy Ariary", DecimalPlaces: 2}
	MWK = CurrencyCode{Code: "MWK", NumericCode: "454", Name: "Malawi Kwacha", DecimalPlaces: 2}
	MYR = CurrencyCode{Code: "MYR", NumericCode: "458", Name: "Malaysian Ringgit", DecimalPlaces: 2}
	MVR = CurrencyCode{Code: "MVR", NumericCode: "462", Name: "Rufiyaa", DecimalPlaces: 2}
	MRU = CurrencyCode{Code: "MRU", NumericCode: "929", Name: "Ouguiya", DecimalPlaces: 2}
	MUR = CurrencyCode{Code: "MUR", NumericCode: "480", Name: "Mauritius Rupee", DecimalPlaces: 2}
	XUA = CurrencyCode{Code: "XUA", NumericCode: "965", Name: "ADB Unit of Account", DecimalPlaces: 0}
	MXN = CurrencyCode{Code: "MXN", NumericCode: "484", Name: "Mexican Peso", DecimalPlaces: 2}
	MXV = CurrencyCode{Code: "MXV", NumericCode: "979", Name: "Mexican Unidad de Inversion (UDI)", DecimalPlaces: 2}
	MDL = CurrencyCode{Code: "MDL", NumericCode: "498", Name: "Moldovan Leu", DecimalPlaces: 2}
	MNT = CurrencyCode{Code: "MNT", NumericCode: "496", Name: "Tugrik", DecimalPlaces: 2}
	MAD = CurrencyCode{Code: "MAD", NumericCode: "504", Name: "Moroccan Dirham", DecimalPlaces: 2}
	MZN = CurrencyCode{Code: "MZN", NumericCode: "943", Name: "Mozambique Metical", DecimalPlaces: 2}
	MMK = CurrencyCode{Code: "MMK", NumericCode: "104", Name: "Kyat", DecimalPlaces: 2}
	NAD = CurrencyCode{Code: "NAD", NumericCode: "516", Name: "Namibia Dollar", DecimalPlaces: 2}
	NPR = CurrencyCode{Code: "NPR", NumericCode: "524", Name: "Nepalese Rupee", DecimalPlaces: 2}
	NIO = CurrencyCode{Code: "NIO", NumericCode: "558", Name: "Cordoba Oro", DecimalPlaces: 2}
	NGN = CurrencyCode{Code: "NGN", NumericCode: "566", Name: "Naira", DecimalPlaces: 2}
	OMR = CurrencyCode{Code: "OMR", NumericCode: "512", Name: "Rial Omani", DecimalPlaces: 3}
	PKR = CurrencyCode{Code: "PKR", NumericCode: "586", Name: "Pakistan Rupee", DecimalPlaces: 2}
	PAB = CurrencyCode{Code: "PAB", NumericCode: "590", Name: "Balboa", DecimalPlaces: 2}
	PGK = CurrencyCode{Code: "PGK", NumericCode: "598", Name: "Kina", DecimalPlaces: 2}
	PYG = CurrencyCode{Code: "PYG", NumericCode: "600", Name: "Guarani", DecimalPlaces: 0}
	PEN = CurrencyCode{Code: "PEN", NumericCode: "604", Name: "Sol", DecimalPlaces: 2}
	PHP = CurrencyCode{Code: "PHP", NumericCode: "608", Name: "Philippine Peso", DecimalPlaces: 2}
	PLN = CurrencyCode{Code: "PLN", NumericCode: "985", Name: "Zloty", DecimalPlaces: 2}
	QAR = CurrencyCode{Code: "QAR", NumericCode: "634", Name: "Qatari Rial", DecimalPlaces: 2}
	RON = CurrencyCode{Code: "RON", NumericCode: "946", Name: "Romanian Leu", DecimalPlaces: 2}
	RUB = CurrencyCode{Code: "RUB", NumericCode: "643", Name: "Russian Ruble", DecimalPlaces: 2}
	RWF = CurrencyCode{Code: "RWF", NumericCode: "646", Name: "Rwanda Franc", DecimalPlaces: 0}
	SHP = CurrencyCode{Code: "SHP", NumericCode: "654", Name: "Saint Helena Pound", DecimalPlaces: 2}
	WST = CurrencyCode{Code: "WST", NumericCode: "882", Name: "Tala", DecimalPlaces: 2}
	STN = CurrencyCode{Code: "STN", NumericCode: "930", Name: "Dobra", DecimalPlaces: 2}
	SAR = CurrencyCode{Code: "SAR", NumericCode: "682", Name: "Saudi Riyal", DecimalPlaces: 2}
	RSD = CurrencyCode{Code: "RSD", NumericCode: "941", Name: "Serbian Dinar", DecimalPlaces: 2}
	SCR = CurrencyCode{Code: "SCR", NumericCode: "690", Name: "Seychelles Rupee", DecimalPlaces: 2}
	SLL = CurrencyCode{Code: "SLL", NumericCode: "694", Name: "Leone", DecimalPlaces: 2}
	SGD = CurrencyCode{Code: "SGD", NumericCode: "702", Name: "Singapore Dollar", DecimalPlaces: 2}
	XSU = CurrencyCode{Code: "XSU", NumericCode: "994", Name: "Sucre", DecimalPlaces: 0}
	SBD = CurrencyCode{Code: "SBD", NumericCode: "090", Name: "Solomon Islands Dollar", DecimalPlaces: 2}
	SOS = CurrencyCode{Code: "SOS", NumericCode: "706", Name: "Somali Shilling", DecimalPlaces: 2}
	SSP = CurrencyCode{Code: "SSP", NumericCode: "728", Name: "South Sudanese Pound", DecimalPlaces: 2}
	LKR = CurrencyCode{Code: "LKR", NumericCode: "144", Name: "Sri Lanka Rupee", DecimalPlaces: 2}
	SDG = CurrencyCode{Code: "SDG", NumericCode: "938", Name: "Sudanese Pound", DecimalPlaces: 2}
	SRD = CurrencyCode{Code: "SRD", NumericCode: "968", Name: "Surinam Dollar", DecimalPlaces: 2}
	SEK = CurrencyCode{Code: "SEK", NumericCode: "752", Name: "Swedish Krona", DecimalPlaces: 2}
	CHE = CurrencyCode{Code: "CHE", NumericCode: "947", Name: "WIR Euro", DecimalPlaces: 2}
	CHW = CurrencyCode{Code: "CHW", NumericCode: "948", Name: "WIR Franc", DecimalPlaces: 2}
	SYP = CurrencyCode{Code: "SYP", NumericCode: "760", Name: "Syrian Pound", DecimalPlaces: 2}
	TWD = CurrencyCode{Code: "TWD", NumericCode: "901", Name: "New Taiwan Dollar", DecimalPlaces: 2}
	TJS = CurrencyCode{Code: "TJS", NumericCode: "972", Name: "Somoni", DecimalPlaces: 2}
	TZS = CurrencyCode{Code: "TZS", NumericCode: "834", Name: "Tanzanian Shilling", DecimalPlaces: 2}
	THB = CurrencyCode{Code: "THB", NumericCode: "764", Name: "Baht", DecimalPlaces: 2}
	TOP = CurrencyCode{Code: "TOP", NumericCode: "776", Name: "Pa'anga", DecimalPlaces: 2}
	TTD = CurrencyCode{Code: "TTD", NumericCode: "780", Name: "Trinidad and Tobago Dollar", DecimalPlaces: 2}
	TND = CurrencyCode{Code: "TND", NumericCode: "788", Name: "Tunisian Dinar", DecimalPlaces: 3}
	TRY = CurrencyCode{Code: "TRY", NumericCode: "949", Name: "Turkish Lira", DecimalPlaces: 2}
	TMT = CurrencyCode{Code: "TMT", NumericCode: "934", Name: "Turkmenistan New Manat", DecimalPlaces: 2}
	UGX = CurrencyCode{Code: "UGX", NumericCode: "800", Name: "Uganda Shilling", DecimalPlaces: 0}
	UAH = CurrencyCode{Code: "UAH", NumericCode: "980", Name: "Hryvnia", DecimalPlaces: 2}
	AED = CurrencyCode{Code: "AED", NumericCode: "784", Name: "UAE Dirham", DecimalPlaces: 2}
	USN = CurrencyCode{Code: "USN", NumericCode: "997", Name: "US Dollar (Next day)", DecimalPlaces: 2}
	UYU = CurrencyCode{Code: "UYU", NumericCode: "858", Name: "Peso Uruguayo", DecimalPlaces: 2}
	UYI = CurrencyCode{Code: "UYI", NumericCode: "940", Name: "Uruguay Peso en Unidades Indexadas (UI)", DecimalPlaces: 0}
	UYW = CurrencyCode{Code: "UYW", NumericCode: "927", Name: "Unidad Previsional", DecimalPlaces: 4}
	UZS = CurrencyCode{Code: "UZS", NumericCode: "860", Name: "Uzbekistan Sum", DecimalPlaces: 2}
	VUV = CurrencyCode{Code: "VUV", NumericCode: "548", Name: "Vatu", DecimalPlaces: 0}
	VES = CurrencyCode{Code: "VES", NumericCode: "928", Name: "Bolívar Soberano", DecimalPlaces: 2}
	VND = CurrencyCode{Code: "VND", NumericCode: "704", Name: "Dong", DecimalPlaces: 0}
	YER = CurrencyCode{Code: "YER", NumericCode: "886", Name: "Yemeni Rial", DecimalPlaces: 2}
	ZMW = CurrencyCode{Code: "ZMW", NumericCode: "967", Name: "Zambian Kwacha", DecimalPlaces: 2}
	ZWL = CurrencyCode{Code: "ZWL", NumericCode: "932", Name: "Zimbabwe Dollar", DecimalPlaces: 2}
	XBA = CurrencyCode{Code: "XBA", NumericCode: "955", Name: "Bond Markets Unit European Composite Unit (EURCO)", DecimalPlaces: 0}
	XBB = CurrencyCode{Code: "XBB", NumericCode: "956", Name: "Bond Markets Unit European Monetary Unit (E.M.U.-6)", DecimalPlaces: 0}
	XBC = CurrencyCode{Code: "XBC", NumericCode: "957", Name: "Bond Markets Unit European Unit of Account 9 (E.U.A.-9)", DecimalPlaces: 0}
	XBD = CurrencyCode{Code: "XBD", NumericCode: "958", Name: "Bond Markets Unit European Unit of Account 17 (E.U.A.-17)", DecimalPlaces: 0}
	XTS = CurrencyCode{Code: "XTS", NumericCode: "963", Name: "Codes specifically reserved for testing purposes", DecimalPlaces: 0}
	XXX = CurrencyCode{Code: "XXX", NumericCode: "999", Name: "The codes assigned for transactions where no currency is involved", DecimalPlaces: 0}
	XAU = CurrencyCode{Code: "XAU", NumericCode: "959", Name: "Gold", DecimalPlaces: 0}
	XPD = CurrencyCode{Code: "XPD", NumericCode: "964", Name: "Palladium", DecimalPlaces: 0}
	XPT = CurrencyCode{Code: "XPT", NumericCode: "962", Name: "Platinum", DecimalPlaces: 0}
	XAG = CurrencyCode{Code: "XAG", NumericCode: "961", Name: "Silver", DecimalPlaces: 0}
)

var lookupTable = map[string]CurrencyCode{
	"AFN": AFN, // Afghani
	"971": AFN, // Afghani
	"EUR": EUR, // Euro
	"978": EUR, // Euro
	"ALL": ALL, // Lek
	"008": ALL, // Lek
	"DZD": DZD, // Algerian Dinar
	"012": DZD, // Algerian Dinar
	"USD": USD, // US Dollar
	"840": USD, // US Dollar
	"AOA": AOA, // Kwanza
	"973": AOA, // Kwanza
	"XCD": XCD, // East Caribbean Dollar
	"951": XCD, // East Caribbean Dollar
	"ARS": ARS, // Argentine Peso
	"032": ARS, // Argentine Peso
	"AMD": AMD, // Armenian Dram
	"051": AMD, // Armenian Dram
	"AWG": AWG, // Aruban Florin
	"533": AWG, // Aruban Florin
	"AUD": AUD, // Australian Dollar
	"036": AUD, // Australian Dollar
	"AZN": AZN, // Azerbaijan Manat
	"944": AZN, // Azerbaijan Manat
	"BSD": BSD, // Bahamian Dollar
	"044": BSD, // Bahamian Dollar
	"BHD": BHD, // Bahraini Dinar
	"048": BHD, // Bahraini Dinar
	"BDT": BDT, // Taka
	"050": BDT, // Taka
	"BBD": BBD, // Barbados Dollar
	"052": BBD, // Barbados Dollar
	"BYN": BYN, // Belarusian Ruble
	"933": BYN, // Belarusian Ruble
	"BZD": BZD, // Belize Dollar
	"084": BZD, // Belize Dollar
	"XOF": XOF, // CFA Franc BCEAO
	"952": XOF, // CFA Franc BCEAO
	"BMD": BMD, // Bermudian Dollar
	"060": BMD, // Bermudian Dollar
	"INR": INR, // Indian Rupee
	"356": INR, // Indian Rupee
	"BTN": BTN, // Ngultrum
	"064": BTN, // Ngultrum
	"BOB": BOB, // Boliviano
	"068": BOB, // Boliviano
	"BOV": BOV, // Mvdol
	"984": BOV, // Mvdol
	"BAM": BAM, // Convertible Mark
	"977": BAM, // Convertible Mark
	"BWP": BWP, // Pula
	"072": BWP, // Pula
	"NOK": NOK, // Norwegian Krone
	"578": NOK, // Norwegian Krone
	"BRL": BRL, // Brazilian Real
	"986": BRL, // Brazilian Real
	"BND": BND, // Brunei Dollar
	"096": BND, // Brunei Dollar
	"BGN": BGN, // Bulgarian Lev
	"975": BGN, // Bulgarian Lev
	"BIF": BIF, // Burundi Franc
	"108": BIF, // Burundi Franc
	"CVE": CVE, // Cabo Verde Escudo
	"132": CVE, // Cabo Verde Escudo
	"KHR": KHR, // Riel
	"116": KHR, // Riel
	"XAF": XAF, // CFA Franc BEAC
	"950": XAF, // CFA Franc BEAC
	"CAD": CAD, // Canadian Dollar
	"124": CAD, // Canadian Dollar
	"KYD": KYD, // Cayman Islands Dollar
	"136": KYD, // Cayman Islands Dollar
	"CLP": CLP, // Chilean Peso
	"152": CLP, // Chilean Peso
	"CLF": CLF, // Unidad de Fomento
	"990": CLF, // Unidad de Fomento
	"CNY": CNY, // Yuan Renminbi
	"156": CNY, // Yuan Renminbi
	"COP": COP, // Colombian Peso
	"170": COP, // Colombian Peso
	"COU": COU, // Unidad de Valor Real
	"970": COU, // Unidad de Valor Real
	"KMF": KMF, // Comorian Franc
	"174": KMF, // Comorian Franc
	"CDF": CDF, // Congolese Franc
	"976": CDF, // Congolese Franc
	"NZD": NZD, // New Zealand Dollar
	"554": NZD, // New Zealand Dollar
	"CRC": CRC, // Costa Rican Colon
	"188": CRC, // Costa Rican Colon
	"HRK": HRK, // Kuna
	"191": HRK, // Kuna
	"CUP": CUP, // Cuban Peso
	"192": CUP, // Cuban Peso
	"CUC": CUC, // Peso Convertible
	"931": CUC, // Peso Convertible
	"ANG": ANG, // Netherlands Antillean Guilder
	"532": ANG, // Netherlands Antillean Guilder
	"CZK": CZK, // Czech Koruna
	"203": CZK, // Czech Koruna
	"DKK": DKK, // Danish Krone
	"208": DKK, // Danish Krone
	"DJF": DJF, // Djibouti Franc
	"262": DJF, // Djibouti Franc
	"DOP": DOP, // Dominican Peso
	"214": DOP, // Dominican Peso
	"EGP": EGP, // Egyptian Pound
	"818": EGP, // Egyptian Pound
	"SVC": SVC, // El Salvador Colon
	"222": SVC, // El Salvador Colon
	"ERN": ERN, // Nakfa
	"232": ERN, // Nakfa
	"SZL": SZL, // Lilangeni
	"748": SZL, // Lilangeni
	"ETB": ETB, // Ethiopian Birr
	"230": ETB, // Ethiopian Birr
	"FKP": FKP, // Falkland Islands Pound
	"238": FKP, // Falkland Islands Pound
	"FJD": FJD, // Fiji Dollar
	"242": FJD, // Fiji Dollar
	"XPF": XPF, // CFP Franc
	"953": XPF, // CFP Franc
	"GMD": GMD, // Dalasi
	"270": GMD, // Dalasi
	"GEL": GEL, // Lari
	"981": GEL, // Lari
	"GHS": GHS, // Ghana Cedi
	"936": GHS, // Ghana Cedi
	"GIP": GIP, // Gibraltar Pound
	"292": GIP, // Gibraltar Pound
	"GTQ": GTQ, // Quetzal
	"320": GTQ, // Quetzal
	"GBP": GBP, // Pound Sterling
	"826": GBP, // Pound Sterling
	"GNF": GNF, // Guinean Franc
	"324": GNF, // Guinean Franc
	"GYD": GYD, // Guyana Dollar
	"328": GYD, // Guyana Dollar
	"HTG": HTG, // Gourde
	"332": HTG, // Gourde
	"HNL": HNL, // Lempira
	"340": HNL, // Lempira
	"HKD": HKD, // Hong Kong Dollar
	"344": HKD, // Hong Kong Dollar
	"HUF": HUF, // Forint
	"348": HUF, // Forint
	"ISK": ISK, // Iceland Krona
	"352": ISK, // Iceland Krona
	"IDR": IDR, // Rupiah
	"360": IDR, // Rupiah
	"XDR": XDR, // SDR (Special Drawing Right)
	"960": XDR, // SDR (Special Drawing Right)
	"IRR": IRR, // Iranian Rial
	"364": IRR, // Iranian Rial
	"IQD": IQD, // Iraqi Dinar
	"368": IQD, // Iraqi Dinar
	"ILS": ILS, // New Israeli Sheqel
	"376": ILS, // New Israeli Sheqel
	"JMD": JMD, // Jamaican Dollar
	"388": JMD, // Jamaican Dollar
	"JPY": JPY, // Yen
	"392": JPY, // Yen
	"JOD": JOD, // Jordanian Dinar
	"400": JOD, // Jordanian Dinar
	"KZT": KZT, // Tenge
	"398": KZT, // Tenge
	"KES": KES, // Kenyan Shilling
	"404": KES, // Kenyan Shilling
	"KPW": KPW, // North Korean Won
	"408": KPW, // North Korean Won
	"KRW": KRW, // Won
	"410": KRW, // Won
	"KWD": KWD, // Kuwaiti Dinar
	"414": KWD, // Kuwaiti Dinar
	"KGS": KGS, // Som
	"417": KGS, // Som
	"LAK": LAK, // Lao Kip
	"418": LAK, // Lao Kip
	"LBP": LBP, // Lebanese Pound
	"422": LBP, // Lebanese Pound
	"LSL": LSL, // Loti
	"426": LSL, // Loti
	"ZAR": ZAR, // Rand
	"710": ZAR, // Rand
	"LRD": LRD, // Liberian Dollar
	"430": LRD, // Liberian Dollar
	"LYD": LYD, // Libyan Dinar
	"434": LYD, // Libyan Dinar
	"CHF": CHF, // Swiss Franc
	"756": CHF, // Swiss Franc
	"MOP": MOP, // Pataca
	"446": MOP, // Pataca
	"MKD": MKD, // Denar
	"807": MKD, // Denar
	"MGA": MGA, // Malagasy Ariary
	"969": MGA, // Malagasy Ariary
	"MWK": MWK, // Malawi Kwacha
	"454": MWK, // Malawi Kwacha
	"MYR": MYR, // Malaysian Ringgit
	"458": MYR, // Malaysian Ringgit
	"MVR": MVR, // Rufiyaa
	"462": MVR, // Rufiyaa
	"MRU": MRU, // Ouguiya
	"929": MRU, // Ouguiya
	"MUR": MUR, // Mauritius Rupee
	"480": MUR, // Mauritius Rupee
	"XUA": XUA, // ADB Unit of Account
	"965": XUA, // ADB Unit of Account
	"MXN": MXN, // Mexican Peso
	"484": MXN, // Mexican Peso
	"MXV": MXV, // Mexican Unidad de Inversion (UDI)
	"979": MXV, // Mexican Unidad de Inversion (UDI)
	"MDL": MDL, // Moldovan Leu
	"498": MDL, // Moldovan Leu
	"MNT": MNT, // Tugrik
	"496": MNT, // Tugrik
	"MAD": MAD, // Moroccan Dirham
	"504": MAD, // Moroccan Dirham
	"MZN": MZN, // Mozambique Metical
	"943": MZN, // Mozambique Metical
	"MMK": MMK, // Kyat
	"104": MMK, // Kyat
	"NAD": NAD, // Namibia Dollar
	"516": NAD, // Namibia Dollar
	"NPR": NPR, // Nepalese Rupee
	"524": NPR, // Nepalese Rupee
	"NIO": NIO, // Cordoba Oro
	"558": NIO, // Cordoba Oro
	"NGN": NGN, // Naira
	"566": NGN, // Naira
	"OMR": OMR, // Rial Omani
	"512": OMR, // Rial Omani
	"PKR": PKR, // Pakistan Rupee
	"586": PKR, // Pakistan Rupee
	"PAB": PAB, // Balboa
	"590": PAB, // Balboa
	"PGK": PGK, // Kina
	"598": PGK, // Kina
	"PYG": PYG, // Guarani
	"600": PYG, // Guarani
	"PEN": PEN, // Sol
	"604": PEN, // Sol
	"PHP": PHP, // Philippine Peso
	"608": PHP, // Philippine Peso
	"PLN": PLN, // Zloty
	"985": PLN, // Zloty
	"QAR": QAR, // Qatari Rial
	"634": QAR, // Qatari Rial
	"RON": RON, // Romanian Leu
	"946": RON, // Romanian Leu
	"RUB": RUB, // Russian Ruble
	"643": RUB, // Russian Ruble
	"RWF": RWF, // Rwanda Franc
	"646": RWF, // Rwanda Franc
	"SHP": SHP, // Saint Helena Pound
	"654": SHP, // Saint Helena Pound
	"WST": WST, // Tala
	"882": WST, // Tala
	"STN": STN, // Dobra
	"930": STN, // Dobra
	"SAR": SAR, // Saudi Riyal
	"682": SAR, // Saudi Riyal
	"RSD": RSD, // Serbian Dinar
	"941": RSD, // Serbian Dinar
	"SCR": SCR, // Seychelles Rupee
	"690": SCR, // Seychelles Rupee
	"SLL": SLL, // Leone
	"694": SLL, // Leone
	"SGD": SGD, // Singapore Dollar
	"702": SGD, // Singapore Dollar
	"XSU": XSU, // Sucre
	"994": XSU, // Sucre
	"SBD": SBD, // Solomon Islands Dollar
	"090": SBD, // Solomon Islands Dollar
	"SOS": SOS, // Somali Shilling
	"706": SOS, // Somali Shilling
	"SSP": SSP, // South Sudanese Pound
	"728": SSP, // South Sudanese Pound
	"LKR": LKR, // Sri Lanka Rupee
	"144": LKR, // Sri Lanka Rupee
	"SDG": SDG, // Sudanese Pound
	"938": SDG, // Sudanese Pound
	"SRD": SRD, // Surinam Dollar
	"968": SRD, // Surinam Dollar
	"SEK": SEK, // Swedish Krona
	"752": SEK, // Swedish Krona
	"CHE": CHE, // WIR Euro
	"947": CHE, // WIR Euro
	"CHW": CHW, // WIR Franc
	"948": CHW, // WIR Franc
	"SYP": SYP, // Syrian Pound
	"760": SYP, // Syrian Pound
	"TWD": TWD, // New Taiwan Dollar
	"901": TWD, // New Taiwan Dollar
	"TJS": TJS, // Somoni
	"972": TJS, // Somoni
	"TZS": TZS, // Tanzanian Shilling
	"834": TZS, // Tanzanian Shilling
	"THB": THB, // Baht
	"764": THB, // Baht
	"TOP": TOP, // Pa'anga
	"776": TOP, // Pa'anga
	"TTD": TTD, // Trinidad and Tobago Dollar
	"780": TTD, // Trinidad and Tobago Dollar
	"TND": TND, // Tunisian Dinar
	"788": TND, // Tunisian Dinar
	"TRY": TRY, // Turkish Lira
	"949": TRY, // Turkish Lira
	"TMT": TMT, // Turkmenistan New Manat
	"934": TMT, // Turkmenistan New Manat
	"UGX": UGX, // Uganda Shilling
	"800": UGX, // Uganda Shilling
	"UAH": UAH, // Hryvnia
	"980": UAH, // Hryvnia
	"AED": AED, // UAE Dirham
	"784": AED, // UAE Dirham
	"USN": USN, // US Dollar (Next day)
	"997": USN, // US Dollar (Next day)
	"UYU": UYU, // Peso Uruguayo
	"858": UYU, // Peso Uruguayo
	"UYI": UYI, // Uruguay Peso en Unidades Indexadas (UI)
	"940": UYI, // Uruguay Peso en Unidades Indexadas (UI)
	"UYW": UYW, // Unidad Previsional
	"927": UYW, // Unidad Previsional
	"UZS": UZS, // Uzbekistan Sum
	"860": UZS, // Uzbekistan Sum
	"VUV": VUV, // Vatu
	"548": VUV, // Vatu
	"VES": VES, // Bolívar Soberano
	"928": VES, // Bolívar Soberano
	"VND": VND, // Dong
	"704": VND, // Dong
	"YER": YER, // Yemeni Rial
	"886": YER, // Yemeni Rial
	"ZMW": ZMW, // Zambian Kwacha
	"967": ZMW, // Zambian Kwacha
	"ZWL": ZWL, // Zimbabwe Dollar
	"932": ZWL, // Zimbabwe Dollar
	"XBA": XBA, // Bond Markets Unit European Composite Unit (EURCO)
	"955": XBA, // Bond Markets Unit European Composite Unit (EURCO)
	"XBB": XBB, // Bond Markets Unit European Monetary Unit (E.M.U.-6)
	"956": XBB, // Bond Markets Unit European Monetary Unit (E.M.U.-6)
	"XBC": XBC, // Bond Markets Unit European Unit of Account 9 (E.U.A.-9)
	"957": XBC, // Bond Markets Unit European Unit of Account 9 (E.U.A.-9)
	"XBD": XBD, // Bond Markets Unit European Unit of Account 17 (E.U.A.-17)
	"958": XBD, // Bond Markets Unit European Unit of Account 17 (E.U.A.-17)
	"XTS": XTS, // Codes specifically reserved for testing purposes
	"963": XTS, // Codes specifically reserved for testing purposes
	"XXX": XXX, // The codes assigned for transactions where no currency is involved
	"999": XXX, // The codes assigned for transactions where no currency is involved
	"XAU": XAU, // Gold
	"959": XAU, // Gold
	"XPD": XPD, // Palladium
	"964": XPD, // Palladium
	"XPT": XPT, // Platinum
	"962": XPT, // Platinum
	"XAG": XAG, // Silver
	"961": XAG, // Silver
}
