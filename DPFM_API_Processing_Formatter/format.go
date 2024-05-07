package dpfm_api_processing_formatter

import (
	dpfm_api_input_reader "data-platform-api-questionnaire-creates-rmq-kube/DPFM_API_Input_Reader"
)

func ConvertToHeaderUpdates(header dpfm_api_input_reader.Header) *HeaderUpdates {
	data := header

	return &HeaderUpdates{
		Questionnaire:			*data.Questionnaire,
	}
}

func ConvertToItemUpdates(header dpfm_api_input_reader.Header, item dpfm_api_input_reader.Item) *ItemUpdates {
	dataHeader := header
	data := item

	return &PartnerUpdates{
		Questionnaire:           *dataHeader.Questionnaire,
		QuestionnaireItem:       data.QuestionnaireItem,
	}
}
