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

func (c *SAPAPICaller) AsyncGetSalesPricingCondition(conditionRecord string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "SlsPrcgCndnRecdSuplmnt":
			func() {
				c.SlsPrcgCndnRecdSuplmnt(conditionRecord)
				wg.Done()
			}()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) SlsPrcgCndnRecdSuplmnt(conditionRecord string) {
	slsPrcgCndnRecdSuplmntData, err := c.callSlsPrcgCndnRecdSuplmnt("A_SlsPrcgCndnRecdSuplmnt", conditionRecord)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(slsPrcgCndnRecdSuplmntData)
	}

	slsPrcgCndnRecdValidityData, err := c.callToSlsPrcgCndnRecdValidity(slsPrcgCndnRecdSuplmntData[0].ToSlsPrcgCndnRecdValidity)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(slsPrcgCndnRecdValidityData)
	}

	slsPrcgConditionRecordData, err := c.callToConditionRecord(slsPrcgCndnRecdValidityData[0].ToSlsPrcgConditionRecord)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(slsPrcgConditionRecordData)
	}
}

func (c *SAPAPICaller) callSlsPrcgCndnRecdSuplmnt(api, conditionRecord string) ([]sap_api_output_formatter.SlsPrcgCndnRecdSuplmnt, error) {
	url := strings.Join([]string{c.baseURL, "API_SLSPRICINGCONDITIONRECORD_SRV", api}, "/")
	param := c.getQueryWithSlsPrcgCndnRecdSuplmnt(map[string]string{}, conditionRecord)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToSlsPrcgCndnRecdSuplmnt(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToSlsPrcgCndnRecdValidity(url string) ([]sap_api_output_formatter.ToSlsPrcgCndnRecdValidity, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToSlsPrcgCndnRecdValidity(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToConditionRecord(url string) (*sap_api_output_formatter.ToSlsPrcgConditionRecord, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToSlsPrcgConditionRecord(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) getQueryWithSlsPrcgCndnRecdSuplmnt(params map[string]string, conditionRecord string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("ConditionRecord eq '%s'", conditionRecord)
	return params
}
