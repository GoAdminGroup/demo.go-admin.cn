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
					Options:  types.FieldOptions{},
				},
				{
					Field:    "age",
					TypeName: "int",
					Head:     "年龄",
					Default:  "11",
					Editable: true,
					FormType: form.Number,
					Value:    "11",
					Options:  types.FieldOptions{},
				},
				{
					Field:    "homepage",
					TypeName: db.Varchar,
					Head:     "主页",
					Default:  "http://google.com",
					Editable: true,
					FormType: form.Url,
					Value:    "http://google.com",
					Options:  types.FieldOptions{},
				},
				{
					Field:    "email",
					TypeName: db.Varchar,
					Head:     "邮箱",
					Default:  "xxxx@xxx.com",
					Editable: true,
					FormType: form.Email,
					Value:    "",
					Options:  types.FieldOptions{},
				},
				{
					Field:    "birthday",
					TypeName: db.Varchar,
					Head:     "生日",
					Default:  "2010-09-05",
					Editable: true,
					FormType: form.Datetime,
					Value:    "2010-09-05",
					Options:  types.FieldOptions{},
				},
				{
					Field:    "password",
					TypeName: db.Varchar,
					Head:     "密码",
					Default:  "",
					Editable: true,
					FormType: form.Password,
					Value:    "",
					Options:  types.FieldOptions{},
				},
				{
					Field:    "ip",
					TypeName: db.Varchar,
					Head:     "Ip",
					Default:  "",
					Editable: true,
					FormType: form.Ip,
					Value:    "",
					Options:  types.FieldOptions{},
				},
				{
					Field:    "currency",
					TypeName: db.Int,
					Head:     "金额",
					Default:  "",
					Editable: true,
					FormType: form.Currency,
					Value:    "",
					Options:  types.FieldOptions{},
				},
				{
					Field:    "content",
					TypeName: db.Text,
					Head:     "内容",
					Default:  "",
					Editable: true,
					FormType: form.RichText,
					Value:    "",
					Options:  types.FieldOptions{},
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
					Options: types.FieldOptions{
						{Text: "website", Value: "0"},
						{Text: "website", Value: "1"},
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
					Options: types.FieldOptions{
						{Text: "苹果", Value: "apple"},
						{Text: "香蕉", Value: "banana"},
						{Text: "西瓜", Value: "watermelon"},
						{Text: "梨", Value: "pear"},
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
					Options: types.FieldOptions{
						{Text: "男生", Value: "0"},
						{Text: "女生", Value: "1"},
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
					Options: types.FieldOptions{
						{Text: "啤酒", Value: "beer"},
						{Text: "果汁", Value: "juice"},
						{Text: "白开水", Value: "water"},
						{Text: "红牛", Value: "red bull"},
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
					Options: types.FieldOptions{
						{Text: "两年", Value: "0"},
						{Text: "三年", Value: "1"},
						{Text: "四年", Value: "2"},
						{Text: "五年", Value: "3"},
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
