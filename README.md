# 使用说明：
> 本项目立项初衷为减少开发人员书写代码量，特别是设计了数据库后还要重新将字段转换为结构体，故此设计了此工具。
> 说白了...... 就是为了偷懒O(∩_∩)O哈哈~

## config 配置文件说明
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
## .exe 文件使用说明
> 下载压缩文件解压即可

## 拉取项目编译
```shell
go mod tidy
go run main.go
```
## 生成文件如下  entity/usersModel.go
> 方便快捷的使用数据库字段
```go 

package entity


type UsersBase struct{
  //account  
	Account	string `gorm:"account" form:"account" json:"account"` 
 	//头像  
	Avatar	string `gorm:"avatar" form:"avatar" json:"avatar"` 
 	//创建时间  
	CreateTime	string `gorm:"create_time" form:"createTime" json:"createTime"` 
 	//创建人  
	CreateUser	int32 `gorm:"create_user" form:"createUser" json:"createUser"` 
 	//主键  
	Id	int32 `gorm:"id" form:"id" json:"id"` 
 	//是否已删除  
	IsDeleted	int32 `gorm:"is_deleted" form:"isDeleted" json:"isDeleted"` 
 	//姓名  
	Name	string `gorm:"name" form:"name" json:"name"` 
 	//昵称  
	Nickname	string `gorm:"nickname" form:"nickname" json:"nickname"` 
  //password  
	Password	string `gorm:"password" form:"password" json:"password"` 
 	//1 为pc 端  
	PlatformId	int32 `gorm:"platform_id" form:"platformId" json:"platformId"` 
 	//角色  
	RoleId	int32 `gorm:"role_id" form:"roleId" json:"roleId"` 
 	//-1 未知 1 男 2 女  
	Sex	int32 `gorm:"sex" form:"sex" json:"sex"` 
 	//1 正常 2 禁用  
	Status	int32 `gorm:"status" form:"status" json:"status"` 
 	//租户ID  
	TenantId	string `gorm:"tenant_id" form:"tenantId" json:"tenantId"` 
 	//修改时间  
	UpdateTime	string `gorm:"update_time" form:"updateTime" json:"updateTime"` 
 	//修改人  
	UpdateUser	int32 `gorm:"update_user" form:"updateUser" json:"updateUser"` 

}

type Users struct{
  //account  
	Account	string `gorm:"account" form:"account" json:"account"` 
 	//头像  
	Avatar	string `gorm:"avatar" form:"avatar" json:"avatar"` 
 	//创建时间  
	CreateTime	string `gorm:"create_time" form:"createTime" json:"createTime"` 
 	//创建人  
	CreateUser	int32 `gorm:"create_user" form:"createUser" json:"createUser"` 
 	//主键  
	Id	int32 `gorm:"id" form:"id" json:"id"` 
 	//是否已删除  
	IsDeleted	int32 `gorm:"is_deleted" form:"isDeleted" json:"isDeleted"` 
 	//姓名  
	Name	string `gorm:"name" form:"name" json:"name"` 
 	//昵称  
	Nickname	string `gorm:"nickname" form:"nickname" json:"nickname"` 
  //password  
	Password	string `gorm:"password" form:"password" json:"password"` 
 	//1 为pc 端  
	PlatformId	int32 `gorm:"platform_id" form:"platformId" json:"platformId"` 
 	//角色  
	RoleId	int32 `gorm:"role_id" form:"roleId" json:"roleId"` 
 	//-1 未知 1 男 2 女  
	Sex	int32 `gorm:"sex" form:"sex" json:"sex"` 
 	//1 正常 2 禁用  
	Status	int32 `gorm:"status" form:"status" json:"status"` 
 	//租户ID  
	TenantId	string `gorm:"tenant_id" form:"tenantId" json:"tenantId"` 
 	//修改时间  
	UpdateTime	string `gorm:"update_time" form:"updateTime" json:"updateTime"` 
 	//修改人  
	UpdateUser	int32 `gorm:"update_user" form:"updateUser" json:"updateUser"` 

}

```


