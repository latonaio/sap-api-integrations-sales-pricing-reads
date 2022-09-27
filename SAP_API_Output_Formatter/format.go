package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-sales-pricing-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToSlsPrcgCndnRecdSuplmnt(raw []byte, l *logger.Logger) ([]SlsPrcgCndnRecdSuplmnt, error) {
	pm := &responses.SlsPrcgCndnRecdSuplmnt{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ConditionRecordSuplmnt. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	conditionRecordSuplmnt := make([]SlsPrcgCndnRecdSuplmnt, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		conditionRecordSuplmnt = append(conditionRecordSuplmnt, SlsPrcgCndnRecdSuplmnt{
			ConditionRecord:              data.ConditionRecord,
			ConditionSequentialNumber:    data.ConditionSequentialNumber,
			ConditionTable:               data.ConditionTable,
			ConditionApplication:         data.ConditionApplication,
			ConditionType:                data.ConditionType,
			ConditionValidityEndDate:     data.ConditionValidityEndDate,
			ConditionValidityStartDate:   data.ConditionValidityStartDate,
			CreatedByUser:                data.CreatedByUser,
			CreationDate:                 data.CreationDate,
			ConditionTextID:              data.ConditionTextID,
			PricingScaleType:             data.PricingScaleType,
			PricingScaleBasis:            data.PricingScaleBasis,
			ConditionScaleQuantity:       data.ConditionScaleQuantity,
			ConditionScaleQuantityUnit:   data.ConditionScaleQuantityUnit,
			ConditionScaleAmount:         data.ConditionScaleAmount,
			ConditionScaleAmountCurrency: data.ConditionScaleAmountCurrency,
			ConditionCalculationType:     data.ConditionCalculationType,
			ConditionRateValue:           data.ConditionRateValue,
			ConditionRateValueUnit:       data.ConditionRateValueUnit,
			ConditionRateRatio:           data.ConditionRateRatio,
			ConditionRateRatioUnit:       data.ConditionRateRatioUnit,
			ConditionRateAmount:          data.ConditionRateAmount,
			ConditionCurrency:            data.ConditionCurrency,
			ConditionQuantity:            data.ConditionQuantity,
			ConditionQuantityUnit:        data.ConditionQuantityUnit,
			ConditionToBaseQtyNmrtr:      data.ConditionToBaseQtyNmrtr,
			ConditionToBaseQtyDnmntr:     data.ConditionToBaseQtyDnmntr,
			BaseUnit:                     data.BaseUnit,
			ConditionLowerLimit:          data.ConditionLowerLimit,
			ConditionLowerLimitAmount:    data.ConditionLowerLimitAmount,
			ConditionLowerLimitRatio:     data.ConditionLowerLimitRatio,
			ConditionUpperLimit:          data.ConditionUpperLimit,
			ConditionUpperLimitAmount:    data.ConditionUpperLimitAmount,
			ConditionUpperLimitRatio:     data.ConditionUpperLimitRatio,
			ConditionAlternativeCurrency: data.ConditionAlternativeCurrency,
			ConditionExclusion:           data.ConditionExclusion,
			ConditionIsDeleted:           data.ConditionIsDeleted,
			AdditionalValueDays:          data.AdditionalValueDays,
			FixedValueDate:               data.FixedValueDate,
			PaymentTerms:                 data.PaymentTerms,
			CndnMaxNumberOfSalesOrders:   data.CndnMaxNumberOfSalesOrders,
			MinimumConditionBasisValue:   data.MinimumConditionBasisValue,
			MaximumConditionBasisValue:   data.MaximumConditionBasisValue,
			MaximumConditionAmount:       data.MaximumConditionAmount,
			IncrementalScale:             data.IncrementalScale,
			PricingScaleLine:             data.PricingScaleLine,
			ConditionReleaseStatus:       data.ConditionReleaseStatus,
			ETag:                         data.ETag,
			ToSlsPrcgCndnRecdValidity:    data.ToSlsPrcgCndnRecdValidity.Deferred.URI,
		})
	}

	return conditionRecordSuplmnt, nil
}

func ConvertToToSlsPrcgCndnRecdValidity(raw []byte, l *logger.Logger) ([]ToSlsPrcgCndnRecdValidity, error) {
	pm := &responses.ToSlsPrcgCndnRecdValidity{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToSlsPrcgCndnRecdValidity. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	slsPrcgCndnRecdValidity := make([]ToSlsPrcgCndnRecdValidity, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		slsPrcgCndnRecdValidity = append(slsPrcgCndnRecdValidity, ToSlsPrcgCndnRecdValidity{
			ConditionRecord:                data.ConditionRecord,
			ConditionValidityEndDate:       data.ConditionValidityEndDate,
			ConditionValidityStartDate:     data.ConditionValidityStartDate,
			ConditionApplication:           data.ConditionApplication,
			ConditionType:                  data.ConditionType,
			ConditionReleaseStatus:         data.ConditionReleaseStatus,
			DepartureCountry:               data.DepartureCountry,
			SalesDocument:                  data.SalesDocument,
			SalesDocumentItem:              data.SalesDocumentItem,
			BillableControl:                data.BillableControl,
			Industry:                       data.Industry,
			CityCode:                       data.CityCode,
			ConditionContract:              data.ConditionContract,
			County:                         data.County,
			EngagementProject:              data.EngagementProject,
			ConfigurationNumber:            data.ConfigurationNumber,
			BRNFDocumentType:               data.BRNFDocumentType,
			BRSpcfcFreeDefinedField1:       data.BRSpcfcFreeDefinedField1,
			BRSpcfcFreeDefinedField2:       data.BRSpcfcFreeDefinedField2,
			BRSpcfcFreeDefinedField3:       data.BRSpcfcFreeDefinedField3,
			InternationalArticleNumber:     data.InternationalArticleNumber,
			TechnicalObjectType:            data.TechnicalObjectType,
			Equipment:                      data.Equipment,
			CustomerHierarchy:              data.CustomerHierarchy,
			IncotermsClassification:        data.IncotermsClassification,
			IncotermsTransferLocation:      data.IncotermsTransferLocation,
			AccountTaxType:                 data.AccountTaxType,
			TxRlvnceClassfctnForArgentina:  data.TxRlvnceClassfctnForArgentina,
			BRTaxCode:                      data.BRTaxCode,
			LocalSalesTaxApplicabilityCode: data.LocalSalesTaxApplicabilityCode,
			CustomerGroup:                  data.CustomerGroup,
			CustomerPriceGroup:             data.CustomerPriceGroup,
			MaterialPricingGroup:           data.MaterialPricingGroup,
			SoldToParty:                    data.SoldToParty,
			BPForSoldToParty:               data.BPForSoldToParty,
			Customer:                       data.Customer,
			BPForCustomer:                  data.BPForCustomer,
			PayerParty:                     data.PayerParty,
			BPForPayerParty:                data.BPForPayerParty,
			ShipToParty:                    data.ShipToParty,
			BPForShipToParty:               data.BPForShipToParty,
			Supplier:                       data.Supplier,
			BPForSupplier:                  data.BPForSupplier,
			DestinationCountry:             data.DestinationCountry,
			MaterialGroup:                  data.MaterialGroup,
			Material:                       data.Material,
			ReturnsRefundExtent:            data.ReturnsRefundExtent,
			AdditionalMaterialGroup1:       data.AdditionalMaterialGroup1,
			AdditionalMaterialGroup2:       data.AdditionalMaterialGroup2,
			AdditionalMaterialGroup3:       data.AdditionalMaterialGroup3,
			AdditionalMaterialGroup4:       data.AdditionalMaterialGroup4,
			AdditionalMaterialGroup5:       data.AdditionalMaterialGroup5,
			Personnel:                      data.Personnel,
			PriceListType:                  data.PriceListType,
			PostalCode:                     data.PostalCode,
			Region:                         data.Region,
			EngagementProjectServiceOrg:    data.EngagementProjectServiceOrg,
			RequirementSegment:             data.RequirementSegment,
			StockSegment:                   data.StockSegment,
			Division:                       data.Division,
			CommodityCode:                  data.CommodityCode,
			ConsumptionTaxCtrlCode:         data.ConsumptionTaxCtrlCode,
			BRSpcfcTaxBasePercentageCode:   data.BRSpcfcTaxBasePercentageCode,
			BRSpcfcTxGrpDynTaxExceptions:   data.BRSpcfcTxGrpDynTaxExceptions,
			CustomerTaxClassification1:     data.CustomerTaxClassification1,
			CustomerTaxClassification2:     data.CustomerTaxClassification2,
			CustomerTaxClassification3:     data.CustomerTaxClassification3,
			CustomerTaxClassification4:     data.CustomerTaxClassification4,
			ProductTaxClassification1:      data.ProductTaxClassification1,
			ProductTaxClassification2:      data.ProductTaxClassification2,
			ProductTaxClassification3:      data.ProductTaxClassification3,
			ProductTaxClassification4:      data.ProductTaxClassification4,
			TradingContract:                data.TradingContract,
			TradingContractItem:            data.TradingContractItem,
			TaxJurisdiction:                data.TaxJurisdiction,
			BRSpcfcTaxDepartureRegion:      data.BRSpcfcTaxDepartureRegion,
			BRSpcfcTaxDestinationRegion:    data.BRSpcfcTaxDestinationRegion,
			MainItemMaterialPricingGroup:   data.MainItemMaterialPricingGroup,
			MainItemPricingRefMaterial:     data.MainItemPricingRefMaterial,
			VariantCondition:               data.VariantCondition,
			ValueAddedServiceChargeCode:    data.ValueAddedServiceChargeCode,
			SDDocument:                     data.SDDocument,
			ReferenceSDDocument:            data.ReferenceSDDocument,
			ReferenceSDDocumentItem:        data.ReferenceSDDocumentItem,
			SalesOffice:                    data.SalesOffice,
			SalesGroup:                     data.SalesGroup,
			SalesOrganization:              data.SalesOrganization,
			SalesOrderSalesOrganization:    data.SalesOrderSalesOrganization,
			OrderQuantityUnit:              data.OrderQuantityUnit,
			DistributionChannel:            data.DistributionChannel,
			TransactionCurrency:            data.TransactionCurrency,
			WBSElementInternalID:           data.WBSElementInternalID,
			WBSElementExternalID:           data.WBSElementExternalID,
			WorkPackage:                    data.WorkPackage,
			Plant:                          data.Plant,
			PlantRegion:                    data.PlantRegion,
			WorkItem:                       data.WorkItem,
			ConditionProcessingStatus:      data.ConditionProcessingStatus,
			PricingDate:                    data.PricingDate,
			ConditionIsExclusive:           data.ConditionIsExclusive,
			ConditionScaleBasisValue:       data.ConditionScaleBasisValue,
			TaxCode:                        data.TaxCode,
			ServiceDocument:                data.ServiceDocument,
			ServiceDocumentItem:            data.ServiceDocumentItem,
			TimeSheetOvertimeCategory:      data.TimeSheetOvertimeCategory,
			SalesSDDocumentCategory:        data.SalesSDDocumentCategory,
			ReturnReason:                   data.ReturnReason,
			ProductHierarchyNode:           data.ProductHierarchyNode,
			CustomerConditionGroup:         data.CustomerConditionGroup,
			ShippingType:                   data.ShippingType,
			SubscriptionContractDuration:   data.SubscriptionContractDuration,
			SubscrpnContrDurationUnit:      data.SubscrpnContrDurationUnit,
			SubscriptionContractLockReason: data.SubscriptionContractLockReason,
			CrsCtlgMappgPriceVersionNumber: data.CrsCtlgMappgPriceVersionNumber,
			PRAContract:                    data.PRAContract,
			Well:                           data.Well,
			WellCompletion:                 data.WellCompletion,
			MeasurementPoint:               data.MeasurementPoint,
			PricingFormulaNumber:           data.PricingFormulaNumber,
			ETag:                           data.ETag,
			ToSlsPrcgConditionRecord:       data.ToSlsPrcgConditionRecord.Deferred.URI,
		})
	}

	return slsPrcgCndnRecdValidity, nil
}

func ConvertToToSlsPrcgConditionRecord(raw []byte, l *logger.Logger) (*ToSlsPrcgConditionRecord, error) {
	pm := &responses.ToSlsPrcgConditionRecord{}

	err := json.Unmarshal(raw, &pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToSlsPrcgConditionRecord. unmarshal error: %w", err)
	}

	return &ToSlsPrcgConditionRecord{
		ConditionRecord:              pm.D.ConditionRecord,
		ConditionSequentialNumber:    pm.D.ConditionSequentialNumber,
		ConditionTable:               pm.D.ConditionTable,
		ConditionApplication:         pm.D.ConditionApplication,
		ConditionType:                pm.D.ConditionType,
		ConditionValidityEndDate:     pm.D.ConditionValidityEndDate,
		ConditionValidityStartDate:   pm.D.ConditionValidityStartDate,
		CreationDate:                 pm.D.CreationDate,
		PricingScaleType:             pm.D.PricingScaleType,
		PricingScaleBasis:            pm.D.PricingScaleBasis,
		ConditionScaleQuantity:       pm.D.ConditionScaleQuantity,
		ConditionScaleQuantityUnit:   pm.D.ConditionScaleQuantityUnit,
		ConditionScaleAmount:         pm.D.ConditionScaleAmount,
		ConditionScaleAmountCurrency: pm.D.ConditionScaleAmountCurrency,
		ConditionCalculationType:     pm.D.ConditionCalculationType,
		ConditionRateValue:           pm.D.ConditionRateValue,
		ConditionRateValueUnit:       pm.D.ConditionRateValueUnit,
		ConditionQuantity:            pm.D.ConditionQuantity,
		ConditionQuantityUnit:        pm.D.ConditionQuantityUnit,
		BaseUnit:                     pm.D.BaseUnit,
		ConditionIsDeleted:           pm.D.ConditionIsDeleted,
		PaymentTerms:                 pm.D.PaymentTerms,
		IncrementalScale:             pm.D.IncrementalScale,
		PricingScaleLine:             pm.D.PricingScaleLine,
		ConditionReleaseStatus:       pm.D.ConditionReleaseStatus,
	}, nil
}
