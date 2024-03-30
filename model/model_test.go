package model

import (
	"bytes"
	"moredoc/conf"
	"os"
	"strings"
	"testing"
	"text/template"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const dataGoTpl = `
package model

func getPermissions() (permissions []Permission) {
	permissions = []Permission{
		{{ range .Permissions }}{Title: "{{.Title}}", Description: "{{.Description}}", Method: "{{.Method}}", Path: "{{.Path}}"},
		{{end}}
	}
	return
}

func getLangs() (langs []Language) {
	langs = []Language{
		{{ range .Languages }}{Language: "{{.Language}}", Code: "{{.Code}}", Enable: {{.Enable}}},
		{{end}}
	}
	return
}
`

func TestGenData(t *testing.T) {
	// 访问的数据库链接地址。 dsn: data source name
	// 【注意】：以哪个数据库的【菜单】和【API】为准，则配置查询哪个数据库，以生成menu和api基础数据
	dsn := "root:@tcp(127.0.0.1)/moredoc?charset=utf8mb4&parseTime=True&loc=Local"
	t.Log("开始生成 data.go 文件")
	t.Log("dsn:", dsn)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn, // DSN
	}), &gorm.Config{})
	if err != nil {
		return
	}

	defer func() {
		d, _ := db.DB()
		d.Close()
	}()

	var (
		permissions         []Permission
		languages           []Language
		permissionTableName = Permission{}.TableName()
		languageTableName   = Language{}.TableName()
	)

	err = db.Table(permissionTableName).Find(&permissions).Error
	if err != nil {
		t.Errorf("生成 data.go 文件失败： %s", err.Error())
		return
	}

	err = db.Table(languageTableName).Find(&languages).Error
	if err != nil {
		t.Errorf("生成 data.go 文件失败： %s", err.Error())
		return
	}

	replacer := strings.NewReplacer("\"", "\\\"")
	for idx, permission := range permissions {
		permissions[idx].Title = replacer.Replace(permission.Title)
		permissions[idx].Description = replacer.Replace(permission.Description)
	}

	tmpl, err := template.New("data.go").Parse(dataGoTpl)
	if err != nil {
		t.Fatal(err.Error())
	}
	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, map[string]interface{}{
		"Permissions": permissions,
		"Languages":   languages,
	}); err != nil {
		t.Fatal(err.Error())
	}

	err = os.WriteFile("data.go", buf.Bytes(), 0666)
	if err != nil {
		t.Fatalf("生成 data.go 文件失败： %s", err.Error())
	}

	t.Log("生成 data.go 文件成功")
}

func TestConvertDocument(t *testing.T) {
	dsn := "root:root@tcp(127.0.0.1)/moredoc?charset=utf8mb4&parseTime=True&loc=Local"
	logger, _ := zap.NewDevelopment()
	dbModel, err := NewDBModel(&conf.Database{
		DSN:     dsn,
		Prefix:  "mnt_",
		ShowSQL: true,
	}, logger)
	if err != nil {
		t.Fatal(err.Error())
	}
	err = dbModel.ConvertDocument()
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log("success")
}
