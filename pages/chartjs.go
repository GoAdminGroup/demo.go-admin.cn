package pages

import (
	"github.com/GoAdminGroup/go-admin/modules/config"
	template2 "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
)

func GetChartJsContent() types.Panel {

	components := template2.Get(config.Get().Theme)

	size := map[string]string{"md": "4", "sm": "3", "xs": "6"}

	col1 := components.Col().SetSize(size)
	col2 := components.Col().SetSize(size)
	col3 := components.Col().SetSize(size)

	row1 := components.Row().SetContent(col1.GetContent() + col2.GetContent() + col3.GetContent())

	return types.Panel{
		Content:     row1.GetContent(),
		Title:       "ChartJs",
		Description: "this is a chart js example",
	}
}
