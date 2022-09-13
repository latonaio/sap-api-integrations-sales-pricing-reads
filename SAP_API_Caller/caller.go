package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	sap_api_output_formatter "sap-api-integrations-sales-pricing-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

type SAPAPICaller struct {
	baseURL         string
	sapClientNumber string
	requestClient   *sap_api_request_client_header_setup.SAPRequestClient
	log             *logger.Logger
}

func NewSAPAPICaller(baseUrl, sapClientNumber string, requestClient *sap_api_request_client_header_setup.SAPRequestClient, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:         baseUrl,
		requestClient:   requestClient,
		sapClientNumber: sapClientNumber,
		log:             l,
	}
}

func (c *SAPAPICaller) AsyncGetSalesPricingCondition(material, distributionChannel, customer, salesOrganization string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "MaterialDistChannel":
			func() {
				c.MaterialDistChannel(material, distributionChannel)
				wg.Done()
			}()
		case "MaterialDistChannelCustomer":
			func() {
				c.MaterialDistChannelCustomer(material, distributionChannel, customer)
				wg.Done()
			}()
		case "MaterialSalesOrgDistChannel":
			func() {
				c.MaterialSalesOrgDistChannel(material, salesOrganization, distributionChannel)
				wg.Done()
			}()
		case "MaterialSalesOrgDistChannelCustomer":
			func() {
				c.MaterialSalesOrgDistChannelCustomer(material, salesOrganization, distributionChannel, customer)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) MaterialDistChannel(material, distributionChannel string) {
	data, err := c.callSalesPricingSrvAPIRequirementMaterialDistChannel("A_SlsPrcgCndnRecdValidity", material, distributionChannel)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)

	conditionRecordData, err := c.callToConditionRecord(data[0].ToConditionRecord)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(conditionRecordData)
}

func (c *SAPAPICaller) callSalesPricingSrvAPIRequirementMaterialDistChannel(api, material, distributionChannel string) ([]sap_api_output_formatter.PricingConditionValidity, error) {
	url := strings.Join([]string{c.baseURL, "API_SLSPRICINGCONDITIONRECORD_SRV", api}, "/")
	param := c.getQueryWithMaterialDistChannel(map[string]string{}, material, distributionChannel)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToPricingConditionValidity(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToConditionRecord(url string) (*sap_api_output_formatter.ToConditionRecord, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToConditionRecord(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) MaterialDistChannelCustomer(material, distributionChannel, customer string) {
	data, err := c.callSalesPricingConditionSrvAPIRequirementMaterialDistChannelCustomer("A_SlsPrcgCndnRecdValidity", material, distributionChannel, customer)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)

	conditionRecordData, err := c.callToConditionRecord(data[0].ToConditionRecord)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(conditionRecordData)

}

func (c *SAPAPICaller) callSalesPricingConditionSrvAPIRequirementMaterialDistChannelCustomer(api, material, distributionChannel, customer string) ([]sap_api_output_formatter.PricingConditionValidity, error) {
	url := strings.Join([]string{c.baseURL, "API_SLSPRICINGCONDITIONRECORD_SRV", api}, "/")

	param := c.getQueryWithMaterialDistChannelCustomer(map[string]string{}, material, distributionChannel, customer)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToPricingConditionValidity(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) MaterialSalesOrgDistChannel(material, salesOrganization, distributionChannel string) {
	data, err := c.callSalesPricingConditionSrvAPIRequirementMaterialSalesOrgDistChannel("A_SlsPrcgCndnRecdValidity", material, salesOrganization, distributionChannel)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)

	conditionRecordData, err := c.callToConditionRecord(data[0].ToConditionRecord)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(conditionRecordData)

}

func (c *SAPAPICaller) callSalesPricingConditionSrvAPIRequirementMaterialSalesOrgDistChannel(api, material, salesOrganization, distributionChannel string) ([]sap_api_output_formatter.PricingConditionValidity, error) {
	url := strings.Join([]string{c.baseURL, "API_SLSPRICINGCONDITIONRECORD_SRV", api}, "/")

	param := c.getQueryWithMaterialSalesOrgDistChannel(map[string]string{}, material, salesOrganization, distributionChannel)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToPricingConditionValidity(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) MaterialSalesOrgDistChannelCustomer(material, salesOrganization, distributionChannel, customer string) {
	data, err := c.callSalesPricingConditionSrvAPIRequirementMaterialSalesOrgDistChannelCustomer("A_SlsPrcgCndnRecdValidity", material, salesOrganization, distributionChannel, customer)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)

	conditionRecordData, err := c.callToConditionRecord(data[0].ToConditionRecord)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(conditionRecordData)

}

func (c *SAPAPICaller) callSalesPricingConditionSrvAPIRequirementMaterialSalesOrgDistChannelCustomer(api, material, salesOrganization, distributionChannel, customer string) ([]sap_api_output_formatter.PricingConditionValidity, error) {
	url := strings.Join([]string{c.baseURL, "API_SLSPRICINGCONDITIONRECORD_SRV", api}, "/")

	param := c.getQueryWithMaterialSalesOrgDistChannelCustomer(map[string]string{}, material, salesOrganization, distributionChannel, customer)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToPricingConditionValidity(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) getQueryWithMaterialDistChannel(params map[string]string, material, distributionChannel string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Material eq '%s' and DistributionChannel eq '%s'", material, distributionChannel)
	return params
}

func (c *SAPAPICaller) getQueryWithMaterialDistChannelCustomer(params map[string]string, material, distributionChannel, customer string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Material eq '%s' and DistributionChannel eq '%s' and Customer eq '%s'", material, distributionChannel, customer)
	return params
}

func (c *SAPAPICaller) getQueryWithMaterialSalesOrgDistChannel(params map[string]string, material, salesOrganization, distributionChannel string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Material eq '%s' and SalesOrganization eq '%s' and DistributionChannel eq '%s'", material, salesOrganization, distributionChannel)
	return params
}

func (c *SAPAPICaller) getQueryWithMaterialSalesOrgDistChannelCustomer(params map[string]string, material, salesOrganization, distributionChannel, customer string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Material eq '%s' and SalesOrganization eq '%s' and DistributionChannel eq '%s' and Customer eq '%s'", material, salesOrganization, distributionChannel, customer)
	return params
}
