package login

import (
	"bytes"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/template/login"
	"html/template"
)

type Login struct{}

func Get() *Login { return new(Login) }

func (l *Login) GetTemplate() (*template.Template, string) {
	tmpl, err := template.New("login_theme1").
		Funcs(login.DefaultFuncMap).
		Parse(List["login/theme1"])

	if err != nil {
		logger.Error("login component, get template error: ", err)
	}

	return tmpl, "login_theme1"
}

func (l *Login) GetAssetList() []string               { return AssetsList }
func (l *Login) GetAsset(name string) ([]byte, error) { return Asset(name[1:]) }
func (l *Login) GetName() string                      { return "login" }
func (l *Login) IsAPage() bool                        { return true }

func (l *Login) GetContent() template.HTML {
	buffer := new(bytes.Buffer)
	tmpl, defineName := l.GetTemplate()
	err := tmpl.ExecuteTemplate(buffer, defineName, l)
	if err != nil {
		logger.Error("login component, compose html error:", err)
	}
	return template.HTML(buffer.String())
}
