package main

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/template/types/action"
	"log"
	"net/http"
	"os"
	"os/signal"

	ada "github.com/GoAdminGroup/go-admin/adapter/gin"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql"
	_ "github.com/GoAdminGroup/themes/sword"

	"github.com/GoAdminGroup/components/echarts"
	"github.com/GoAdminGroup/demo/login"
	"github.com/GoAdminGroup/demo/pages"
	"github.com/GoAdminGroup/demo/tables"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	eng := engine.Default()

	template.AddLoginComp(login.GetLoginComponent())
	template.AddComp(chartjs.NewChart())
	template.AddComp(echarts.NewChart())

	rootPath := "/data/www/go-admin"
	//rootPath = "."

	cfg := config.ReadFromJson(rootPath + "/config.json")
	cfg.CustomFootHtml = template.HTML(`<div style="display:none;">
    <script type="text/javascript" src="https://s9.cnzz.com/z_stat.php?id=1278156902&web_id=1278156902"></script>
</div>`)
	cfg.CustomHeadHtml = template.HTML(`<link rel="icon" type="image/png" sizes="32x32" href="//quick.go-admin.cn/official/assets/imgs/icons.ico/favicon-32x32.png">
        <link rel="icon" type="image/png" sizes="96x96" href="//quick.go-admin.cn/official/assets/imgs/icons.ico/favicon-64x64.png">
        <link rel="icon" type="image/png" sizes="16x16" href="//quick.go-admin.cn/official/assets/imgs/icons.ico/favicon-16x16.png">`)

	cfg.Animation = config.PageAnimation{
		Type:     "fadeInUp",
		Duration: 0.9,
	}

	if err := eng.AddConfig(cfg).
		AddGenerators(tables.Generators).
		AddGenerator("user", tables.GetUserTable).
		AddNavButtons("网站信息", "", action.PopUp("/website/info", "网站信息",
			func(ctx *context.Context) (success bool, msg string, data interface{}) {
				return true, "ok", `<p>网站由 <a href="https://github.com/chenhg5">cg33<a/> 创造</p>`
			})).
		AddNavButtons("用户管理", "", action.Jump("/admin/info/manager")).
		Use(r); err != nil {
		panic(err)
	}

	r.Static("/uploads", rootPath+"/uploads")

	// you can custom your pages like:

	r.GET("/admin", ada.Content(pages.GetDashBoard2Content))
	r.GET("/admin/echarts", ada.Content(pages.GetDashBoard3Content))
	r.GET("/admin/table", ada.Content(pages.GetTableContent))

	r.GET("/admin/form1", ada.Content(pages.GetForm1Content))
	eng.Data("POST", "/admin/form/update", func(ctx *context.Context) {
		fmt.Println("ctx.PostForm()", ctx.PostForm())
		ctx.PjaxUrl("/admin")
	})

	r.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "/admin")
	})

	go func() {
		_ = r.Run(":9033")
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Print("closing database connection")
	eng.MysqlConnection().Close()
}
