package main

import (
	"github.com/GoAdminGroup/components/echarts"
	_ "github.com/GoAdminGroup/go-admin/adapter/gin"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	_ "github.com/GoAdminGroup/themes/adminlte"
	_ "github.com/GoAdminGroup/themes/sword"
	"log"
	"os"
	"os/signal"

	"github.com/GoAdminGroup/demo/login"
	"github.com/GoAdminGroup/demo/pages"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/examples/datamodel"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/plugins/admin"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/gin-gonic/gin"
	template2 "html/template"
	"net/http"
)

func main() {
	r := gin.Default()

	eng := engine.Default()

	adminPlugin := admin.NewAdmin(datamodel.Generators)

	// add generator, first parameter is the url prefix of table when visit.
	// example:
	//
	// "user" => http://localhost:9033/admin/info/user
	//
	adminPlugin.AddGenerator("user", datamodel.GetUserTable)

	template.AddLoginComp(login.GetLoginComponent())
	template.AddComp(chartjs.NewChart())
	template.AddComp(echarts.NewChart())

	rootPath := "/data/www/go-admin"
	//rootPath = "."

	cfg := config.ReadFromJson(rootPath + "/config.json")
	cfg.CustomFootHtml = template2.HTML(`<div style="display:none;">
    <script type="text/javascript" src="https://s9.cnzz.com/z_stat.php?id=1278156902&web_id=1278156902"></script>
</div>`)
	cfg.CustomHeadHtml = template2.HTML(`<link rel="icon" type="image/png" sizes="32x32" href="//quick.go-admin.cn/official/assets/imgs/icons.ico/favicon-32x32.png">
        <link rel="icon" type="image/png" sizes="96x96" href="//quick.go-admin.cn/official/assets/imgs/icons.ico/favicon-64x64.png">
        <link rel="icon" type="image/png" sizes="16x16" href="//quick.go-admin.cn/official/assets/imgs/icons.ico/favicon-16x16.png">`)

	if err := eng.AddConfig(cfg).AddPlugins(adminPlugin).Use(r); err != nil {
		panic(err)
	}

	r.Static("/uploads", rootPath+"/uploads")

	// you can custom your pages like:

	r.GET("/admin", func(ctx *gin.Context) {
		eng.Content(ctx, func(ctx interface{}) (types.Panel, error) {
			return pages.GetDashBoard2Content()
		})
	})

	r.GET("/admin/form1", func(ctx *gin.Context) {
		eng.Content(ctx, func(ctx interface{}) (types.Panel, error) {
			return pages.GetForm1Content()
		})
	})

	r.GET("/admin/echarts", func(ctx *gin.Context) {
		eng.Content(ctx, func(ctx interface{}) (types.Panel, error) {
			return pages.GetDashBoard3Content()
		})
	})

	r.POST("/admin/popup", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": "<h2>hello world</h2>",
		})
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
