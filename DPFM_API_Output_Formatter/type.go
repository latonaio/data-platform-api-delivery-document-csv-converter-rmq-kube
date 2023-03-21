package dpfm_api_output_formatter

type DeliveryDocumentSDC struct {
	ConnectionKey       string             `json:"connection_key"`
	Result              bool               `json:"result"`
	RedisKey            string             `json:"redis_key"`
	Filepath            string             `json:"filepath"`
	APIStatusCode       int                `json:"api_status_code"`
	RuntimeSessionID    string             `json:"runtime_session_id"`
	BusinessPartnerID   *int               `json:"business_partner"`
	ServiceLabel        string             `json:"service_label"`
	APIType             string             `json:"api_type"`
	DataConcatenation   *DataConcatenation `json:"DataConcatenation"`
	APISchema           string             `json:"api_schema"`
	Accepter            []string           `json:"accepter"`
	Deleted             bool               `json:"deleted"`
	APIProcessingResult bool               `json:"api_processing_result"`
	APIProcessingError  *string            `json:"api_processing_error"`
}

type DataConcatenation struct {
	Header  DeliveryDocumentHeader    `json:"DeliveryDocumentHeader"`
	Item    []DeliveryDocumentItem    `json:"DeliveryDocumentItem"`
	Address []DeliveryDocumentAddress `json:"DeliveryDocumentAddress"`
	Partner []DeliveryDocumentPartner `json:"DeliveryDocumentPartner"`
}

type InputParameters struct {
	DeliverToParty            *[]*int    `json:"DeliverToParty"`
	DeliverToPartyTo          *int       `json:"DeliverToPartyTo"`
	DeliverToPartyFrom        *int       `json:"DeliverToPartyFrom"`
	DeliverFromParty          *[]*int    `json:"DeliverFromParty"`
	DeliverFromPartyTo        *int       `json:"DeliverFromPartyTo"`
	DeliverFromPartyFrom      *int       `json:"DeliverFromPartyFrom"`
	DeliverToPlant            *[]*string `json:"DeliverToPlant"`
	DeliverToPlantTo          *string    `json:"DeliverToPlantTo"`
	DeliverToPlantFrom        *string    `json:"DeliverToPlantFrom"`
	DeliverFromPlant          *[]*string `json:"DeliverFromPlant"`
	DeliverFromPlantTo        *string    `json:"DeliverFromPlantTo"`
	DeliverFromPlantFrom      *string    `json:"DeliverFromPlantFrom"`
	ConfirmedDeliveryDate     *[]*string `json:"ConfirmedDeliveryDate"`
	ConfirmedDeliveryDateTo   *string    `json:"ConfirmedDeliveryDateTo"`
	ConfirmedDeliveryDateFrom *string    `json:"ConfirmedDeliveryDateFrom"`
}

type DeliveryDocumentHeader struct {
	DeliveryDocument                       int      `json:"DeliveryDocument" csv:"1"`
	SupplyChainRelationshipID              *int     `json:"SupplyChainRelationshipID" csv:"2"`
	SupplyChainRelationshipDeliveryID      *int     `json:"SupplyChainRelationshipDeliveryID" csv:"3"`
	SupplyChainRelationshipDeliveryPlantID *int     `json:"SupplyChainRelationshipDeliveryPlantID" csv:"4"`
	SupplyChainRelationshipBillingID       *int     `json:"SupplyChainRelationshipBillingID" csv:"5"`
	SupplyChainRelationshipPaymentID       *int     `json:"SupplyChainRelationshipPaymentID" csv:"6"`
	Buyer                                  *int     `json:"Buyer" csv:"7"`
	Seller                                 *int     `json:"Seller" csv:"8"`
	DeliverToParty                         *int     `json:"DeliverToParty" csv:"9"`
	DeliverFromParty                       *int     `json:"DeliverFromParty" csv:"10"`
	DeliverToPlant                         *string  `json:"DeliverToPlant" csv:"11"`
	DeliverFromPlant                       *string  `json:"DeliverFromPlant" csv:"12"`
	BillToParty                            *int     `json:"BillToParty" csv:"13"`
	BillFromParty                          *int     `json:"BillFromParty" csv:"14"`
	BillToCountry                          *string  `json:"BillToCountry" csv:"15"`
	BillFromCountry                        *string  `json:"BillFromCountry" csv:"16"`
	Payer                                  *int     `json:"Payer" csv:"17"`
	Payee                                  *int     `json:"Payee" csv:"18"`
	IsExportImport                         *bool    `json:"IsExportImport" csv:"19"`
	DeliverToPlantTimeZone                 *string  `json:"DeliverToPlantTimeZone" csv:"20"`
	DeliverFromPlantTimeZone               *string  `json:"DeliverFromPlantTimeZone" csv:"21"`
	ReferenceDocument                      *int     `json:"ReferenceDocument" csv:"22"`
	ReferenceDocumentItem                  *int     `json:"ReferenceDocumentItem" csv:"23"`
	OrderID                                *int     `json:"OrderID" csv:"24"`
	OrderItem                              *int     `json:"OrderItem" csv:"25"`
	ContractType                           *string  `json:"ContractType" csv:"26"`
	OrderValidityStartDate                 *string  `json:"OrderValidityStartDate" csv:"27"`
	OrderValidityEndDate                   *string  `json:"OrderValidityEndDate" csv:"28"`
	DocumentDate                           *string  `json:"DocumentDate" csv:"29"`
	PlannedGoodsIssueDate                  *string  `json:"PlannedGoodsIssueDate" csv:"30"`
	PlannedGoodsIssueTime                  *string  `json:"PlannedGoodsIssueTime" csv:"31"`
	PlannedGoodsReceiptDate                *string  `json:"PlannedGoodsReceiptDate" csv:"32"`
	PlannedGoodsReceiptTime                *string  `json:"PlannedGoodsReceiptTime" csv:"33"`
	InvoiceDocumentDate                    *string  `json:"InvoiceDocumentDate" csv:"34"`
	HeaderCompleteDeliveryIsDefined        *bool    `json:"HeaderCompleteDeliveryIsDefined" csv:"35"`
	HeaderDeliveryStatus                   *string  `json:"HeaderDeliveryStatus" csv:"36"`
	CreationDate                           *string  `json:"CreationDate" csv:"37"`
	CreationTime                           *string  `json:"CreationTime" csv:"38"`
	LastChangeDate                         *string  `json:"LastChangeDate" csv:"39"`
	LastChangeTime                         *string  `json:"LastChangeTime" csv:"40"`
	GoodsIssueOrReceiptSlipNumber          *string  `json:"GoodsIssueOrReceiptSlipNumber" csv:"41"`
	HeaderBillingStatus                    *string  `json:"HeaderBillingStatus" csv:"42"`
	HeaderBillingConfStatus                *string  `json:"HeaderBillingConfStatus" csv:"43"`
	HeaderBillingBlockStatus               *bool    `json:"HeaderBillingBlockStatus" csv:"44"`
	HeaderGrossWeight                      *float32 `json:"HeaderGrossWeight" csv:"45"`
	HeaderNetWeight                        *float32 `json:"HeaderNetWeight" csv:"46"`
	HeaderWeightUnit                       *string  `json:"HeaderWeightUnit" csv:"47"`
	Incoterms                              *string  `json:"Incoterms" csv:"48"`
	TransactionCurrency                    *string  `json:"TransactionCurrency" csv:"49"`
	HeaderDeliveryBlockStatus              *bool    `json:"HeaderDeliveryBlockStatus" csv:"50"`
	HeaderIssuingBlockStatus               *bool    `json:"HeaderIssuingBlockStatus" csv:"51"`
	HeaderReceivingBlockStatus             *bool    `json:"HeaderReceivingBlockStatus" csv:"52"`
	IsCancelled                            *bool    `json:"IsCancelled" csv:"53"`
	IsMarkedForDeletion                    *bool    `json:"IsMarkedForDeletion" csv:"54"`
}

type DeliveryDocumentItem struct {
	DeliveryDocument                              int      `json:"DeliveryDocument" csv:"1"`
	DeliveryDocumentItem                          int      `json:"DeliveryDocumentItem" csv:"55"`
	DeliveryDocumentItemCategory                  *string  `json:"DeliveryDocumentItemCategory" csv:"56"`
	SupplyChainRelationshipID                     *int     `json:"SupplyChainRelationshipID" csv:"57"`
	SupplyChainRelationshipDeliveryID             *int     `json:"SupplyChainRelationshipDeliveryID" csv:"58"`
	SupplyChainRelationshipDeliveryPlantID        *int     `json:"SupplyChainRelationshipDeliveryPlantID" csv:"59"`
	SupplyChainRelationshipStockConfPlantID       *int     `json:"SupplyChainRelationshipStockConfPlantID" csv:"60"`
	SupplyChainRelationshipProductionPlantID      *int     `json:"SupplyChainRelationshipProductionPlantID" csv:"61"`
	SupplyChainRelationshipBillingID              *int     `json:"SupplyChainRelationshipBillingID" csv:"62"`
	SupplyChainRelationshipPaymentID              *int     `json:"SupplyChainRelationshipPaymentID" csv:"63"`
	Buyer                                         *int     `json:"Buyer" csv:"64"`
	Seller                                        *int     `json:"Seller" csv:"65"`
	DeliverToParty                                *int     `json:"DeliverToParty" csv:"66"`
	DeliverFromParty                              *int     `json:"DeliverFromParty" csv:"67"`
	DeliverToPlant                                *string  `json:"DeliverToPlant" csv:"68"`
	DeliverFromPlant                              *string  `json:"DeliverFromPlant" csv:"69"`
	BillToParty                                   *int     `json:"BillToParty" csv:"70"`
	BillFromParty                                 *int     `json:"BillFromParty" csv:"71"`
	BillToCountry                                 *string  `json:"BillToCountry" csv:"72"`
	BillFromCountry                               *string  `json:"BillFromCountry" csv:"73"`
	Payer                                         *int     `json:"Payer" csv:"74"`
	Payee                                         *int     `json:"Payee" csv:"75"`
	DeliverToPlantStorageLocation                 *string  `json:"DeliverToPlantStorageLocation" csv:"76"`
	ProductIsBatchManagedInDeliverToPlant         *bool    `json:"ProductIsBatchManagedInDeliverToPlant" csv:"77"`
	BatchMgmtPolicyInDeliverToPlant               *string  `json:"BatchMgmtPolicyInDeliverToPlant" csv:"78"`
	DeliverToPlantBatch                           *string  `json:"DeliverToPlantBatch" csv:"79"`
	DeliverToPlantBatchValidityStartDate          *string  `json:"DeliverToPlantBatchValidityStartDate" csv:"80"`
	DeliverToPlantBatchValidityStartTime          *string  `json:"DeliverToPlantBatchValidityStartTime" csv:"81"`
	DeliverToPlantBatchValidityEndDate            *string  `json:"DeliverToPlantBatchValidityEndDate" csv:"82"`
	DeliverToPlantBatchValidityEndTime            *string  `json:"DeliverToPlantBatchValidityEndTime" csv:"83"`
	DeliverFromPlantStorageLocation               *string  `json:"DeliverFromPlantStorageLocation" csv:"84"`
	ProductIsBatchManagedInDeliverFromPlant       *bool    `json:"ProductIsBatchManagedInDeliverFromPlant" csv:"85"`
	BatchMgmtPolicyInDeliverFromPlant             *string  `json:"BatchMgmtPolicyInDeliverFromPlant" csv:"86"`
	DeliverFromPlantBatch                         *string  `json:"DeliverFromPlantBatch" csv:"87"`
	DeliverFromPlantBatchValidityStartDate        *string  `json:"DeliverFromPlantBatchValidityStartDate" csv:"88"`
	DeliverFromPlantBatchValidityStartTime        *string  `json:"DeliverFromPlantBatchValidityStartTime" csv:"89"`
	DeliverFromPlantBatchValidityEndDate          *string  `json:"DeliverFromPlantBatchValidityEndDate" csv:"90"`
	DeliverFromPlantBatchValidityEndTime          *string  `json:"DeliverFromPlantBatchValidityEndTime" csv:"91"`
	StockConfirmationBusinessPartner              *int     `json:"StockConfirmationBusinessPartner" csv:"92"`
	StockConfirmationPlant                        *string  `json:"StockConfirmationPlant" csv:"93"`
	ProductIsBatchManagedInStockConfirmationPlant *bool    `json:"ProductIsBatchManagedInStockConfirmationPlant" csv:"94"`
	BatchMgmtPolicyInStockConfirmationPlant       *string  `json:"BatchMgmtPolicyInStockConfirmationPlant" csv:"95"`
	StockConfirmationPlantBatch                   *string  `json:"StockConfirmationPlantBatch" csv:"96"`
	StockConfirmationPlantBatchValidityStartDate  *string  `json:"StockConfirmationPlantBatchValidityStartDate" csv:"97"`
	StockConfirmationPlantBatchValidityStartTime  *string  `json:"StockConfirmationPlantBatchValidityStartTime" csv:"98"`
	StockConfirmationPlantBatchValidityEndDate    *string  `json:"StockConfirmationPlantBatchValidityEndDate" csv:"99"`
	StockConfirmationPlantBatchValidityEndTime    *string  `json:"StockConfirmationPlantBatchValidityEndTime" csv:"100"`
	StockConfirmationPolicy                       *string  `json:"StockConfirmationPolicy" csv:"101"`
	StockConfirmationStatus                       *string  `json:"StockConfirmationStatus" csv:"102"`
	ProductionPlantBusinessPartner                *int     `json:"ProductionPlantBusinessPartner" csv:"103"`
	ProductionPlant                               *string  `json:"ProductionPlant" csv:"104"`
	ProductionPlantStorageLocation                *string  `json:"ProductionPlantStorageLocation" csv:"105"`
	ProductIsBatchManagedInProductionPlant        *bool    `json:"ProductIsBatchManagedInProductionPlant" csv:"106"`
	BatchMgmtPolicyInProductionPlant              *string  `json:"BatchMgmtPolicyInProductionPlant" csv:"107"`
	ProductionPlantBatch                          *string  `json:"ProductionPlantBatch" csv:"108"`
	ProductionPlantBatchValidityStartDate         *string  `json:"ProductionPlantBatchValidityStartDate" csv:"109"`
	ProductionPlantBatchValidityStartTime         *string  `json:"ProductionPlantBatchValidityStartTime" csv:"110"`
	ProductionPlantBatchValidityEndDate           *string  `json:"ProductionPlantBatchValidityEndDate" csv:"111"`
	ProductionPlantBatchValidityEndTime           *string  `json:"ProductionPlantBatchValidityEndTime" csv:"112"`
	InspectionPlan                                *int     `json:"InspectionPlan" csv:"113"`
	InspectionPlant                               *string  `json:"InspectionPlant" csv:"114"`
	InspectionOrder                               *int     `json:"InspectionOrder" csv:"115"`
	DeliveryDocumentItemText                      *string  `json:"DeliveryDocumentItemText" csv:"116"`
	DeliveryDocumentItemTextByBuyer               *string  `json:"DeliveryDocumentItemTextByBuyer" csv:"117"`
	DeliveryDocumentItemTextBySeller              *string  `json:"DeliveryDocumentItemTextBySeller" csv:"118"`
	Product                                       *string  `json:"Product" csv:"119"`
	ProductStandardID                             *string  `json:"ProductStandardID" csv:"120"`
	ProductGroup                                  *string  `json:"ProductGroup" csv:"121"`
	BaseUnit                                      *string  `json:"BaseUnit" csv:"122"`
	OriginalQuantityInBaseUnit                    *float32 `json:"OriginalQuantityInBaseUnit" csv:"123"`
	DeliveryUnit                                  *string  `json:"DeliveryUnit" csv:"124"`
	PlannedGoodsIssueDate                         *string  `json:"PlannedGoodsIssueDate" csv:"125"`
	PlannedGoodsIssueTime                         *string  `json:"PlannedGoodsIssueTime" csv:"126"`
	PlannedGoodsReceiptDate                       *string  `json:"PlannedGoodsReceiptDate" csv:"127"`
	PlannedGoodsReceiptTime                       *string  `json:"PlannedGoodsReceiptTime" csv:"128"`
	PlannedGoodsIssueQuantity                     *float32 `json:"PlannedGoodsIssueQuantity" csv:"129"`
	PlannedGoodsIssueQtyInBaseUnit                *float32 `json:"PlannedGoodsIssueQtyInBaseUnit" csv:"130"`
	PlannedGoodsReceiptQuantity                   *float32 `json:"PlannedGoodsReceiptQuantity" csv:"131"`
	PlannedGoodsReceiptQtyInBaseUnit              *float32 `json:"PlannedGoodsReceiptQtyInBaseUnit" csv:"132"`
	ActualGoodsIssueDate                          *string  `json:"ActualGoodsIssueDate" csv:"133"`
	ActualGoodsIssueTime                          *string  `json:"ActualGoodsIssueTime" csv:"134"`
	ActualGoodsReceiptDate                        *string  `json:"ActualGoodsReceiptDate" csv:"135"`
	ActualGoodsReceiptTime                        *string  `json:"ActualGoodsReceiptTime" csv:"136"`
	ActualGoodsIssueQuantity                      *float32 `json:"ActualGoodsIssueQuantity" csv:"137"`
	ActualGoodsIssueQtyInBaseUnit                 *float32 `json:"ActualGoodsIssueQtyInBaseUnit" csv:"138"`
	ActualGoodsReceiptQuantity                    *float32 `json:"ActualGoodsReceiptQuantity" csv:"139"`
	ActualGoodsReceiptQtyInBaseUnit               *float32 `json:"ActualGoodsReceiptQtyInBaseUnit" csv:"140"`
	CreationDate                                  *string  `json:"CreationDate" csv:"141"`
	CreationTime                                  *string  `json:"CreationTime" csv:"142"`
	LastChangeDate                                *string  `json:"LastChangeDate" csv:"143"`
	LastChangeTime                                *string  `json:"LastChangeTime" csv:"144"`
	ItemBillingStatus                             *string  `json:"ItemBillingStatus" csv:"145"`
	ItemCompleteDeliveryIsDefined                 *bool    `json:"ItemCompleteDeliveryIsDefined" csv:"146"`
	ItemGrossWeight                               *float32 `json:"ItemGrossWeight" csv:"147"`
	ItemNetWeight                                 *float32 `json:"ItemNetWeight" csv:"148"`
	ItemWeightUnit                                *string  `json:"ItemWeightUnit" csv:"149"`
	InternalCapacityQuantity                      *float32 `json:"InternalCapacityQuantity" csv:"150"`
	InternalCapacityQuantityUnit                  *string  `json:"InternalCapacityQuantityUnit" csv:"151"`
	ItemIsBillingRelevant                         *bool    `json:"ItemIsBillingRelevant" csv:"152"`
	NetAmount                                     *float32 `json:"NetAmount" csv:"153"`
	TaxAmount                                     *float32 `json:"TaxAmount" csv:"154"`
	GrossAmount                                   *float32 `json:"GrossAmount" csv:"155"`
	OrderID                                       *int     `json:"OrderID" csv:"156"`
	OrderItem                                     *int     `json:"OrderItem" csv:"157"`
	OrderType                                     *string  `json:"OrderType" csv:"158"`
	ContractType                                  *string  `json:"ContractType" csv:"159"`
	OrderValidityStartDate                        *string  `json:"OrderValidityStartDate" csv:"160"`
	OrderValidityEndDate                          *string  `json:"OrderValidityEndDate" csv:"161"`
	PaymentTerms                                  *string  `json:"PaymentTerms" csv:"162"`
	DueCalculationBaseDate                        *string  `json:"DueCalculationBaseDate" csv:"163"`
	PaymentDueDate                                *string  `json:"PaymentDueDate" csv:"164"`
	NetPaymentDays                                *int     `json:"NetPaymentDays" csv:"165"`
	PaymentMethod                                 *string  `json:"PaymentMethod" csv:"166"`
	InvoicePeriodStartDate                        *string  `json:"InvoicePeriodStartDate" csv:"167"`
	InvoicePeriodEndDate                          *string  `json:"InvoicePeriodEndDate" csv:"168"`
	ConfirmedDeliveryDate                         *string  `json:"ConfirmedDeliveryDate" csv:"169"`
	Project                                       *string  `json:"Project" csv:"170"`
	ReferenceDocument                             *int     `json:"ReferenceDocument" csv:"171"`
	ReferenceDocumentItem                         *int     `json:"ReferenceDocumentItem" csv:"172"`
	TransactionTaxClassification                  *string  `json:"TransactionTaxClassification" csv:"173"`
	ProductTaxClassificationBillToCountry         *string  `json:"ProductTaxClassificationBillToCountry" csv:"174"`
	ProductTaxClassificationBillFromCountry       *string  `json:"ProductTaxClassificationBillFromCountry" csv:"175"`
	DefinedTaxClassification                      *string  `json:"DefinedTaxClassification" csv:"176"`
	AccountAssignmentGroup                        *string  `json:"AccountAssignmentGroup" csv:"177"`
	ProductAccountAssignmentGroup                 *string  `json:"ProductAccountAssignmentGroup" csv:"178"`
	TaxCode                                       *string  `json:"TaxCode" csv:"179"`
	TaxRate                                       *float32 `json:"TaxRate" csv:"180"`
	CountryOfOrigin                               *string  `json:"CountryOfOrigin" csv:"181"`
	CountryOfOriginLanguage                       *string  `json:"CountryOfOriginLanguage" csv:"182"`
	ItemDeliveryBlockStatus                       *bool    `json:"ItemDeliveryBlockStatus" csv:"183"`
	ItemIssuingBlockStatus                        *bool    `json:"ItemIssuingBlockStatus" csv:"184"`
	ItemReceivingBlockStatus                      *bool    `json:"ItemReceivingBlockStatus" csv:"185"`
	ItemBillingBlockStatus                        *bool    `json:"ItemBillingBlockStatus" csv:"186"`
	IsCancelled                                   *bool    `json:"IsCancelled" csv:"187"`
	IsMarkedForDeletion                           *bool    `json:"IsMarkedForDeletion" csv:"188"`
}

type DeliveryDocumentPartner struct {
	DeliveryDocument        int     `json:"DeliveryDocument" csv:"1"`
	PartnerFunction         string  `json:"PartnerFunction" csv:"189"`
	BusinessPartner         int     `json:"BusinessPartner" csv:"190"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName" csv:"191"`
	BusinessPartnerName     *string `json:"BusinessPartnerName" csv:"192"`
	Organization            *string `json:"Organization" csv:"193"`
	Country                 *string `json:"Country" csv:"194"`
	Language                *string `json:"Language" csv:"195"`
	Currency                *string `json:"Currency" csv:"196"`
	ExternalDocumentID      *string `json:"ExternalDocumentID" csv:"197"`
	AddressID               *int    `json:"AddressID" csv:"198"`
}

type DeliveryDocumentAddress struct {
	DeliveryDocument int     `json:"DeliveryDocument" csv:"1"`
	AddressID        int     `json:"AddressID" csv:"199"`
	PostalCode       *string `json:"PostalCode" csv:"200"`
	LocalRegion      *string `json:"LocalRegion" csv:"201"`
	Country          *string `json:"Country" csv:"202"`
	District         *string `json:"District" csv:"203"`
	StreetName       *string `json:"StreetName" csv:"204"`
	CityName         *string `json:"CityName" csv:"205"`
	Building         *string `json:"Building" csv:"206"`
	Floor            *int    `json:"Floor" csv:"207"`
	Room             *int    `json:"Room" csv:"208"`
}
