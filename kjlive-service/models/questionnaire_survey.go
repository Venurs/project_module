package models

import (
	"kjlive-service/redis"
    "fmt"
)

type QuestionnaireSerializer struct {
	TemplateId   int    `orm:"column(template_id)"`
	TemplateName string `orm:"column(tempalte_name)"`
	QuestionId   int    `orm:"column(question_id)"`
	Question     string `orm:"column(question)"`
	QuestionType int    `orm:"column(question_type)"`
	IsRequuired  bool   `orm:"column(is_required)"`
	OptionalId   int    `orm:"column(optional_id)"`
	Optional     string `orm:"column(optional)"`
}

type OptionalSerializer struct {
	Id         int    `json:"id"`
	OptionName string `json:"option_name"`
}

type QuestionSerializer struct {
	Id           int                  `json:"id"`
	QuestionName string               `json:"question_name"`
	QuestionType int                  `json:"question_type"`
	IsRequired   bool                 `json:"is_required"`
	Options      []OptionalSerializer `json:"options"`
}

type TemplateSerializer struct {
	Id        int                       `json:"id"`
	Name      string                    `json:"name"`
	Questions []QuestionnaireSerializer `json:"questions"`
}

func GetUserQuestionnaireSurveyStatus(userid int, templateid int8) (IsSurvey int) {
	res := redis.GetSurveyStatus("USER-"+fmt.Sprintf("%d", userid)  +"-QUESTIONNAIRE-"+fmt.Sprintf("%d", templateid))
	return res
}

//func GetSurveyDetail(templateid int8) {
//	var questionnaires []QuestionnaireSerializer
//
//	o := orm.NewOrm()
//
//	sql := `
//		SELECT T0.id template_id,
//			T0.name template_name,
//         	T1.id question_id,
//         	T1.question_name question,
//         	T1.question_type question_type,
//         	T1.is_required is_required,
//         	T2.id optional_id,
//         	T2.option_name optional
//		FROM questionnaire_question T1
//				RIGHT JOIN questionnaire_template T0 on T1.template_name_id= T0.id
//				LEFT JOIN questionnaire_options T2 ON T1.id= T2.question_id
//   		WHERE T0.id= ?
//	`
//	o.Raw(sql, templateid).QueryRows(&questionnaires)
//
//	//tempQuestionsId := -1
//	//var tempQuestion QuestionSerializer
//	//for _, item := range (questionnaires) {
//	//	if tempQuestionsId != item.QuestionId {
//	//
//	//		if tempQuestionsId == -1 {
//	//			tempQuestion = QuestionSerializer{
//	//				Id: item.QuestionId,
//	//				QuestionType: item.QuestionType,
//	//				QuestionName: item.Question,
//	//				IsRequired: item.IsRequuired}
//	//		} else {
//	//			tempQuestion = QuestionSerializer{
//	//				Id: item.QuestionId,
//	//				QuestionType: item.QuestionType,
//	//				QuestionName: item.Question,
//	//				IsRequired: item.IsRequuired}
//	//		}
//	//		tempQuestionsId = item.QuestionId
//	//	} else {
//	//	}
//	//}
//
//}
