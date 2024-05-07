package dpfm_api_processing_formatter

type HeaderUpdates struct {
	Questionnaire			int	`json:"Questionnaire"`
}

type ItemUpdates struct {
	Questionnaire			int	`json:"Questionnaire"`
	QuestionnaireItem		int	`json:"QuestionnaireItem"`
}
