package sap_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	SalesOrder    struct {
		SalesOrder     string `json:"document_no"`
		Plant          string `json:"deliver_to"`
		OrderQuantity  string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		NetPriceAmount string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema                string   `json:"api_schema"`
	Accepter                 []string `json:"accepter"`
	MaterialCode             string   `json:"material_code"`
	Supplier                 string   `json:"plant/supplier"`
	Stock                    string   `json:"stock"`
	SalesOrderType           string   `json:"document_type"`
	SONumber                 string   `json:"document_no"`
	ScheduleLineDeliveryDate string   `json:"planned_date"`
	ValidatedDate            string   `json:"validated_date"`
	Deleted                  bool     `json:"deleted"`
}

type SDC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	GLAccount   struct {
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
		Text struct {
			Language					string `json:"Language"`
			GLAccountName				string `json:"GLAccountName"`
			GLAccountLongName			string `json:"GLAccountLongName"`
			LastChangeDateTime			string `json:"LastChangeDateTime"`
		} `json:"Text"`
	} `json:"GLAccount"`
	APISchema 	    string   `json:"api_schema"`
	Accepter  	    []string `json:"accepter"`
	GLAccountCode	string   `json:"gl_account_code"`
	Deleted		    bool     `json:"deleted"`
	}
