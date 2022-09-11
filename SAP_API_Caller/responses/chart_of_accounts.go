package responses

type ChartOfAccounts struct {
	D struct {
		Count   string `json:"__count"`
		Results []struct {
			ChartOfAccounts             string `json:"ChartOfAccounts"`
			GLAccount                   string `json:"GLAccount"`
			IsBalanceSheetAccount       bool   `json:"IsBalanceSheetAccount"`
			GLAccountGroup              string `json:"GLAccountGroup"`
			CorporateGroupAccount       string `json:"CorporateGroupAccount"`
			ProfitLossAccountType       string `json:"ProfitLossAccountType"`
			SampleGLAccount             string `json:"SampleGLAccount"`
			AccountIsMarkedForDeletion  bool   `json:"AccountIsMarkedForDeletion"`
			AccountIsBlockedForCreation bool   `json:"AccountIsBlockedForCreation"`
			AccountIsBlockedForPosting  bool   `json:"AccountIsBlockedForPosting"`
			AccountIsBlockedForPlanning bool   `json:"AccountIsBlockedForPlanning"`
			PartnerCompany              string `json:"PartnerCompany"`
			FunctionalArea              string `json:"FunctionalArea"`
			CreationDate                string `json:"CreationDate"`
			CreatedByUser               string `json:"CreatedByUser"`
			LastChangeDateTime          string `json:"LastChangeDateTime"`
			GLAccountType               string `json:"GLAccountType"`
			GLAccountExternal           string `json:"GLAccountExternal"`
			IsProfitLossAccount         bool   `json:"IsProfitLossAccount"`
			ToText struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_Text"`
		} `json:"results"`
	} `json:"d"`
}
