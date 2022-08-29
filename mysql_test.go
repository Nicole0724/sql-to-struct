package sqltostruct

import (
	"fmt"
	"testing"
)

func TestGetTable(t *testing.T) {
	basePath := "D:\\work_space\\open_source\\go-libs\\sql-to-struct\\entity\\"
	var sqlToStructParams SqlToStructParams
	sqlToStructParams.SqlConfig.Type = "mysql"
	sqlToStructParams.SqlConfig.Host = "127.0.0.1"
	sqlToStructParams.SqlConfig.Port = "3306"
	sqlToStructParams.SqlConfig.Account = "root"
	sqlToStructParams.SqlConfig.Password = "root"
	sqlToStructParams.SqlConfig.DataBaseName = "go_frame"

	sqlToStructParams.PackageName = "entity" //"model"
	sqlToStructParams.TableName = "dict"
	sqlToStructParams.SaveFileName = "dictModel"
	sqlToStructParams.SaveFilePwd = fmt.Sprintf("%s", basePath)

	sqlToStructParams.SqlTwoStruct()

}

func TestTempMain(t *testing.T) {
	TempMain()
}
