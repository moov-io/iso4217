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

// Generated on 2023-06-05T05:08:58Z by raphaelsantodomingo, any modifications will be overwritten
package iso4217

type CurrencyCode struct {
	Code, Name string

	// DecimalPlaces represents the integer value of a currency's minor unit.
	// DecimalPlaces is 0 if the currency doesn't have a minor unit.
	DecimalPlaces int
}

func (cc CurrencyCode) String() string {
	return cc.Code
}

func (cc CurrencyCode) Valid() bool {
	_, exists := Lookup(cc.Code)
	return exists
}

var (
	AFN = CurrencyCode{Code: "AFN", Name: "Afghani", DecimalPlaces: 2}
	EUR = CurrencyCode{Code: "EUR", Name: "Euro", DecimalPlaces: 2}
	ALL = CurrencyCode{Code: "ALL", Name: "Lek", DecimalPlaces: 2}
	DZD = CurrencyCode{Code: "DZD", Name: "Algerian Dinar", DecimalPlaces: 2}
	USD = CurrencyCode{Code: "USD", Name: "US Dollar", DecimalPlaces: 2}
	AOA = CurrencyCode{Code: "AOA", Name: "Kwanza", DecimalPlaces: 2}
	XCD = CurrencyCode{Code: "XCD", Name: "East Caribbean Dollar", DecimalPlaces: 2}
	ARS = CurrencyCode{Code: "ARS", Name: "Argentine Peso", DecimalPlaces: 2}
	AMD = CurrencyCode{Code: "AMD", Name: "Armenian Dram", DecimalPlaces: 2}
	AWG = CurrencyCode{Code: "AWG", Name: "Aruban Florin", DecimalPlaces: 2}
	AUD = CurrencyCode{Code: "AUD", Name: "Australian Dollar", DecimalPlaces: 2}
	AZN = CurrencyCode{Code: "AZN", Name: "Azerbaijan Manat", DecimalPlaces: 2}
	BSD = CurrencyCode{Code: "BSD", Name: "Bahamian Dollar", DecimalPlaces: 2}
	BHD = CurrencyCode{Code: "BHD", Name: "Bahraini Dinar", DecimalPlaces: 3}
	BDT = CurrencyCode{Code: "BDT", Name: "Taka", DecimalPlaces: 2}
	BBD = CurrencyCode{Code: "BBD", Name: "Barbados Dollar", DecimalPlaces: 2}
	BYN = CurrencyCode{Code: "BYN", Name: "Belarusian Ruble", DecimalPlaces: 2}
	BZD = CurrencyCode{Code: "BZD", Name: "Belize Dollar", DecimalPlaces: 2}
	XOF = CurrencyCode{Code: "XOF", Name: "CFA Franc BCEAO", DecimalPlaces: 0}
	BMD = CurrencyCode{Code: "BMD", Name: "Bermudian Dollar", DecimalPlaces: 2}
	INR = CurrencyCode{Code: "INR", Name: "Indian Rupee", DecimalPlaces: 2}
	BTN = CurrencyCode{Code: "BTN", Name: "Ngultrum", DecimalPlaces: 2}
	BOB = CurrencyCode{Code: "BOB", Name: "Boliviano", DecimalPlaces: 2}
	BOV = CurrencyCode{Code: "BOV", Name: "Mvdol", DecimalPlaces: 2}
	BAM = CurrencyCode{Code: "BAM", Name: "Convertible Mark", DecimalPlaces: 2}
	BWP = CurrencyCode{Code: "BWP", Name: "Pula", DecimalPlaces: 2}
	NOK = CurrencyCode{Code: "NOK", Name: "Norwegian Krone", DecimalPlaces: 2}
	BRL = CurrencyCode{Code: "BRL", Name: "Brazilian Real", DecimalPlaces: 2}
	BND = CurrencyCode{Code: "BND", Name: "Brunei Dollar", DecimalPlaces: 2}
	BGN = CurrencyCode{Code: "BGN", Name: "Bulgarian Lev", DecimalPlaces: 2}
	BIF = CurrencyCode{Code: "BIF", Name: "Burundi Franc", DecimalPlaces: 0}
	CVE = CurrencyCode{Code: "CVE", Name: "Cabo Verde Escudo", DecimalPlaces: 2}
	KHR = CurrencyCode{Code: "KHR", Name: "Riel", DecimalPlaces: 2}
	XAF = CurrencyCode{Code: "XAF", Name: "CFA Franc BEAC", DecimalPlaces: 0}
	CAD = CurrencyCode{Code: "CAD", Name: "Canadian Dollar", DecimalPlaces: 2}
	KYD = CurrencyCode{Code: "KYD", Name: "Cayman Islands Dollar", DecimalPlaces: 2}
	CLP = CurrencyCode{Code: "CLP", Name: "Chilean Peso", DecimalPlaces: 0}
	CLF = CurrencyCode{Code: "CLF", Name: "Unidad de Fomento", DecimalPlaces: 4}
	CNY = CurrencyCode{Code: "CNY", Name: "Yuan Renminbi", DecimalPlaces: 2}
	COP = CurrencyCode{Code: "COP", Name: "Colombian Peso", DecimalPlaces: 2}
	COU = CurrencyCode{Code: "COU", Name: "Unidad de Valor Real", DecimalPlaces: 2}
	KMF = CurrencyCode{Code: "KMF", Name: "Comorian Franc", DecimalPlaces: 0}
	CDF = CurrencyCode{Code: "CDF", Name: "Congolese Franc", DecimalPlaces: 2}
	NZD = CurrencyCode{Code: "NZD", Name: "New Zealand Dollar", DecimalPlaces: 2}
	CRC = CurrencyCode{Code: "CRC", Name: "Costa Rican Colon", DecimalPlaces: 2}
	HRK = CurrencyCode{Code: "HRK", Name: "Kuna", DecimalPlaces: 2}
	CUP = CurrencyCode{Code: "CUP", Name: "Cuban Peso", DecimalPlaces: 2}
	CUC = CurrencyCode{Code: "CUC", Name: "Peso Convertible", DecimalPlaces: 2}
	ANG = CurrencyCode{Code: "ANG", Name: "Netherlands Antillean Guilder", DecimalPlaces: 2}
	CZK = CurrencyCode{Code: "CZK", Name: "Czech Koruna", DecimalPlaces: 2}
	DKK = CurrencyCode{Code: "DKK", Name: "Danish Krone", DecimalPlaces: 2}
	DJF = CurrencyCode{Code: "DJF", Name: "Djibouti Franc", DecimalPlaces: 0}
	DOP = CurrencyCode{Code: "DOP", Name: "Dominican Peso", DecimalPlaces: 2}
	EGP = CurrencyCode{Code: "EGP", Name: "Egyptian Pound", DecimalPlaces: 2}
	SVC = CurrencyCode{Code: "SVC", Name: "El Salvador Colon", DecimalPlaces: 2}
	ERN = CurrencyCode{Code: "ERN", Name: "Nakfa", DecimalPlaces: 2}
	SZL = CurrencyCode{Code: "SZL", Name: "Lilangeni", DecimalPlaces: 2}
	ETB = CurrencyCode{Code: "ETB", Name: "Ethiopian Birr", DecimalPlaces: 2}
	FKP = CurrencyCode{Code: "FKP", Name: "Falkland Islands Pound", DecimalPlaces: 2}
	FJD = CurrencyCode{Code: "FJD", Name: "Fiji Dollar", DecimalPlaces: 2}
	XPF = CurrencyCode{Code: "XPF", Name: "CFP Franc", DecimalPlaces: 0}
	GMD = CurrencyCode{Code: "GMD", Name: "Dalasi", DecimalPlaces: 2}
	GEL = CurrencyCode{Code: "GEL", Name: "Lari", DecimalPlaces: 2}
	GHS = CurrencyCode{Code: "GHS", Name: "Ghana Cedi", DecimalPlaces: 2}
	GIP = CurrencyCode{Code: "GIP", Name: "Gibraltar Pound", DecimalPlaces: 2}
	GTQ = CurrencyCode{Code: "GTQ", Name: "Quetzal", DecimalPlaces: 2}
	GBP = CurrencyCode{Code: "GBP", Name: "Pound Sterling", DecimalPlaces: 2}
	GNF = CurrencyCode{Code: "GNF", Name: "Guinean Franc", DecimalPlaces: 0}
	GYD = CurrencyCode{Code: "GYD", Name: "Guyana Dollar", DecimalPlaces: 2}
	HTG = CurrencyCode{Code: "HTG", Name: "Gourde", DecimalPlaces: 2}
	HNL = CurrencyCode{Code: "HNL", Name: "Lempira", DecimalPlaces: 2}
	HKD = CurrencyCode{Code: "HKD", Name: "Hong Kong Dollar", DecimalPlaces: 2}
	HUF = CurrencyCode{Code: "HUF", Name: "Forint", DecimalPlaces: 2}
	ISK = CurrencyCode{Code: "ISK", Name: "Iceland Krona", DecimalPlaces: 0}
	IDR = CurrencyCode{Code: "IDR", Name: "Rupiah", DecimalPlaces: 2}
	XDR = CurrencyCode{Code: "XDR", Name: "SDR (Special Drawing Right)", DecimalPlaces: 0}
	IRR = CurrencyCode{Code: "IRR", Name: "Iranian Rial", DecimalPlaces: 2}
	IQD = CurrencyCode{Code: "IQD", Name: "Iraqi Dinar", DecimalPlaces: 3}
	ILS = CurrencyCode{Code: "ILS", Name: "New Israeli Sheqel", DecimalPlaces: 2}
	JMD = CurrencyCode{Code: "JMD", Name: "Jamaican Dollar", DecimalPlaces: 2}
	JPY = CurrencyCode{Code: "JPY", Name: "Yen", DecimalPlaces: 0}
	JOD = CurrencyCode{Code: "JOD", Name: "Jordanian Dinar", DecimalPlaces: 3}
	KZT = CurrencyCode{Code: "KZT", Name: "Tenge", DecimalPlaces: 2}
	KES = CurrencyCode{Code: "KES", Name: "Kenyan Shilling", DecimalPlaces: 2}
	KPW = CurrencyCode{Code: "KPW", Name: "North Korean Won", DecimalPlaces: 2}
	KRW = CurrencyCode{Code: "KRW", Name: "Won", DecimalPlaces: 0}
	KWD = CurrencyCode{Code: "KWD", Name: "Kuwaiti Dinar", DecimalPlaces: 3}
	KGS = CurrencyCode{Code: "KGS", Name: "Som", DecimalPlaces: 2}
	LAK = CurrencyCode{Code: "LAK", Name: "Lao Kip", DecimalPlaces: 2}
	LBP = CurrencyCode{Code: "LBP", Name: "Lebanese Pound", DecimalPlaces: 2}
	LSL = CurrencyCode{Code: "LSL", Name: "Loti", DecimalPlaces: 2}
	ZAR = CurrencyCode{Code: "ZAR", Name: "Rand", DecimalPlaces: 2}
	LRD = CurrencyCode{Code: "LRD", Name: "Liberian Dollar", DecimalPlaces: 2}
	LYD = CurrencyCode{Code: "LYD", Name: "Libyan Dinar", DecimalPlaces: 3}
	CHF = CurrencyCode{Code: "CHF", Name: "Swiss Franc", DecimalPlaces: 2}
	MOP = CurrencyCode{Code: "MOP", Name: "Pataca", DecimalPlaces: 2}
	MKD = CurrencyCode{Code: "MKD", Name: "Denar", DecimalPlaces: 2}
	MGA = CurrencyCode{Code: "MGA", Name: "Malagasy Ariary", DecimalPlaces: 2}
	MWK = CurrencyCode{Code: "MWK", Name: "Malawi Kwacha", DecimalPlaces: 2}
	MYR = CurrencyCode{Code: "MYR", Name: "Malaysian Ringgit", DecimalPlaces: 2}
	MVR = CurrencyCode{Code: "MVR", Name: "Rufiyaa", DecimalPlaces: 2}
	MRU = CurrencyCode{Code: "MRU", Name: "Ouguiya", DecimalPlaces: 2}
	MUR = CurrencyCode{Code: "MUR", Name: "Mauritius Rupee", DecimalPlaces: 2}
	XUA = CurrencyCode{Code: "XUA", Name: "ADB Unit of Account", DecimalPlaces: 0}
	MXN = CurrencyCode{Code: "MXN", Name: "Mexican Peso", DecimalPlaces: 2}
	MXV = CurrencyCode{Code: "MXV", Name: "Mexican Unidad de Inversion (UDI)", DecimalPlaces: 2}
	MDL = CurrencyCode{Code: "MDL", Name: "Moldovan Leu", DecimalPlaces: 2}
	MNT = CurrencyCode{Code: "MNT", Name: "Tugrik", DecimalPlaces: 2}
	MAD = CurrencyCode{Code: "MAD", Name: "Moroccan Dirham", DecimalPlaces: 2}
	MZN = CurrencyCode{Code: "MZN", Name: "Mozambique Metical", DecimalPlaces: 2}
	MMK = CurrencyCode{Code: "MMK", Name: "Kyat", DecimalPlaces: 2}
	NAD = CurrencyCode{Code: "NAD", Name: "Namibia Dollar", DecimalPlaces: 2}
	NPR = CurrencyCode{Code: "NPR", Name: "Nepalese Rupee", DecimalPlaces: 2}
	NIO = CurrencyCode{Code: "NIO", Name: "Cordoba Oro", DecimalPlaces: 2}
	NGN = CurrencyCode{Code: "NGN", Name: "Naira", DecimalPlaces: 2}
	OMR = CurrencyCode{Code: "OMR", Name: "Rial Omani", DecimalPlaces: 3}
	PKR = CurrencyCode{Code: "PKR", Name: "Pakistan Rupee", DecimalPlaces: 2}
	PAB = CurrencyCode{Code: "PAB", Name: "Balboa", DecimalPlaces: 2}
	PGK = CurrencyCode{Code: "PGK", Name: "Kina", DecimalPlaces: 2}
	PYG = CurrencyCode{Code: "PYG", Name: "Guarani", DecimalPlaces: 0}
	PEN = CurrencyCode{Code: "PEN", Name: "Sol", DecimalPlaces: 2}
	PHP = CurrencyCode{Code: "PHP", Name: "Philippine Peso", DecimalPlaces: 2}
	PLN = CurrencyCode{Code: "PLN", Name: "Zloty", DecimalPlaces: 2}
	QAR = CurrencyCode{Code: "QAR", Name: "Qatari Rial", DecimalPlaces: 2}
	RON = CurrencyCode{Code: "RON", Name: "Romanian Leu", DecimalPlaces: 2}
	RUB = CurrencyCode{Code: "RUB", Name: "Russian Ruble", DecimalPlaces: 2}
	RWF = CurrencyCode{Code: "RWF", Name: "Rwanda Franc", DecimalPlaces: 0}
	SHP = CurrencyCode{Code: "SHP", Name: "Saint Helena Pound", DecimalPlaces: 2}
	WST = CurrencyCode{Code: "WST", Name: "Tala", DecimalPlaces: 2}
	STN = CurrencyCode{Code: "STN", Name: "Dobra", DecimalPlaces: 2}
	SAR = CurrencyCode{Code: "SAR", Name: "Saudi Riyal", DecimalPlaces: 2}
	RSD = CurrencyCode{Code: "RSD", Name: "Serbian Dinar", DecimalPlaces: 2}
	SCR = CurrencyCode{Code: "SCR", Name: "Seychelles Rupee", DecimalPlaces: 2}
	SLL = CurrencyCode{Code: "SLL", Name: "Leone", DecimalPlaces: 2}
	SGD = CurrencyCode{Code: "SGD", Name: "Singapore Dollar", DecimalPlaces: 2}
	XSU = CurrencyCode{Code: "XSU", Name: "Sucre", DecimalPlaces: 0}
	SBD = CurrencyCode{Code: "SBD", Name: "Solomon Islands Dollar", DecimalPlaces: 2}
	SOS = CurrencyCode{Code: "SOS", Name: "Somali Shilling", DecimalPlaces: 2}
	SSP = CurrencyCode{Code: "SSP", Name: "South Sudanese Pound", DecimalPlaces: 2}
	LKR = CurrencyCode{Code: "LKR", Name: "Sri Lanka Rupee", DecimalPlaces: 2}
	SDG = CurrencyCode{Code: "SDG", Name: "Sudanese Pound", DecimalPlaces: 2}
	SRD = CurrencyCode{Code: "SRD", Name: "Surinam Dollar", DecimalPlaces: 2}
	SEK = CurrencyCode{Code: "SEK", Name: "Swedish Krona", DecimalPlaces: 2}
	CHE = CurrencyCode{Code: "CHE", Name: "WIR Euro", DecimalPlaces: 2}
	CHW = CurrencyCode{Code: "CHW", Name: "WIR Franc", DecimalPlaces: 2}
	SYP = CurrencyCode{Code: "SYP", Name: "Syrian Pound", DecimalPlaces: 2}
	TWD = CurrencyCode{Code: "TWD", Name: "New Taiwan Dollar", DecimalPlaces: 2}
	TJS = CurrencyCode{Code: "TJS", Name: "Somoni", DecimalPlaces: 2}
	TZS = CurrencyCode{Code: "TZS", Name: "Tanzanian Shilling", DecimalPlaces: 2}
	THB = CurrencyCode{Code: "THB", Name: "Baht", DecimalPlaces: 2}
	TOP = CurrencyCode{Code: "TOP", Name: "Pa'anga", DecimalPlaces: 2}
	TTD = CurrencyCode{Code: "TTD", Name: "Trinidad and Tobago Dollar", DecimalPlaces: 2}
	TND = CurrencyCode{Code: "TND", Name: "Tunisian Dinar", DecimalPlaces: 3}
	TRY = CurrencyCode{Code: "TRY", Name: "Turkish Lira", DecimalPlaces: 2}
	TMT = CurrencyCode{Code: "TMT", Name: "Turkmenistan New Manat", DecimalPlaces: 2}
	UGX = CurrencyCode{Code: "UGX", Name: "Uganda Shilling", DecimalPlaces: 0}
	UAH = CurrencyCode{Code: "UAH", Name: "Hryvnia", DecimalPlaces: 2}
	AED = CurrencyCode{Code: "AED", Name: "UAE Dirham", DecimalPlaces: 2}
	USN = CurrencyCode{Code: "USN", Name: "US Dollar (Next day)", DecimalPlaces: 2}
	UYU = CurrencyCode{Code: "UYU", Name: "Peso Uruguayo", DecimalPlaces: 2}
	UYI = CurrencyCode{Code: "UYI", Name: "Uruguay Peso en Unidades Indexadas (UI)", DecimalPlaces: 0}
	UYW = CurrencyCode{Code: "UYW", Name: "Unidad Previsional", DecimalPlaces: 4}
	UZS = CurrencyCode{Code: "UZS", Name: "Uzbekistan Sum", DecimalPlaces: 2}
	VUV = CurrencyCode{Code: "VUV", Name: "Vatu", DecimalPlaces: 0}
	VES = CurrencyCode{Code: "VES", Name: "Bolívar Soberano", DecimalPlaces: 2}
	VND = CurrencyCode{Code: "VND", Name: "Dong", DecimalPlaces: 0}
	YER = CurrencyCode{Code: "YER", Name: "Yemeni Rial", DecimalPlaces: 2}
	ZMW = CurrencyCode{Code: "ZMW", Name: "Zambian Kwacha", DecimalPlaces: 2}
	ZWL = CurrencyCode{Code: "ZWL", Name: "Zimbabwe Dollar", DecimalPlaces: 2}
	XBA = CurrencyCode{Code: "XBA", Name: "Bond Markets Unit European Composite Unit (EURCO)", DecimalPlaces: 0}
	XBB = CurrencyCode{Code: "XBB", Name: "Bond Markets Unit European Monetary Unit (E.M.U.-6)", DecimalPlaces: 0}
	XBC = CurrencyCode{Code: "XBC", Name: "Bond Markets Unit European Unit of Account 9 (E.U.A.-9)", DecimalPlaces: 0}
	XBD = CurrencyCode{Code: "XBD", Name: "Bond Markets Unit European Unit of Account 17 (E.U.A.-17)", DecimalPlaces: 0}
	XTS = CurrencyCode{Code: "XTS", Name: "Codes specifically reserved for testing purposes", DecimalPlaces: 0}
	XXX = CurrencyCode{Code: "XXX", Name: "The codes assigned for transactions where no currency is involved", DecimalPlaces: 0}
	XAU = CurrencyCode{Code: "XAU", Name: "Gold", DecimalPlaces: 0}
	XPD = CurrencyCode{Code: "XPD", Name: "Palladium", DecimalPlaces: 0}
	XPT = CurrencyCode{Code: "XPT", Name: "Platinum", DecimalPlaces: 0}
	XAG = CurrencyCode{Code: "XAG", Name: "Silver", DecimalPlaces: 0}
)

var lookupTable = map[string]CurrencyCode{
	"AFN": AFN, // Afghani
	"EUR": EUR, // Euro
	"ALL": ALL, // Lek
	"DZD": DZD, // Algerian Dinar
	"USD": USD, // US Dollar
	"AOA": AOA, // Kwanza
	"XCD": XCD, // East Caribbean Dollar
	"ARS": ARS, // Argentine Peso
	"AMD": AMD, // Armenian Dram
	"AWG": AWG, // Aruban Florin
	"AUD": AUD, // Australian Dollar
	"AZN": AZN, // Azerbaijan Manat
	"BSD": BSD, // Bahamian Dollar
	"BHD": BHD, // Bahraini Dinar
	"BDT": BDT, // Taka
	"BBD": BBD, // Barbados Dollar
	"BYN": BYN, // Belarusian Ruble
	"BZD": BZD, // Belize Dollar
	"XOF": XOF, // CFA Franc BCEAO
	"BMD": BMD, // Bermudian Dollar
	"INR": INR, // Indian Rupee
	"BTN": BTN, // Ngultrum
	"BOB": BOB, // Boliviano
	"BOV": BOV, // Mvdol
	"BAM": BAM, // Convertible Mark
	"BWP": BWP, // Pula
	"NOK": NOK, // Norwegian Krone
	"BRL": BRL, // Brazilian Real
	"BND": BND, // Brunei Dollar
	"BGN": BGN, // Bulgarian Lev
	"BIF": BIF, // Burundi Franc
	"CVE": CVE, // Cabo Verde Escudo
	"KHR": KHR, // Riel
	"XAF": XAF, // CFA Franc BEAC
	"CAD": CAD, // Canadian Dollar
	"KYD": KYD, // Cayman Islands Dollar
	"CLP": CLP, // Chilean Peso
	"CLF": CLF, // Unidad de Fomento
	"CNY": CNY, // Yuan Renminbi
	"COP": COP, // Colombian Peso
	"COU": COU, // Unidad de Valor Real
	"KMF": KMF, // Comorian Franc
	"CDF": CDF, // Congolese Franc
	"NZD": NZD, // New Zealand Dollar
	"CRC": CRC, // Costa Rican Colon
	"HRK": HRK, // Kuna
	"CUP": CUP, // Cuban Peso
	"CUC": CUC, // Peso Convertible
	"ANG": ANG, // Netherlands Antillean Guilder
	"CZK": CZK, // Czech Koruna
	"DKK": DKK, // Danish Krone
	"DJF": DJF, // Djibouti Franc
	"DOP": DOP, // Dominican Peso
	"EGP": EGP, // Egyptian Pound
	"SVC": SVC, // El Salvador Colon
	"ERN": ERN, // Nakfa
	"SZL": SZL, // Lilangeni
	"ETB": ETB, // Ethiopian Birr
	"FKP": FKP, // Falkland Islands Pound
	"FJD": FJD, // Fiji Dollar
	"XPF": XPF, // CFP Franc
	"GMD": GMD, // Dalasi
	"GEL": GEL, // Lari
	"GHS": GHS, // Ghana Cedi
	"GIP": GIP, // Gibraltar Pound
	"GTQ": GTQ, // Quetzal
	"GBP": GBP, // Pound Sterling
	"GNF": GNF, // Guinean Franc
	"GYD": GYD, // Guyana Dollar
	"HTG": HTG, // Gourde
	"HNL": HNL, // Lempira
	"HKD": HKD, // Hong Kong Dollar
	"HUF": HUF, // Forint
	"ISK": ISK, // Iceland Krona
	"IDR": IDR, // Rupiah
	"XDR": XDR, // SDR (Special Drawing Right)
	"IRR": IRR, // Iranian Rial
	"IQD": IQD, // Iraqi Dinar
	"ILS": ILS, // New Israeli Sheqel
	"JMD": JMD, // Jamaican Dollar
	"JPY": JPY, // Yen
	"JOD": JOD, // Jordanian Dinar
	"KZT": KZT, // Tenge
	"KES": KES, // Kenyan Shilling
	"KPW": KPW, // North Korean Won
	"KRW": KRW, // Won
	"KWD": KWD, // Kuwaiti Dinar
	"KGS": KGS, // Som
	"LAK": LAK, // Lao Kip
	"LBP": LBP, // Lebanese Pound
	"LSL": LSL, // Loti
	"ZAR": ZAR, // Rand
	"LRD": LRD, // Liberian Dollar
	"LYD": LYD, // Libyan Dinar
	"CHF": CHF, // Swiss Franc
	"MOP": MOP, // Pataca
	"MKD": MKD, // Denar
	"MGA": MGA, // Malagasy Ariary
	"MWK": MWK, // Malawi Kwacha
	"MYR": MYR, // Malaysian Ringgit
	"MVR": MVR, // Rufiyaa
	"MRU": MRU, // Ouguiya
	"MUR": MUR, // Mauritius Rupee
	"XUA": XUA, // ADB Unit of Account
	"MXN": MXN, // Mexican Peso
	"MXV": MXV, // Mexican Unidad de Inversion (UDI)
	"MDL": MDL, // Moldovan Leu
	"MNT": MNT, // Tugrik
	"MAD": MAD, // Moroccan Dirham
	"MZN": MZN, // Mozambique Metical
	"MMK": MMK, // Kyat
	"NAD": NAD, // Namibia Dollar
	"NPR": NPR, // Nepalese Rupee
	"NIO": NIO, // Cordoba Oro
	"NGN": NGN, // Naira
	"OMR": OMR, // Rial Omani
	"PKR": PKR, // Pakistan Rupee
	"PAB": PAB, // Balboa
	"PGK": PGK, // Kina
	"PYG": PYG, // Guarani
	"PEN": PEN, // Sol
	"PHP": PHP, // Philippine Peso
	"PLN": PLN, // Zloty
	"QAR": QAR, // Qatari Rial
	"RON": RON, // Romanian Leu
	"RUB": RUB, // Russian Ruble
	"RWF": RWF, // Rwanda Franc
	"SHP": SHP, // Saint Helena Pound
	"WST": WST, // Tala
	"STN": STN, // Dobra
	"SAR": SAR, // Saudi Riyal
	"RSD": RSD, // Serbian Dinar
	"SCR": SCR, // Seychelles Rupee
	"SLL": SLL, // Leone
	"SGD": SGD, // Singapore Dollar
	"XSU": XSU, // Sucre
	"SBD": SBD, // Solomon Islands Dollar
	"SOS": SOS, // Somali Shilling
	"SSP": SSP, // South Sudanese Pound
	"LKR": LKR, // Sri Lanka Rupee
	"SDG": SDG, // Sudanese Pound
	"SRD": SRD, // Surinam Dollar
	"SEK": SEK, // Swedish Krona
	"CHE": CHE, // WIR Euro
	"CHW": CHW, // WIR Franc
	"SYP": SYP, // Syrian Pound
	"TWD": TWD, // New Taiwan Dollar
	"TJS": TJS, // Somoni
	"TZS": TZS, // Tanzanian Shilling
	"THB": THB, // Baht
	"TOP": TOP, // Pa'anga
	"TTD": TTD, // Trinidad and Tobago Dollar
	"TND": TND, // Tunisian Dinar
	"TRY": TRY, // Turkish Lira
	"TMT": TMT, // Turkmenistan New Manat
	"UGX": UGX, // Uganda Shilling
	"UAH": UAH, // Hryvnia
	"AED": AED, // UAE Dirham
	"USN": USN, // US Dollar (Next day)
	"UYU": UYU, // Peso Uruguayo
	"UYI": UYI, // Uruguay Peso en Unidades Indexadas (UI)
	"UYW": UYW, // Unidad Previsional
	"UZS": UZS, // Uzbekistan Sum
	"VUV": VUV, // Vatu
	"VES": VES, // Bolívar Soberano
	"VND": VND, // Dong
	"YER": YER, // Yemeni Rial
	"ZMW": ZMW, // Zambian Kwacha
	"ZWL": ZWL, // Zimbabwe Dollar
	"XBA": XBA, // Bond Markets Unit European Composite Unit (EURCO)
	"XBB": XBB, // Bond Markets Unit European Monetary Unit (E.M.U.-6)
	"XBC": XBC, // Bond Markets Unit European Unit of Account 9 (E.U.A.-9)
	"XBD": XBD, // Bond Markets Unit European Unit of Account 17 (E.U.A.-17)
	"XTS": XTS, // Codes specifically reserved for testing purposes
	"XXX": XXX, // The codes assigned for transactions where no currency is involved
	"XAU": XAU, // Gold
	"XPD": XPD, // Palladium
	"XPT": XPT, // Platinum
	"XAG": XAG, // Silver
}
