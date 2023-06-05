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

// Generated on 2023-06-05T03:51:27Z by raphaelsantodomingo, any modifications will be overwritten
package iso4217

type CurrencyCode struct {
	Code, Name, MicroUnit string
}

func (cc CurrencyCode) String() string {
	return cc.Code
}

func (cc CurrencyCode) Valid() bool {
	_, exists := Lookup(cc.Code)
	return exists
}

var (
	AFN = CurrencyCode{Code: "AFN", Name: "Afghani", MicroUnit: "2"}
	EUR = CurrencyCode{Code: "EUR", Name: "Euro", MicroUnit: "2"}
	ALL = CurrencyCode{Code: "ALL", Name: "Lek", MicroUnit: "2"}
	DZD = CurrencyCode{Code: "DZD", Name: "Algerian Dinar", MicroUnit: "2"}
	USD = CurrencyCode{Code: "USD", Name: "US Dollar", MicroUnit: "2"}
	AOA = CurrencyCode{Code: "AOA", Name: "Kwanza", MicroUnit: "2"}
	XCD = CurrencyCode{Code: "XCD", Name: "East Caribbean Dollar", MicroUnit: "2"}
	ARS = CurrencyCode{Code: "ARS", Name: "Argentine Peso", MicroUnit: "2"}
	AMD = CurrencyCode{Code: "AMD", Name: "Armenian Dram", MicroUnit: "2"}
	AWG = CurrencyCode{Code: "AWG", Name: "Aruban Florin", MicroUnit: "2"}
	AUD = CurrencyCode{Code: "AUD", Name: "Australian Dollar", MicroUnit: "2"}
	AZN = CurrencyCode{Code: "AZN", Name: "Azerbaijan Manat", MicroUnit: "2"}
	BSD = CurrencyCode{Code: "BSD", Name: "Bahamian Dollar", MicroUnit: "2"}
	BHD = CurrencyCode{Code: "BHD", Name: "Bahraini Dinar", MicroUnit: "3"}
	BDT = CurrencyCode{Code: "BDT", Name: "Taka", MicroUnit: "2"}
	BBD = CurrencyCode{Code: "BBD", Name: "Barbados Dollar", MicroUnit: "2"}
	BYN = CurrencyCode{Code: "BYN", Name: "Belarusian Ruble", MicroUnit: "2"}
	BZD = CurrencyCode{Code: "BZD", Name: "Belize Dollar", MicroUnit: "2"}
	XOF = CurrencyCode{Code: "XOF", Name: "CFA Franc BCEAO", MicroUnit: "0"}
	BMD = CurrencyCode{Code: "BMD", Name: "Bermudian Dollar", MicroUnit: "2"}
	INR = CurrencyCode{Code: "INR", Name: "Indian Rupee", MicroUnit: "2"}
	BTN = CurrencyCode{Code: "BTN", Name: "Ngultrum", MicroUnit: "2"}
	BOB = CurrencyCode{Code: "BOB", Name: "Boliviano", MicroUnit: "2"}
	BOV = CurrencyCode{Code: "BOV", Name: "Mvdol", MicroUnit: "2"}
	BAM = CurrencyCode{Code: "BAM", Name: "Convertible Mark", MicroUnit: "2"}
	BWP = CurrencyCode{Code: "BWP", Name: "Pula", MicroUnit: "2"}
	NOK = CurrencyCode{Code: "NOK", Name: "Norwegian Krone", MicroUnit: "2"}
	BRL = CurrencyCode{Code: "BRL", Name: "Brazilian Real", MicroUnit: "2"}
	BND = CurrencyCode{Code: "BND", Name: "Brunei Dollar", MicroUnit: "2"}
	BGN = CurrencyCode{Code: "BGN", Name: "Bulgarian Lev", MicroUnit: "2"}
	BIF = CurrencyCode{Code: "BIF", Name: "Burundi Franc", MicroUnit: "0"}
	CVE = CurrencyCode{Code: "CVE", Name: "Cabo Verde Escudo", MicroUnit: "2"}
	KHR = CurrencyCode{Code: "KHR", Name: "Riel", MicroUnit: "2"}
	XAF = CurrencyCode{Code: "XAF", Name: "CFA Franc BEAC", MicroUnit: "0"}
	CAD = CurrencyCode{Code: "CAD", Name: "Canadian Dollar", MicroUnit: "2"}
	KYD = CurrencyCode{Code: "KYD", Name: "Cayman Islands Dollar", MicroUnit: "2"}
	CLP = CurrencyCode{Code: "CLP", Name: "Chilean Peso", MicroUnit: "0"}
	CLF = CurrencyCode{Code: "CLF", Name: "Unidad de Fomento", MicroUnit: "4"}
	CNY = CurrencyCode{Code: "CNY", Name: "Yuan Renminbi", MicroUnit: "2"}
	COP = CurrencyCode{Code: "COP", Name: "Colombian Peso", MicroUnit: "2"}
	COU = CurrencyCode{Code: "COU", Name: "Unidad de Valor Real", MicroUnit: "2"}
	KMF = CurrencyCode{Code: "KMF", Name: "Comorian Franc", MicroUnit: "0"}
	CDF = CurrencyCode{Code: "CDF", Name: "Congolese Franc", MicroUnit: "2"}
	NZD = CurrencyCode{Code: "NZD", Name: "New Zealand Dollar", MicroUnit: "2"}
	CRC = CurrencyCode{Code: "CRC", Name: "Costa Rican Colon", MicroUnit: "2"}
	HRK = CurrencyCode{Code: "HRK", Name: "Kuna", MicroUnit: "2"}
	CUP = CurrencyCode{Code: "CUP", Name: "Cuban Peso", MicroUnit: "2"}
	CUC = CurrencyCode{Code: "CUC", Name: "Peso Convertible", MicroUnit: "2"}
	ANG = CurrencyCode{Code: "ANG", Name: "Netherlands Antillean Guilder", MicroUnit: "2"}
	CZK = CurrencyCode{Code: "CZK", Name: "Czech Koruna", MicroUnit: "2"}
	DKK = CurrencyCode{Code: "DKK", Name: "Danish Krone", MicroUnit: "2"}
	DJF = CurrencyCode{Code: "DJF", Name: "Djibouti Franc", MicroUnit: "0"}
	DOP = CurrencyCode{Code: "DOP", Name: "Dominican Peso", MicroUnit: "2"}
	EGP = CurrencyCode{Code: "EGP", Name: "Egyptian Pound", MicroUnit: "2"}
	SVC = CurrencyCode{Code: "SVC", Name: "El Salvador Colon", MicroUnit: "2"}
	ERN = CurrencyCode{Code: "ERN", Name: "Nakfa", MicroUnit: "2"}
	SZL = CurrencyCode{Code: "SZL", Name: "Lilangeni", MicroUnit: "2"}
	ETB = CurrencyCode{Code: "ETB", Name: "Ethiopian Birr", MicroUnit: "2"}
	FKP = CurrencyCode{Code: "FKP", Name: "Falkland Islands Pound", MicroUnit: "2"}
	FJD = CurrencyCode{Code: "FJD", Name: "Fiji Dollar", MicroUnit: "2"}
	XPF = CurrencyCode{Code: "XPF", Name: "CFP Franc", MicroUnit: "0"}
	GMD = CurrencyCode{Code: "GMD", Name: "Dalasi", MicroUnit: "2"}
	GEL = CurrencyCode{Code: "GEL", Name: "Lari", MicroUnit: "2"}
	GHS = CurrencyCode{Code: "GHS", Name: "Ghana Cedi", MicroUnit: "2"}
	GIP = CurrencyCode{Code: "GIP", Name: "Gibraltar Pound", MicroUnit: "2"}
	GTQ = CurrencyCode{Code: "GTQ", Name: "Quetzal", MicroUnit: "2"}
	GBP = CurrencyCode{Code: "GBP", Name: "Pound Sterling", MicroUnit: "2"}
	GNF = CurrencyCode{Code: "GNF", Name: "Guinean Franc", MicroUnit: "0"}
	GYD = CurrencyCode{Code: "GYD", Name: "Guyana Dollar", MicroUnit: "2"}
	HTG = CurrencyCode{Code: "HTG", Name: "Gourde", MicroUnit: "2"}
	HNL = CurrencyCode{Code: "HNL", Name: "Lempira", MicroUnit: "2"}
	HKD = CurrencyCode{Code: "HKD", Name: "Hong Kong Dollar", MicroUnit: "2"}
	HUF = CurrencyCode{Code: "HUF", Name: "Forint", MicroUnit: "2"}
	ISK = CurrencyCode{Code: "ISK", Name: "Iceland Krona", MicroUnit: "0"}
	IDR = CurrencyCode{Code: "IDR", Name: "Rupiah", MicroUnit: "2"}
	XDR = CurrencyCode{Code: "XDR", Name: "SDR (Special Drawing Right)", MicroUnit: "-"}
	IRR = CurrencyCode{Code: "IRR", Name: "Iranian Rial", MicroUnit: "2"}
	IQD = CurrencyCode{Code: "IQD", Name: "Iraqi Dinar", MicroUnit: "3"}
	ILS = CurrencyCode{Code: "ILS", Name: "New Israeli Sheqel", MicroUnit: "2"}
	JMD = CurrencyCode{Code: "JMD", Name: "Jamaican Dollar", MicroUnit: "2"}
	JPY = CurrencyCode{Code: "JPY", Name: "Yen", MicroUnit: "0"}
	JOD = CurrencyCode{Code: "JOD", Name: "Jordanian Dinar", MicroUnit: "3"}
	KZT = CurrencyCode{Code: "KZT", Name: "Tenge", MicroUnit: "2"}
	KES = CurrencyCode{Code: "KES", Name: "Kenyan Shilling", MicroUnit: "2"}
	KPW = CurrencyCode{Code: "KPW", Name: "North Korean Won", MicroUnit: "2"}
	KRW = CurrencyCode{Code: "KRW", Name: "Won", MicroUnit: "0"}
	KWD = CurrencyCode{Code: "KWD", Name: "Kuwaiti Dinar", MicroUnit: "3"}
	KGS = CurrencyCode{Code: "KGS", Name: "Som", MicroUnit: "2"}
	LAK = CurrencyCode{Code: "LAK", Name: "Lao Kip", MicroUnit: "2"}
	LBP = CurrencyCode{Code: "LBP", Name: "Lebanese Pound", MicroUnit: "2"}
	LSL = CurrencyCode{Code: "LSL", Name: "Loti", MicroUnit: "2"}
	ZAR = CurrencyCode{Code: "ZAR", Name: "Rand", MicroUnit: "2"}
	LRD = CurrencyCode{Code: "LRD", Name: "Liberian Dollar", MicroUnit: "2"}
	LYD = CurrencyCode{Code: "LYD", Name: "Libyan Dinar", MicroUnit: "3"}
	CHF = CurrencyCode{Code: "CHF", Name: "Swiss Franc", MicroUnit: "2"}
	MOP = CurrencyCode{Code: "MOP", Name: "Pataca", MicroUnit: "2"}
	MKD = CurrencyCode{Code: "MKD", Name: "Denar", MicroUnit: "2"}
	MGA = CurrencyCode{Code: "MGA", Name: "Malagasy Ariary", MicroUnit: "2"}
	MWK = CurrencyCode{Code: "MWK", Name: "Malawi Kwacha", MicroUnit: "2"}
	MYR = CurrencyCode{Code: "MYR", Name: "Malaysian Ringgit", MicroUnit: "2"}
	MVR = CurrencyCode{Code: "MVR", Name: "Rufiyaa", MicroUnit: "2"}
	MRU = CurrencyCode{Code: "MRU", Name: "Ouguiya", MicroUnit: "2"}
	MUR = CurrencyCode{Code: "MUR", Name: "Mauritius Rupee", MicroUnit: "2"}
	XUA = CurrencyCode{Code: "XUA", Name: "ADB Unit of Account", MicroUnit: "-"}
	MXN = CurrencyCode{Code: "MXN", Name: "Mexican Peso", MicroUnit: "2"}
	MXV = CurrencyCode{Code: "MXV", Name: "Mexican Unidad de Inversion (UDI)", MicroUnit: "2"}
	MDL = CurrencyCode{Code: "MDL", Name: "Moldovan Leu", MicroUnit: "2"}
	MNT = CurrencyCode{Code: "MNT", Name: "Tugrik", MicroUnit: "2"}
	MAD = CurrencyCode{Code: "MAD", Name: "Moroccan Dirham", MicroUnit: "2"}
	MZN = CurrencyCode{Code: "MZN", Name: "Mozambique Metical", MicroUnit: "2"}
	MMK = CurrencyCode{Code: "MMK", Name: "Kyat", MicroUnit: "2"}
	NAD = CurrencyCode{Code: "NAD", Name: "Namibia Dollar", MicroUnit: "2"}
	NPR = CurrencyCode{Code: "NPR", Name: "Nepalese Rupee", MicroUnit: "2"}
	NIO = CurrencyCode{Code: "NIO", Name: "Cordoba Oro", MicroUnit: "2"}
	NGN = CurrencyCode{Code: "NGN", Name: "Naira", MicroUnit: "2"}
	OMR = CurrencyCode{Code: "OMR", Name: "Rial Omani", MicroUnit: "3"}
	PKR = CurrencyCode{Code: "PKR", Name: "Pakistan Rupee", MicroUnit: "2"}
	PAB = CurrencyCode{Code: "PAB", Name: "Balboa", MicroUnit: "2"}
	PGK = CurrencyCode{Code: "PGK", Name: "Kina", MicroUnit: "2"}
	PYG = CurrencyCode{Code: "PYG", Name: "Guarani", MicroUnit: "0"}
	PEN = CurrencyCode{Code: "PEN", Name: "Sol", MicroUnit: "2"}
	PHP = CurrencyCode{Code: "PHP", Name: "Philippine Peso", MicroUnit: "2"}
	PLN = CurrencyCode{Code: "PLN", Name: "Zloty", MicroUnit: "2"}
	QAR = CurrencyCode{Code: "QAR", Name: "Qatari Rial", MicroUnit: "2"}
	RON = CurrencyCode{Code: "RON", Name: "Romanian Leu", MicroUnit: "2"}
	RUB = CurrencyCode{Code: "RUB", Name: "Russian Ruble", MicroUnit: "2"}
	RWF = CurrencyCode{Code: "RWF", Name: "Rwanda Franc", MicroUnit: "0"}
	SHP = CurrencyCode{Code: "SHP", Name: "Saint Helena Pound", MicroUnit: "2"}
	WST = CurrencyCode{Code: "WST", Name: "Tala", MicroUnit: "2"}
	STN = CurrencyCode{Code: "STN", Name: "Dobra", MicroUnit: "2"}
	SAR = CurrencyCode{Code: "SAR", Name: "Saudi Riyal", MicroUnit: "2"}
	RSD = CurrencyCode{Code: "RSD", Name: "Serbian Dinar", MicroUnit: "2"}
	SCR = CurrencyCode{Code: "SCR", Name: "Seychelles Rupee", MicroUnit: "2"}
	SLL = CurrencyCode{Code: "SLL", Name: "Leone", MicroUnit: "2"}
	SGD = CurrencyCode{Code: "SGD", Name: "Singapore Dollar", MicroUnit: "2"}
	XSU = CurrencyCode{Code: "XSU", Name: "Sucre", MicroUnit: "-"}
	SBD = CurrencyCode{Code: "SBD", Name: "Solomon Islands Dollar", MicroUnit: "2"}
	SOS = CurrencyCode{Code: "SOS", Name: "Somali Shilling", MicroUnit: "2"}
	SSP = CurrencyCode{Code: "SSP", Name: "South Sudanese Pound", MicroUnit: "2"}
	LKR = CurrencyCode{Code: "LKR", Name: "Sri Lanka Rupee", MicroUnit: "2"}
	SDG = CurrencyCode{Code: "SDG", Name: "Sudanese Pound", MicroUnit: "2"}
	SRD = CurrencyCode{Code: "SRD", Name: "Surinam Dollar", MicroUnit: "2"}
	SEK = CurrencyCode{Code: "SEK", Name: "Swedish Krona", MicroUnit: "2"}
	CHE = CurrencyCode{Code: "CHE", Name: "WIR Euro", MicroUnit: "2"}
	CHW = CurrencyCode{Code: "CHW", Name: "WIR Franc", MicroUnit: "2"}
	SYP = CurrencyCode{Code: "SYP", Name: "Syrian Pound", MicroUnit: "2"}
	TWD = CurrencyCode{Code: "TWD", Name: "New Taiwan Dollar", MicroUnit: "2"}
	TJS = CurrencyCode{Code: "TJS", Name: "Somoni", MicroUnit: "2"}
	TZS = CurrencyCode{Code: "TZS", Name: "Tanzanian Shilling", MicroUnit: "2"}
	THB = CurrencyCode{Code: "THB", Name: "Baht", MicroUnit: "2"}
	TOP = CurrencyCode{Code: "TOP", Name: "Pa'anga", MicroUnit: "2"}
	TTD = CurrencyCode{Code: "TTD", Name: "Trinidad and Tobago Dollar", MicroUnit: "2"}
	TND = CurrencyCode{Code: "TND", Name: "Tunisian Dinar", MicroUnit: "3"}
	TRY = CurrencyCode{Code: "TRY", Name: "Turkish Lira", MicroUnit: "2"}
	TMT = CurrencyCode{Code: "TMT", Name: "Turkmenistan New Manat", MicroUnit: "2"}
	UGX = CurrencyCode{Code: "UGX", Name: "Uganda Shilling", MicroUnit: "0"}
	UAH = CurrencyCode{Code: "UAH", Name: "Hryvnia", MicroUnit: "2"}
	AED = CurrencyCode{Code: "AED", Name: "UAE Dirham", MicroUnit: "2"}
	USN = CurrencyCode{Code: "USN", Name: "US Dollar (Next day)", MicroUnit: "2"}
	UYU = CurrencyCode{Code: "UYU", Name: "Peso Uruguayo", MicroUnit: "2"}
	UYI = CurrencyCode{Code: "UYI", Name: "Uruguay Peso en Unidades Indexadas (UI)", MicroUnit: "0"}
	UYW = CurrencyCode{Code: "UYW", Name: "Unidad Previsional", MicroUnit: "4"}
	UZS = CurrencyCode{Code: "UZS", Name: "Uzbekistan Sum", MicroUnit: "2"}
	VUV = CurrencyCode{Code: "VUV", Name: "Vatu", MicroUnit: "0"}
	VES = CurrencyCode{Code: "VES", Name: "Bolívar Soberano", MicroUnit: "2"}
	VND = CurrencyCode{Code: "VND", Name: "Dong", MicroUnit: "0"}
	YER = CurrencyCode{Code: "YER", Name: "Yemeni Rial", MicroUnit: "2"}
	ZMW = CurrencyCode{Code: "ZMW", Name: "Zambian Kwacha", MicroUnit: "2"}
	ZWL = CurrencyCode{Code: "ZWL", Name: "Zimbabwe Dollar", MicroUnit: "2"}
	XBA = CurrencyCode{Code: "XBA", Name: "Bond Markets Unit European Composite Unit (EURCO)", MicroUnit: "-"}
	XBB = CurrencyCode{Code: "XBB", Name: "Bond Markets Unit European Monetary Unit (E.M.U.-6)", MicroUnit: "-"}
	XBC = CurrencyCode{Code: "XBC", Name: "Bond Markets Unit European Unit of Account 9 (E.U.A.-9)", MicroUnit: "-"}
	XBD = CurrencyCode{Code: "XBD", Name: "Bond Markets Unit European Unit of Account 17 (E.U.A.-17)", MicroUnit: "-"}
	XTS = CurrencyCode{Code: "XTS", Name: "Codes specifically reserved for testing purposes", MicroUnit: "-"}
	XXX = CurrencyCode{Code: "XXX", Name: "The codes assigned for transactions where no currency is involved", MicroUnit: "-"}
	XAU = CurrencyCode{Code: "XAU", Name: "Gold", MicroUnit: "-"}
	XPD = CurrencyCode{Code: "XPD", Name: "Palladium", MicroUnit: "-"}
	XPT = CurrencyCode{Code: "XPT", Name: "Platinum", MicroUnit: "-"}
	XAG = CurrencyCode{Code: "XAG", Name: "Silver", MicroUnit: "-"}
)

var lookupTable = map[string]CurrencyCode{
	"AFN": AFN, // Afghani // 2
	"EUR": EUR, // Euro // 2
	"ALL": ALL, // Lek // 2
	"DZD": DZD, // Algerian Dinar // 2
	"USD": USD, // US Dollar // 2
	"AOA": AOA, // Kwanza // 2
	"XCD": XCD, // East Caribbean Dollar // 2
	"ARS": ARS, // Argentine Peso // 2
	"AMD": AMD, // Armenian Dram // 2
	"AWG": AWG, // Aruban Florin // 2
	"AUD": AUD, // Australian Dollar // 2
	"AZN": AZN, // Azerbaijan Manat // 2
	"BSD": BSD, // Bahamian Dollar // 2
	"BHD": BHD, // Bahraini Dinar // 3
	"BDT": BDT, // Taka // 2
	"BBD": BBD, // Barbados Dollar // 2
	"BYN": BYN, // Belarusian Ruble // 2
	"BZD": BZD, // Belize Dollar // 2
	"XOF": XOF, // CFA Franc BCEAO // 0
	"BMD": BMD, // Bermudian Dollar // 2
	"INR": INR, // Indian Rupee // 2
	"BTN": BTN, // Ngultrum // 2
	"BOB": BOB, // Boliviano // 2
	"BOV": BOV, // Mvdol // 2
	"BAM": BAM, // Convertible Mark // 2
	"BWP": BWP, // Pula // 2
	"NOK": NOK, // Norwegian Krone // 2
	"BRL": BRL, // Brazilian Real // 2
	"BND": BND, // Brunei Dollar // 2
	"BGN": BGN, // Bulgarian Lev // 2
	"BIF": BIF, // Burundi Franc // 0
	"CVE": CVE, // Cabo Verde Escudo // 2
	"KHR": KHR, // Riel // 2
	"XAF": XAF, // CFA Franc BEAC // 0
	"CAD": CAD, // Canadian Dollar // 2
	"KYD": KYD, // Cayman Islands Dollar // 2
	"CLP": CLP, // Chilean Peso // 0
	"CLF": CLF, // Unidad de Fomento // 4
	"CNY": CNY, // Yuan Renminbi // 2
	"COP": COP, // Colombian Peso // 2
	"COU": COU, // Unidad de Valor Real // 2
	"KMF": KMF, // Comorian Franc // 0
	"CDF": CDF, // Congolese Franc // 2
	"NZD": NZD, // New Zealand Dollar // 2
	"CRC": CRC, // Costa Rican Colon // 2
	"HRK": HRK, // Kuna // 2
	"CUP": CUP, // Cuban Peso // 2
	"CUC": CUC, // Peso Convertible // 2
	"ANG": ANG, // Netherlands Antillean Guilder // 2
	"CZK": CZK, // Czech Koruna // 2
	"DKK": DKK, // Danish Krone // 2
	"DJF": DJF, // Djibouti Franc // 0
	"DOP": DOP, // Dominican Peso // 2
	"EGP": EGP, // Egyptian Pound // 2
	"SVC": SVC, // El Salvador Colon // 2
	"ERN": ERN, // Nakfa // 2
	"SZL": SZL, // Lilangeni // 2
	"ETB": ETB, // Ethiopian Birr // 2
	"FKP": FKP, // Falkland Islands Pound // 2
	"FJD": FJD, // Fiji Dollar // 2
	"XPF": XPF, // CFP Franc // 0
	"GMD": GMD, // Dalasi // 2
	"GEL": GEL, // Lari // 2
	"GHS": GHS, // Ghana Cedi // 2
	"GIP": GIP, // Gibraltar Pound // 2
	"GTQ": GTQ, // Quetzal // 2
	"GBP": GBP, // Pound Sterling // 2
	"GNF": GNF, // Guinean Franc // 0
	"GYD": GYD, // Guyana Dollar // 2
	"HTG": HTG, // Gourde // 2
	"HNL": HNL, // Lempira // 2
	"HKD": HKD, // Hong Kong Dollar // 2
	"HUF": HUF, // Forint // 2
	"ISK": ISK, // Iceland Krona // 0
	"IDR": IDR, // Rupiah // 2
	"XDR": XDR, // SDR (Special Drawing Right) // -
	"IRR": IRR, // Iranian Rial // 2
	"IQD": IQD, // Iraqi Dinar // 3
	"ILS": ILS, // New Israeli Sheqel // 2
	"JMD": JMD, // Jamaican Dollar // 2
	"JPY": JPY, // Yen // 0
	"JOD": JOD, // Jordanian Dinar // 3
	"KZT": KZT, // Tenge // 2
	"KES": KES, // Kenyan Shilling // 2
	"KPW": KPW, // North Korean Won // 2
	"KRW": KRW, // Won // 0
	"KWD": KWD, // Kuwaiti Dinar // 3
	"KGS": KGS, // Som // 2
	"LAK": LAK, // Lao Kip // 2
	"LBP": LBP, // Lebanese Pound // 2
	"LSL": LSL, // Loti // 2
	"ZAR": ZAR, // Rand // 2
	"LRD": LRD, // Liberian Dollar // 2
	"LYD": LYD, // Libyan Dinar // 3
	"CHF": CHF, // Swiss Franc // 2
	"MOP": MOP, // Pataca // 2
	"MKD": MKD, // Denar // 2
	"MGA": MGA, // Malagasy Ariary // 2
	"MWK": MWK, // Malawi Kwacha // 2
	"MYR": MYR, // Malaysian Ringgit // 2
	"MVR": MVR, // Rufiyaa // 2
	"MRU": MRU, // Ouguiya // 2
	"MUR": MUR, // Mauritius Rupee // 2
	"XUA": XUA, // ADB Unit of Account // -
	"MXN": MXN, // Mexican Peso // 2
	"MXV": MXV, // Mexican Unidad de Inversion (UDI) // 2
	"MDL": MDL, // Moldovan Leu // 2
	"MNT": MNT, // Tugrik // 2
	"MAD": MAD, // Moroccan Dirham // 2
	"MZN": MZN, // Mozambique Metical // 2
	"MMK": MMK, // Kyat // 2
	"NAD": NAD, // Namibia Dollar // 2
	"NPR": NPR, // Nepalese Rupee // 2
	"NIO": NIO, // Cordoba Oro // 2
	"NGN": NGN, // Naira // 2
	"OMR": OMR, // Rial Omani // 3
	"PKR": PKR, // Pakistan Rupee // 2
	"PAB": PAB, // Balboa // 2
	"PGK": PGK, // Kina // 2
	"PYG": PYG, // Guarani // 0
	"PEN": PEN, // Sol // 2
	"PHP": PHP, // Philippine Peso // 2
	"PLN": PLN, // Zloty // 2
	"QAR": QAR, // Qatari Rial // 2
	"RON": RON, // Romanian Leu // 2
	"RUB": RUB, // Russian Ruble // 2
	"RWF": RWF, // Rwanda Franc // 0
	"SHP": SHP, // Saint Helena Pound // 2
	"WST": WST, // Tala // 2
	"STN": STN, // Dobra // 2
	"SAR": SAR, // Saudi Riyal // 2
	"RSD": RSD, // Serbian Dinar // 2
	"SCR": SCR, // Seychelles Rupee // 2
	"SLL": SLL, // Leone // 2
	"SGD": SGD, // Singapore Dollar // 2
	"XSU": XSU, // Sucre // -
	"SBD": SBD, // Solomon Islands Dollar // 2
	"SOS": SOS, // Somali Shilling // 2
	"SSP": SSP, // South Sudanese Pound // 2
	"LKR": LKR, // Sri Lanka Rupee // 2
	"SDG": SDG, // Sudanese Pound // 2
	"SRD": SRD, // Surinam Dollar // 2
	"SEK": SEK, // Swedish Krona // 2
	"CHE": CHE, // WIR Euro // 2
	"CHW": CHW, // WIR Franc // 2
	"SYP": SYP, // Syrian Pound // 2
	"TWD": TWD, // New Taiwan Dollar // 2
	"TJS": TJS, // Somoni // 2
	"TZS": TZS, // Tanzanian Shilling // 2
	"THB": THB, // Baht // 2
	"TOP": TOP, // Pa'anga // 2
	"TTD": TTD, // Trinidad and Tobago Dollar // 2
	"TND": TND, // Tunisian Dinar // 3
	"TRY": TRY, // Turkish Lira // 2
	"TMT": TMT, // Turkmenistan New Manat // 2
	"UGX": UGX, // Uganda Shilling // 0
	"UAH": UAH, // Hryvnia // 2
	"AED": AED, // UAE Dirham // 2
	"USN": USN, // US Dollar (Next day) // 2
	"UYU": UYU, // Peso Uruguayo // 2
	"UYI": UYI, // Uruguay Peso en Unidades Indexadas (UI) // 0
	"UYW": UYW, // Unidad Previsional // 4
	"UZS": UZS, // Uzbekistan Sum // 2
	"VUV": VUV, // Vatu // 0
	"VES": VES, // Bolívar Soberano // 2
	"VND": VND, // Dong // 0
	"YER": YER, // Yemeni Rial // 2
	"ZMW": ZMW, // Zambian Kwacha // 2
	"ZWL": ZWL, // Zimbabwe Dollar // 2
	"XBA": XBA, // Bond Markets Unit European Composite Unit (EURCO) // -
	"XBB": XBB, // Bond Markets Unit European Monetary Unit (E.M.U.-6) // -
	"XBC": XBC, // Bond Markets Unit European Unit of Account 9 (E.U.A.-9) // -
	"XBD": XBD, // Bond Markets Unit European Unit of Account 17 (E.U.A.-17) // -
	"XTS": XTS, // Codes specifically reserved for testing purposes // -
	"XXX": XXX, // The codes assigned for transactions where no currency is involved // -
	"XAU": XAU, // Gold // -
	"XPD": XPD, // Palladium // -
	"XPT": XPT, // Platinum // -
	"XAG": XAG, // Silver // -
}
