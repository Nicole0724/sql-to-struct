# 使用说明：
> 本项目立项初衷为减少开发人员书写代码量，特别是设计了数据库后还要重新将字段转换为结构体，故此设计了此工具。
> 说白了...... 就是为了偷懒O(∩_∩)O哈哈~

## 其他推荐

gitee 新增方案如下：<br>
通过mod管理，然后走testing 方式来使用<br><br>
仓库lib 地址：https://gitee.com/nicole-go-libs/sql-to-struct-lib <br>
lib 使用demo：https://gitee.com/nicole-go-libs/sql-to-struct-lib-demo
```shell
go get gitee.com/nicole-go-libs/sql-to-struct-lib
```

## 更新相关
2022-11-28 V1.0.5
> 新增 通过 表名 模糊匹配来转换功能

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

type DictBase struct{
	//字典码 
	Code	string `gorm:"code" form:"code" json:"code"` 
	//字典值 
	DictKey	string `gorm:"dict_key" form:"dictKey" json:"dictKey"` 
	//字典名称 
	DictValue	string `gorm:"dict_value" form:"dictValue" json:"dictValue"` 
	//主键 
	Id	int32 `gorm:"id" form:"id" json:"id"` 
	//是否已封存 
	IsSealed	int32 `gorm:"is_sealed" form:"isSealed" json:"isSealed"` 
	//父主键 
	ParentId	int32 `gorm:"parent_id" form:"parentId" json:"parentId"` 
	//字典备注 
	Remark	string `gorm:"remark" form:"remark" json:"remark"` 
	//排序 
	Sort	int32 `gorm:"sort" form:"sort" json:"sort"` 
}

type Dict struct{
 	//字典码  
	Code	string `gorm:"code" form:"code" json:"code"` 
 	//创建时间  
	CreateTime	string `gorm:"create_time" form:"createTime" json:"createTime"` 
 	//创建人  
	CreateUser	int32 `gorm:"create_user" form:"createUser" json:"createUser"` 
 	//字典值  
	DictKey	string `gorm:"dict_key" form:"dictKey" json:"dictKey"` 
 	//字典名称  
	DictValue	string `gorm:"dict_value" form:"dictValue" json:"dictValue"` 
 	//主键  
	Id	int32 `gorm:"id" form:"id" json:"id"` 
 	//是否已删除  
	IsDeleted	int32 `gorm:"is_deleted" form:"isDeleted" json:"isDeleted"` 
 	//是否已封存  
	IsSealed	int32 `gorm:"is_sealed" form:"isSealed" json:"isSealed"` 
 	//父主键  
	ParentId	int32 `gorm:"parent_id" form:"parentId" json:"parentId"` 
 	//字典备注  
	Remark	string `gorm:"remark" form:"remark" json:"remark"` 
 	//排序  
	Sort	int32 `gorm:"sort" form:"sort" json:"sort"` 
 	//修改时间  
	UpdateTime	string `gorm:"update_time" form:"updateTime" json:"updateTime"` 
 	//修改人  
	UpdateUser	int32 `gorm:"update_user" form:"updateUser" json:"updateUser"` 

}



```



