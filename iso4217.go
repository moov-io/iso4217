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

// Generated on 2023-06-02T09:32:24Z by adam, any modifications will be overwritten
package iso4217

type CurrencyCode struct {
	Code, Name string
}

func (cc CurrencyCode) String() string {
	return cc.Code
}

func (cc CurrencyCode) Valid() bool {
	_, exists := Lookup(cc.Code)
	return exists
}

var (
	AFN = CurrencyCode{Code: "AFN", Name: "Afghani"}
	EUR = CurrencyCode{Code: "EUR", Name: "Euro"}
	ALL = CurrencyCode{Code: "ALL", Name: "Lek"}
	DZD = CurrencyCode{Code: "DZD", Name: "Algerian Dinar"}
	USD = CurrencyCode{Code: "USD", Name: "US Dollar"}
	AOA = CurrencyCode{Code: "AOA", Name: "Kwanza"}
	XCD = CurrencyCode{Code: "XCD", Name: "East Caribbean Dollar"}
	ARS = CurrencyCode{Code: "ARS", Name: "Argentine Peso"}
	AMD = CurrencyCode{Code: "AMD", Name: "Armenian Dram"}
	AWG = CurrencyCode{Code: "AWG", Name: "Aruban Florin"}
	AUD = CurrencyCode{Code: "AUD", Name: "Australian Dollar"}
	AZN = CurrencyCode{Code: "AZN", Name: "Azerbaijan Manat"}
	BSD = CurrencyCode{Code: "BSD", Name: "Bahamian Dollar"}
	BHD = CurrencyCode{Code: "BHD", Name: "Bahraini Dinar"}
	BDT = CurrencyCode{Code: "BDT", Name: "Taka"}
	BBD = CurrencyCode{Code: "BBD", Name: "Barbados Dollar"}
	BYN = CurrencyCode{Code: "BYN", Name: "Belarusian Ruble"}
	BZD = CurrencyCode{Code: "BZD", Name: "Belize Dollar"}
	XOF = CurrencyCode{Code: "XOF", Name: "CFA Franc BCEAO"}
	BMD = CurrencyCode{Code: "BMD", Name: "Bermudian Dollar"}
	INR = CurrencyCode{Code: "INR", Name: "Indian Rupee"}
	BTN = CurrencyCode{Code: "BTN", Name: "Ngultrum"}
	BOB = CurrencyCode{Code: "BOB", Name: "Boliviano"}
	BOV = CurrencyCode{Code: "BOV", Name: "Mvdol"}
	BAM = CurrencyCode{Code: "BAM", Name: "Convertible Mark"}
	BWP = CurrencyCode{Code: "BWP", Name: "Pula"}
	NOK = CurrencyCode{Code: "NOK", Name: "Norwegian Krone"}
	BRL = CurrencyCode{Code: "BRL", Name: "Brazilian Real"}
	BND = CurrencyCode{Code: "BND", Name: "Brunei Dollar"}
	BGN = CurrencyCode{Code: "BGN", Name: "Bulgarian Lev"}
	BIF = CurrencyCode{Code: "BIF", Name: "Burundi Franc"}
	CVE = CurrencyCode{Code: "CVE", Name: "Cabo Verde Escudo"}
	KHR = CurrencyCode{Code: "KHR", Name: "Riel"}
	XAF = CurrencyCode{Code: "XAF", Name: "CFA Franc BEAC"}
	CAD = CurrencyCode{Code: "CAD", Name: "Canadian Dollar"}
	KYD = CurrencyCode{Code: "KYD", Name: "Cayman Islands Dollar"}
	CLP = CurrencyCode{Code: "CLP", Name: "Chilean Peso"}
	CLF = CurrencyCode{Code: "CLF", Name: "Unidad de Fomento"}
	CNY = CurrencyCode{Code: "CNY", Name: "Yuan Renminbi"}
	COP = CurrencyCode{Code: "COP", Name: "Colombian Peso"}
	COU = CurrencyCode{Code: "COU", Name: "Unidad de Valor Real"}
	KMF = CurrencyCode{Code: "KMF", Name: "Comorian Franc"}
	CDF = CurrencyCode{Code: "CDF", Name: "Congolese Franc"}
	NZD = CurrencyCode{Code: "NZD", Name: "New Zealand Dollar"}
	CRC = CurrencyCode{Code: "CRC", Name: "Costa Rican Colon"}
	HRK = CurrencyCode{Code: "HRK", Name: "Kuna"}
	CUP = CurrencyCode{Code: "CUP", Name: "Cuban Peso"}
	CUC = CurrencyCode{Code: "CUC", Name: "Peso Convertible"}
	ANG = CurrencyCode{Code: "ANG", Name: "Netherlands Antillean Guilder"}
	CZK = CurrencyCode{Code: "CZK", Name: "Czech Koruna"}
	DKK = CurrencyCode{Code: "DKK", Name: "Danish Krone"}
	DJF = CurrencyCode{Code: "DJF", Name: "Djibouti Franc"}
	DOP = CurrencyCode{Code: "DOP", Name: "Dominican Peso"}
	EGP = CurrencyCode{Code: "EGP", Name: "Egyptian Pound"}
	SVC = CurrencyCode{Code: "SVC", Name: "El Salvador Colon"}
	ERN = CurrencyCode{Code: "ERN", Name: "Nakfa"}
	SZL = CurrencyCode{Code: "SZL", Name: "Lilangeni"}
	ETB = CurrencyCode{Code: "ETB", Name: "Ethiopian Birr"}
	FKP = CurrencyCode{Code: "FKP", Name: "Falkland Islands Pound"}
	FJD = CurrencyCode{Code: "FJD", Name: "Fiji Dollar"}
	XPF = CurrencyCode{Code: "XPF", Name: "CFP Franc"}
	GMD = CurrencyCode{Code: "GMD", Name: "Dalasi"}
	GEL = CurrencyCode{Code: "GEL", Name: "Lari"}
	GHS = CurrencyCode{Code: "GHS", Name: "Ghana Cedi"}
	GIP = CurrencyCode{Code: "GIP", Name: "Gibraltar Pound"}
	GTQ = CurrencyCode{Code: "GTQ", Name: "Quetzal"}
	GBP = CurrencyCode{Code: "GBP", Name: "Pound Sterling"}
	GNF = CurrencyCode{Code: "GNF", Name: "Guinean Franc"}
	GYD = CurrencyCode{Code: "GYD", Name: "Guyana Dollar"}
	HTG = CurrencyCode{Code: "HTG", Name: "Gourde"}
	HNL = CurrencyCode{Code: "HNL", Name: "Lempira"}
	HKD = CurrencyCode{Code: "HKD", Name: "Hong Kong Dollar"}
	HUF = CurrencyCode{Code: "HUF", Name: "Forint"}
	ISK = CurrencyCode{Code: "ISK", Name: "Iceland Krona"}
	IDR = CurrencyCode{Code: "IDR", Name: "Rupiah"}
	XDR = CurrencyCode{Code: "XDR", Name: "SDR (Special Drawing Right)"}
	IRR = CurrencyCode{Code: "IRR", Name: "Iranian Rial"}
	IQD = CurrencyCode{Code: "IQD", Name: "Iraqi Dinar"}
	ILS = CurrencyCode{Code: "ILS", Name: "New Israeli Sheqel"}
	JMD = CurrencyCode{Code: "JMD", Name: "Jamaican Dollar"}
	JPY = CurrencyCode{Code: "JPY", Name: "Yen"}
	JOD = CurrencyCode{Code: "JOD", Name: "Jordanian Dinar"}
	KZT = CurrencyCode{Code: "KZT", Name: "Tenge"}
	KES = CurrencyCode{Code: "KES", Name: "Kenyan Shilling"}
	KPW = CurrencyCode{Code: "KPW", Name: "North Korean Won"}
	KRW = CurrencyCode{Code: "KRW", Name: "Won"}
	KWD = CurrencyCode{Code: "KWD", Name: "Kuwaiti Dinar"}
	KGS = CurrencyCode{Code: "KGS", Name: "Som"}
	LAK = CurrencyCode{Code: "LAK", Name: "Lao Kip"}
	LBP = CurrencyCode{Code: "LBP", Name: "Lebanese Pound"}
	LSL = CurrencyCode{Code: "LSL", Name: "Loti"}
	ZAR = CurrencyCode{Code: "ZAR", Name: "Rand"}
	LRD = CurrencyCode{Code: "LRD", Name: "Liberian Dollar"}
	LYD = CurrencyCode{Code: "LYD", Name: "Libyan Dinar"}
	CHF = CurrencyCode{Code: "CHF", Name: "Swiss Franc"}
	MOP = CurrencyCode{Code: "MOP", Name: "Pataca"}
	MKD = CurrencyCode{Code: "MKD", Name: "Denar"}
	MGA = CurrencyCode{Code: "MGA", Name: "Malagasy Ariary"}
	MWK = CurrencyCode{Code: "MWK", Name: "Malawi Kwacha"}
	MYR = CurrencyCode{Code: "MYR", Name: "Malaysian Ringgit"}
	MVR = CurrencyCode{Code: "MVR", Name: "Rufiyaa"}
	MRU = CurrencyCode{Code: "MRU", Name: "Ouguiya"}
	MUR = CurrencyCode{Code: "MUR", Name: "Mauritius Rupee"}
	XUA = CurrencyCode{Code: "XUA", Name: "ADB Unit of Account"}
	MXN = CurrencyCode{Code: "MXN", Name: "Mexican Peso"}
	MXV = CurrencyCode{Code: "MXV", Name: "Mexican Unidad de Inversion (UDI)"}
	MDL = CurrencyCode{Code: "MDL", Name: "Moldovan Leu"}
	MNT = CurrencyCode{Code: "MNT", Name: "Tugrik"}
	MAD = CurrencyCode{Code: "MAD", Name: "Moroccan Dirham"}
	MZN = CurrencyCode{Code: "MZN", Name: "Mozambique Metical"}
	MMK = CurrencyCode{Code: "MMK", Name: "Kyat"}
	NAD = CurrencyCode{Code: "NAD", Name: "Namibia Dollar"}
	NPR = CurrencyCode{Code: "NPR", Name: "Nepalese Rupee"}
	NIO = CurrencyCode{Code: "NIO", Name: "Cordoba Oro"}
	NGN = CurrencyCode{Code: "NGN", Name: "Naira"}
	OMR = CurrencyCode{Code: "OMR", Name: "Rial Omani"}
	PKR = CurrencyCode{Code: "PKR", Name: "Pakistan Rupee"}
	PAB = CurrencyCode{Code: "PAB", Name: "Balboa"}
	PGK = CurrencyCode{Code: "PGK", Name: "Kina"}
	PYG = CurrencyCode{Code: "PYG", Name: "Guarani"}
	PEN = CurrencyCode{Code: "PEN", Name: "Sol"}
	PHP = CurrencyCode{Code: "PHP", Name: "Philippine Peso"}
	PLN = CurrencyCode{Code: "PLN", Name: "Zloty"}
	QAR = CurrencyCode{Code: "QAR", Name: "Qatari Rial"}
	RON = CurrencyCode{Code: "RON", Name: "Romanian Leu"}
	RUB = CurrencyCode{Code: "RUB", Name: "Russian Ruble"}
	RWF = CurrencyCode{Code: "RWF", Name: "Rwanda Franc"}
	SHP = CurrencyCode{Code: "SHP", Name: "Saint Helena Pound"}
	WST = CurrencyCode{Code: "WST", Name: "Tala"}
	STN = CurrencyCode{Code: "STN", Name: "Dobra"}
	SAR = CurrencyCode{Code: "SAR", Name: "Saudi Riyal"}
	RSD = CurrencyCode{Code: "RSD", Name: "Serbian Dinar"}
	SCR = CurrencyCode{Code: "SCR", Name: "Seychelles Rupee"}
	SLL = CurrencyCode{Code: "SLL", Name: "Leone"}
	SGD = CurrencyCode{Code: "SGD", Name: "Singapore Dollar"}
	XSU = CurrencyCode{Code: "XSU", Name: "Sucre"}
	SBD = CurrencyCode{Code: "SBD", Name: "Solomon Islands Dollar"}
	SOS = CurrencyCode{Code: "SOS", Name: "Somali Shilling"}
	SSP = CurrencyCode{Code: "SSP", Name: "South Sudanese Pound"}
	LKR = CurrencyCode{Code: "LKR", Name: "Sri Lanka Rupee"}
	SDG = CurrencyCode{Code: "SDG", Name: "Sudanese Pound"}
	SRD = CurrencyCode{Code: "SRD", Name: "Surinam Dollar"}
	SEK = CurrencyCode{Code: "SEK", Name: "Swedish Krona"}
	CHE = CurrencyCode{Code: "CHE", Name: "WIR Euro"}
	CHW = CurrencyCode{Code: "CHW", Name: "WIR Franc"}
	SYP = CurrencyCode{Code: "SYP", Name: "Syrian Pound"}
	TWD = CurrencyCode{Code: "TWD", Name: "New Taiwan Dollar"}
	TJS = CurrencyCode{Code: "TJS", Name: "Somoni"}
	TZS = CurrencyCode{Code: "TZS", Name: "Tanzanian Shilling"}
	THB = CurrencyCode{Code: "THB", Name: "Baht"}
	TOP = CurrencyCode{Code: "TOP", Name: "Pa'anga"}
	TTD = CurrencyCode{Code: "TTD", Name: "Trinidad and Tobago Dollar"}
	TND = CurrencyCode{Code: "TND", Name: "Tunisian Dinar"}
	TRY = CurrencyCode{Code: "TRY", Name: "Turkish Lira"}
	TMT = CurrencyCode{Code: "TMT", Name: "Turkmenistan New Manat"}
	UGX = CurrencyCode{Code: "UGX", Name: "Uganda Shilling"}
	UAH = CurrencyCode{Code: "UAH", Name: "Hryvnia"}
	AED = CurrencyCode{Code: "AED", Name: "UAE Dirham"}
	USN = CurrencyCode{Code: "USN", Name: "US Dollar (Next day)"}
	UYU = CurrencyCode{Code: "UYU", Name: "Peso Uruguayo"}
	UYI = CurrencyCode{Code: "UYI", Name: "Uruguay Peso en Unidades Indexadas (UI)"}
	UYW = CurrencyCode{Code: "UYW", Name: "Unidad Previsional"}
	UZS = CurrencyCode{Code: "UZS", Name: "Uzbekistan Sum"}
	VUV = CurrencyCode{Code: "VUV", Name: "Vatu"}
	VES = CurrencyCode{Code: "VES", Name: "Bolívar Soberano"}
	VND = CurrencyCode{Code: "VND", Name: "Dong"}
	YER = CurrencyCode{Code: "YER", Name: "Yemeni Rial"}
	ZMW = CurrencyCode{Code: "ZMW", Name: "Zambian Kwacha"}
	ZWL = CurrencyCode{Code: "ZWL", Name: "Zimbabwe Dollar"}
	XBA = CurrencyCode{Code: "XBA", Name: "Bond Markets Unit European Composite Unit (EURCO)"}
	XBB = CurrencyCode{Code: "XBB", Name: "Bond Markets Unit European Monetary Unit (E.M.U.-6)"}
	XBC = CurrencyCode{Code: "XBC", Name: "Bond Markets Unit European Unit of Account 9 (E.U.A.-9)"}
	XBD = CurrencyCode{Code: "XBD", Name: "Bond Markets Unit European Unit of Account 17 (E.U.A.-17)"}
	XTS = CurrencyCode{Code: "XTS", Name: "Codes specifically reserved for testing purposes"}
	XXX = CurrencyCode{Code: "XXX", Name: "The codes assigned for transactions where no currency is involved"}
	XAU = CurrencyCode{Code: "XAU", Name: "Gold"}
	XPD = CurrencyCode{Code: "XPD", Name: "Palladium"}
	XPT = CurrencyCode{Code: "XPT", Name: "Platinum"}
	XAG = CurrencyCode{Code: "XAG", Name: "Silver"}
	AFA = CurrencyCode{Code: "AFA", Name: "Afghani"}
	FIM = CurrencyCode{Code: "FIM", Name: "Markka"}
	ALK = CurrencyCode{Code: "ALK", Name: "Old Lek"}
	ADP = CurrencyCode{Code: "ADP", Name: "Andorran Peseta"}
	ESP = CurrencyCode{Code: "ESP", Name: "Spanish Peseta"}
	FRF = CurrencyCode{Code: "FRF", Name: "French Franc"}
	AOK = CurrencyCode{Code: "AOK", Name: "Kwanza"}
	AON = CurrencyCode{Code: "AON", Name: "New Kwanza"}
	AOR = CurrencyCode{Code: "AOR", Name: "Kwanza Reajustado"}
	ARA = CurrencyCode{Code: "ARA", Name: "Austral"}
	ARP = CurrencyCode{Code: "ARP", Name: "Peso Argentino"}
	ARY = CurrencyCode{Code: "ARY", Name: "Peso"}
	RUR = CurrencyCode{Code: "RUR", Name: "Russian Ruble"}
	ATS = CurrencyCode{Code: "ATS", Name: "Schilling"}
	AYM = CurrencyCode{Code: "AYM", Name: "Azerbaijan Manat"}
	AZM = CurrencyCode{Code: "AZM", Name: "Azerbaijanian Manat"}
	BYB = CurrencyCode{Code: "BYB", Name: "Belarusian Ruble"}
	BYR = CurrencyCode{Code: "BYR", Name: "Belarusian Ruble"}
	BEC = CurrencyCode{Code: "BEC", Name: "Convertible Franc"}
	BEF = CurrencyCode{Code: "BEF", Name: "Belgian Franc"}
	BEL = CurrencyCode{Code: "BEL", Name: "Financial Franc"}
	BOP = CurrencyCode{Code: "BOP", Name: "Peso boliviano"}
	BAD = CurrencyCode{Code: "BAD", Name: "Dinar"}
	BRB = CurrencyCode{Code: "BRB", Name: "Cruzeiro"}
	BRC = CurrencyCode{Code: "BRC", Name: "Cruzado"}
	BRE = CurrencyCode{Code: "BRE", Name: "Cruzeiro"}
	BRN = CurrencyCode{Code: "BRN", Name: "New Cruzado"}
	BRR = CurrencyCode{Code: "BRR", Name: "Cruzeiro Real"}
	BGJ = CurrencyCode{Code: "BGJ", Name: "Lev A/52"}
	BGK = CurrencyCode{Code: "BGK", Name: "Lev A/62"}
	BGL = CurrencyCode{Code: "BGL", Name: "Lev"}
	BUK = CurrencyCode{Code: "BUK", Name: "Kyat"}
	HRD = CurrencyCode{Code: "HRD", Name: "Croatian Dinar"}
	CYP = CurrencyCode{Code: "CYP", Name: "Cyprus Pound"}
	CSJ = CurrencyCode{Code: "CSJ", Name: "Krona A/53"}
	CSK = CurrencyCode{Code: "CSK", Name: "Koruna"}
	ECS = CurrencyCode{Code: "ECS", Name: "Sucre"}
	ECV = CurrencyCode{Code: "ECV", Name: "Unidad de Valor Constante (UVC)"}
	GQE = CurrencyCode{Code: "GQE", Name: "Ekwele"}
	EEK = CurrencyCode{Code: "EEK", Name: "Kroon"}
	XEU = CurrencyCode{Code: "XEU", Name: "European Currency Unit (E.C.U)"}
	GEK = CurrencyCode{Code: "GEK", Name: "Georgian Coupon"}
	DDM = CurrencyCode{Code: "DDM", Name: "Mark der DDR"}
	DEM = CurrencyCode{Code: "DEM", Name: "Deutsche Mark"}
	GHC = CurrencyCode{Code: "GHC", Name: "Cedi"}
	GHP = CurrencyCode{Code: "GHP", Name: "Ghana Cedi"}
	GRD = CurrencyCode{Code: "GRD", Name: "Drachma"}
	GNE = CurrencyCode{Code: "GNE", Name: "Syli"}
	GNS = CurrencyCode{Code: "GNS", Name: "Syli"}
	GWE = CurrencyCode{Code: "GWE", Name: "Guinea Escudo"}
	GWP = CurrencyCode{Code: "GWP", Name: "Guinea-Bissau Peso"}
	ITL = CurrencyCode{Code: "ITL", Name: "Italian Lira"}
	ISJ = CurrencyCode{Code: "ISJ", Name: "Old Krona"}
	IEP = CurrencyCode{Code: "IEP", Name: "Irish Pound"}
	ILP = CurrencyCode{Code: "ILP", Name: "Pound"}
	ILR = CurrencyCode{Code: "ILR", Name: "Old Shekel"}
	LAJ = CurrencyCode{Code: "LAJ", Name: "Pathet Lao Kip"}
	LVL = CurrencyCode{Code: "LVL", Name: "Latvian Lats"}
	LVR = CurrencyCode{Code: "LVR", Name: "Latvian Ruble"}
	LSM = CurrencyCode{Code: "LSM", Name: "Loti"}
	ZAL = CurrencyCode{Code: "ZAL", Name: "Financial Rand"}
	LTL = CurrencyCode{Code: "LTL", Name: "Lithuanian Litas"}
	LTT = CurrencyCode{Code: "LTT", Name: "Talonas"}
	LUC = CurrencyCode{Code: "LUC", Name: "Luxembourg Convertible Franc"}
	LUF = CurrencyCode{Code: "LUF", Name: "Luxembourg Franc"}
	LUL = CurrencyCode{Code: "LUL", Name: "Luxembourg Financial Franc"}
	MGF = CurrencyCode{Code: "MGF", Name: "Malagasy Franc"}
	MVQ = CurrencyCode{Code: "MVQ", Name: "Maldive Rupee"}
	MLF = CurrencyCode{Code: "MLF", Name: "Mali Franc"}
	MTL = CurrencyCode{Code: "MTL", Name: "Maltese Lira"}
	MTP = CurrencyCode{Code: "MTP", Name: "Maltese Pound"}
	MRO = CurrencyCode{Code: "MRO", Name: "Ouguiya"}
	MXP = CurrencyCode{Code: "MXP", Name: "Mexican Peso"}
	MZE = CurrencyCode{Code: "MZE", Name: "Mozambique Escudo"}
	MZM = CurrencyCode{Code: "MZM", Name: "Mozambique Metical"}
	NLG = CurrencyCode{Code: "NLG", Name: "Netherlands Guilder"}
	NIC = CurrencyCode{Code: "NIC", Name: "Cordoba"}
	PEH = CurrencyCode{Code: "PEH", Name: "Sol"}
	PEI = CurrencyCode{Code: "PEI", Name: "Inti"}
	PES = CurrencyCode{Code: "PES", Name: "Sol"}
	PLZ = CurrencyCode{Code: "PLZ", Name: "Zloty"}
	PTE = CurrencyCode{Code: "PTE", Name: "Portuguese Escudo"}
	ROK = CurrencyCode{Code: "ROK", Name: "Leu A/52"}
	ROL = CurrencyCode{Code: "ROL", Name: "Old Leu"}
	STD = CurrencyCode{Code: "STD", Name: "Dobra"}
	CSD = CurrencyCode{Code: "CSD", Name: "Serbian Dinar"}
	SKK = CurrencyCode{Code: "SKK", Name: "Slovak Koruna"}
	SIT = CurrencyCode{Code: "SIT", Name: "Tolar"}
	RHD = CurrencyCode{Code: "RHD", Name: "Rhodesian Dollar"}
	ESA = CurrencyCode{Code: "ESA", Name: "Spanish Peseta"}
	ESB = CurrencyCode{Code: "ESB", Name: "A Account (convertible Peseta Account)"}
	SDD = CurrencyCode{Code: "SDD", Name: "Sudanese Dinar"}
	SDP = CurrencyCode{Code: "SDP", Name: "Sudanese Pound"}
	SRG = CurrencyCode{Code: "SRG", Name: "Surinam Guilder"}
	CHC = CurrencyCode{Code: "CHC", Name: "WIR Franc (for electronic)"}
	TJR = CurrencyCode{Code: "TJR", Name: "Tajik Ruble"}
	TPE = CurrencyCode{Code: "TPE", Name: "Timor Escudo"}
	TRL = CurrencyCode{Code: "TRL", Name: "Old Turkish Lira"}
	TMM = CurrencyCode{Code: "TMM", Name: "Turkmenistan Manat"}
	UGS = CurrencyCode{Code: "UGS", Name: "Uganda Shilling"}
	UGW = CurrencyCode{Code: "UGW", Name: "Old Shilling"}
	UAK = CurrencyCode{Code: "UAK", Name: "Karbovanet"}
	SUR = CurrencyCode{Code: "SUR", Name: "Rouble"}
	USS = CurrencyCode{Code: "USS", Name: "US Dollar (Same day)"}
	UYN = CurrencyCode{Code: "UYN", Name: "Old Uruguay Peso"}
	UYP = CurrencyCode{Code: "UYP", Name: "Uruguayan Peso"}
	VEB = CurrencyCode{Code: "VEB", Name: "Bolivar"}
	VEF = CurrencyCode{Code: "VEF", Name: "Bolivar Fuerte"}
	VNC = CurrencyCode{Code: "VNC", Name: "Old Dong"}
	YDD = CurrencyCode{Code: "YDD", Name: "Yemeni Dinar"}
	YUD = CurrencyCode{Code: "YUD", Name: "New Yugoslavian Dinar"}
	YUM = CurrencyCode{Code: "YUM", Name: "New Dinar"}
	YUN = CurrencyCode{Code: "YUN", Name: "Yugoslavian Dinar"}
	ZRN = CurrencyCode{Code: "ZRN", Name: "New Zaire"}
	ZRZ = CurrencyCode{Code: "ZRZ", Name: "Zaire"}
	ZMK = CurrencyCode{Code: "ZMK", Name: "Zambian Kwacha"}
	ZWC = CurrencyCode{Code: "ZWC", Name: "Rhodesian Dollar"}
	ZWD = CurrencyCode{Code: "ZWD", Name: "Zimbabwe Dollar (old)"}
	ZWN = CurrencyCode{Code: "ZWN", Name: "Zimbabwe Dollar (new)"}
	ZWR = CurrencyCode{Code: "ZWR", Name: "Zimbabwe Dollar"}
	XFO = CurrencyCode{Code: "XFO", Name: "Gold-Franc"}
	XRE = CurrencyCode{Code: "XRE", Name: "RINET Funds Code"}
	XFU = CurrencyCode{Code: "XFU", Name: "UIC-Franc"}
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
	"AFA": AFA, // Afghani
	"FIM": FIM, // Markka
	"ALK": ALK, // Old Lek
	"ADP": ADP, // Andorran Peseta
	"ESP": ESP, // Spanish Peseta
	"FRF": FRF, // French Franc
	"AOK": AOK, // Kwanza
	"AON": AON, // New Kwanza
	"AOR": AOR, // Kwanza Reajustado
	"ARA": ARA, // Austral
	"ARP": ARP, // Peso Argentino
	"ARY": ARY, // Peso
	"RUR": RUR, // Russian Ruble
	"ATS": ATS, // Schilling
	"AYM": AYM, // Azerbaijan Manat
	"AZM": AZM, // Azerbaijanian Manat
	"BYB": BYB, // Belarusian Ruble
	"BYR": BYR, // Belarusian Ruble
	"BEC": BEC, // Convertible Franc
	"BEF": BEF, // Belgian Franc
	"BEL": BEL, // Financial Franc
	"BOP": BOP, // Peso boliviano
	"BAD": BAD, // Dinar
	"BRB": BRB, // Cruzeiro
	"BRC": BRC, // Cruzado
	"BRE": BRE, // Cruzeiro
	"BRN": BRN, // New Cruzado
	"BRR": BRR, // Cruzeiro Real
	"BGJ": BGJ, // Lev A/52
	"BGK": BGK, // Lev A/62
	"BGL": BGL, // Lev
	"BUK": BUK, // Kyat
	"HRD": HRD, // Croatian Dinar
	"CYP": CYP, // Cyprus Pound
	"CSJ": CSJ, // Krona A/53
	"CSK": CSK, // Koruna
	"ECS": ECS, // Sucre
	"ECV": ECV, // Unidad de Valor Constante (UVC)
	"GQE": GQE, // Ekwele
	"EEK": EEK, // Kroon
	"XEU": XEU, // European Currency Unit (E.C.U)
	"GEK": GEK, // Georgian Coupon
	"DDM": DDM, // Mark der DDR
	"DEM": DEM, // Deutsche Mark
	"GHC": GHC, // Cedi
	"GHP": GHP, // Ghana Cedi
	"GRD": GRD, // Drachma
	"GNE": GNE, // Syli
	"GNS": GNS, // Syli
	"GWE": GWE, // Guinea Escudo
	"GWP": GWP, // Guinea-Bissau Peso
	"ITL": ITL, // Italian Lira
	"ISJ": ISJ, // Old Krona
	"IEP": IEP, // Irish Pound
	"ILP": ILP, // Pound
	"ILR": ILR, // Old Shekel
	"LAJ": LAJ, // Pathet Lao Kip
	"LVL": LVL, // Latvian Lats
	"LVR": LVR, // Latvian Ruble
	"LSM": LSM, // Loti
	"ZAL": ZAL, // Financial Rand
	"LTL": LTL, // Lithuanian Litas
	"LTT": LTT, // Talonas
	"LUC": LUC, // Luxembourg Convertible Franc
	"LUF": LUF, // Luxembourg Franc
	"LUL": LUL, // Luxembourg Financial Franc
	"MGF": MGF, // Malagasy Franc
	"MVQ": MVQ, // Maldive Rupee
	"MLF": MLF, // Mali Franc
	"MTL": MTL, // Maltese Lira
	"MTP": MTP, // Maltese Pound
	"MRO": MRO, // Ouguiya
	"MXP": MXP, // Mexican Peso
	"MZE": MZE, // Mozambique Escudo
	"MZM": MZM, // Mozambique Metical
	"NLG": NLG, // Netherlands Guilder
	"NIC": NIC, // Cordoba
	"PEH": PEH, // Sol
	"PEI": PEI, // Inti
	"PES": PES, // Sol
	"PLZ": PLZ, // Zloty
	"PTE": PTE, // Portuguese Escudo
	"ROK": ROK, // Leu A/52
	"ROL": ROL, // Old Leu
	"STD": STD, // Dobra
	"CSD": CSD, // Serbian Dinar
	"SKK": SKK, // Slovak Koruna
	"SIT": SIT, // Tolar
	"RHD": RHD, // Rhodesian Dollar
	"ESA": ESA, // Spanish Peseta
	"ESB": ESB, // A Account (convertible Peseta Account)
	"SDD": SDD, // Sudanese Dinar
	"SDP": SDP, // Sudanese Pound
	"SRG": SRG, // Surinam Guilder
	"CHC": CHC, // WIR Franc (for electronic)
	"TJR": TJR, // Tajik Ruble
	"TPE": TPE, // Timor Escudo
	"TRL": TRL, // Old Turkish Lira
	"TMM": TMM, // Turkmenistan Manat
	"UGS": UGS, // Uganda Shilling
	"UGW": UGW, // Old Shilling
	"UAK": UAK, // Karbovanet
	"SUR": SUR, // Rouble
	"USS": USS, // US Dollar (Same day)
	"UYN": UYN, // Old Uruguay Peso
	"UYP": UYP, // Uruguayan Peso
	"VEB": VEB, // Bolivar
	"VEF": VEF, // Bolivar Fuerte
	"VNC": VNC, // Old Dong
	"YDD": YDD, // Yemeni Dinar
	"YUD": YUD, // New Yugoslavian Dinar
	"YUM": YUM, // New Dinar
	"YUN": YUN, // Yugoslavian Dinar
	"ZRN": ZRN, // New Zaire
	"ZRZ": ZRZ, // Zaire
	"ZMK": ZMK, // Zambian Kwacha
	"ZWC": ZWC, // Rhodesian Dollar
	"ZWD": ZWD, // Zimbabwe Dollar (old)
	"ZWN": ZWN, // Zimbabwe Dollar (new)
	"ZWR": ZWR, // Zimbabwe Dollar
	"XFO": XFO, // Gold-Franc
	"XRE": XRE, // RINET Funds Code
	"XFU": XFU, // UIC-Franc
}
