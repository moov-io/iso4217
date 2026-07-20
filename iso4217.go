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
	AED = CurrencyCode{Code: "AED", NumericCode: "784", Name: "UAE Dirham", DecimalPlaces: 2}
	AFN = CurrencyCode{Code: "AFN", NumericCode: "971", Name: "Afghani", DecimalPlaces: 2}
	ALL = CurrencyCode{Code: "ALL", NumericCode: "008", Name: "Lek", DecimalPlaces: 2}
	AMD = CurrencyCode{Code: "AMD", NumericCode: "051", Name: "Armenian Dram", DecimalPlaces: 2}
	ANG = CurrencyCode{Code: "ANG", NumericCode: "532", Name: "Netherlands Antillean Guilder", DecimalPlaces: 2}
	AOA = CurrencyCode{Code: "AOA", NumericCode: "973", Name: "Kwanza", DecimalPlaces: 2}
	ARS = CurrencyCode{Code: "ARS", NumericCode: "032", Name: "Argentine Peso", DecimalPlaces: 2}
	AUD = CurrencyCode{Code: "AUD", NumericCode: "036", Name: "Australian Dollar", DecimalPlaces: 2}
	AWG = CurrencyCode{Code: "AWG", NumericCode: "533", Name: "Aruban Florin", DecimalPlaces: 2}
	AZN = CurrencyCode{Code: "AZN", NumericCode: "944", Name: "Azerbaijan Manat", DecimalPlaces: 2}
	BAM = CurrencyCode{Code: "BAM", NumericCode: "977", Name: "Convertible Mark", DecimalPlaces: 2}
	BBD = CurrencyCode{Code: "BBD", NumericCode: "052", Name: "Barbados Dollar", DecimalPlaces: 2}
	BDT = CurrencyCode{Code: "BDT", NumericCode: "050", Name: "Taka", DecimalPlaces: 2}
	BGN = CurrencyCode{Code: "BGN", NumericCode: "975", Name: "Bulgarian Lev", DecimalPlaces: 2}
	BHD = CurrencyCode{Code: "BHD", NumericCode: "048", Name: "Bahraini Dinar", DecimalPlaces: 3}
	BIF = CurrencyCode{Code: "BIF", NumericCode: "108", Name: "Burundi Franc", DecimalPlaces: 0}
	BMD = CurrencyCode{Code: "BMD", NumericCode: "060", Name: "Bermudian Dollar", DecimalPlaces: 2}
	BND = CurrencyCode{Code: "BND", NumericCode: "096", Name: "Brunei Dollar", DecimalPlaces: 2}
	BOB = CurrencyCode{Code: "BOB", NumericCode: "068", Name: "Boliviano", DecimalPlaces: 2}
	BOV = CurrencyCode{Code: "BOV", NumericCode: "984", Name: "Mvdol", DecimalPlaces: 2}
	BRL = CurrencyCode{Code: "BRL", NumericCode: "986", Name: "Brazilian Real", DecimalPlaces: 2}
	BSD = CurrencyCode{Code: "BSD", NumericCode: "044", Name: "Bahamian Dollar", DecimalPlaces: 2}
	BTN = CurrencyCode{Code: "BTN", NumericCode: "064", Name: "Ngultrum", DecimalPlaces: 2}
	BWP = CurrencyCode{Code: "BWP", NumericCode: "072", Name: "Pula", DecimalPlaces: 2}
	BYN = CurrencyCode{Code: "BYN", NumericCode: "933", Name: "Belarusian Ruble", DecimalPlaces: 2}
	BZD = CurrencyCode{Code: "BZD", NumericCode: "084", Name: "Belize Dollar", DecimalPlaces: 2}
	CAD = CurrencyCode{Code: "CAD", NumericCode: "124", Name: "Canadian Dollar", DecimalPlaces: 2}
	CDF = CurrencyCode{Code: "CDF", NumericCode: "976", Name: "Congolese Franc", DecimalPlaces: 2}
	CHE = CurrencyCode{Code: "CHE", NumericCode: "947", Name: "WIR Euro", DecimalPlaces: 2}
	CHF = CurrencyCode{Code: "CHF", NumericCode: "756", Name: "Swiss Franc", DecimalPlaces: 2}
	CHW = CurrencyCode{Code: "CHW", NumericCode: "948", Name: "WIR Franc", DecimalPlaces: 2}
	CLF = CurrencyCode{Code: "CLF", NumericCode: "990", Name: "Unidad de Fomento", DecimalPlaces: 4}
	CLP = CurrencyCode{Code: "CLP", NumericCode: "152", Name: "Chilean Peso", DecimalPlaces: 0}
	CNH = CurrencyCode{Code: "CNH", NumericCode: "157", Name: "Yuan Fen", DecimalPlaces: 2}
	CNY = CurrencyCode{Code: "CNY", NumericCode: "156", Name: "Yuan Renminbi", DecimalPlaces: 2}
	COP = CurrencyCode{Code: "COP", NumericCode: "170", Name: "Colombian Peso", DecimalPlaces: 2}
	COU = CurrencyCode{Code: "COU", NumericCode: "970", Name: "Unidad de Valor Real", DecimalPlaces: 2}
	CRC = CurrencyCode{Code: "CRC", NumericCode: "188", Name: "Costa Rican Colon", DecimalPlaces: 2}
	CUC = CurrencyCode{Code: "CUC", NumericCode: "931", Name: "Peso Convertible", DecimalPlaces: 2}
	CUP = CurrencyCode{Code: "CUP", NumericCode: "192", Name: "Cuban Peso", DecimalPlaces: 2}
	CVE = CurrencyCode{Code: "CVE", NumericCode: "132", Name: "Cabo Verde Escudo", DecimalPlaces: 2}
	CZK = CurrencyCode{Code: "CZK", NumericCode: "203", Name: "Czech Koruna", DecimalPlaces: 2}
	DJF = CurrencyCode{Code: "DJF", NumericCode: "262", Name: "Djibouti Franc", DecimalPlaces: 0}
	DKK = CurrencyCode{Code: "DKK", NumericCode: "208", Name: "Danish Krone", DecimalPlaces: 2}
	DOP = CurrencyCode{Code: "DOP", NumericCode: "214", Name: "Dominican Peso", DecimalPlaces: 2}
	DZD = CurrencyCode{Code: "DZD", NumericCode: "012", Name: "Algerian Dinar", DecimalPlaces: 2}
	EGP = CurrencyCode{Code: "EGP", NumericCode: "818", Name: "Egyptian Pound", DecimalPlaces: 2}
	ERN = CurrencyCode{Code: "ERN", NumericCode: "232", Name: "Nakfa", DecimalPlaces: 2}
	ETB = CurrencyCode{Code: "ETB", NumericCode: "230", Name: "Ethiopian Birr", DecimalPlaces: 2}
	EUR = CurrencyCode{Code: "EUR", NumericCode: "978", Name: "Euro", DecimalPlaces: 2}
	FJD = CurrencyCode{Code: "FJD", NumericCode: "242", Name: "Fiji Dollar", DecimalPlaces: 2}
	FKP = CurrencyCode{Code: "FKP", NumericCode: "238", Name: "Falkland Islands Pound", DecimalPlaces: 2}
	GBP = CurrencyCode{Code: "GBP", NumericCode: "826", Name: "Pound Sterling", DecimalPlaces: 2}
	GEL = CurrencyCode{Code: "GEL", NumericCode: "981", Name: "Lari", DecimalPlaces: 2}
	GHS = CurrencyCode{Code: "GHS", NumericCode: "936", Name: "Ghana Cedi", DecimalPlaces: 2}
	GIP = CurrencyCode{Code: "GIP", NumericCode: "292", Name: "Gibraltar Pound", DecimalPlaces: 2}
	GMD = CurrencyCode{Code: "GMD", NumericCode: "270", Name: "Dalasi", DecimalPlaces: 2}
	GNF = CurrencyCode{Code: "GNF", NumericCode: "324", Name: "Guinean Franc", DecimalPlaces: 0}
	GTQ = CurrencyCode{Code: "GTQ", NumericCode: "320", Name: "Quetzal", DecimalPlaces: 2}
	GYD = CurrencyCode{Code: "GYD", NumericCode: "328", Name: "Guyana Dollar", DecimalPlaces: 2}
	HKD = CurrencyCode{Code: "HKD", NumericCode: "344", Name: "Hong Kong Dollar", DecimalPlaces: 2}
	HNL = CurrencyCode{Code: "HNL", NumericCode: "340", Name: "Lempira", DecimalPlaces: 2}
	HRK = CurrencyCode{Code: "HRK", NumericCode: "191", Name: "Kuna", DecimalPlaces: 2}
	HTG = CurrencyCode{Code: "HTG", NumericCode: "332", Name: "Gourde", DecimalPlaces: 2}
	HUF = CurrencyCode{Code: "HUF", NumericCode: "348", Name: "Forint", DecimalPlaces: 2}
	IDR = CurrencyCode{Code: "IDR", NumericCode: "360", Name: "Rupiah", DecimalPlaces: 2}
	ILS = CurrencyCode{Code: "ILS", NumericCode: "376", Name: "New Israeli Sheqel", DecimalPlaces: 2}
	INR = CurrencyCode{Code: "INR", NumericCode: "356", Name: "Indian Rupee", DecimalPlaces: 2}
	IQD = CurrencyCode{Code: "IQD", NumericCode: "368", Name: "Iraqi Dinar", DecimalPlaces: 3}
	IRR = CurrencyCode{Code: "IRR", NumericCode: "364", Name: "Iranian Rial", DecimalPlaces: 2}
	ISK = CurrencyCode{Code: "ISK", NumericCode: "352", Name: "Iceland Krona", DecimalPlaces: 0}
	JMD = CurrencyCode{Code: "JMD", NumericCode: "388", Name: "Jamaican Dollar", DecimalPlaces: 2}
	JOD = CurrencyCode{Code: "JOD", NumericCode: "400", Name: "Jordanian Dinar", DecimalPlaces: 3}
	JPY = CurrencyCode{Code: "JPY", NumericCode: "392", Name: "Yen", DecimalPlaces: 0}
	KES = CurrencyCode{Code: "KES", NumericCode: "404", Name: "Kenyan Shilling", DecimalPlaces: 2}
	KGS = CurrencyCode{Code: "KGS", NumericCode: "417", Name: "Som", DecimalPlaces: 2}
	KHR = CurrencyCode{Code: "KHR", NumericCode: "116", Name: "Riel", DecimalPlaces: 2}
	KMF = CurrencyCode{Code: "KMF", NumericCode: "174", Name: "Comorian Franc", DecimalPlaces: 0}
	KPW = CurrencyCode{Code: "KPW", NumericCode: "408", Name: "North Korean Won", DecimalPlaces: 2}
	KRW = CurrencyCode{Code: "KRW", NumericCode: "410", Name: "Won", DecimalPlaces: 0}
	KWD = CurrencyCode{Code: "KWD", NumericCode: "414", Name: "Kuwaiti Dinar", DecimalPlaces: 3}
	KYD = CurrencyCode{Code: "KYD", NumericCode: "136", Name: "Cayman Islands Dollar", DecimalPlaces: 2}
	KZT = CurrencyCode{Code: "KZT", NumericCode: "398", Name: "Tenge", DecimalPlaces: 2}
	LAK = CurrencyCode{Code: "LAK", NumericCode: "418", Name: "Lao Kip", DecimalPlaces: 2}
	LBP = CurrencyCode{Code: "LBP", NumericCode: "422", Name: "Lebanese Pound", DecimalPlaces: 2}
	LKR = CurrencyCode{Code: "LKR", NumericCode: "144", Name: "Sri Lanka Rupee", DecimalPlaces: 2}
	LRD = CurrencyCode{Code: "LRD", NumericCode: "430", Name: "Liberian Dollar", DecimalPlaces: 2}
	LSL = CurrencyCode{Code: "LSL", NumericCode: "426", Name: "Loti", DecimalPlaces: 2}
	LYD = CurrencyCode{Code: "LYD", NumericCode: "434", Name: "Libyan Dinar", DecimalPlaces: 3}
	MAD = CurrencyCode{Code: "MAD", NumericCode: "504", Name: "Moroccan Dirham", DecimalPlaces: 2}
	MDL = CurrencyCode{Code: "MDL", NumericCode: "498", Name: "Moldovan Leu", DecimalPlaces: 2}
	MGA = CurrencyCode{Code: "MGA", NumericCode: "969", Name: "Malagasy Ariary", DecimalPlaces: 2}
	MKD = CurrencyCode{Code: "MKD", NumericCode: "807", Name: "Denar", DecimalPlaces: 2}
	MMK = CurrencyCode{Code: "MMK", NumericCode: "104", Name: "Kyat", DecimalPlaces: 2}
	MNT = CurrencyCode{Code: "MNT", NumericCode: "496", Name: "Tugrik", DecimalPlaces: 2}
	MOP = CurrencyCode{Code: "MOP", NumericCode: "446", Name: "Pataca", DecimalPlaces: 2}
	MRU = CurrencyCode{Code: "MRU", NumericCode: "929", Name: "Ouguiya", DecimalPlaces: 2}
	MUR = CurrencyCode{Code: "MUR", NumericCode: "480", Name: "Mauritius Rupee", DecimalPlaces: 2}
	MVR = CurrencyCode{Code: "MVR", NumericCode: "462", Name: "Rufiyaa", DecimalPlaces: 2}
	MWK = CurrencyCode{Code: "MWK", NumericCode: "454", Name: "Malawi Kwacha", DecimalPlaces: 2}
	MXN = CurrencyCode{Code: "MXN", NumericCode: "484", Name: "Mexican Peso", DecimalPlaces: 2}
	MXV = CurrencyCode{Code: "MXV", NumericCode: "979", Name: "Mexican Unidad de Inversion (UDI)", DecimalPlaces: 2}
	MYR = CurrencyCode{Code: "MYR", NumericCode: "458", Name: "Malaysian Ringgit", DecimalPlaces: 2}
	MZN = CurrencyCode{Code: "MZN", NumericCode: "943", Name: "Mozambique Metical", DecimalPlaces: 2}
	NAD = CurrencyCode{Code: "NAD", NumericCode: "516", Name: "Namibia Dollar", DecimalPlaces: 2}
	NGN = CurrencyCode{Code: "NGN", NumericCode: "566", Name: "Naira", DecimalPlaces: 2}
	NIO = CurrencyCode{Code: "NIO", NumericCode: "558", Name: "Cordoba Oro", DecimalPlaces: 2}
	NOK = CurrencyCode{Code: "NOK", NumericCode: "578", Name: "Norwegian Krone", DecimalPlaces: 2}
	NPR = CurrencyCode{Code: "NPR", NumericCode: "524", Name: "Nepalese Rupee", DecimalPlaces: 2}
	NZD = CurrencyCode{Code: "NZD", NumericCode: "554", Name: "New Zealand Dollar", DecimalPlaces: 2}
	OMR = CurrencyCode{Code: "OMR", NumericCode: "512", Name: "Rial Omani", DecimalPlaces: 3}
	PAB = CurrencyCode{Code: "PAB", NumericCode: "590", Name: "Balboa", DecimalPlaces: 2}
	PEN = CurrencyCode{Code: "PEN", NumericCode: "604", Name: "Sol", DecimalPlaces: 2}
	PGK = CurrencyCode{Code: "PGK", NumericCode: "598", Name: "Kina", DecimalPlaces: 2}
	PHP = CurrencyCode{Code: "PHP", NumericCode: "608", Name: "Philippine Peso", DecimalPlaces: 2}
	PKR = CurrencyCode{Code: "PKR", NumericCode: "586", Name: "Pakistan Rupee", DecimalPlaces: 2}
	PLN = CurrencyCode{Code: "PLN", NumericCode: "985", Name: "Zloty", DecimalPlaces: 2}
	PYG = CurrencyCode{Code: "PYG", NumericCode: "600", Name: "Guarani", DecimalPlaces: 0}
	QAR = CurrencyCode{Code: "QAR", NumericCode: "634", Name: "Qatari Rial", DecimalPlaces: 2}
	RON = CurrencyCode{Code: "RON", NumericCode: "946", Name: "Romanian Leu", DecimalPlaces: 2}
	RSD = CurrencyCode{Code: "RSD", NumericCode: "941", Name: "Serbian Dinar", DecimalPlaces: 2}
	RUB = CurrencyCode{Code: "RUB", NumericCode: "643", Name: "Russian Ruble", DecimalPlaces: 2}
	RWF = CurrencyCode{Code: "RWF", NumericCode: "646", Name: "Rwanda Franc", DecimalPlaces: 0}
	SAR = CurrencyCode{Code: "SAR", NumericCode: "682", Name: "Saudi Riyal", DecimalPlaces: 2}
	SBD = CurrencyCode{Code: "SBD", NumericCode: "090", Name: "Solomon Islands Dollar", DecimalPlaces: 2}
	SCR = CurrencyCode{Code: "SCR", NumericCode: "690", Name: "Seychelles Rupee", DecimalPlaces: 2}
	SDG = CurrencyCode{Code: "SDG", NumericCode: "938", Name: "Sudanese Pound", DecimalPlaces: 2}
	SEK = CurrencyCode{Code: "SEK", NumericCode: "752", Name: "Swedish Krona", DecimalPlaces: 2}
	SGD = CurrencyCode{Code: "SGD", NumericCode: "702", Name: "Singapore Dollar", DecimalPlaces: 2}
	SHP = CurrencyCode{Code: "SHP", NumericCode: "654", Name: "Saint Helena Pound", DecimalPlaces: 2}
	SLE = CurrencyCode{Code: "SLE", NumericCode: "925", Name: "Leone", DecimalPlaces: 2}
	SLL = CurrencyCode{Code: "SLL", NumericCode: "694", Name: "Leone", DecimalPlaces: 2}
	SOS = CurrencyCode{Code: "SOS", NumericCode: "706", Name: "Somali Shilling", DecimalPlaces: 2}
	SRD = CurrencyCode{Code: "SRD", NumericCode: "968", Name: "Surinam Dollar", DecimalPlaces: 2}
	SSP = CurrencyCode{Code: "SSP", NumericCode: "728", Name: "South Sudanese Pound", DecimalPlaces: 2}
	STN = CurrencyCode{Code: "STN", NumericCode: "930", Name: "Dobra", DecimalPlaces: 2}
	SVC = CurrencyCode{Code: "SVC", NumericCode: "222", Name: "El Salvador Colon", DecimalPlaces: 2}
	SYP = CurrencyCode{Code: "SYP", NumericCode: "760", Name: "Syrian Pound", DecimalPlaces: 2}
	SZL = CurrencyCode{Code: "SZL", NumericCode: "748", Name: "Lilangeni", DecimalPlaces: 2}
	THB = CurrencyCode{Code: "THB", NumericCode: "764", Name: "Baht", DecimalPlaces: 2}
	TJS = CurrencyCode{Code: "TJS", NumericCode: "972", Name: "Somoni", DecimalPlaces: 2}
	TMT = CurrencyCode{Code: "TMT", NumericCode: "934", Name: "Turkmenistan New Manat", DecimalPlaces: 2}
	TND = CurrencyCode{Code: "TND", NumericCode: "788", Name: "Tunisian Dinar", DecimalPlaces: 3}
	TOP = CurrencyCode{Code: "TOP", NumericCode: "776", Name: "Pa'anga", DecimalPlaces: 2}
	TRY = CurrencyCode{Code: "TRY", NumericCode: "949", Name: "Turkish Lira", DecimalPlaces: 2}
	TTD = CurrencyCode{Code: "TTD", NumericCode: "780", Name: "Trinidad and Tobago Dollar", DecimalPlaces: 2}
	TWD = CurrencyCode{Code: "TWD", NumericCode: "901", Name: "New Taiwan Dollar", DecimalPlaces: 2}
	TZS = CurrencyCode{Code: "TZS", NumericCode: "834", Name: "Tanzanian Shilling", DecimalPlaces: 2}
	UAH = CurrencyCode{Code: "UAH", NumericCode: "980", Name: "Hryvnia", DecimalPlaces: 2}
	UGX = CurrencyCode{Code: "UGX", NumericCode: "800", Name: "Uganda Shilling", DecimalPlaces: 0}
	USD = CurrencyCode{Code: "USD", NumericCode: "840", Name: "US Dollar", DecimalPlaces: 2}
	USN = CurrencyCode{Code: "USN", NumericCode: "997", Name: "US Dollar (Next day)", DecimalPlaces: 2}
	UYI = CurrencyCode{Code: "UYI", NumericCode: "940", Name: "Uruguay Peso en Unidades Indexadas (UI)", DecimalPlaces: 0}
	UYU = CurrencyCode{Code: "UYU", NumericCode: "858", Name: "Peso Uruguayo", DecimalPlaces: 2}
	UYW = CurrencyCode{Code: "UYW", NumericCode: "927", Name: "Unidad Previsional", DecimalPlaces: 4}
	UZS = CurrencyCode{Code: "UZS", NumericCode: "860", Name: "Uzbekistan Sum", DecimalPlaces: 2}
	VED = CurrencyCode{Code: "VED", NumericCode: "926", Name: "Bolívar Soberano", DecimalPlaces: 2}
	VES = CurrencyCode{Code: "VES", NumericCode: "928", Name: "Bolívar Soberano", DecimalPlaces: 2}
	VND = CurrencyCode{Code: "VND", NumericCode: "704", Name: "Dong", DecimalPlaces: 0}
	VUV = CurrencyCode{Code: "VUV", NumericCode: "548", Name: "Vatu", DecimalPlaces: 0}
	WST = CurrencyCode{Code: "WST", NumericCode: "882", Name: "Tala", DecimalPlaces: 2}
	XAD = CurrencyCode{Code: "XAD", NumericCode: "396", Name: "Arab Accounting Dinar", DecimalPlaces: 2}
	XAF = CurrencyCode{Code: "XAF", NumericCode: "950", Name: "CFA Franc BEAC", DecimalPlaces: 0}
	XAG = CurrencyCode{Code: "XAG", NumericCode: "961", Name: "Silver", DecimalPlaces: 0}
	XAU = CurrencyCode{Code: "XAU", NumericCode: "959", Name: "Gold", DecimalPlaces: 0}
	XBA = CurrencyCode{Code: "XBA", NumericCode: "955", Name: "Bond Markets Unit European Composite Unit (EURCO)", DecimalPlaces: 0}
	XBB = CurrencyCode{Code: "XBB", NumericCode: "956", Name: "Bond Markets Unit European Monetary Unit (E.M.U.-6)", DecimalPlaces: 0}
	XBC = CurrencyCode{Code: "XBC", NumericCode: "957", Name: "Bond Markets Unit European Unit of Account 9 (E.U.A.-9)", DecimalPlaces: 0}
	XBD = CurrencyCode{Code: "XBD", NumericCode: "958", Name: "Bond Markets Unit European Unit of Account 17 (E.U.A.-17)", DecimalPlaces: 0}
	XCD = CurrencyCode{Code: "XCD", NumericCode: "951", Name: "East Caribbean Dollar", DecimalPlaces: 2}
	XCG = CurrencyCode{Code: "XCG", NumericCode: "532", Name: "Caribbean Guilder", DecimalPlaces: 2}
	XDR = CurrencyCode{Code: "XDR", NumericCode: "960", Name: "SDR (Special Drawing Right)", DecimalPlaces: 0}
	XOF = CurrencyCode{Code: "XOF", NumericCode: "952", Name: "CFA Franc BCEAO", DecimalPlaces: 0}
	XPD = CurrencyCode{Code: "XPD", NumericCode: "964", Name: "Palladium", DecimalPlaces: 0}
	XPF = CurrencyCode{Code: "XPF", NumericCode: "953", Name: "CFP Franc", DecimalPlaces: 0}
	XPT = CurrencyCode{Code: "XPT", NumericCode: "962", Name: "Platinum", DecimalPlaces: 0}
	XSU = CurrencyCode{Code: "XSU", NumericCode: "994", Name: "Sucre", DecimalPlaces: 0}
	XTS = CurrencyCode{Code: "XTS", NumericCode: "963", Name: "Codes specifically reserved for testing purposes", DecimalPlaces: 0}
	XUA = CurrencyCode{Code: "XUA", NumericCode: "965", Name: "ADB Unit of Account", DecimalPlaces: 0}
	XXX = CurrencyCode{Code: "XXX", NumericCode: "999", Name: "The codes assigned for transactions where no currency is involved", DecimalPlaces: 0}
	YER = CurrencyCode{Code: "YER", NumericCode: "886", Name: "Yemeni Rial", DecimalPlaces: 2}
	ZAR = CurrencyCode{Code: "ZAR", NumericCode: "710", Name: "Rand", DecimalPlaces: 2}
	ZMW = CurrencyCode{Code: "ZMW", NumericCode: "967", Name: "Zambian Kwacha", DecimalPlaces: 2}
	ZWG = CurrencyCode{Code: "ZWG", NumericCode: "924", Name: "Zimbabwe Gold", DecimalPlaces: 2}
	ZWL = CurrencyCode{Code: "ZWL", NumericCode: "932", Name: "Zimbabwe Dollar", DecimalPlaces: 2}
)

var lookupTable = map[string]CurrencyCode{
	"008": ALL, // Lek
	"012": DZD, // Algerian Dinar
	"032": ARS, // Argentine Peso
	"036": AUD, // Australian Dollar
	"044": BSD, // Bahamian Dollar
	"048": BHD, // Bahraini Dinar
	"050": BDT, // Taka
	"051": AMD, // Armenian Dram
	"052": BBD, // Barbados Dollar
	"060": BMD, // Bermudian Dollar
	"064": BTN, // Ngultrum
	"068": BOB, // Boliviano
	"072": BWP, // Pula
	"084": BZD, // Belize Dollar
	"090": SBD, // Solomon Islands Dollar
	"096": BND, // Brunei Dollar
	"104": MMK, // Kyat
	"108": BIF, // Burundi Franc
	"116": KHR, // Riel
	"124": CAD, // Canadian Dollar
	"132": CVE, // Cabo Verde Escudo
	"136": KYD, // Cayman Islands Dollar
	"144": LKR, // Sri Lanka Rupee
	"152": CLP, // Chilean Peso
	"156": CNY, // Yuan Renminbi
	"157": CNH, // Yuan Fen
	"170": COP, // Colombian Peso
	"174": KMF, // Comorian Franc
	"188": CRC, // Costa Rican Colon
	"191": HRK, // Kuna
	"192": CUP, // Cuban Peso
	"203": CZK, // Czech Koruna
	"208": DKK, // Danish Krone
	"214": DOP, // Dominican Peso
	"222": SVC, // El Salvador Colon
	"230": ETB, // Ethiopian Birr
	"232": ERN, // Nakfa
	"238": FKP, // Falkland Islands Pound
	"242": FJD, // Fiji Dollar
	"262": DJF, // Djibouti Franc
	"270": GMD, // Dalasi
	"292": GIP, // Gibraltar Pound
	"320": GTQ, // Quetzal
	"324": GNF, // Guinean Franc
	"328": GYD, // Guyana Dollar
	"332": HTG, // Gourde
	"340": HNL, // Lempira
	"344": HKD, // Hong Kong Dollar
	"348": HUF, // Forint
	"352": ISK, // Iceland Krona
	"356": INR, // Indian Rupee
	"360": IDR, // Rupiah
	"364": IRR, // Iranian Rial
	"368": IQD, // Iraqi Dinar
	"376": ILS, // New Israeli Sheqel
	"388": JMD, // Jamaican Dollar
	"392": JPY, // Yen
	"396": XAD, // Arab Accounting Dinar
	"398": KZT, // Tenge
	"400": JOD, // Jordanian Dinar
	"404": KES, // Kenyan Shilling
	"408": KPW, // North Korean Won
	"410": KRW, // Won
	"414": KWD, // Kuwaiti Dinar
	"417": KGS, // Som
	"418": LAK, // Lao Kip
	"422": LBP, // Lebanese Pound
	"426": LSL, // Loti
	"430": LRD, // Liberian Dollar
	"434": LYD, // Libyan Dinar
	"446": MOP, // Pataca
	"454": MWK, // Malawi Kwacha
	"458": MYR, // Malaysian Ringgit
	"462": MVR, // Rufiyaa
	"480": MUR, // Mauritius Rupee
	"484": MXN, // Mexican Peso
	"496": MNT, // Tugrik
	"498": MDL, // Moldovan Leu
	"504": MAD, // Moroccan Dirham
	"512": OMR, // Rial Omani
	"516": NAD, // Namibia Dollar
	"524": NPR, // Nepalese Rupee
	"532": XCG, // Caribbean Guilder (532 reassigned from the withdrawn ANG)
	"533": AWG, // Aruban Florin
	"548": VUV, // Vatu
	"554": NZD, // New Zealand Dollar
	"558": NIO, // Cordoba Oro
	"566": NGN, // Naira
	"578": NOK, // Norwegian Krone
	"586": PKR, // Pakistan Rupee
	"590": PAB, // Balboa
	"598": PGK, // Kina
	"600": PYG, // Guarani
	"604": PEN, // Sol
	"608": PHP, // Philippine Peso
	"634": QAR, // Qatari Rial
	"643": RUB, // Russian Ruble
	"646": RWF, // Rwanda Franc
	"654": SHP, // Saint Helena Pound
	"682": SAR, // Saudi Riyal
	"690": SCR, // Seychelles Rupee
	"694": SLL, // Leone
	"702": SGD, // Singapore Dollar
	"704": VND, // Dong
	"706": SOS, // Somali Shilling
	"710": ZAR, // Rand
	"728": SSP, // South Sudanese Pound
	"748": SZL, // Lilangeni
	"752": SEK, // Swedish Krona
	"756": CHF, // Swiss Franc
	"760": SYP, // Syrian Pound
	"764": THB, // Baht
	"776": TOP, // Pa'anga
	"780": TTD, // Trinidad and Tobago Dollar
	"784": AED, // UAE Dirham
	"788": TND, // Tunisian Dinar
	"800": UGX, // Uganda Shilling
	"807": MKD, // Denar
	"818": EGP, // Egyptian Pound
	"826": GBP, // Pound Sterling
	"834": TZS, // Tanzanian Shilling
	"840": USD, // US Dollar
	"858": UYU, // Peso Uruguayo
	"860": UZS, // Uzbekistan Sum
	"882": WST, // Tala
	"886": YER, // Yemeni Rial
	"901": TWD, // New Taiwan Dollar
	"924": ZWG, // Zimbabwe Gold
	"925": SLE, // Leone
	"926": VED, // Bolívar Soberano
	"927": UYW, // Unidad Previsional
	"928": VES, // Bolívar Soberano
	"929": MRU, // Ouguiya
	"930": STN, // Dobra
	"931": CUC, // Peso Convertible
	"932": ZWL, // Zimbabwe Dollar
	"933": BYN, // Belarusian Ruble
	"934": TMT, // Turkmenistan New Manat
	"936": GHS, // Ghana Cedi
	"938": SDG, // Sudanese Pound
	"940": UYI, // Uruguay Peso en Unidades Indexadas (UI)
	"941": RSD, // Serbian Dinar
	"943": MZN, // Mozambique Metical
	"944": AZN, // Azerbaijan Manat
	"946": RON, // Romanian Leu
	"947": CHE, // WIR Euro
	"948": CHW, // WIR Franc
	"949": TRY, // Turkish Lira
	"950": XAF, // CFA Franc BEAC
	"951": XCD, // East Caribbean Dollar
	"952": XOF, // CFA Franc BCEAO
	"953": XPF, // CFP Franc
	"955": XBA, // Bond Markets Unit European Composite Unit (EURCO)
	"956": XBB, // Bond Markets Unit European Monetary Unit (E.M.U.-6)
	"957": XBC, // Bond Markets Unit European Unit of Account 9 (E.U.A.-9)
	"958": XBD, // Bond Markets Unit European Unit of Account 17 (E.U.A.-17)
	"959": XAU, // Gold
	"960": XDR, // SDR (Special Drawing Right)
	"961": XAG, // Silver
	"962": XPT, // Platinum
	"963": XTS, // Codes specifically reserved for testing purposes
	"964": XPD, // Palladium
	"965": XUA, // ADB Unit of Account
	"967": ZMW, // Zambian Kwacha
	"968": SRD, // Surinam Dollar
	"969": MGA, // Malagasy Ariary
	"970": COU, // Unidad de Valor Real
	"971": AFN, // Afghani
	"972": TJS, // Somoni
	"973": AOA, // Kwanza
	"975": BGN, // Bulgarian Lev
	"976": CDF, // Congolese Franc
	"977": BAM, // Convertible Mark
	"978": EUR, // Euro
	"979": MXV, // Mexican Unidad de Inversion (UDI)
	"980": UAH, // Hryvnia
	"981": GEL, // Lari
	"984": BOV, // Mvdol
	"985": PLN, // Zloty
	"986": BRL, // Brazilian Real
	"990": CLF, // Unidad de Fomento
	"994": XSU, // Sucre
	"997": USN, // US Dollar (Next day)
	"999": XXX, // The codes assigned for transactions where no currency is involved
	"AED": AED, // UAE Dirham
	"AFN": AFN, // Afghani
	"ALL": ALL, // Lek
	"AMD": AMD, // Armenian Dram
	"ANG": ANG, // Netherlands Antillean Guilder (withdrawn; succeeded by XCG)
	"AOA": AOA, // Kwanza
	"ARS": ARS, // Argentine Peso
	"AUD": AUD, // Australian Dollar
	"AWG": AWG, // Aruban Florin
	"AZN": AZN, // Azerbaijan Manat
	"BAM": BAM, // Convertible Mark
	"BBD": BBD, // Barbados Dollar
	"BDT": BDT, // Taka
	"BGN": BGN, // Bulgarian Lev
	"BHD": BHD, // Bahraini Dinar
	"BIF": BIF, // Burundi Franc
	"BMD": BMD, // Bermudian Dollar
	"BND": BND, // Brunei Dollar
	"BOB": BOB, // Boliviano
	"BOV": BOV, // Mvdol
	"BRL": BRL, // Brazilian Real
	"BSD": BSD, // Bahamian Dollar
	"BTN": BTN, // Ngultrum
	"BWP": BWP, // Pula
	"BYN": BYN, // Belarusian Ruble
	"BZD": BZD, // Belize Dollar
	"CAD": CAD, // Canadian Dollar
	"CDF": CDF, // Congolese Franc
	"CHE": CHE, // WIR Euro
	"CHF": CHF, // Swiss Franc
	"CHW": CHW, // WIR Franc
	"CLF": CLF, // Unidad de Fomento
	"CLP": CLP, // Chilean Peso
	"CNH": CNH, // Yuan Fen
	"CNY": CNY, // Yuan Renminbi
	"COP": COP, // Colombian Peso
	"COU": COU, // Unidad de Valor Real
	"CRC": CRC, // Costa Rican Colon
	"CUC": CUC, // Peso Convertible
	"CUP": CUP, // Cuban Peso
	"CVE": CVE, // Cabo Verde Escudo
	"CZK": CZK, // Czech Koruna
	"DJF": DJF, // Djibouti Franc
	"DKK": DKK, // Danish Krone
	"DOP": DOP, // Dominican Peso
	"DZD": DZD, // Algerian Dinar
	"EGP": EGP, // Egyptian Pound
	"ERN": ERN, // Nakfa
	"ETB": ETB, // Ethiopian Birr
	"EUR": EUR, // Euro
	"FJD": FJD, // Fiji Dollar
	"FKP": FKP, // Falkland Islands Pound
	"GBP": GBP, // Pound Sterling
	"GEL": GEL, // Lari
	"GHS": GHS, // Ghana Cedi
	"GIP": GIP, // Gibraltar Pound
	"GMD": GMD, // Dalasi
	"GNF": GNF, // Guinean Franc
	"GTQ": GTQ, // Quetzal
	"GYD": GYD, // Guyana Dollar
	"HKD": HKD, // Hong Kong Dollar
	"HNL": HNL, // Lempira
	"HRK": HRK, // Kuna
	"HTG": HTG, // Gourde
	"HUF": HUF, // Forint
	"IDR": IDR, // Rupiah
	"ILS": ILS, // New Israeli Sheqel
	"INR": INR, // Indian Rupee
	"IQD": IQD, // Iraqi Dinar
	"IRR": IRR, // Iranian Rial
	"ISK": ISK, // Iceland Krona
	"JMD": JMD, // Jamaican Dollar
	"JOD": JOD, // Jordanian Dinar
	"JPY": JPY, // Yen
	"KES": KES, // Kenyan Shilling
	"KGS": KGS, // Som
	"KHR": KHR, // Riel
	"KMF": KMF, // Comorian Franc
	"KPW": KPW, // North Korean Won
	"KRW": KRW, // Won
	"KWD": KWD, // Kuwaiti Dinar
	"KYD": KYD, // Cayman Islands Dollar
	"KZT": KZT, // Tenge
	"LAK": LAK, // Lao Kip
	"LBP": LBP, // Lebanese Pound
	"LKR": LKR, // Sri Lanka Rupee
	"LRD": LRD, // Liberian Dollar
	"LSL": LSL, // Loti
	"LYD": LYD, // Libyan Dinar
	"MAD": MAD, // Moroccan Dirham
	"MDL": MDL, // Moldovan Leu
	"MGA": MGA, // Malagasy Ariary
	"MKD": MKD, // Denar
	"MMK": MMK, // Kyat
	"MNT": MNT, // Tugrik
	"MOP": MOP, // Pataca
	"MRU": MRU, // Ouguiya
	"MUR": MUR, // Mauritius Rupee
	"MVR": MVR, // Rufiyaa
	"MWK": MWK, // Malawi Kwacha
	"MXN": MXN, // Mexican Peso
	"MXV": MXV, // Mexican Unidad de Inversion (UDI)
	"MYR": MYR, // Malaysian Ringgit
	"MZN": MZN, // Mozambique Metical
	"NAD": NAD, // Namibia Dollar
	"NGN": NGN, // Naira
	"NIO": NIO, // Cordoba Oro
	"NOK": NOK, // Norwegian Krone
	"NPR": NPR, // Nepalese Rupee
	"NZD": NZD, // New Zealand Dollar
	"OMR": OMR, // Rial Omani
	"PAB": PAB, // Balboa
	"PEN": PEN, // Sol
	"PGK": PGK, // Kina
	"PHP": PHP, // Philippine Peso
	"PKR": PKR, // Pakistan Rupee
	"PLN": PLN, // Zloty
	"PYG": PYG, // Guarani
	"QAR": QAR, // Qatari Rial
	"RON": RON, // Romanian Leu
	"RSD": RSD, // Serbian Dinar
	"RUB": RUB, // Russian Ruble
	"RWF": RWF, // Rwanda Franc
	"SAR": SAR, // Saudi Riyal
	"SBD": SBD, // Solomon Islands Dollar
	"SCR": SCR, // Seychelles Rupee
	"SDG": SDG, // Sudanese Pound
	"SEK": SEK, // Swedish Krona
	"SGD": SGD, // Singapore Dollar
	"SHP": SHP, // Saint Helena Pound
	"SLE": SLE, // Leone
	"SLL": SLL, // Leone
	"SOS": SOS, // Somali Shilling
	"SRD": SRD, // Surinam Dollar
	"SSP": SSP, // South Sudanese Pound
	"STN": STN, // Dobra
	"SVC": SVC, // El Salvador Colon
	"SYP": SYP, // Syrian Pound
	"SZL": SZL, // Lilangeni
	"THB": THB, // Baht
	"TJS": TJS, // Somoni
	"TMT": TMT, // Turkmenistan New Manat
	"TND": TND, // Tunisian Dinar
	"TOP": TOP, // Pa'anga
	"TRY": TRY, // Turkish Lira
	"TTD": TTD, // Trinidad and Tobago Dollar
	"TWD": TWD, // New Taiwan Dollar
	"TZS": TZS, // Tanzanian Shilling
	"UAH": UAH, // Hryvnia
	"UGX": UGX, // Uganda Shilling
	"USD": USD, // US Dollar
	"USN": USN, // US Dollar (Next day)
	"UYI": UYI, // Uruguay Peso en Unidades Indexadas (UI)
	"UYU": UYU, // Peso Uruguayo
	"UYW": UYW, // Unidad Previsional
	"UZS": UZS, // Uzbekistan Sum
	"VED": VED, // Bolívar Soberano
	"VES": VES, // Bolívar Soberano
	"VND": VND, // Dong
	"VUV": VUV, // Vatu
	"WST": WST, // Tala
	"XAD": XAD, // Arab Accounting Dinar
	"XAF": XAF, // CFA Franc BEAC
	"XAG": XAG, // Silver
	"XAU": XAU, // Gold
	"XBA": XBA, // Bond Markets Unit European Composite Unit (EURCO)
	"XBB": XBB, // Bond Markets Unit European Monetary Unit (E.M.U.-6)
	"XBC": XBC, // Bond Markets Unit European Unit of Account 9 (E.U.A.-9)
	"XBD": XBD, // Bond Markets Unit European Unit of Account 17 (E.U.A.-17)
	"XCD": XCD, // East Caribbean Dollar
	"XCG": XCG, // Caribbean Guilder
	"XDR": XDR, // SDR (Special Drawing Right)
	"XOF": XOF, // CFA Franc BCEAO
	"XPD": XPD, // Palladium
	"XPF": XPF, // CFP Franc
	"XPT": XPT, // Platinum
	"XSU": XSU, // Sucre
	"XTS": XTS, // Codes specifically reserved for testing purposes
	"XUA": XUA, // ADB Unit of Account
	"XXX": XXX, // The codes assigned for transactions where no currency is involved
	"YER": YER, // Yemeni Rial
	"ZAR": ZAR, // Rand
	"ZMW": ZMW, // Zambian Kwacha
	"ZWG": ZWG, // Zimbabwe Gold
	"ZWL": ZWL, // Zimbabwe Dollar
}
