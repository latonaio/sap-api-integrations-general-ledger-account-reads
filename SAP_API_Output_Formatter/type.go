package sap_api_output_formatter

type GLAccount struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Product       string `json:"Product"`
	APISchema     string `json:"api_schema"`
	GLAccount     string `json:"gl_account"`
	Deleted       string `json:"deleted"`
}

type ChartOfAccounts struct {
	ChartOfAccounts				string `json:"ChartOfAccounts"`
	GLAccount					string `json:"GLAccount"`
	IsBalanceSheetAccount 		bool   `json:"IsBalanceSheetAccount"`
	GLAccountGroup				string `json:"GLAccountGroup"`
	CorporateGroupAccount		string `json:"CorporateGroupAccount"`
	ProfitLossAccountType		string `json:"ProfitLossAccountType"`
    SampleGLAccount				string `json:"SampleGLAccount"`
	AccountIsMarkedForDeletion 	bool   `json:"AccountIsMarkedForDeletion"`
	AccountIsBlockedForCreation bool   `json:"AccountIsBlockedForCreation"`
	AccountIsBlockedForPosting  bool   `json:"AccountIsBlockedForPosting"`
	AccountIsBlockedForPlanning bool   `json:"AccountIsBlockedForPlanning"`
	PartnerCompany				string `json:"PartnerCompany"`
	FunctionalArea				string `json:"FunctionalArea"`
	CreationDate				string `json:"CreationDate"`
	CreatedByUser				string `json:"CreatedByUser"`
	LastChangeDateTime			string `json:"LastChangeDateTime"`
	GLAccountType				string `json:"GLAccountType"`
	GLAccountExternal			string `json:"GLAccountExternal"`
	IsProfitLossAccount			bool   `json:"IsProfitLossAccount"`
	ToText						string `json:"to_Text"`
}

type Text struct {
	ChartOfAccounts				string `json:"ChartOfAccounts"`
	GLAccount					string `json:"GLAccount"`
	Language					string `json:"Language"`
	GLAccountName				string `json:"GLAccountName"`
	GLAccountLongName			string `json:"GLAccountLongName"`
	LastChangeDateTime			string `json:"LastChangeDateTime"`
}
