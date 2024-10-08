package pcib

import "strings"

type Scheme string

const (
	SchemeVisa       Scheme = "visa"
	SchemeMastercard Scheme = "mastercard"
	SchemeAmex       Scheme = "amex"
	SchemeDiscover   Scheme = "discover"
	SchemeDiners     Scheme = "diners"
	SchemeJCB        Scheme = "jcb"
	SchemeUnionPay   Scheme = "unionpay"
	SchemeUnknown    Scheme = "unknown"
)

func GetScheme(in string) Scheme {
	switch strings.ToLower(in) {
	case string(SchemeVisa):
		return SchemeVisa
	case string(SchemeMastercard), "master_card", "mc":
		return SchemeMastercard
	case string(SchemeAmex), "american_express", "americanexpress":
		return SchemeAmex
	case string(SchemeDiscover):
		return SchemeDiscover
	case string(SchemeDiners), "diners_club", "dinersclub":
		return SchemeDiners
	case string(SchemeJCB), "japanese_credit_bank":
		return SchemeJCB
	case string(SchemeUnionPay), "union_pay", "china_union_pay":
		return SchemeUnionPay
	}
	return SchemeUnknown
}

type Funding string

const (
	FundingCredit        Funding = "credit"
	FundingDebit         Funding = "debit"
	FundingPrepaid       Funding = "prepaid"
	FundingCharge        Funding = "charge"
	FundingDeferredDebit Funding = "deferred-debit"
	FundingUnknown       Funding = "unknown"
)

type CardProduct string

var VisaCardProduct = map[CardProduct]string{
	"A":  "Visa Traditional",
	"B":  "Visa Traditional Rewards",
	"C":  "Visa Signature",
	"D":  "Visa Signature Preferred",
	"F":  "Visa Classic",
	"G":  "Visa Business",
	"G1": "Visa Signature Business",
	"G3": "Visa Business Enhanced",
	"G4": "Visa Infinite Business",
	"G5": "Visa Business Rewards",
	"I":  "Visa Infinite",
	"I1": "Visa Infinite Privilege",
	"I2": "Visa UHNW",
	"J3": "Visa Prepaid Healthcare",
	"K":  "Visa Corporate",
	"K1": "Visa Government Corporate T&E",
	"L":  "Electron",
	"N":  "Visa Platinum",
	"N1": "Visa Rewards",
	"N2": "Visa Select",
	"O":  "Reserved",
	"P":  "Visa Gold",
	"Q":  "Private Label",
	"Q2": "Private Label Basic",
	"Q3": "Private Label Enhanced",
	"Q4": "Private Label Standard",
	"Q5": "Private Label Specialized",
	"Q6": "Private Label Premium",
	"R":  "Proprietary",
	"S":  "Visa Purchasing",
	"S1": "Visa Purchasing",
	"S2": "Visa Government Purchasing",
	"S3": "Visa Government Purchasing with Fleet",
	"S4": "Visa Commercial Agriculture",
	"S5": "Visa Commercial Transport",
	"S6": "Visa Business Loan",
	"Sl": "Visa Purchasing with Fleet",
	"U":  "Visa TravelMoney",
	"V":  "Visa V Pay",
	"X":  "Visa B2B",
}

var MastercardCardProduct = map[CardProduct]string{
	"ACS":     "ACH for Consumer",
	"BPC":     "Bill Pay Commercial",
	"BPD":     "Mastercard World Business Debit",
	"CIR":     "Cirrus Card",
	"DBK":     "DIGITAL MASTERCARD BLACK",
	"DCG":     "DIGITAL GOLD",
	"DCS":     "DIGITAL STANDARD",
	"DPL":     "DIGITAL PLATINUM",
	"MAB":     "World Elite Mastercard for Business",
	"MAC":     "Mastercard Corporate World Elite Card",
	"Maestro": "Branded Prepaid business programs issued in Europe",
	"MAP":     "Mastercard Commercial Payment Account",
	"MAQ":     "Mastercard Prepaid Commercial Payment Account",
	"MBA":     "MasterCard B2B Product 2",
	"MBB":     "Mastercard Prepaid Consumer",
	"MBC":     "MasterCard Prepaid Voucher",
	"MBD":     "Debit MasterCard Professional Card",
	"MBE":     "MasterCard Electronic BusinessCard Card",
	"MBF":     "MasterCard Alimentacao (Food)",
	"MBG":     "MasterCard B2B Product 3",
	"MBH":     "MasterCard B2B Product 4",
	"MBI":     "MasterCard B2B Product 5",
	"MBJ":     "MasterCard B2B Product 6",
	"MBK":     "MasterCard Black",
	"MBM":     "MasterCard Refeicao (Meal)",
	"MBP":     "MasterCard Corporate Prepaid",
	"MBS":     "MasterCard B2B Product 1",
	"MBW":     "World MasterCard Black Edition",
	"MCB":     "MasterCard BusinessCard Card",
	"MCC":     "Mixed Product",
	"MCE":     "MasterCard Electronic",
	"MCF":     "MasterCard Corporate Fleet Card",
	"MCG":     "Gold MasterCard",
	"MCM":     "MasterCard Corporate Meeting Card",
	"MCO":     "MasterCard Corporate Card",
	"MCP":     "Mastercard Corporate Purchasing Card",
	"MCS":     "Mastercard Standard",
	"MCT":     "Titanium MasterCard",
	"MCU":     "MasterCard Unembossed Card",
	"MCW":     "World MasterCard Card",
	"MDB":     "Debit MasterCard BusinessCard Card",
	"MDG":     "Debit Gold MasterCard",
	"MDH":     "World Debit MasterCard Embossed",
	"MDJ":     "Debit MasterCard (enhanced)",
	"MDO":     "Debit MasterCard Other Programs",
	"MDP":     "Debit Platinum MasterCard",
	"MDS":     "Debit MasterCard",
	"MDT":     "Business Debit MasterCard",
	"MDU":     "Debit MasterCard Unembossed",
	"MDW":     "World Elite Debit MasterCard",
	"MEB":     "MasterCard Executive BusinessCard Card",
	"MEO":     "MasterCard Corporate Executive Card",
	"MEP":     "Premium Debit MasterCard",
	"MES":     "Mastercard Enterprise Solutions",
	"MET":     "Titanium Debit MasterCard",
	"MFR":     "MasterCard Commercial Reward Funding",
	"MGF":     "MasterCard Government Commercial Card",
	"MGP":     "Mastercard Prepaid Gold Payroll",
	"MGS":     "Platinum Mastercard Prepaid General Spend",
	"MHA":     "MasterCard Healthcare Prepaid Non",
	"MHB":     "MasterCard HSA Substantiated",
	"MHH":     "MasterCard HSA Non",
	"MHP":     "HELOC Platinum MasterCard",
	"MHS":     "HELOC Standard MasterCard",
	"MIC":     "ISIC MasterCard Student Card",
	"MIP":     "ISIC MasterCard Prepaid Student Card",
	"MIS":     "ISIC Debit MasterCard Student Card",
	"MIU":     "Debit MasterCard Unembossed",
	"MLA":     "MasterCard Central Travel Solutions Air",
	"MLB":     "MasterCard Brazil Prepaid Benefits",
	"MLC":     "MasterCard Micro",
	"MLD":     "MasterCard Distribution Card",
	"MLE":     "MasterCard Pedagio Prepaid Card",
	"MLF":     "MasterCard Agro",
	"MLL":     "MasterCard Central Travel Solutions Land",
	"MNF":     "MasterCard Public Sector Commercial Card",
	"MNW":     "World MasterCard Card",
	"MOW":     "World Maestro",
	"MPA":     "Prepaid MasterCard Payroll Card",
	"MPB":     "MasterCard Preferred BusinessCard Card",
	"MPC":     "MasterCard Professional Card",
	"MPD":     "Mastercard Flex Prepaid",
	"MPF":     "Prepaid MasterCard Gift Card",
	"MPG":     "Prepaid MasterCard General Spend Card",
	"MPJ":     "Prepaid MC Debit Voucher Meal/Food Card",
	"MPL":     "Platinum MasterCard",
	"MPM":     "Prepaid MasterCard Consumer Incentive Card",
	"MPN":     "Prepaid MasterCard Insurance Card",
	"MPO":     "Prepaid MasterCard Other Card",
	"MPP":     "MasterCard Prepaid Embossed",
	"MPR":     "Prepaid MasterCard Travel Card",
	"MPV":     "Prepaid MasterCard Government Card",
	"MPW":     "Prepaid MasterCard Workplace B2B Solutions",
	"MPX":     "Prepaid MasterCard Flex Benefit Card",
	"MPY":     "Prepaid MasterCard Employee Incentive Card",
	"MRC":     "Prepaid MasterCard Electronic Card",
	"MRD":     "Platinum Debit Mastercard Prepaid General Spend",
	"MRF":     "Standard Deferred",
	"MRG":     "Prepaid MasterCard Card",
	"MRH":     "MasterCard Prepaid Platinum Travel Card",
	"MRJ":     "Prepaid MasterCard Voucher Meal/Food Card",
	"MRK":     "Prepaid MC Public Sector Commercial Card",
	"MRL":     "Mastercard Prepaid Business Preferred",
	"MRO":     "MasterCard Rewards Only",
	"MRW":     "Prepaid MasterCard BusinessCard Card",
	"MSB":     "Maestro Small Business",
	"MSI":     "Maestro Card",
	"MTP":     "MasterCard Platinum Prepaid Travel",
	"MUP":     "Platinum Debit MasterCard Unembossed",
	"MUS":     "Prepaid MasterCard Unembossed",
	"MVA":     "Mastercard B2B VIP 1",
	"MVB":     "Mastercard B2B VIP 2",
	"MVC":     "Mastercard B2B VIP 3",
	"MVD":     "Mastercard B2B VIP 4",
	"MVE":     "Mastercard B2B VIP 5",
	"MVF":     "Mastercard B2B VIP 6",
	"MVG":     "Mastercard B2B VIP 7",
	"MVH":     "Mastercard B2B VIP 8",
	"MVJ":     "Mastercard B2B VIP 10",
	"MVK":     "Mastercard B2B VIP 11",
	"MVL":     "Mastercard B2B VIP 12",
	"MVM":     "Mastercard B2B VIP 13",
	"MVN":     "Mastercard B2B VIP 14",
	"MWB":     "World MasterCard for Business",
	"MWE":     "World Elite MasterCard Card",
	"MWF":     "Mastercard Humanitarian Prepaid",
	"MWO":     "MasterCard Corporate World Card",
	"MWP":     "Mastercard World Prepaid",
	"MXG":     "Digital Enablement Program",
	"OLB":     "Maestro Small Business",
	"OLC":     "Prepaid MasterCard Commercial Card",
	"OLR":     "Prepaid MasterCard Consumer Card",
	"OLS":     "Maestro",
	"PVL":     "Private Label",
	"PVT":     "Private Label Trade Program",
	"SAP":     "MasterCard Salary Platinum",
	"SAS":     "MasterCard Salary Standard",
	"SUR":     "Prepaid MasterCard Unembossed",
	"TCB":     "Mastercard BusinessCard Card",
	"TCC":     "MasterCard Mixed Product",
	"TCG":     "Gold MasterCard",
	"TCO":     "Mastercard Corporate Card",
	"TCS":     "MasterCard Standard",
	"TCW":     "World Elite Mastercard",
	"TIC":     "ISIC MasterCard Student Card",
	"TIU":     "MasterCard Unembossed",
	"TNW":     "World MasterCard",
	"TPC":     "MasterCard Professional Card",
	"TPL":     "Platinum MasterCard",
	"TWB":     "World MasterCard Black Edition",
	"WBE":     "World MasterCard Black Edition",
	"WDR":     "World Debit MasterCard Rewards",
	"WMR":     "World MasterCard Rewards",
	"WPD":     "World Prepaid Debit",
}
