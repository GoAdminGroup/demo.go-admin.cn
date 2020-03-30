package main

import (
	"github.com/GoAdminGroup/demo/tables"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/tests"
	"github.com/GoAdminGroup/go-admin/tests/common"
	"github.com/GoAdminGroup/go-admin/tests/frameworks/gin"
	"github.com/GoAdminGroup/go-admin/tests/web"
	"github.com/gavv/httpexpect"
	"net/http"
	"testing"
)

// 黑盒测试
func TestDemoBlackBox(t *testing.T) {
	tests.BlackBoxTestSuit(t, gin.NewHandler, config.DatabaseList{
		"default": config.Database{
			Host:   "127.0.0.1",
			Port:   "3306",
			User:   "root",
			Pwd:    "root",
			Name:   "go_admin_demo_test", // 注意：测试数据库名必须包含"test"
			Driver: "mysql",
		},
	}, tables.Generators, func(cfg config.DatabaseList) {
		// 框架自带数据清理
		tests.Cleaner(cfg)
		// 以下清理自己的数据：
		// ...
	}, func(e *httpexpect.Expect) {
		// 框架自带内置表测试
		common.Test(e)
		// 以下写API测试：
		// 更多用法：https://github.com/gavv/httpexpect
		// ...
		e.POST("/signin").Expect().Status(http.StatusOK)
	})
}

// 浏览器验收测试
func TestDemoUserAcceptance(t *testing.T) {
	web.UserAcceptanceTestSuit(t, func(t *testing.T, page *web.Page) {
		// 写浏览器测试，基于chromedriver
		// 更多用法：https://github.com/sclevine/agouti
		page.NavigateTo("http://127.0.0.1:9033/admin")
		page.Contain("username")
	}, func(quit chan struct{}) {
		// 启动服务器
	}, true)
}
