package pages

import (
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	template2 "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetForm1Content() (types.Panel, error) {

	components := template2.Get(config.Get().Theme)

	aform := components.Form().
		SetTabHeaders([]string{"input", "select"}).
		SetTabContents([][]types.FormField{
			{
				{
					Field:    "name",
					TypeName: db.Varchar,
					Head:     "Name",
					Default:  "jane",
					Editable: true,
					FormType: form.Text,
					Value:    "jane",
					Options:  []map[string]string{},
				},
				{
					Field:    "age",
					TypeName: "int",
					Head:     "Age",
					Default:  "11",
					Editable: true,
					FormType: form.Number,
					Value:    "11",
					Options:  []map[string]string{},
				},
				{
					Field:    "homepage",
					TypeName: db.Varchar,
					Head:     "HomePage",
					Default:  "http://google.com",
					Editable: true,
					FormType: form.Url,
					Value:    "http://google.com",
					Options:  []map[string]string{},
				},
				{
					Field:    "email",
					TypeName: db.Varchar,
					Head:     "Email",
					Default:  "xxxx@xxx.com",
					Editable: true,
					FormType: form.Email,
					Value:    "",
					Options:  []map[string]string{},
				},
				{
					Field:    "birthday",
					TypeName: db.Varchar,
					Head:     "Birthday",
					Default:  "2010-09-05",
					Editable: true,
					FormType: form.Datetime,
					Value:    "2010-09-05",
					Options:  []map[string]string{},
				},
				{
					Field:    "password",
					TypeName: db.Varchar,
					Head:     "Password",
					Default:  "",
					Editable: true,
					FormType: form.Password,
					Value:    "",
					Options:  []map[string]string{},
				},
				{
					Field:    "ip",
					TypeName: db.Varchar,
					Head:     "Ip",
					Default:  "",
					Editable: true,
					FormType: form.Ip,
					Value:    "",
					Options:  []map[string]string{},
				},
				{
					Field:    "currency",
					TypeName: db.Int,
					Head:     "Currency",
					Default:  "",
					Editable: true,
					FormType: form.Currency,
					Value:    "",
					Options:  []map[string]string{},
				},
				{
					Field:    "content",
					TypeName: db.Text,
					Head:     "Content",
					Default:  "",
					Editable: true,
					FormType: form.RichText,
					Value:    "",
					Options:  []map[string]string{},
				},
			},
			{
				{
					Field:    "fruit",
					TypeName: db.Varchar,
					Head:     "Fruit",
					Default:  "",
					Editable: true,
					FormType: form.SelectBox,
					Value:    "",
					Options: []map[string]string{
						{
							"field": "apple",
							"value": "apple",
						}, {
							"field": "banana",
							"value": "banana",
						}, {
							"field": "watermelon",
							"value": "watermelon",
						}, {
							"field": "pear",
							"value": "pear",
						},
					},
					FieldDisplay: types.FieldDisplay{
						Display: func(value types.FieldModel) interface{} {
							return []string{"pear"}
						},
					},
				},
				{
					Field:    "gender",
					TypeName: db.Tinyint,
					Head:     "Gender",
					Default:  "0",
					Editable: true,
					FormType: form.Radio,
					Value:    "",
					Options: []map[string]string{
						{
							"field":    "gender",
							"label":    "male",
							"value":    "0",
							"selected": "true",
						},
						{
							"field":    "gender",
							"label":    "female",
							"value":    "1",
							"selected": "false",
						},
					},
				},
				{
					Field:    "drink",
					TypeName: db.Varchar,
					Head:     "Drink",
					Default:  "",
					Editable: true,
					FormType: form.Select,
					Value:    "",
					Options: []map[string]string{
						{
							"field": "beer",
							"value": "beer",
						}, {
							"field": "juice",
							"value": "juice",
						}, {
							"field": "water",
							"value": "water",
						}, {
							"field": "red bull",
							"value": "red bull",
						},
					},
					FieldDisplay: types.FieldDisplay{
						Display: func(value types.FieldModel) interface{} {
							return []string{"beer"}
						},
					},
				},
				{
					Field:    "experience",
					TypeName: db.Tinyint,
					Head:     "Work experience",
					Default:  "",
					Editable: true,
					FormType: form.SelectSingle,
					Value:    "",
					Options: []map[string]string{
						{
							"field": "two years",
							"value": "0",
						}, {
							"field": "three years",
							"value": "1",
						}, {
							"field": "four years",
							"value": "2",
						}, {
							"field": "five years",
							"value": "3",
						},
					},
					FieldDisplay: types.FieldDisplay{
						Display: func(value types.FieldModel) interface{} {
							return []string{"two years"}
						},
					},
				},
			},
		}).
		SetPrefix(config.Get().PrefixFixSlash()).
		SetUrl("/").
		SetTitle("Form").
		SetToken(auth.TokenHelper.AddToken()).
		SetInfoUrl("/admin").
		GetContent()

	return types.Panel{
		Content:     aform,
		Title:       "表单",
		Description: "表单例子",
	}, nil
}
