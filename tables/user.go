package tables

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/icon"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/action"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	selection "github.com/GoAdminGroup/go-admin/template/types/form/select"
	editType "github.com/GoAdminGroup/go-admin/template/types/table"
)

// GetUserTable return the model of table user.
func GetUserTable(ctx *context.Context) (userTable table.Table) {

	userTable = table.NewDefaultTable(table.Config{
		Driver:     db.DriverMysql,
		CanAdd:     true,
		Editable:   true,
		Deletable:  true,
		Exportable: true,
		Connection: table.DefaultConnectionName,
		PrimaryKey: table.PrimaryKey{
			Type: db.Int,
			Name: table.DefaultPrimaryKeyName,
		},
	})

	info := userTable.GetInfo().SetFilterFormLayout(form.LayoutThreeCol).ExportValue()
	info.AddField("ID", "id", db.Int).FieldSortable()
	info.AddField("Name", "name", db.Varchar).FieldEditAble(editType.Text).
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("Gender", "gender", db.Tinyint).FieldDisplay(func(model types.FieldModel) interface{} {
		if model.Value == "0" {
			return "men"
		}
		if model.Value == "1" {
			return "women"
		}
		return "unknown"
	}).FieldEditAble(editType.Switch).FieldEditOptions(types.FieldOptions{
		{Value: "0", Text: "男"},
		{Value: "1", Text: "女"},
	}).FieldFilterable(types.FilterType{FormType: form.SelectSingle}).FieldFilterOptions(types.FieldOptions{
		{Value: "0", Text: "men"},
		{Value: "1", Text: "women"},
	})
	info.AddColumn("个性", func(value types.FieldModel) interface{} {
		return "帅气"
	})
	info.AddColumnButtons("查看更多", types.GetColumnButton("更多", icon.Info,
		action.PopUp("/see/more/example", "更多", func(ctx *context.Context) (success bool, msg string, data interface{}) {
			return true, "ok", "<h1>详情</h1><p>balabala</p><p>此功能v1.2.7开放</p>"
		})))
	info.AddField("Phone", "phone", db.Varchar).FieldFilterable()
	info.AddField("City", "city", db.Varchar).FieldFilterable().
		FieldEditAble(editType.Select).FieldEditOptions(types.FieldOptions{
		{Value: "guangzhou", Text: "广州"},
		{Value: "shanghai", Text: "上海"},
		{Value: "beijing", Text: "北京"},
		{Value: "shenzhen", Text: "深圳"},
	})
	info.AddField("Avatar", "avatar", db.Varchar).FieldDisplay(func(value types.FieldModel) interface{} {
		return template.Default().Image().
			SetSrc(`//quick.go-admin.cn/demo/assets/dist/img/gopher_avatar.png`).
			SetHeight("120").SetWidth("120").WithModal().GetContent()
	})
	info.AddField("CreatedAt", "created_at", db.Timestamp).
		FieldFilterable(types.FilterType{FormType: form.DatetimeRange})
	info.AddField("UpdatedAt", "updated_at", db.Timestamp).FieldEditAble(editType.Datetime)

	info.AddActionButton("google", action.Jump("https://google.com"))
	info.AddActionButton("审批", action.Ajax("/admin/audit",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			return true, "审批成功，奥利给", ""
		}))
	info.AddActionButton("预览", action.PopUp("/admin/preview", "预览",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			return true, "", "<h2>预览内容</h2>"
		}))
	info.AddButton("jump", icon.User, action.JumpInNewTab("/admin/info/authors", "作者"))
	info.AddButton("popup", icon.Terminal, action.PopUp("/admin/popup", "Popup Example",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			return true, "", "<h2>hello world</h2>"
		}))
	info.AddButton("iframe", icon.Tv, action.PopUpWithIframe("/admin/iframe", "Iframe Example",
		action.IframeData{Src: "/admin/info/profile/new"}, "900px", "600px"))
	info.AddButton("ajax", icon.Android, action.Ajax("/admin/ajax",
		func(ctx *context.Context) (success bool, msg string, data interface{}) {
			return true, "请求成功，奥利给", ""
		}))

	info.AddSelectBox("性别", types.FieldOptions{
		{Value: "0", Text: "男"},
		{Value: "1", Text: "女"},
	}, action.FieldFilter("gender"))

	info.SetTable("users").SetTitle("Users").SetDescription("Users")

	formList := userTable.GetForm()
	formList.AddField("ID", "id", db.Int, form.Default).FieldNotAllowEdit().FieldNotAllowAdd()
	formList.AddField("IP", "ip", db.Varchar, form.Text)
	formList.AddField("姓名", "name", db.Varchar, form.Text)
	formList.AddField("性别", "gender", db.Tinyint, form.Radio).
		FieldOptions(types.FieldOptions{
			{Text: "男", Value: "0"},
			{Text: "女", Value: "1"},
		}).FieldDefault("0")
	formList.AddField("电话", "phone", db.Varchar, form.Text)
	formList.AddField("国家", "country", db.Tinyint, form.SelectSingle).
		FieldOptions(types.FieldOptions{
			{Text: "中国", Value: "0"},
			{Text: "美国", Value: "1"},
			{Text: "英国", Value: "2"},
			{Text: "加拿大", Value: "3"},
		}).FieldDefault("0").FieldOnChooseAjax("city", "/choose/country",
		func(ctx *context.Context) (bool, string, interface{}) {
			country := ctx.FormValue("value")
			var data = make(selection.Options, 0)
			switch country {
			case "0":
				data = selection.Options{
					{Text: "北京", ID: "beijing"},
					{Text: "上海", ID: "shangHai"},
					{Text: "广州", ID: "guangZhou"},
					{Text: "深圳", ID: "shenZhen"},
				}
			case "1":
				data = selection.Options{
					{Text: "洛杉矶", ID: "los angeles"},
					{Text: "华盛顿", ID: "washington, dc"},
					{Text: "纽约", ID: "new york"},
					{Text: "拉斯维加斯", ID: "las vegas"},
				}
			case "2":
				data = selection.Options{
					{Text: "伦敦", ID: "london"},
					{Text: "剑桥", ID: "cambridge"},
					{Text: "曼切斯特", ID: "manchester"},
					{Text: "利物浦", ID: "liverpool"},
				}
			case "3":
				data = selection.Options{
					{Text: "温哥华", ID: "vancouver"},
					{Text: "多伦多", ID: "toronto"},
				}
			default:
				data = selection.Options{
					{Text: "北京", ID: "beijing"},
					{Text: "上海", ID: "shangHai"},
					{Text: "广州", ID: "guangZhou"},
					{Text: "深圳", ID: "shenZhen"},
				}
			}
			return true, "ok", data
		})
	formList.AddField("城市", "city", db.Varchar, form.SelectSingle).
		FieldOptionInitFn(func(val types.FieldModel) types.FieldOptions {

			if val.Value == "" {
				return types.FieldOptions{
					{Text: "北京", Value: "beijing"},
					{Text: "上海", Value: "shangHai"},
					{Text: "广州", Value: "guangZhou"},
					{Text: "深圳", Value: "shenZhen"},
				}
			}

			return types.FieldOptions{
				{Value: val.Value, Text: val.Value, Selected: true},
			}
		}).FieldOnChooseAjax("district", "/choose/city",
		func(ctx *context.Context) (bool, string, interface{}) {
			country := ctx.FormValue("value")
			var data = make(selection.Options, 0)
			switch country {
			case "beijing":
				data = selection.Options{
					{Text: "朝阳", ID: "chaoyang"},
					{Text: "海淀", ID: "haidian"},
				}
			case "shangHai":
				data = selection.Options{
					{Text: "杨浦", ID: "yangpu"},
					{Text: "浦东", ID: "pudong"},
				}
			default:
				data = selection.Options{
					{Text: "南区", ID: "southern"},
					{Text: "北区", ID: "north"},
				}
			}
			return true, "ok", data
		})
	formList.AddField("地区", "district", db.Varchar, form.SelectSingle).
		FieldOptionInitFn(func(val types.FieldModel) types.FieldOptions {

			if val.Value == "" {
				return types.FieldOptions{
					{Text: "南区", Value: "southern"},
					{Text: "北区", Value: "north"},
				}
			}

			return types.FieldOptions{
				{Value: val.Value, Text: val.Value, Selected: true},
			}
		})
	formList.AddField("自定义字段", "role", db.Varchar, form.Text).
		FieldPostFilterFn(func(value types.PostFieldModel) interface{} {
			fmt.Println("user custom field", value)
			return ""
		})

	formList.AddField("UpdatedAt", "updated_at", db.Timestamp, form.Default).FieldNotAllowAdd()
	formList.AddField("CreatedAt", "created_at", db.Timestamp, form.Default).FieldNotAllowAdd()

	userTable.GetForm().SetTabGroups(types.
		NewTabGroups("id", "ip", "name", "gender", "country", "city", "district").
		AddGroup("phone", "role", "created_at", "updated_at")).
		SetTabHeaders("profile1", "profile2")

	formList.SetTable("users").SetTitle("Users").SetDescription("Users")

	formList.SetPostHook(func(values form2.Values) error {
		fmt.Println("userTable.GetForm().PostHook", values)
		return nil
	})

	return
}
