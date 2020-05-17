package pages

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/language"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	template2 "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/icon"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"github.com/gin-gonic/gin"
	"html/template"
)

func GetForm1Content(ctx *gin.Context) (types.Panel, error) {

	components := template2.Get(config.GetTheme())

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
	panel.AddField("名字", "name", db.Varchar, form.Text).
		FieldFoot(seeCodeHTML(`formList.AddField("名字", "name", db.Varchar, form.Text)`))
	panel.AddField("年龄", "age", db.Int, form.Number).
		FieldFoot(seeCodeHTML(`formList.AddField("年龄", "age", db.Int, form.Number)`))
	panel.AddField("主页", "homepage", db.Varchar, form.Url).FieldDefault("http://google.com").
		FieldFoot(seeCodeHTML(`formList.AddField("主页", "homepage", db.Varchar, form.Url).FieldDefault("http://google.com")`))
	panel.AddField("邮箱", "email", db.Varchar, form.Email).FieldDefault("xxxx@xxx.com").
		FieldFoot(seeCodeHTML(`formList.AddField("邮箱", "email", db.Varchar, form.Email).FieldDefault("xxxx@xxx.com")`))
	panel.AddField("生日", "birthday", db.Varchar, form.Date).FieldDefault("2010-09-03 18:09:05").
		FieldFoot(seeCodeHTML(`formList.AddField("生日", "birthday", db.Varchar, form.Date).FieldDefault("2010-09-03 18:09:05")`))
	panel.AddField("时间", "time", db.Varchar, form.Datetime).FieldDefault("2010-09-05").
		FieldFoot(seeCodeHTML(`formList.AddField("时间", "time", db.Varchar, form.Datetime).FieldDefault("2010-09-05")`))
	panel.AddField("时间范围", "time_range", db.Varchar, form.DatetimeRange).
		FieldFoot(seeCodeHTML(`formList.AddField("时间范围", "time_range", db.Varchar, form.DatetimeRange)`))
	panel.AddField("日期范围", "date_range", db.Varchar, form.DateRange).
		FieldFoot(seeCodeHTML(`formList.AddField("日期范围", "date_range", db.Varchar, form.DateRange)`))
	panel.AddField("密码", "password", db.Varchar, form.Password).FieldDivider("我是分割线").
		FieldFoot(seeCodeHTML(`formList.AddField("密码", "password", db.Varchar, form.Password).FieldDivider("我是分割线")`, true))
	panel.AddField("IP", "ip", db.Varchar, form.Ip).
		FieldFoot(seeCodeHTML(`formList.AddField("IP", "ip", db.Varchar, form.Ip)`))
	panel.AddField("证件", "certificate", db.Varchar, form.Multifile).FieldOptionExt(map[string]interface{}{
		"maxFileCount": 10,
	}).
		FieldFoot(seeCodeHTML(`formList.AddField("证件", "certificate", db.Varchar, form.Multifile).FieldOptionExt(map[string]interface{}{
		"maxFileCount": 10,
	})`))
	panel.AddField("金额", "currency", db.Int, form.Currency).
		FieldFoot(seeCodeHTML(`formList.AddField("金额", "currency", db.Int, form.Currency)`))
	panel.AddField("比例", "rate", db.Int, form.Rate).
		FieldFoot(seeCodeHTML(`formList.AddField("比例", "rate", db.Int, form.Rate)`))
	panel.AddField("奖金", "reward", db.Int, form.Slider).FieldOptionExt(map[string]interface{}{
		"max":     1000,
		"min":     1,
		"step":    1,
		"postfix": "元",
	}).
		FieldFoot(seeCodeHTML(`formList.AddField("奖金", "reward", db.Int, form.Slider).FieldOptionExt(map[string]interface{}{
		"max":     1000,
		"min":     1,
		"step":    1,
		"postfix": "元",
	})`))
	panel.AddField("内容", "content", db.Text, form.RichText).
		FieldDefault(`<h1>343434</h1><p>34344433434</p><ol><li>23234</li><li>2342342342</li><li>asdfads</li></ol><ul><li>3434334</li><li>34343343434</li><li>44455</li></ul><p><span style="color: rgb(194, 79, 74);">343434</span></p><p><span style="background-color: rgb(194, 79, 74); color: rgb(0, 0, 0);">434434433434</span></p><table border="0" width="100%" cellpadding="0" cellspacing="0"><tbody><tr><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td></tr><tr><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td></tr><tr><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td></tr><tr><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td></tr></tbody></table><p><br></p><p><span style="color: rgb(194, 79, 74);"><br></span></p>`).
		FieldDivider("二号分割线").
		FieldFoot(seeCodeHTML(`formList.AddField("内容", "content", db.Text, form.RichText).
		FieldDefault(`+"`"+`<h1>343434</h1><p>34344433434</p><ol><li>23234</li><li>2342342342</li><li>asdfads</li></ol><ul><li>3434334</li><li>34343343434</li><li>44455</li></ul><p><span style="color: rgb(194, 79, 74);">343434</span></p><p><span style="background-color: rgb(194, 79, 74); color: rgb(0, 0, 0);">434434433434</span></p><table border="0" width="100%" cellpadding="0" cellspacing="0"><tbody><tr><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td></tr><tr><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td></tr><tr><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td></tr><tr><td>&nbsp;</td><td>&nbsp;</td><td>&nbsp;</td></tr></tbody></table><p><br></p><p><span style="color: rgb(194, 79, 74);"><br></span></p>`+"`"+`).
		FieldDivider("二号分割线")`, true))
	panel.AddField("代码", "code", db.Text, form.Code).FieldDefault(`package main

import "fmt"

func main() {
	fmt.Println("hello GoAdmin!")
}
`).
		FieldFoot(seeCodeHTML(`formList.AddField("代码", "code", db.Text, form.Code).FieldDefault(` + "`" + `package main

import "fmt"

func main() {
	fmt.Println("hello GoAdmin!")
}` + "`)"))

	panel.AddField("站点开关", "website", db.Tinyint, form.Switch).
		FieldHelpMsg("站点关闭后将不能访问，后台可正常登录").
		FieldOptions(types.FieldOptions{
			{Value: "0"},
			{Value: "1"},
		}).
		FieldFoot(seeCodeHTML(`formList.AddField("站点开关", "website", db.Tinyint, form.Switch).
		FieldHelpMsg("站点关闭后将不能访问，后台可正常登录").
		FieldOptions(types.FieldOptions{
			{Value: "0"},
			{Value: "1"},
		})`))
	panel.AddField("水果", "fruit", db.Varchar, form.SelectBox).
		FieldOptions(types.FieldOptions{
			{Text: "苹果", Value: "apple"},
			{Text: "香蕉", Value: "banana"},
			{Text: "西瓜", Value: "watermelon"},
			{Text: "梨", Value: "pear"},
		}).
		FieldDisplay(func(value types.FieldModel) interface{} {
			return []string{"梨"}
		}).
		FieldFoot(seeCodeHTML(`formList.AddField("水果", "fruit", db.Varchar, form.SelectBox).
		FieldOptions(types.FieldOptions{
			{Text: "苹果", Value: "apple"},
			{Text: "香蕉", Value: "banana"},
			{Text: "西瓜", Value: "watermelon"},
			{Text: "梨", Value: "pear"},
		}).
		FieldDisplay(func(value types.FieldModel) interface{} {
			return []string{"梨"}
		})`))
	panel.AddField("性别", "gender", db.Tinyint, form.Radio).
		FieldOptions(types.FieldOptions{
			{Text: "男生", Value: "0"},
			{Text: "女生", Value: "1"},
		}).
		FieldFoot(seeCodeHTML(`formList.AddField("性别", "gender", db.Tinyint, form.Radio).
		FieldOptions(types.FieldOptions{
			{Text: "男生", Value: "0"},
			{Text: "女生", Value: "1"},
		})`))
	panel.AddField("饮料", "drink", db.Tinyint, form.Select).
		FieldOptions(types.FieldOptions{
			{Text: "啤酒", Value: "beer"},
			{Text: "果汁", Value: "juice"},
			{Text: "白开水", Value: "water"},
			{Text: "红牛", Value: "red bull"},
		}).FieldDefault("beer").
		FieldFoot(seeCodeHTML(`formList.AddField("饮料", "drink", db.Tinyint, form.Select).
		FieldOptions(types.FieldOptions{
			{Text: "啤酒", Value: "beer"},
			{Text: "果汁", Value: "juice"},
			{Text: "白开水", Value: "water"},
			{Text: "红牛", Value: "red bull"},
		}).FieldDefault("beer")`))
	panel.AddField("工作经验", "experience", db.Tinyint, form.SelectSingle).
		FieldOptions(types.FieldOptions{
			{Text: "两年", Value: "0"},
			{Text: "三年", Value: "1"},
			{Text: "四年", Value: "2"},
			{Text: "五年", Value: "3"},
		}).FieldDefault("beer").
		FieldFoot(seeCodeHTML(`formList.AddField("工作经验", "experience", db.Tinyint, form.SelectSingle).
		FieldOptions(types.FieldOptions{
			{Text: "两年", Value: "0"},
			{Text: "三年", Value: "1"},
			{Text: "四年", Value: "2"},
			{Text: "五年", Value: "3"},
		}).FieldDefault("beer")`))
	panel.AddField("零食", "snacks", db.Varchar, form.Checkbox).
		FieldOptions(types.FieldOptions{
			{Text: "麦片", Value: "0"},
			{Text: "薯条", Value: "1"},
			{Text: "辣条", Value: "2"},
			{Text: "雪糕", Value: "3"},
		}).
		FieldFoot(seeCodeHTML(`formList.AddField("零食", "snacks", db.Varchar, form.Checkbox).
		FieldOptions(types.FieldOptions{
			{Text: "麦片", Value: "0"},
			{Text: "薯条", Value: "1"},
			{Text: "辣条", Value: "2"},
			{Text: "雪糕", Value: "3"},
		})`))
	panel.AddField("猫", "cat", db.Varchar, form.CheckboxStacked).
		FieldOptions(types.FieldOptions{
			{Text: "卡菲猫", Value: "0"},
			{Text: "英短", Value: "1"},
			{Text: "美短", Value: "2"},
		}).
		FieldFoot(seeCodeHTML(`formList.AddField("猫", "cat", db.Varchar, form.CheckboxStacked).
		FieldOptions(types.FieldOptions{
			{Text: "卡菲猫", Value: "0"},
			{Text: "英短", Value: "1"},
			{Text: "美短", Value: "2"},
		})`))
	panel.AddRow(func(pa *types.FormPanel) {
		panel.AddField("省份", "province", db.Tinyint, form.SelectSingle).
			FieldOptions(types.FieldOptions{
				{Text: "北京", Value: "0"},
				{Text: "上海", Value: "1"},
				{Text: "广东", Value: "2"},
				{Text: "重庆", Value: "3"},
			}).FieldRowWidth(2)
		panel.AddField("城市", "city", db.Tinyint, form.SelectSingle).
			FieldOptions(types.FieldOptions{
				{Text: "北京", Value: "0"},
				{Text: "上海", Value: "1"},
				{Text: "广州", Value: "2"},
				{Text: "深圳", Value: "3"},
			}).FieldRowWidth(3).FieldHeadWidth(2).FieldInputWidth(10)
		panel.AddField("区域", "district", db.Tinyint, form.SelectSingle).
			FieldOptions(types.FieldOptions{
				{Text: "朝阳区", Value: "0"},
				{Text: "海珠区", Value: "1"},
				{Text: "浦东新区", Value: "2"},
				{Text: "宝安区", Value: "3"},
			}).FieldRowWidth(3).FieldHeadWidth(2).FieldInputWidth(9)
	}).FieldFoot(seeCodeHTML(`panel.AddRow(func(pa *types.FormPanel) {
		panel.AddField("省份", "province", db.Tinyint, form.SelectSingle).
			FieldOptions(types.FieldOptions{
				{Text: "北京", Value: "0"},
				{Text: "上海", Value: "1"},
				{Text: "广东", Value: "2"},
				{Text: "重庆", Value: "3"},
			}).FieldRowWidth(2)
		panel.AddField("城市", "city", db.Tinyint, form.SelectSingle).
			FieldOptions(types.FieldOptions{
				{Text: "北京", Value: "0"},
				{Text: "上海", Value: "1"},
				{Text: "广州", Value: "2"},
				{Text: "深圳", Value: "3"},
			}).FieldRowWidth(3).FieldHeadWidth(2).FieldInputWidth(10)
		panel.AddField("区域", "district", db.Tinyint, form.SelectSingle).
			FieldOptions(types.FieldOptions{
				{Text: "朝阳区", Value: "0"},
				{Text: "海珠区", Value: "1"},
				{Text: "浦东新区", Value: "2"},
				{Text: "宝安区", Value: "3"},
			}).FieldRowWidth(3).FieldHeadWidth(2).FieldInputWidth(9)
	})`))
	panel.AddField("雇员", "employee", db.Varchar, form.Array).
		FieldFoot(seeCodeHTML(`formList.AddField("雇员", "employee", db.Varchar, form.Array)`))
	panel.AddTable("设置", "setting", func(panel *types.FormPanel) {
		panel.AddField("Key", "key", db.Varchar, form.Text).FieldHideLabel()
		panel.AddField("Value", "value", db.Varchar, form.Text).FieldHideLabel()
	}).
		FieldFoot(seeCodeHTML(`formList.AddTable("设置", "setting", func(panel *types.FormPanel) {
		panel.AddField("Key", "key", db.Varchar, form.Text).FieldHideLabel()
		panel.AddField("Value", "value", db.Varchar, form.Text).FieldHideLabel()
	})`))
	panel.SetTabGroups(types.TabGroups{
		{"name", "age", "homepage", "email", "birthday", "time", "time_range", "date_range", "password", "ip",
			"certificate", "currency", "rate", "reward", "content", "code"},
		{"website", "snacks", "fruit", "gender", "cat", "drink", "province", "city", "district", "experience"},
		{"employee", "setting"},
	})
	panel.SetTabHeaders("输入", "选择", "多项")

	fields, headers := panel.GroupField()

	aform := components.Form().
		SetTabHeaders(headers).
		SetTabContents(fields).
		SetPrefix(config.PrefixFixSlash()).
		SetUrl("/admin/form/update").
		SetTitle("表单例子").
		SetHiddenFields(map[string]string{
			form2.PreviousKey: "/admin",
		}).
		SetOperationFooter(col1 + col2)

	popup := components.Popup().SetID("code_modal").
		SetHideFooter().
		SetTitle("代码").
		SetHeight("300px").
		SetBody(template.HTML("")).
		GetContent()

	return types.Panel{
		Content: components.Box().
			SetHeader(aform.GetDefaultBoxHeader(true)).
			WithHeadBorder().
			SetBody(aform.GetContent()).
			GetContent() + popup,
		Title:       "表单",
		Description: "表单例子",
		CSS:         `.modal.fade.in{z-index:10002}`,
		JS: `
$(".see-code").on("click", function(){
	$('#code_modal .modal-body').html('<textarea style="width: 100%;height: 100%;font-size: 17px;">' + $(this).parent().next().html() + "</textarea>");
	$("#code_modal").modal();
})
`,
	}, nil
}

func seeCodeHTML(data string, divide ...bool) template.HTML {
	if len(divide) > 0 && divide[0] {
		return template.HTML(fmt.Sprintf(`<div style="margin-top: 24px;"><a class="see-code" href="javascript:;">查看代码</a></div><div style="display:none;">%s</div>`, data))
	}
	return template.HTML(fmt.Sprintf(`<div><a class="see-code" href="javascript:;">查看代码</a></div><div style="display:none;">%s</div>`, data))
}
