package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	template2 "html/template"
)

func GetEmployeeTable(ctx *context.Context) table.Table {

	employeeTable := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := employeeTable.GetInfo().HideFilterArea()

	info.AddField("ID", "id", db.Int).FieldFilterable()
	info.AddField("姓名", "name", db.Varchar)
	info.AddField("性别", "gender", db.Tinyint).FieldDisplay(func(value types.FieldModel) interface{} {
		if value.Value == "0" {
			return "男"
		}
		return "女"
	})
	info.AddField("部门", "department", db.Varchar)
	info.AddField("手机", "phone", db.Varchar)
	info.AddField("职位", "job", db.Varchar)

	department := ctx.Query("department")

	info.SetTable("employee").SetTitle("雇员").SetDescription("雇员").
		SetWrapper(func(content template2.HTML) template2.HTML {
			col1 := `<div style="margin-left:243px;">` + content + `</div>`

			tree := template.Default().TreeView().SetTree(types.TreeViewData{
				Data: types.TreeViewItems{
					{
						Text: "趣杰网络",
						Href: "/admin/info/employee?__go_admin_no_animation_=true",
						Nodes: types.TreeViewItems{
							{
								Text: "技术",
								State: types.TreeViewItemState{
									Expanded: department == "前端" || department == "中台" || department == "后端",
								},
								Nodes: types.TreeViewItems{
									{
										Text: "前端",
										Href: "/admin/info/employee?department=前端&__go_admin_no_animation_=true",
										State: types.TreeViewItemState{
											Checked:  department == "前端",
											Selected: department == "前端",
										},
									}, {
										Text: "中台",
										Href: "/admin/info/employee?department=中台&__go_admin_no_animation_=true",
										State: types.TreeViewItemState{
											Checked:  department == "中台",
											Selected: department == "中台",
										},
									}, {
										Text: "后端",
										Href: "/admin/info/employee?department=后端&__go_admin_no_animation_=true",
										State: types.TreeViewItemState{
											Checked:  department == "后端",
											Selected: department == "后端",
										},
									},
								},
							}, {
								Text: "销售",
								Href: "/admin/info/employee?department=销售&__go_admin_no_animation_=true",
								State: types.TreeViewItemState{
									Checked:  department == "销售",
									Selected: department == "销售",
								},
							}, {
								Text: "前台",
								Href: "/admin/info/employee?department=前台&__go_admin_no_animation_=true",
								State: types.TreeViewItemState{
									Checked:  department == "前台",
									Selected: department == "前台",
								},
							}, {
								Text: "人力",
								Href: "/admin/info/employee?department=人力&__go_admin_no_animation_=true",
								State: types.TreeViewItemState{
									Checked:  department == "人力",
									Selected: department == "人力",
								},
							},
						},
					},
				},
				ExpandIcon:        "fa fa-angle-right",
				CollapseIcon:      "fa fa-angle-down",
				SelectedBackColor: "#fbfbfb",
				SelectedColor:     "#333333",
				EnableLinks:       true,
			}).GetContent()

			col2 := `<div style="position: absolute;width:230px;">` + template.Default().Box().SetHeader("组织结构").
				WithHeadBorder().SetBody(tree).GetContent() + `</div>`
			return `<div style="width:100%;">` + col2 + col1 + `</div>`
		})

	formList := employeeTable.GetForm()

	formList.AddField("ID", "id", db.Int, form.Default).FieldNotAllowAdd()
	formList.AddField("姓名", "name", db.Varchar, form.Text)
	formList.AddField("性别", "gender", db.Tinyint, form.Number)
	formList.AddField("部门", "department", db.Varchar, form.Text)
	formList.AddField("手机", "phone", db.Varchar, form.Text)
	formList.AddField("职位", "job", db.Varchar, form.Text)

	formList.SetTable("employee").SetTitle("雇员").SetDescription("雇员")

	return employeeTable
}
