package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey       string      `json:"connection_key"`
	RedisKey            string      `json:"redis_key"`
	Filepath            string      `json:"filepath"`
	APIStatusCode       int         `json:"api_status_code"`
	RuntimeSessionID    string      `json:"runtime_session_id"`
	BusinessPartnerID   *int        `json:"business_partner"`
	ServiceLabel        string      `json:"service_label"`
	APIType             string      `json:"api_type"`
	Message             interface{} `json:"message"`
	APISchema           string      `json:"api_schema"`
	Accepter            []string    `json:"accepter"`
	Deleted             bool        `json:"deleted"`
	SQLUpdateResult     *bool       `json:"sql_update_result"`
	SQLUpdateError      string      `json:"sql_update_error"`
	SubfuncResult       *bool       `json:"subfunc_result"`
	SubfuncError        string      `json:"subfunc_error"`
	ExconfResult        *bool       `json:"exconf_result"`
	ExconfError         string      `json:"exconf_error"`
	APIProcessingResult *bool       `json:"api_processing_result"`
	APIProcessingError  string      `json:"api_processing_error"`
}

type Message struct {
	Header             	*Header	`json:"Header"`
	Item            	*[]Item	`json:"Item"`
}

type Header struct {
	Questionnaire			int		`json:"Questionnaire"`
	QuestionnaireOwner		string	`json:"QuestionnaireOwner"`
	QuestionnaireType		string	`json:"QuestionnaireType"`
	QuestionnaireTemplate	string	`json:"QuestionnaireTemplate"`
	QuestionnaireDate		string	`json:"QuestionnaireDate"`
	QuestionnaireTime		string	`json:"QuestionnaireTime"`
	Respondent				int		`json:"Respondent"`
	QuestionnaireObjectType	string	`json:"QuestionnaireObjectType"`
	QuestionnaireObject		int		`json:"QuestionnaireObject"`
	CreationDate			string	`json:"CreationDate"`
	CreationTime			string	`json:"CreationTime"`
	IsMarkedForDeletion		*bool	`json:"IsMarkedForDeletion"`
}

type Item struct {
	Questionnaire					int		`json:"Questionnaire"`
	QuestionnaireItem				int		`json:"QuestionnaireItem"`
	QuestionnaireItemDescription	string	`json:"QuestionnaireItemDescription"`
	QuestionnaireItemFormType		string	`json:"QuestionnaireItemFormType"`
	QuestionnaireItemReplyType		string	`json:"QuestionnaireItemReplyType"`
	QuestionnaireItemReplyByYesNo	*bool	`json:"QuestionnaireItemReplyByYesNo"`
	QuestionnaireItemReplyByNumber	*int	`json:"QuestionnaireItemReplyByNumber"`
	QuestionnaireItemReplyByText	*string	`json:"QuestionnaireItemReplyByText"`
	CreationDate					string	`json:"CreationDate"`
	CreationTime					string	`json:"CreationTime"`
	IsMarkedForDeletion				*bool	`json:"IsMarkedForDeletion"`
}
