package pages

import (
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/language"
	template2 "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetForm1Content() (types.Panel, error) {

	components := template2.Get(config.Get().Theme)

	col1 := components.Col().SetSize(map[string]string{"md": "2"}).GetContent()
	btn1 := components.Button().SetType("submit").
		SetContent(language.GetFromHtml("Save")).
		SetThemePrimary().
		SetOrientationRight().
		SetLoadingText(`<i class='fa fa-spinner fa-spin '></i> Save`).
		GetContent()
	btn2 := components.Button().SetType("reset").
		SetContent(language.GetFromHtml("Reset")).
		SetThemeWarning().
		SetOrientationLeft().
		GetContent()
	col2 := components.Col().SetSize(map[string]string{"md": "8"}).
		SetContent(btn1 + btn2).GetContent()

	aform := components.Form().
		SetTabHeaders([]string{"input", "select"}).
		SetTabContents([][]types.FormField{
			{
				{
					Field:    "name",
					TypeName: db.Varchar,
					Head:     "名字",
					Default:  "张三",
					Editable: true,
					FormType: form.Text,
					Value:    "张三",
					Options:  []map[string]string{},
				},
				{
					Field:    "age",
					TypeName: "int",
					Head:     "年龄",
					Default:  "11",
					Editable: true,
					FormType: form.Number,
					Value:    "11",
					Options:  []map[string]string{},
				},
				{
					Field:    "homepage",
					TypeName: db.Varchar,
					Head:     "主页",
					Default:  "http://google.com",
					Editable: true,
					FormType: form.Url,
					Value:    "http://google.com",
					Options:  []map[string]string{},
				},
				{
					Field:    "email",
					TypeName: db.Varchar,
					Head:     "邮箱",
					Default:  "xxxx@xxx.com",
					Editable: true,
					FormType: form.Email,
					Value:    "",
					Options:  []map[string]string{},
				},
				{
					Field:    "birthday",
					TypeName: db.Varchar,
					Head:     "生日",
					Default:  "2010-09-05",
					Editable: true,
					FormType: form.Datetime,
					Value:    "2010-09-05",
					Options:  []map[string]string{},
				},
				{
					Field:    "password",
					TypeName: db.Varchar,
					Head:     "密码",
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
					Head:     "金额",
					Default:  "",
					Editable: true,
					FormType: form.Currency,
					Value:    "",
					Options:  []map[string]string{},
				},
				{
					Field:    "content",
					TypeName: db.Text,
					Head:     "内容",
					Default:  "",
					Editable: true,
					FormType: form.RichText,
					Value:    "",
					Options:  []map[string]string{},
				},
			},
			{
				{
					Field:    "website",
					TypeName: db.Tinyint,
					Head:     "站点开关",
					HelpMsg:  "站点关闭后将不能访问，后台可正常登录",
					Default:  "",
					Editable: true,
					FormType: form.Switch,
					Value:    "",
					Options: []map[string]string{
						{
							"field":    "website",
							"value":    "0",
							"selected": "checked",
						},
						{
							"field":    "website",
							"value":    "1",
							"selected": "",
						},
					},
				},
				{
					Field:    "fruit",
					TypeName: db.Varchar,
					Head:     "水果",
					Default:  "",
					Editable: true,
					FormType: form.SelectBox,
					Value:    "",
					Options: []map[string]string{
						{
							"field": "苹果",
							"value": "apple",
						}, {
							"field": "香蕉",
							"value": "banana",
						}, {
							"field": "西瓜",
							"value": "watermelon",
						}, {
							"field": "梨",
							"value": "pear",
						},
					},
					FieldDisplay: types.FieldDisplay{
						Display: func(value types.FieldModel) interface{} {
							return []string{"梨"}
						},
					},
				},
				{
					Field:    "gender",
					TypeName: db.Tinyint,
					Head:     "性别",
					Default:  "0",
					Editable: true,
					FormType: form.Radio,
					Value:    "",
					Options: []map[string]string{
						{
							"field":    "gender",
							"label":    "男生",
							"value":    "0",
							"selected": "true",
						},
						{
							"field":    "gender",
							"label":    "女生",
							"value":    "1",
							"selected": "false",
						},
					},
				},
				{
					Field:    "drink",
					TypeName: db.Varchar,
					Head:     "饮料",
					Default:  "beer",
					Editable: true,
					FormType: form.Select,
					Value:    "",
					Options: []map[string]string{
						{
							"field": "啤酒",
							"value": "beer",
						}, {
							"field": "果汁",
							"value": "juice",
						}, {
							"field": "白开水",
							"value": "water",
						}, {
							"field": "红牛",
							"value": "red bull",
						},
					},
				},
				{
					Field:    "experience",
					TypeName: db.Tinyint,
					Head:     "工作经验",
					Default:  "",
					Editable: true,
					FormType: form.SelectSingle,
					Value:    "",
					Options: []map[string]string{
						{
							"field": "两年",
							"value": "0",
						}, {
							"field": "三年",
							"value": "1",
						}, {
							"field": "四年",
							"value": "2",
						}, {
							"field": "五年",
							"value": "3",
						},
					},
				},
			},
		}).
		SetPrefix(config.Get().PrefixFixSlash()).
		SetUrl("/").
		SetTitle("Form").
		SetInfoUrl("/admin").
		SetOperationFooter(col1 + col2)

	return types.Panel{
		Content: components.Box().
			SetHeader(aform.GetBoxHeader()).
			WithHeadBorder().
			SetBody(aform.GetContent()).
			GetContent(),
		Title:       "表单",
		Description: "表单例子",
	}, nil
}
