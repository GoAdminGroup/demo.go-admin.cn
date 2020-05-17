package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/icon"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/action"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

// GetAuthorsTable return the model of table author.
func GetAuthorsTable(ctx *context.Context) (authorsTable table.Table) {

	authorsTable = table.NewDefaultTable(table.DefaultConfig())

	// connect your custom connection
	// authorsTable = table.NewDefaultTable(table.DefaultConfigWithDriverAndConnection("mysql", "admin"))

	info := authorsTable.GetInfo()
	info.AddField("ID", "id", db.Int).FieldSortable()
	info.AddField("First Name", "first_name", db.Varchar).FieldHide()
	info.AddField("Last Name", "last_name", db.Varchar).FieldHide()
	info.AddField("姓名", "name", db.Varchar).FieldDisplay(func(value types.FieldModel) interface{} {
		first, _ := value.Row["first_name"].(string)
		last, _ := value.Row["last_name"].(string)
		return first + " " + last
	})
	info.AddField("邮箱", "email", db.Varchar)
	info.AddField("生日", "birthdate", db.Date)
	info.AddField("加入时间", "added", db.Timestamp)

	info.AddButton("文章列表", icon.Tv, action.PopUpWithIframe("/authors/list", "文章",
		action.IframeData{Src: "/admin/info/posts"}, "900px", "560px"))
	info.SetTable("authors").SetTitle("作者").SetDescription("作者")

	formList := authorsTable.GetForm()
	formList.AddField("ID", "id", db.Int, form.Default).FieldNotAllowEdit().FieldNotAllowAdd()
	formList.AddField("名", "first_name", db.Varchar, form.Text)
	formList.AddField("姓", "last_name", db.Varchar, form.Text)
	formList.AddField("邮箱", "email", db.Varchar, form.Text)
	formList.AddField("生日", "birthdate", db.Date, form.Text)
	formList.AddField("加入时间", "added", db.Timestamp, form.Text)

	formList.SetTable("authors").SetTitle("作者").SetDescription("作者")

	return
}
