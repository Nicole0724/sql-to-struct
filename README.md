# 使用说明：
## .exe 文件使用说明, config 配置文件说明
```yaml
PackageName: entity
IsAllDataBase: true
TableName: dict
SaveFileName: dictModel
SaveFilePwd: "D:\\work_space\\open_source\\go-libs\\sql-to-struct\\entity\\" #绝对路径方式
#SaveFilePwd: './entity/' #相对路径方式
SqlConfig:
  Type: mysql # 数据库类型
  Host: 127.0.0.1 # 填写你的数据库账号
  Port: 3306 # 端口号
  Account: root # 账号
  Password: root #密码
  DataBaseName: go_frame # 数据库名称
```