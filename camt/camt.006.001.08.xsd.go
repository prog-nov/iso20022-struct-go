// Code generated by xgen. DO NOT EDIT.

package schema

// Document ...
type Document *Document

// AccountIdentification4Choice ...
type AccountIdentification4Choice struct {
	IBAN string                         `xml:"IBAN"`
	Othr *GenericAccountIdentification1 `xml:"Othr"`
}

// AccountSchemeName1Choice ...
type AccountSchemeName1Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// ActiveCurrencyAndAmountSimpleType ...
type ActiveCurrencyAndAmountSimpleType float64

// ActiveCurrencyAndAmount ...
type ActiveCurrencyAndAmount struct {
	CcyAttr string  `xml:"Ccy,attr"`
	Value   float64 `xml:",chardata"`
}

// ActiveCurrencyCode ...
type ActiveCurrencyCode string

// ActiveOrHistoricCurrencyAndAmountSimpleType ...
type ActiveOrHistoricCurrencyAndAmountSimpleType float64

// ActiveOrHistoricCurrencyAndAmount ...
type ActiveOrHistoricCurrencyAndAmount struct {
	CcyAttr string  `xml:"Ccy,attr"`
	Value   float64 `xml:",chardata"`
}

// ActiveOrHistoricCurrencyCode ...
type ActiveOrHistoricCurrencyCode string

// AddressType2Code ...
type AddressType2Code string

// AddressType3Choice ...
type AddressType3Choice struct {
	Cd    string                   `xml:"Cd"`
	Prtry *GenericIdentification30 `xml:"Prtry"`
}

// Amount2Choice ...
type Amount2Choice struct {
	AmtWthtCcy float64                  `xml:"AmtWthtCcy"`
	AmtWthCcy  *ActiveCurrencyAndAmount `xml:"AmtWthCcy"`
}

// Amount3Choice ...
type Amount3Choice struct {
	AmtWthCcy  *ActiveOrHistoricCurrencyAndAmount `xml:"AmtWthCcy"`
	AmtWthtCcy float64                            `xml:"AmtWthtCcy"`
}

// AnyBICDec2014Identifier ...
type AnyBICDec2014Identifier string

// BICFIDec2014Identifier ...
type BICFIDec2014Identifier string

// BranchAndFinancialInstitutionIdentification6 ...
type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId *FinancialInstitutionIdentification18 `xml:"FinInstnId"`
	BrnchId    *BranchData3                          `xml:"BrnchId"`
}

// BranchData3 ...
type BranchData3 struct {
	Id      string           `xml:"Id"`
	LEI     string           `xml:"LEI"`
	Nm      string           `xml:"Nm"`
	PstlAdr *PostalAddress24 `xml:"PstlAdr"`
}

// CancelledStatusReason1Code ...
type CancelledStatusReason1Code string

// CashAccount39 ...
type CashAccount39 struct {
	Id   *AccountIdentification4Choice                 `xml:"Id"`
	Tp   *CashAccountType2Choice                       `xml:"Tp"`
	Ccy  string                                        `xml:"Ccy"`
	Nm   string                                        `xml:"Nm"`
	Prxy *ProxyAccountIdentification1                  `xml:"Prxy"`
	Ownr *PartyIdentification135                       `xml:"Ownr"`
	Svcr *BranchAndFinancialInstitutionIdentification6 `xml:"Svcr"`
}

// CashAccountAndEntry3 ...
type CashAccountAndEntry3 struct {
	Acct *CashAccount39 `xml:"Acct"`
	Ntry *CashEntry2    `xml:"Ntry"`
}

// CashAccountType2Choice ...
type CashAccountType2Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// CashEntry2 ...
type CashEntry2 struct {
	Amt          *ActiveCurrencyAndAmount `xml:"Amt"`
	Dt           *DateAndDateTime2Choice  `xml:"Dt"`
	Sts          string                   `xml:"Sts"`
	Id           string                   `xml:"Id"`
	StmtId       string                   `xml:"StmtId"`
	AcctSvcrRef  float64                  `xml:"AcctSvcrRef"`
	AddtlNtryInf []string                 `xml:"AddtlNtryInf"`
}

// ClearingSystemIdentification2Choice ...
type ClearingSystemIdentification2Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// ClearingSystemMemberIdentification2 ...
type ClearingSystemMemberIdentification2 struct {
	ClrSysId *ClearingSystemIdentification2Choice `xml:"ClrSysId"`
	MmbId    string                               `xml:"MmbId"`
}

// Contact4 ...
type Contact4 struct {
	NmPrfx    string           `xml:"NmPrfx"`
	Nm        string           `xml:"Nm"`
	PhneNb    string           `xml:"PhneNb"`
	MobNb     string           `xml:"MobNb"`
	FaxNb     string           `xml:"FaxNb"`
	EmailAdr  string           `xml:"EmailAdr"`
	EmailPurp string           `xml:"EmailPurp"`
	JobTitl   string           `xml:"JobTitl"`
	Rspnsblty string           `xml:"Rspnsblty"`
	Dept      string           `xml:"Dept"`
	Othr      []*OtherContact1 `xml:"Othr"`
	PrefrdMtd string           `xml:"PrefrdMtd"`
}

// CountryCode ...
type CountryCode string

// CreditDebitCode ...
type CreditDebitCode string

// DateAndDateTime2Choice ...
type DateAndDateTime2Choice struct {
	Dt   string `xml:"Dt"`
	DtTm string `xml:"DtTm"`
}

// DateAndPlaceOfBirth1 ...
type DateAndPlaceOfBirth1 struct {
	BirthDt     string `xml:"BirthDt"`
	PrvcOfBirth string `xml:"PrvcOfBirth"`
	CityOfBirth string `xml:"CityOfBirth"`
	CtryOfBirth string `xml:"CtryOfBirth"`
}

// DateTimePeriod1 ...
type DateTimePeriod1 struct {
	FrDtTm string `xml:"FrDtTm"`
	ToDtTm string `xml:"ToDtTm"`
}

// DateTimePeriod1Choice ...
type DateTimePeriod1Choice struct {
	FrDtTm string           `xml:"FrDtTm"`
	ToDtTm string           `xml:"ToDtTm"`
	DtTmRg *DateTimePeriod1 `xml:"DtTmRg"`
}

// DecimalNumber ...
type DecimalNumber float64

// EntryStatus1Code ...
type EntryStatus1Code string

// EntryTypeIdentifier ...
type EntryTypeIdentifier string

// ErrorHandling3Choice ...
type ErrorHandling3Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// ErrorHandling5 ...
type ErrorHandling5 struct {
	Err  *ErrorHandling3Choice `xml:"Err"`
	Desc string                `xml:"Desc"`
}

// Exact4AlphaNumericText ...
type Exact4AlphaNumericText string

// ExternalAccountIdentification1Code ...
type ExternalAccountIdentification1Code string

// ExternalCashAccountType1Code ...
type ExternalCashAccountType1Code string

// ExternalClearingSystemIdentification1Code ...
type ExternalClearingSystemIdentification1Code string

// ExternalEnquiryRequestType1Code ...
type ExternalEnquiryRequestType1Code string

// ExternalFinancialInstitutionIdentification1Code ...
type ExternalFinancialInstitutionIdentification1Code string

// ExternalMarketInfrastructure1Code ...
type ExternalMarketInfrastructure1Code string

// ExternalOrganisationIdentification1Code ...
type ExternalOrganisationIdentification1Code string

// ExternalPaymentControlRequestType1Code ...
type ExternalPaymentControlRequestType1Code string

// ExternalPersonIdentification1Code ...
type ExternalPersonIdentification1Code string

// ExternalProxyAccountType1Code ...
type ExternalProxyAccountType1Code string

// ExternalSystemErrorHandling1Code ...
type ExternalSystemErrorHandling1Code string

// FinalStatus1Code ...
type FinalStatus1Code string

// FinancialIdentificationSchemeName1Choice ...
type FinancialIdentificationSchemeName1Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// FinancialInstitutionIdentification18 ...
type FinancialInstitutionIdentification18 struct {
	BICFI       string                               `xml:"BICFI"`
	ClrSysMmbId *ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId"`
	LEI         string                               `xml:"LEI"`
	Nm          string                               `xml:"Nm"`
	PstlAdr     *PostalAddress24                     `xml:"PstlAdr"`
	Othr        *GenericFinancialIdentification1     `xml:"Othr"`
}

// GenericAccountIdentification1 ...
type GenericAccountIdentification1 struct {
	Id      string                    `xml:"Id"`
	SchmeNm *AccountSchemeName1Choice `xml:"SchmeNm"`
	Issr    string                    `xml:"Issr"`
}

// GenericFinancialIdentification1 ...
type GenericFinancialIdentification1 struct {
	Id      string                                    `xml:"Id"`
	SchmeNm *FinancialIdentificationSchemeName1Choice `xml:"SchmeNm"`
	Issr    string                                    `xml:"Issr"`
}

// GenericIdentification1 ...
type GenericIdentification1 struct {
	Id      string `xml:"Id"`
	SchmeNm string `xml:"SchmeNm"`
	Issr    string `xml:"Issr"`
}

// GenericIdentification30 ...
type GenericIdentification30 struct {
	Id      string `xml:"Id"`
	Issr    string `xml:"Issr"`
	SchmeNm string `xml:"SchmeNm"`
}

// GenericOrganisationIdentification1 ...
type GenericOrganisationIdentification1 struct {
	Id      string                                       `xml:"Id"`
	SchmeNm *OrganisationIdentificationSchemeName1Choice `xml:"SchmeNm"`
	Issr    string                                       `xml:"Issr"`
}

// GenericPersonIdentification1 ...
type GenericPersonIdentification1 struct {
	Id      string                                 `xml:"Id"`
	SchmeNm *PersonIdentificationSchemeName1Choice `xml:"SchmeNm"`
	Issr    string                                 `xml:"Issr"`
}

// IBAN2007Identifier ...
type IBAN2007Identifier string

// ISODate ...
type ISODate string

// ISODateTime ...
type ISODateTime string

// ImpliedCurrencyAndAmount ...
type ImpliedCurrencyAndAmount float64

// LEIIdentifier ...
type LEIIdentifier string

// LongPaymentIdentification2 ...
type LongPaymentIdentification2 struct {
	TxId           string                                        `xml:"TxId"`
	UETR           string                                        `xml:"UETR"`
	IntrBkSttlmAmt float64                                       `xml:"IntrBkSttlmAmt"`
	IntrBkSttlmDt  string                                        `xml:"IntrBkSttlmDt"`
	PmtMtd         *PaymentOrigin1Choice                         `xml:"PmtMtd"`
	InstgAgt       *BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt"`
	InstdAgt       *BranchAndFinancialInstitutionIdentification6 `xml:"InstdAgt"`
	NtryTp         string                                        `xml:"NtryTp"`
	EndToEndId     string                                        `xml:"EndToEndId"`
}

// MarketInfrastructureIdentification1Choice ...
type MarketInfrastructureIdentification1Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// Max10Text ...
type Max10Text string

// Max128Text ...
type Max128Text string

// Max140Text ...
type Max140Text string

// Max15NumericText ...
type Max15NumericText string

// Max16Text ...
type Max16Text string

// Max20000Text ...
type Max20000Text string

// Max2048Text ...
type Max2048Text string

// Max256Text ...
type Max256Text string

// Max34Text ...
type Max34Text string

// Max350Text ...
type Max350Text string

// Max35Text ...
type Max35Text string

// Max3NumericText ...
type Max3NumericText string

// Max4AlphaNumericText ...
type Max4AlphaNumericText string

// Max4Text ...
type Max4Text string

// Max5NumericText ...
type Max5NumericText string

// Max70Text ...
type Max70Text string

// MessageHeader8 ...
type MessageHeader8 struct {
	MsgId       string                  `xml:"MsgId"`
	CreDtTm     string                  `xml:"CreDtTm"`
	MsgPgntn    *Pagination1            `xml:"MsgPgntn"`
	OrgnlBizQry *OriginalBusinessQuery1 `xml:"OrgnlBizQry"`
	ReqTp       *RequestType4Choice     `xml:"ReqTp"`
	QryNm       string                  `xml:"QryNm"`
}

// NamePrefix2Code ...
type NamePrefix2Code string

// Number ...
type Number float64

// NumberAndSumOfTransactions2 ...
type NumberAndSumOfTransactions2 struct {
	NbOfNtries    string  `xml:"NbOfNtries"`
	Sum           float64 `xml:"Sum"`
	TtlNetNtryAmt float64 `xml:"TtlNetNtryAmt"`
	CdtDbtInd     string  `xml:"CdtDbtInd"`
}

// OrganisationIdentification29 ...
type OrganisationIdentification29 struct {
	AnyBIC string                                `xml:"AnyBIC"`
	LEI    string                                `xml:"LEI"`
	Othr   []*GenericOrganisationIdentification1 `xml:"Othr"`
}

// OrganisationIdentificationSchemeName1Choice ...
type OrganisationIdentificationSchemeName1Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// OriginalBusinessQuery1 ...
type OriginalBusinessQuery1 struct {
	MsgId   string `xml:"MsgId"`
	MsgNmId string `xml:"MsgNmId"`
	CreDtTm string `xml:"CreDtTm"`
}

// OtherContact1 ...
type OtherContact1 struct {
	ChanlTp string `xml:"ChanlTp"`
	Id      string `xml:"Id"`
}

// Pagination1 ...
type Pagination1 struct {
	PgNb      string `xml:"PgNb"`
	LastPgInd bool   `xml:"LastPgInd"`
}

// Party38Choice ...
type Party38Choice struct {
	OrgId  *OrganisationIdentification29 `xml:"OrgId"`
	PrvtId *PersonIdentification13       `xml:"PrvtId"`
}

// Party40Choice ...
type Party40Choice struct {
	Pty *PartyIdentification135                       `xml:"Pty"`
	Agt *BranchAndFinancialInstitutionIdentification6 `xml:"Agt"`
}

// PartyIdentification135 ...
type PartyIdentification135 struct {
	Nm        string           `xml:"Nm"`
	PstlAdr   *PostalAddress24 `xml:"PstlAdr"`
	Id        *Party38Choice   `xml:"Id"`
	CtryOfRes string           `xml:"CtryOfRes"`
	CtctDtls  *Contact4        `xml:"CtctDtls"`
}

// PaymentCommon4 ...
type PaymentCommon4 struct {
	PmtFr       *System2                `xml:"PmtFr"`
	PmtTo       *System2                `xml:"PmtTo"`
	CmonSts     []*PaymentStatus6       `xml:"CmonSts"`
	ReqdExctnDt *DateAndDateTime2Choice `xml:"ReqdExctnDt"`
	NtryDt      *DateAndDateTime2Choice `xml:"NtryDt"`
	CdtDbtInd   string                  `xml:"CdtDbtInd"`
	PmtMtd      *PaymentOrigin1Choice   `xml:"PmtMtd"`
}

// PaymentIdentification6Choice ...
type PaymentIdentification6Choice struct {
	TxId      string                           `xml:"TxId"`
	QId       *QueueTransactionIdentification1 `xml:"QId"`
	LngBizId  *LongPaymentIdentification2      `xml:"LngBizId"`
	ShrtBizId *ShortPaymentIdentification2     `xml:"ShrtBizId"`
	PrtryId   string                           `xml:"PrtryId"`
}

// PaymentInstruction32 ...
type PaymentInstruction32 struct {
	MsgId          string                    `xml:"MsgId"`
	ReqdExctnDt    *DateAndDateTime2Choice   `xml:"ReqdExctnDt"`
	Sts            []*PaymentStatus6         `xml:"Sts"`
	InstdAmt       *Amount3Choice            `xml:"InstdAmt"`
	IntrBkSttlmAmt *Amount2Choice            `xml:"IntrBkSttlmAmt"`
	Purp           string                    `xml:"Purp"`
	PmtMtd         *PaymentOrigin1Choice     `xml:"PmtMtd"`
	Prty           *Priority1Choice          `xml:"Prty"`
	PrcgVldtyTm    *DateTimePeriod1Choice    `xml:"PrcgVldtyTm"`
	InstrCpy       string                    `xml:"InstrCpy"`
	Tp             *PaymentType4Choice       `xml:"Tp"`
	GnrtdOrdr      bool                      `xml:"GnrtdOrdr"`
	TxId           string                    `xml:"TxId"`
	IntrBkSttlmDt  string                    `xml:"IntrBkSttlmDt"`
	EndToEndId     string                    `xml:"EndToEndId"`
	Pties          *PaymentTransactionParty3 `xml:"Pties"`
}

// PaymentInstrument1Code ...
type PaymentInstrument1Code string

// PaymentOrigin1Choice ...
type PaymentOrigin1Choice struct {
	FINMT    string `xml:"FINMT"`
	XMLMsgNm string `xml:"XMLMsgNm"`
	Prtry    string `xml:"Prtry"`
	Instrm   string `xml:"Instrm"`
}

// PaymentStatus6 ...
type PaymentStatus6 struct {
	Cd   *PaymentStatusCode6Choice     `xml:"Cd"`
	DtTm *DateAndDateTime2Choice       `xml:"DtTm"`
	Rsn  []*PaymentStatusReason1Choice `xml:"Rsn"`
}

// PaymentStatusCode6Choice ...
type PaymentStatusCode6Choice struct {
	Pdg   string `xml:"Pdg"`
	Fnl   string `xml:"Fnl"`
	RTGS  string `xml:"RTGS"`
	Sttlm string `xml:"Sttlm"`
	Prtry string `xml:"Prtry"`
}

// PaymentStatusReason1Choice ...
type PaymentStatusReason1Choice struct {
	Umtchd       string                           `xml:"Umtchd"`
	Canc         string                           `xml:"Canc"`
	Sspd         string                           `xml:"Sspd"`
	PdgFlngSttlm string                           `xml:"PdgFlngSttlm"`
	PdgSttlm     string                           `xml:"PdgSttlm"`
	PrtryRjctn   *ProprietaryStatusJustification2 `xml:"PrtryRjctn"`
	Prtry        string                           `xml:"Prtry"`
}

// PaymentTransactionParty3 ...
type PaymentTransactionParty3 struct {
	InstgAgt         *BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt"`
	InstdAgt         *BranchAndFinancialInstitutionIdentification6 `xml:"InstdAgt"`
	UltmtDbtr        *Party40Choice                                `xml:"UltmtDbtr"`
	Dbtr             *Party40Choice                                `xml:"Dbtr"`
	DbtrAgt          *BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt"`
	InstgRmbrsmntAgt *BranchAndFinancialInstitutionIdentification6 `xml:"InstgRmbrsmntAgt"`
	InstdRmbrsmntAgt *BranchAndFinancialInstitutionIdentification6 `xml:"InstdRmbrsmntAgt"`
	IntrmyAgt1       *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt1"`
	IntrmyAgt2       *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt2"`
	IntrmyAgt3       *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt3"`
	CdtrAgt          *BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt"`
	Cdtr             *Party40Choice                                `xml:"Cdtr"`
	UltmtCdtr        *Party40Choice                                `xml:"UltmtCdtr"`
}

// PaymentType3Code ...
type PaymentType3Code string

// PaymentType4Choice ...
type PaymentType4Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// PendingFailingSettlement1Code ...
type PendingFailingSettlement1Code string

// PendingSettlement2Code ...
type PendingSettlement2Code string

// PendingStatus4Code ...
type PendingStatus4Code string

// PersonIdentification13 ...
type PersonIdentification13 struct {
	DtAndPlcOfBirth *DateAndPlaceOfBirth1           `xml:"DtAndPlcOfBirth"`
	Othr            []*GenericPersonIdentification1 `xml:"Othr"`
}

// PersonIdentificationSchemeName1Choice ...
type PersonIdentificationSchemeName1Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// PhoneNumber ...
type PhoneNumber string

// PostalAddress24 ...
type PostalAddress24 struct {
	AdrTp       *AddressType3Choice `xml:"AdrTp"`
	Dept        string              `xml:"Dept"`
	SubDept     string              `xml:"SubDept"`
	StrtNm      string              `xml:"StrtNm"`
	BldgNb      string              `xml:"BldgNb"`
	BldgNm      string              `xml:"BldgNm"`
	Flr         string              `xml:"Flr"`
	PstBx       string              `xml:"PstBx"`
	Room        string              `xml:"Room"`
	PstCd       string              `xml:"PstCd"`
	TwnNm       string              `xml:"TwnNm"`
	TwnLctnNm   string              `xml:"TwnLctnNm"`
	DstrctNm    string              `xml:"DstrctNm"`
	CtrySubDvsn string              `xml:"CtrySubDvsn"`
	Ctry        string              `xml:"Ctry"`
	AdrLine     []string            `xml:"AdrLine"`
}

// PreferredContactMethod1Code ...
type PreferredContactMethod1Code string

// Priority1Choice ...
type Priority1Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// Priority5Code ...
type Priority5Code string

// ProprietaryStatusJustification2 ...
type ProprietaryStatusJustification2 struct {
	PrtryStsRsn string `xml:"PrtryStsRsn"`
	Rsn         string `xml:"Rsn"`
}

// ProxyAccountIdentification1 ...
type ProxyAccountIdentification1 struct {
	Tp *ProxyAccountType1Choice `xml:"Tp"`
	Id string                   `xml:"Id"`
}

// ProxyAccountType1Choice ...
type ProxyAccountType1Choice struct {
	Cd    string `xml:"Cd"`
	Prtry string `xml:"Prtry"`
}

// QueueTransactionIdentification1 ...
type QueueTransactionIdentification1 struct {
	QId    string `xml:"QId"`
	PosInQ string `xml:"PosInQ"`
}

// RequestType4Choice ...
type RequestType4Choice struct {
	PmtCtrl string                  `xml:"PmtCtrl"`
	Enqry   string                  `xml:"Enqry"`
	Prtry   *GenericIdentification1 `xml:"Prtry"`
}

// ReturnTransactionV08 ...
type ReturnTransactionV08 struct {
	MsgHdr      *MessageHeader8                  `xml:"MsgHdr"`
	RptOrErr    *TransactionReportOrError4Choice `xml:"RptOrErr"`
	SplmtryData []*SupplementaryData1            `xml:"SplmtryData"`
}

// SecuritiesTransactionReferences1 ...
type SecuritiesTransactionReferences1 struct {
	AcctOwnrTxId      string `xml:"AcctOwnrTxId"`
	AcctSvcrTxId      string `xml:"AcctSvcrTxId"`
	MktInfrstrctrTxId string `xml:"MktInfrstrctrTxId"`
	PrcgId            string `xml:"PrcgId"`
}

// ShortPaymentIdentification2 ...
type ShortPaymentIdentification2 struct {
	TxId          string                                        `xml:"TxId"`
	IntrBkSttlmDt string                                        `xml:"IntrBkSttlmDt"`
	InstgAgt      *BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt"`
}

// SupplementaryData1 ...
type SupplementaryData1 struct {
	PlcAndNm string                      `xml:"PlcAndNm"`
	Envlp    *SupplementaryDataEnvelope1 `xml:"Envlp"`
}

// SupplementaryDataEnvelope1 ...
type SupplementaryDataEnvelope1 struct {
}

// SuspendedStatusReason1Code ...
type SuspendedStatusReason1Code string

// System2 ...
type System2 struct {
	SysId  *MarketInfrastructureIdentification1Choice    `xml:"SysId"`
	MmbId  *BranchAndFinancialInstitutionIdentification6 `xml:"MmbId"`
	Ctry   string                                        `xml:"Ctry"`
	AcctId *AccountIdentification4Choice                 `xml:"AcctId"`
}

// Transaction66 ...
type Transaction66 struct {
	PmtTo        *System2                          `xml:"PmtTo"`
	PmtFr        *System2                          `xml:"PmtFr"`
	CdtDbtInd    string                            `xml:"CdtDbtInd"`
	Pmt          *PaymentInstruction32             `xml:"Pmt"`
	AcctNtry     *CashAccountAndEntry3             `xml:"AcctNtry"`
	SctiesTxRefs *SecuritiesTransactionReferences1 `xml:"SctiesTxRefs"`
}

// TransactionOrError4Choice ...
type TransactionOrError4Choice struct {
	Tx     *Transaction66    `xml:"Tx"`
	BizErr []*ErrorHandling5 `xml:"BizErr"`
}

// TransactionReport5 ...
type TransactionReport5 struct {
	PmtId   *PaymentIdentification6Choice `xml:"PmtId"`
	TxOrErr *TransactionOrError4Choice    `xml:"TxOrErr"`
}

// TransactionReportOrError4Choice ...
type TransactionReportOrError4Choice struct {
	BizRpt  *Transactions8    `xml:"BizRpt"`
	OprlErr []*ErrorHandling5 `xml:"OprlErr"`
}

// Transactions8 ...
type Transactions8 struct {
	PmtCmonInf *PaymentCommon4              `xml:"PmtCmonInf"`
	TxsSummry  *NumberAndSumOfTransactions2 `xml:"TxsSummry"`
	TxRpt      []*TransactionReport5        `xml:"TxRpt"`
}

// TrueFalseIndicator ...
type TrueFalseIndicator bool

// UUIDv4Identifier ...
type UUIDv4Identifier string

// UnmatchedStatusReason1Code ...
type UnmatchedStatusReason1Code string

// YesNoIndicator ...
type YesNoIndicator bool
