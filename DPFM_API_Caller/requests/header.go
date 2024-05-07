package requests

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
