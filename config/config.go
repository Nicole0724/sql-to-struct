package config

var (
	Conf *SqlToStructParams
)

type SqlConnectConfig struct {
	Type         string `yaml:"Type"`
	Host         string `yaml:"Host"`
	Port         string `yaml:"Port"`
	Account      string `yaml:"Account"`
	Password     string `yaml:"Password"`
	DataBaseName string `yaml:"DataBaseName"`
}

type SqlToStructParams struct {
	SqlConfig     SqlConnectConfig `yaml:"SqlConfig"`
	PackageName   string           `yaml:"PackageName"`
	IsAllDataBase bool             `yaml:"IsAllDataBase"`
	TableName     string           `yaml:"TableName"`
	LikeTableName string           `yaml:"LikeTableName"`
	SaveFileName  string           `yaml:"SaveFileName"`
	SaveFilePwd   string           `yaml:"SaveFilePwd"`
}
