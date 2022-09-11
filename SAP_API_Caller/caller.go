package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	sap_api_output_formatter "sap-api-integrations-general-leader-account-reads/SAP_API_Output_Formatter"
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

func (c *SAPAPICaller) AsyncGetGLAccount(chartOfAccounts, gLAccount, language, gLAccountName string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "ChartOfAccounts":
			func() {
				c.ChartOfAccounts(chartOfAccounts, gLAccount)
				wg.Done()
			}()
		case "Text":
			func() {
				c.Text(chartOfAccounts, language, gLAccountName)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) ChartOfAccounts(chartOfAccounts, gLAccount string) {
	chartOfAccountsData, err := c.callGLAccountSrvAPIRequirementChartOfAccounts("A_GLAccountInChartOfAccounts", chartOfAccounts, gLAccount)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(chartOfAccountsData)

	textData, err := c.callToText(chartOfAccountsData[0].ToText)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(textData)
}

func (c *SAPAPICaller) callGLAccountSrvAPIRequirementChartOfAccounts(api, chartOfAccounts, gLAccount string) ([]sap_api_output_formatter.ChartOfAccounts, error) {
	url := strings.Join([]string{c.baseURL, "API_GLACCOUNTINCHARTOFACCOUNTS_SRV", api}, "/")
	param := c.getQueryWithChartOfAccounts(map[string]string{}, chartOfAccounts, gLAccount)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToChartOfAccounts(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToText(url string) ([]sap_api_output_formatter.Text, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToText(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Text(chartOfAccounts, language, gLAccountName string) {
	data, err := c.callGLAccountSrvAPIRequirementText("A_GLAccountText", chartOfAccounts, language, gLAccountName)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callGLAccountSrvAPIRequirementText(api, chartOfAccounts, language, gLAccountName string) ([]sap_api_output_formatter.Text, error) {
	url := strings.Join([]string{c.baseURL, "API_GLACCOUNTINCHARTOFACCOUNTS_SRV", api}, "/")

	param := c.getQueryWithText(map[string]string{}, chartOfAccounts, language, gLAccountName)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToText(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}


func (c *SAPAPICaller) getQueryWithChartOfAccounts(params map[string]string, chartOfAccounts, gLAccount string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("ChartOfAccounts eq '%s' and GLAccount eq '%s'", chartOfAccounts, gLAccount)
	return params
}

func (c *SAPAPICaller) getQueryWithText(params map[string]string, chartOfAccounts, language, gLAccountName string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("ChartOfAccounts eq '%s' and Language eq '%s' and substringof('%s', GLAccountName)", chartOfAccounts, language, gLAccountName)
	return params
}
