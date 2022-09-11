package responses

type ToText struct {
	D struct {
		Results []struct {
			ChartOfAccounts    string `json:"ChartOfAccounts"`
			GLAccount          string `json:"GLAccount"`
			Language           string `json:"Language"`
			GLAccountName      string `json:"GLAccountName"`
			GLAccountLongName  string `json:"GLAccountLongName"`
			LastChangeDateTime string `json:"LastChangeDateTime"`
		} `json:"results"`
	} `json:"d"`
}