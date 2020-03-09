package pages

import (
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/language"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	template2 "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/icon"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"github.com/gin-gonic/gin"
)

func GetForm1Content(ctx *gin.Context) (types.Panel, error) {

	components := template2.Get(config.Get().Theme)

	col1 := components.Col().GetContent()
	btn1 := components.Button().SetType("submit").
		SetContent(language.GetFromHtml("Save")).
		SetThemePrimary().
		SetOrientationRight().
		SetLoadingText(icon.Icon("fa-spinner fa-spin", 2) + `Save`).
		GetContent()
	btn2 := components.Button().SetType("reset").
		SetContent(language.GetFromHtml("Reset")).
		SetThemeWarning().
		SetOrientationLeft().
		GetContent()
	col2 := components.Col().SetSize(types.SizeMD(8)).
		SetContent(btn1 + btn2).GetContent()

	var panel = types.NewFormPanel()
	panel.AddField("名字", "name", db.Varchar, form.Text)
	panel.AddField("年龄", "age", db.Int, form.Number)
	panel.AddField("主页", "homepage", db.Varchar, form.Url).FieldDefault("http://google.com")
	panel.AddField("邮箱", "email", db.Varchar, form.Email).FieldDefault("xxxx@xxx.com")
	panel.AddField("生日", "birthday", db.Varchar, form.Datetime).FieldDefault("2010-09-05")
	panel.AddField("密码", "password", db.Varchar, form.Password)
	panel.AddField("IP", "ip", db.Varchar, form.Ip)
	panel.AddField("证件", "certificate", db.Varchar, form.Multifile).FieldOptionExt(map[string]interface{}{
		"maxFileCount": 10,
	})
	panel.AddField("金额", "currency", db.Int, form.Currency)
	panel.AddField("内容", "content", db.Text, form.RichText).
		FieldDefault(`<h1>343434</h1><p>34344433434</p><ol><li>23234</li><li>2342342342</li><li>asdfads</li></ol><ul><li>3434334</li><li>34343343434</li><li>44455</li></ul><p><span style="color: rgb(194, 79, 74);">343434</span></p><p><span style="background-color: rgb(194, 79, 74); color: rgb(0, 0, 0);">434434433434</span></p><table border="0" width="100%" cellpadding="0" cellspacing="0"><tbody><tr><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td></tr><tr><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td></tr><tr><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td></tr><tr><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td></tr></tbody></table><p><br></p><p><span style="color: rgb(194, 79, 74);"><br></span></p>`)

	panel.AddField("站点开关", "website", db.Tinyint, form.Switch).
		FieldHelpMsg("站点关闭后将不能访问，后台可正常登录").
		FieldOptions(types.FieldOptions{
			{Value: "0"},
			{Value: "1"},
		})
	panel.AddField("水果", "fruit", db.Varchar, form.SelectBox).
		FieldOptions(types.FieldOptions{
			{Text: "苹果", Value: "apple"},
			{Text: "香蕉", Value: "banana"},
			{Text: "西瓜", Value: "watermelon"},
			{Text: "梨", Value: "pear"},
		}).
		FieldDisplay(func(value types.FieldModel) interface{} {
			return []string{"梨"}
		})
	panel.AddField("性别", "gender", db.Tinyint, form.Radio).
		FieldOptions(types.FieldOptions{
			{Text: "男生", Value: "0"},
			{Text: "女生", Value: "1"},
		})
	panel.AddField("饮料", "drink", db.Tinyint, form.Select).
		FieldOptions(types.FieldOptions{
			{Text: "啤酒", Value: "beer"},
			{Text: "果汁", Value: "juice"},
			{Text: "白开水", Value: "water"},
			{Text: "红牛", Value: "red bull"},
		}).FieldDefault("beer")
	panel.AddField("工作经验", "experience", db.Tinyint, form.SelectSingle).
		FieldOptions(types.FieldOptions{
			{Text: "两年", Value: "0"},
			{Text: "三年", Value: "1"},
			{Text: "四年", Value: "2"},
			{Text: "五年", Value: "3"},
		}).FieldDefault("beer")
	panel.SetTabGroups(types.TabGroups{
		{"name", "age", "homepage", "email", "birthday", "password", "ip", "certificate", "currency", "content"},
		{"website", "fruit", "gender", "drink", "experience"},
	})
	panel.SetTabHeaders("input", "select")

	fields, headers := panel.GroupField()

	aform := components.Form().
		SetTabHeaders(headers).
		SetTabContents(fields).
		SetPrefix(config.Get().PrefixFixSlash()).
		SetUrl("/admin/form/update").
		SetTitle("Form").
		SetHiddenFields(map[string]string{
			form2.PreviousKey: "/admin",
		}).
		SetOperationFooter(col1 + col2)

	return types.Panel{
		Content: components.Box().
			SetHeader(aform.GetDefaultBoxHeader()).
			WithHeadBorder().
			SetBody(aform.GetContent()).
			GetContent(),
		Title:       "表单",
		Description: "表单例子",
	}, nil
}
