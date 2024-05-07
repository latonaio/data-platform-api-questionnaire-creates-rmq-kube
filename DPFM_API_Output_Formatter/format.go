package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "data-platform-api-questionnaire-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "data-platform-api-questionnaire-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"data-platform-api-questionnaire-creates-rmq-kube/sub_func_complementer"
	"encoding/json"

	"golang.org/x/xerrors"
)

func ConvertToHeaderCreates(subfuncSDC *sub_func_complementer.SDC) (*Header, error) {
	data := subfuncSDC.Message.Header

	header, err := TypeConverter[*Header](data)
	if err != nil {
		return nil, err
	}

	return header, nil
}

func ConvertToItemCreates(subfuncSDC *sub_func_complementer.SDC) (*[]Item, error) {
	items := make([]Item, 0)

	for _, data := range *subfuncSDC.Message.Item {
		item, err := TypeConverter[*Item](data)
		if err != nil {
			return nil, err
		}

		items = append(items, *item)
	}

	return &items, nil
}

func ConvertToHeaderUpdates(headerData dpfm_api_input_reader.Header) (*Header, error) {
	data := headerData

	header, err := TypeConverter[*Header](data)
	if err != nil {
		return nil, err
	}

	return header, nil
}

func ConvertToItemUpdates(itemUpdates *[]dpfm_api_processing_formatter.ItemUpdates) (*[]Item, error) {
	items := make([]Item, 0)

	for _, data := range *itemUpdates {
		item, err := TypeConverter[*Item](data)
		if err != nil {
			return nil, err
		}

		items = append(items, *item)
	}

	return &items, nil
}

func ConvertToHeader(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
) *sub_func_complementer.SDC {
	subfuncSDC.Message.Header = &sub_func_complementer.Header{
		Questionnaire:				*input.Header.Questionnaire,
		QuestionnaireOwner:			input.Header.QuestionnaireOwner,
		QuestionnaireType:			input.Header.QuestionnaireType,
		QuestionnaireTemplate:		input.Header.QuestionnaireTemplate,
		QuestionnaireDate:			input.Header.QuestionnaireDate,
		QuestionnaireTime:			input.Header.QuestionnaireTime,
		Respondent:					input.Header.Respondent,
		QuestionnaireObjectType:	input.Header.QuestionnaireObjectType,
		QuestionnaireObject:		input.Header.QuestionnaireObject,
		CreationDate:				input.Header.CreationDate,
		CreationTime:				input.Header.CreationTime,
		IsMarkedForDeletion:		input.Header.IsMarkedForDeletion,
	}

	return subfuncSDC
}

func ConvertToItem(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
) *sub_func_complementer.SDC {
	var items []sub_func_complementer.Item

	items = append(
		items,
		sub_func_complementer.Item{
			Questionnaire:					*input.Header.Questionnaire,
			QuestionnaireItem:				input.Header.Item[0].QuestionnaireItem,
			QuestionnaireItemDescription:	input.Header.Item[0].QuestionnaireItemDescription,
			QuestionnaireItemFormType:		input.Header.Item[0].QuestionnaireItemFormType,
			QuestionnaireItemReplyType:		input.Header.Item[0].QuestionnaireItemReplyType,
			QuestionnaireItemReplyByYesNo:	input.Header.Item[0].QuestionnaireItemReplyByYesNo,
			QuestionnaireItemReplyByNumber:	input.Header.Item[0].QuestionnaireItemReplyByNumber,
			QuestionnaireItemReplyByText:	input.Header.Item[0].QuestionnaireItemReplyByText,
			CreationDate:					input.Header.Item[0].CreationDate,
			CreationTime:					input.Header.Item[0].CreationTime,
			IsMarkedForDeletion:			input.Header.Item[0].IsMarkedForDeletion,
		},
	)

	subfuncSDC.Message.Item = &items

	return subfuncSDC
}

func TypeConverter[T any](data interface{}) (T, error) {
	var dist T
	b, err := json.Marshal(data)
	if err != nil {
		return dist, xerrors.Errorf("Marshal error: %w", err)
	}
	err = json.Unmarshal(b, &dist)
	if err != nil {
		return dist, xerrors.Errorf("Unmarshal error: %w", err)
	}
	return dist, nil
}
