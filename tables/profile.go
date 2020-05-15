package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"path/filepath"
	"strings"
)

func GetProfileTable(ctx *context.Context) table.Table {

	profile := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := profile.GetInfo().HideFilterArea()
	info.AddField("ID", "id", db.Int).FieldFilterable()
	info.AddField("UUID", "uuid", db.Varchar).FieldCopyable()
	info.AddField("通关", "pass", db.Tinyint).FieldBool("1", "0")
	info.AddField("照片", "photos", db.Varchar).FieldCarousel(func(value string) []string {
		return strings.Split(value, ",")
	}, 150, 100)
	info.AddField("完成状态", "finish_state", db.Tinyint).
		FieldDisplay(func(value types.FieldModel) interface{} {
			if value.Value == "0" {
				return "第一步"
			}
			if value.Value == "1" {
				return "第二步"
			}
			if value.Value == "2" {
				return "第三步"
			}
			return "未知"
		}).
		FieldDot(map[string]types.FieldDotColor{
			"第一步": types.FieldDotColorDanger,
			"第二步": types.FieldDotColorInfo,
			"第三步": types.FieldDotColorPrimary,
		}, types.FieldDotColorDanger)
	info.AddField("完成进度", "finish_progress", db.Int).FieldProgressBar()
	info.AddField("简历", "resume", db.Varchar).
		FieldDisplay(func(value types.FieldModel) interface{} {
			return filepath.Base(value.Value)
		}).
		FieldDownLoadable("http://yinyanghu.github.io/files/")
	info.AddField("简历大小", "resume_size", db.Int).FieldFileSize()

	info.SetTable("profile").SetTitle("信息").SetDescription("信息")

	formList := profile.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default)
	formList.AddField("UUID", "uuid", db.Varchar, form.Text)
	formList.AddField("照片", "photos", db.Varchar, form.Text)
	formList.AddField("简历", "resume", db.Varchar, form.Text)
	formList.AddField("简历大小", "resume_size", db.Int, form.Number)
	formList.AddField("完成状态", "finish_state", db.Tinyint, form.Number)
	formList.AddField("完成进度", "finish_progress", db.Int, form.Number)
	formList.AddField("通关", "pass", db.Tinyint, form.Number)

	formList.SetTable("profile").SetTitle("信息").SetDescription("信息")

	return profile
}
