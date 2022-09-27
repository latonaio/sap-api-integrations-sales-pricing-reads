package responses

type ToSlsPrcgCndnRecdValidity struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ConditionRecord                string `json:"ConditionRecord"`
			ConditionValidityEndDate       string `json:"ConditionValidityEndDate"`
			ConditionValidityStartDate     string `json:"ConditionValidityStartDate"`
			ConditionApplication           string `json:"ConditionApplication"`
			ConditionType                  string `json:"ConditionType"`
			ConditionReleaseStatus         string `json:"ConditionReleaseStatus"`
			DepartureCountry               string `json:"DepartureCountry"`
			SalesDocument                  string `json:"SalesDocument"`
			SalesDocumentItem              string `json:"SalesDocumentItem"`
			BillableControl                string `json:"BillableControl"`
			Industry                       string `json:"Industry"`
			CityCode                       string `json:"CityCode"`
			ConditionContract              string `json:"ConditionContract"`
			County                         string `json:"County"`
			EngagementProject              string `json:"EngagementProject"`
			ConfigurationNumber            string `json:"ConfigurationNumber"`
			BRNFDocumentType               string `json:"BR_NFDocumentType"`
			BRSpcfcFreeDefinedField1       string `json:"BRSpcfcFreeDefinedField1"`
			BRSpcfcFreeDefinedField2       string `json:"BRSpcfcFreeDefinedField2"`
			BRSpcfcFreeDefinedField3       string `json:"BRSpcfcFreeDefinedField3"`
			InternationalArticleNumber     string `json:"InternationalArticleNumber"`
			TechnicalObjectType            string `json:"TechnicalObjectType"`
			Equipment                      string `json:"Equipment"`
			CustomerHierarchy              string `json:"CustomerHierarchy"`
			IncotermsClassification        string `json:"IncotermsClassification"`
			IncotermsTransferLocation      string `json:"IncotermsTransferLocation"`
			AccountTaxType                 string `json:"AccountTaxType"`
			TxRlvnceClassfctnForArgentina  string `json:"TxRlvnceClassfctnForArgentina"`
			BRTaxCode                      string `json:"BR_TaxCode"`
			LocalSalesTaxApplicabilityCode string `json:"LocalSalesTaxApplicabilityCode"`
			CustomerGroup                  string `json:"CustomerGroup"`
			CustomerPriceGroup             string `json:"CustomerPriceGroup"`
			MaterialPricingGroup           string `json:"MaterialPricingGroup"`
			SoldToParty                    string `json:"SoldToParty"`
			BPForSoldToParty               string `json:"BPForSoldToParty"`
			Customer                       string `json:"Customer"`
			BPForCustomer                  string `json:"BPForCustomer"`
			PayerParty                     string `json:"PayerParty"`
			BPForPayerParty                string `json:"BPForPayerParty"`
			ShipToParty                    string `json:"ShipToParty"`
			BPForShipToParty               string `json:"BPForShipToParty"`
			Supplier                       string `json:"Supplier"`
			BPForSupplier                  string `json:"BPForSupplier"`
			DestinationCountry             string `json:"DestinationCountry"`
			MaterialGroup                  string `json:"MaterialGroup"`
			Material                       string `json:"Material"`
			ReturnsRefundExtent            string `json:"ReturnsRefundExtent"`
			AdditionalMaterialGroup1       string `json:"AdditionalMaterialGroup1"`
			AdditionalMaterialGroup2       string `json:"AdditionalMaterialGroup2"`
			AdditionalMaterialGroup3       string `json:"AdditionalMaterialGroup3"`
			AdditionalMaterialGroup4       string `json:"AdditionalMaterialGroup4"`
			AdditionalMaterialGroup5       string `json:"AdditionalMaterialGroup5"`
			Personnel                      string `json:"Personnel"`
			PriceListType                  string `json:"PriceListType"`
			PostalCode                     string `json:"PostalCode"`
			Region                         string `json:"Region"`
			EngagementProjectServiceOrg    string `json:"EngagementProjectServiceOrg"`
			RequirementSegment             string `json:"RequirementSegment"`
			StockSegment                   string `json:"StockSegment"`
			Division                       string `json:"Division"`
			CommodityCode                  string `json:"CommodityCode"`
			ConsumptionTaxCtrlCode         string `json:"ConsumptionTaxCtrlCode"`
			BRSpcfcTaxBasePercentageCode   string `json:"BRSpcfcTaxBasePercentageCode"`
			BRSpcfcTxGrpDynTaxExceptions   string `json:"BRSpcfcTxGrpDynTaxExceptions"`
			CustomerTaxClassification1     string `json:"CustomerTaxClassification1"`
			CustomerTaxClassification2     string `json:"CustomerTaxClassification2"`
			CustomerTaxClassification3     string `json:"CustomerTaxClassification3"`
			CustomerTaxClassification4     string `json:"CustomerTaxClassification4"`
			ProductTaxClassification1      string `json:"ProductTaxClassification1"`
			ProductTaxClassification2      string `json:"ProductTaxClassification2"`
			ProductTaxClassification3      string `json:"ProductTaxClassification3"`
			ProductTaxClassification4      string `json:"ProductTaxClassification4"`
			TradingContract                string `json:"TradingContract"`
			TradingContractItem            string `json:"TradingContractItem"`
			TaxJurisdiction                string `json:"TaxJurisdiction"`
			BRSpcfcTaxDepartureRegion      string `json:"BRSpcfcTaxDepartureRegion"`
			BRSpcfcTaxDestinationRegion    string `json:"BRSpcfcTaxDestinationRegion"`
			MainItemMaterialPricingGroup   string `json:"MainItemMaterialPricingGroup"`
			MainItemPricingRefMaterial     string `json:"MainItemPricingRefMaterial"`
			VariantCondition               string `json:"VariantCondition"`
			ValueAddedServiceChargeCode    string `json:"ValueAddedServiceChargeCode"`
			SDDocument                     string `json:"SDDocument"`
			ReferenceSDDocument            string `json:"ReferenceSDDocument"`
			ReferenceSDDocumentItem        string `json:"ReferenceSDDocumentItem"`
			SalesOffice                    string `json:"SalesOffice"`
			SalesGroup                     string `json:"SalesGroup"`
			SalesOrganization              string `json:"SalesOrganization"`
			SalesOrderSalesOrganization    string `json:"SalesOrderSalesOrganization"`
			OrderQuantityUnit              string `json:"OrderQuantityUnit"`
			DistributionChannel            string `json:"DistributionChannel"`
			TransactionCurrency            string `json:"TransactionCurrency"`
			WBSElementInternalID           string `json:"WBSElementInternalID"`
			WBSElementExternalID           string `json:"WBSElementExternalID"`
			WorkPackage                    string `json:"WorkPackage"`
			Plant                          string `json:"Plant"`
			PlantRegion                    string `json:"PlantRegion"`
			WorkItem                       string `json:"WorkItem"`
			ConditionProcessingStatus      string `json:"ConditionProcessingStatus"`
			PricingDate                    string `json:"PricingDate"`
			ConditionIsExclusive           bool   `json:"ConditionIsExclusive"`
			ConditionScaleBasisValue       string `json:"ConditionScaleBasisValue"`
			TaxCode                        string `json:"TaxCode"`
			ServiceDocument                string `json:"ServiceDocument"`
			ServiceDocumentItem            string `json:"ServiceDocumentItem"`
			TimeSheetOvertimeCategory      string `json:"TimeSheetOvertimeCategory"`
			SalesSDDocumentCategory        string `json:"SalesSDDocumentCategory"`
			ReturnReason                   string `json:"ReturnReason"`
			ProductHierarchyNode           string `json:"ProductHierarchyNode"`
			CustomerConditionGroup         string `json:"CustomerConditionGroup"`
			ShippingType                   string `json:"ShippingType"`
			SubscriptionContractDuration   string `json:"SubscriptionContractDuration"`
			SubscrpnContrDurationUnit      string `json:"SubscrpnContrDurationUnit"`
			SubscriptionContractLockReason string `json:"SubscriptionContractLockReason"`
			CrsCtlgMappgPriceVersionNumber string `json:"CrsCtlgMappgPriceVersionNumber"`
			PRAContract                    string `json:"PRAContract"`
			Well                           string `json:"Well"`
			WellCompletion                 string `json:"WellCompletion"`
			MeasurementPoint               string `json:"MeasurementPoint"`
			PricingFormulaNumber           string `json:"PricingFormulaNumber"`
			ETag                           string `json:"ETag"`
			ToSlsPrcgConditionRecord       struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_SlsPrcgConditionRecord"`
		} `json:"results"`
	} `json:"d"`
}