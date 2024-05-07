package requests

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
