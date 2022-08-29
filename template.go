package sqltostruct

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"text/template"
)

const structTpl = `
package {{.PackageName}}


type {{.DB.TableName | ToBigCamelCase }}Base struct{
{{ range .DB.Columns }} {{$length := len .Comment }}{{if gt $length 0}}	//{{ .Comment }} {{else}} //{{.Name}} {{end}}{{$typeLen := len .Type }} {{if gt $typeLen 0}}
	{{.Name | ToBigCamelCase}}	{{.Type}} {{.Tag}} {{else}}{{.Name}}{{end}}
{{end}}
}

type {{.DB.TableName | ToBigCamelCase }} struct{
{{ range .DB.Columns }} {{$length := len .Comment }}{{if gt $length 0}}	//{{ .Comment }} {{else}} //{{.Name}} {{end}}{{$typeLen := len .Type }} {{if gt $typeLen 0}}
	{{.Name | ToBigCamelCase}}	{{.Type}} {{.Tag}} {{else}}{{.Name}}{{end}}
{{end}}
}

`

type StructTemplate struct {
	structTpl string
}

type StructColumn struct {
	Name    string
	Type    string
	Tag     string
	Comment string
}

type StructTemplateDB struct {
	TableName string
	Columns   []*StructColumn
}

type StructTemplateDBPackage struct {
	PackageName string
	DB          StructTemplateDB
}

func NewStructTemplate() *StructTemplate {
	return &StructTemplate{structTpl: structTpl}
}

func (t *StructTemplate) AssemblyColumns(tbColumns []*TableColumn) []*StructColumn {
	tplColumns := make([]*StructColumn, 0, len(tbColumns))
	for _, column := range tbColumns {
		//`gorm:"create_user" form:"createUser" json:"createUser"`
		tag := fmt.Sprintf("`"+"gorm:"+"\"%s\""+" form:"+"\"%s\""+" json:"+"\"%s\""+"`", column.ColumnName, ToCamelCase(column.ColumnName), ToCamelCase(column.ColumnName))
		tplColumns = append(tplColumns, &StructColumn{
			Type:    DBTypeToStructType[column.DataType],
			Tag:     tag,
			Comment: column.ColumnComment,
			Name:    column.ColumnName,
		})
	}
	return tplColumns
}

func (t *StructTemplate) Generate(packageName string, tableName string, tplColumns []*StructColumn, filePwd string, fileName string) error {
	//解析模板文件
	// t1, err :=template.ParseFiles("./test.tpl")
	//template.Must 检测模板是否正确，例如大括号是否匹配，注释是否正确的关闭，变量是否正确的书写
	tpl := template.Must(template.New(packageName).Funcs(template.FuncMap{
		"ToBigCamelCase": ToBigCamelCase,
		"ToCamelCase":    ToBigCamelCase,
	}).Parse(t.structTpl))

	tplDB := StructTemplateDBPackage{
		PackageName: packageName,
		DB: StructTemplateDB{
			TableName: tableName,
			Columns:   tplColumns,
		},
	}

	//创建文件句柄
	f, er := os.Create(fmt.Sprintf("%s%s.go", filePwd, fileName))
	if er != nil {
		panic(er)
	}

	/*
		func (t Template) ExecuteTemplate(wrio.Writer, namestring, data interface{})error
		渲染模板,第一个参数是io类型,可以选择写入文件或者是控制台os.Stdout
	*/
	err := tpl.Execute(f, tplDB)
	if err != nil {
		return err
	}

	return nil
}

// ToBigCamelCase 大驼峰
func ToBigCamelCase(s string) string {
	words := regexp.MustCompile("-|_").Split(s, -1)

	for i, w := range words {
		words[i] = strings.Title(w)
	}
	return strings.Join(words, "")
}

// ToCamelCase 小驼峰
func ToCamelCase(s string) string {
	words := regexp.MustCompile("-|_").Split(s, -1)

	for i, w := range words[1:] {
		words[i+1] = strings.Title(w)
	}
	return strings.Join(words, "")
}
