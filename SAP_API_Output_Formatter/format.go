package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-general-leader-account-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToChartOfAccounts(raw []byte, l *logger.Logger) ([]ChartOfAccounts, error) {
	pm := &responses.ChartOfAccounts{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ChartOfAccounts. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	chartOfAccounts := make([]ChartOfAccounts, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		chartOfAccounts = append(chartOfAccounts, ChartOfAccounts{
			ChartOfAccounts:				data.ChartOfAccounts,
			GLAccount:						data.GLAccount,
			IsBalanceSheetAccount:			data.IsBalanceSheetAccount,
			GLAccountGroup:					data.GLAccountGroup,
			CorporateGroupAccount:			data.CorporateGroupAccount,
			ProfitLossAccountType:			data.ProfitLossAccountType,
    		SampleGLAccount:				data.SampleGLAccount,
			AccountIsMarkedForDeletion: 	data.AccountIsMarkedForDeletion,
			AccountIsBlockedForCreation:	data.AccountIsBlockedForCreation,
			AccountIsBlockedForPosting:  	data.AccountIsBlockedForPosting,
			AccountIsBlockedForPlanning: 	data.AccountIsBlockedForPlanning,
			PartnerCompany:					data.PartnerCompany,
			FunctionalArea:					data.FunctionalArea,
			CreationDate:					data.CreationDate,
			CreatedByUser:					data.CreatedByUser,
			LastChangeDateTime:				data.LastChangeDateTime,
			GLAccountType:					data.GLAccountType,
			GLAccountExternal:				data.GLAccountExternal,
			IsProfitLossAccount:			data.IsProfitLossAccount,
			ToText:          				data.ToText.Deferred.URI,
		})
	}

	return chartOfAccounts, nil
}

func ConvertToText(raw []byte, l *logger.Logger) ([]Text, error) {
	pm := &responses.Text{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Text. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	text := make([]Text, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		text = append(text, Text{
			ChartOfAccounts:		data.ChartOfAccounts,
			GLAccount:				data.GLAccount,
			Language:				data.Language,
			GLAccountName:			data.GLAccountName,
			GLAccountLongName:		data.GLAccountLongName,
			LastChangeDateTime:		data.LastChangeDateTime,
		})
	}

	return text, nil
}
