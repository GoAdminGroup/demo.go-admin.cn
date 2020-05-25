package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	editType "github.com/GoAdminGroup/go-admin/template/types/table"
	template2 "html/template"
)

// GetPostsTable return the model of table posts.
func GetPostsTable(ctx *context.Context) (postsTable table.Table) {

	postsTable = table.NewDefaultTable(table.DefaultConfig().SetExportable(true))

	info := postsTable.GetInfo()
	info.AddField("ID", "id", db.Int).FieldSortable()
	info.AddField("标题", "title", db.Varchar)
	info.AddField("作者ID", "author_id", db.Int).FieldDisplay(func(value types.FieldModel) interface{} {
		return template.Default().
			Link().
			SetURL("/admin/info/authors/detail?__goadmin_detail_pk=" + value.Value).
			SetContent(template2.HTML(value.Value)).
			OpenInNewTab().
			SetTabTitle(template.HTML("Author Detail(" + value.Value + ")")).
			GetContent()
	})
	info.AddField("作者姓名", "name", db.Varchar).FieldDisplay(func(value types.FieldModel) interface{} {
		first, _ := value.Row["authors_goadmin_join_first_name"].(string)
		last, _ := value.Row["authors_goadmin_join_last_name"].(string)
		return first + " " + last
	})
	info.AddField("AuthorFirstName", "first_name", db.Varchar).FieldJoin(types.Join{
		Field:     "author_id",
		JoinField: "id",
		Table:     "authors",
	}).FieldHide()
	info.AddField("AuthorLastName", "last_name", db.Varchar).FieldJoin(types.Join{
		Field:     "author_id",
		JoinField: "id",
		Table:     "authors",
	}).FieldHide()
	info.AddField("简介", "description", db.Varchar).FieldWidth(230)
	info.AddField("内容", "content", db.Varchar).FieldEditAble(editType.Textarea).FieldWidth(230)
	info.AddField("日期", "date", db.Varchar).FieldWidth(120)

	info.SetTable("posts").SetTitle("文章").SetDescription("文章")

	formList := postsTable.GetForm()
	formList.AddField("ID", "id", db.Int, form.Default).FieldNotAllowEdit().FieldNotAllowAdd()
	formList.AddField("标题", "title", db.Varchar, form.Text)
	formList.AddField("简介", "description", db.Varchar, form.Text)
	formList.AddField("内容", "content", db.Varchar, form.RichText).FieldEnableFileUpload()
	formList.AddField("日期", "date", db.Varchar, form.Date)

	formList.SetWrapper(func(content template2.HTML) template2.HTML {
		tableContent := template.Default().Table().SetThead(types.Thead{
			{Head: "总阅读量"},
			{Head: "总订阅量"},
			{Head: "今日总浏览人数"},
			{Head: "新增用户数"},
			{Head: "留存"},
		}).SetInfoList([]map[string]types.InfoItem{
			{
				"总阅读量":    {Content: "1223"},
				"总订阅量":    {Content: "1433"},
				"今日总浏览人数": {Content: "230"},
				"新增用户数":   {Content: "20"},
				"留存":      {Content: "50%"},
			},
		}).GetContent()
		return template.Default().Box().
			SetBody(tableContent).
			SetNoPadding().
			WithHeadBorder().
			GetContent() + content
	})
	formList.EnableAjax("成功", "失败")

	formList.SetTable("posts").SetTitle("文章").SetDescription("文章")

	return
}
