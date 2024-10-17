# iris-ridge

[![MIT licensed][3]][4]

[3]: https://img.shields.io/badge/license-MIT-blue.svg
[4]: LICENSE

---
### 介绍
Iris-ridge 是基于 go-iris 编写的一个快速、简单且非常高效的面向对象的 Web 开发框架。该框架可以作为敏捷 WEB 应用开发和简化企业应用开发的基础。

---
### 目录结构
支持多应用模式部署，所以实际的目录结构取决于你采用的是单应用还是多应用模式，分别说明如下。

---
#### 单应用模式
默认代码架构目录就是一个单应用模式
```
Iris-ridge 架构目录
├─app            应用目录
│  ├─controller      控制器目录
│  ├─dao             数据访问对象目录
│  ├─dto             数据传输对象目录
│  ├─model           数据模型目录
│  ├─pkg             引用第三方包目录
│  ├─service         接口服务层目录
│  ├─utils           工具类函数目录
│  ├─ ...            更多类库目录
│  │
│  └─app.php         应用入口文件
│
├─common          应用抽象目录(与 app 目录类似)
│  ├─ ...            关联抽象类库目录(省略...)
│  ├─global          全局变量定义目录
│  ├─lib             内部类库文件目录
│  ├─share           内部分享类库目录
│  │
│  └─...             更多类库目录
│ 
├─config          配置目录
│  ├─abstract        配置抽象目录
│  │  ├─structs          yml 配置文件所对应的 struct 文件目录
│  │  ├─abstract.go      配置抽象和初始化函数文件
│  │  ├─database.yml     数据库配置文件
│  │  ├─migration.yml    迁移配置文件
│  │  │ 
│  │  └─...              更多 yml 配置文件
│  ├─dev             yml 配置文件目录（开发环境）
│  ├─prod            yml 配置文件目录（生产环境）
│  ├─ ...            更多 yml 配置文件目录(按环境自定义)
│  │
│  └─app.yml         应用配置文件
│
├─docs            swagger 目录
│
├─extend          应用拓展组件目录
│  ├─autoload        启动组件目录    
│  │  ├─custom           自定义引用组件目录
│  │  ├─indoor           内部引用组件目录
│  │  └─autoload.go      启动组件入口文件
│  │  
│  └─...             更多应用拓展组件目录
│
├─migration       迁移功能目录
│  ├─scripts         迁移脚本目录
│  └─migration.go       迁移(.go)抽象类文件
│
├─script          应用脚本目录
│  ├─migrate         迁移包程序目录(关联根目录 goose 程序)   
│  └─...             更多 yml 配置文件   
│
├─public          WEB目录（对外访问目录）
├─route           自定义路由目录
├─tool            自定义工具方法目录     
├─view            视图目录   
├─.gitignore      git 忽略配置管理文件
├─go.mod          依赖管理文件    
├─goose           业务迁移命令行入口文件
├─LICENSE         授权说明文件 
├─main.go         入口文件
├─README.md       README 文件
```
---
#### 多应用模式（扩展）
如果你需要一个多应用的项目架构，目录结构可以参考下面的结构进行调整（关于配置文件的详细结构参考后面章节）。
PS：与“单应用模式”一致的目录忽略...
```
Iris-ridge 架构目录
├─app            应用目录
│  ├─app_name        应用目录
│  │  ├─controller      控制器目录
│  │  ├─dao             数据访问对象目录
│  │  ├─dto             数据传输对象目录
│  │  ├─model           数据模型目录
│  │  ├─pkg             引用第三方包目录
│  │  ├─service         接口服务层目录
│  │  ├─utils           工具类函数目录
│  │  │
│  │  └─ ...            更多类库目录
│  │  
│  └─app.php      应用入口文件
│
├─...            更多程序目录(与"单应用模式"一致)
```
---
### config 配置
根目录下 config 目录中配置文件说明，当然你也可以根据自身业务需求增加自定义的配置文件。

---
#### app.yml 入口配置文件
如下配置项明确了应用启动后读取 config 目录下哪个子目录的配置文件
```
# 环境设置
env: dev 

...
```
---
#### abstract 目录说明
如下配置中 abstract 目录中 structs 中除 app.go 对应 app.yml 配置文件，其它 struct 文件与 abstract 目录下 yml 文件一一对应。如需按各类环境调整相关配置，只需要将 abstract 目录下 yml 文件拷贝到对应环境文件夹中即可，如：dev（开发环境）、prod（生产环境）...... 同时，需要修改 app.yml 中 env 配置项。如果需要新增 yml 文件，可以直接在 abstract 目录中直接新建，并在 structs 目录中创建与 yml 文件对应的 struct 文件。
```
├─config        配置目录
│  ├─abstract       默认目录
│  │  ├─structs          yml 配置文件对应 struct 文件目录 
│  │  │  ├─app.go            app.yml 对应 struct 文件
│  │  │  ├─database.go       database.yml 对应 struct 文件
│  │  │  ├─migration.go      migration.yml 对应 struct 文件
│  │  │  │
│  │  │  └─...              更多 yml 配置文件对应 struct 文件
│  │  ├─abstract.go      config 配置 init 函数文件
│  │  ├─database.yml     数据库配置文件
│  │  ├─migration.yml    迁移配置文件
│  │  │ 
│  │  └─...              更多 yml 配置文件
│  ├─dev             yml 配置文件目录（开发环境）
│  ├─prod            yml 配置文件目录（生产环境）
│  ├─ ...            更多 yml 配置文件目录(按环境自定义)
│  │
│  └─app.yml         应用配置文件
│                  
```
---
### 多应用模式 - 配置
针对启用“多应用模式”，需要调整根目录下 app 目录中 app.gp 文件中的相关引用。
```
// 单应用模式下
import (
    ...

    // for init controller route and table model
    // import all module's controllers
    _ "ridge/app/controller"
    
    // import all module's models
    _ "ridge/app/model"

    ...
）

// 多应用模式 [app_name - 应用名称]
import (
    ...

    // for init controller route and table model
    // import all module's controllers
    _ "ridge/app/<app_name>/controller"
    
    // import all module's models
    _ "ridge/app/<app_name>/model"

    ...
）
```
同时，可以针对应用名称<app_name>对应用模块下 controller 层的动态路由前缀进行调整，其调整文件为应用文件夹下 controller 目录中 controller.go 文件。
```
// 单应用模式下(默认)
// 目录 app/controller/controller.go
// Party custom routing prefix of "app" module
func (c *Controller) Party() string {
	return "app"
}

// 多应用模式 [app_name - 应用名称]
// 目录 app/<app_name>/controller/controller.go
// Party custom routing prefix of "app_name" module
func (c *Controller) Party() string {
	return "app_name"
}
```
---
### 数据迁移
数据迁移依赖 gorm 和 goose 组件，其中默认自动迁移 gorm 依赖表 model 结构体的迁移，同时支持 goose 自动迁移模式或命令行迁移模式。
##### gorm - 表结构创建
```
// 在应用目录下 model 目录中创建对应表 model文件
// 可以参考 app/model/userModel.go 文件
```
#### goose - 数据迁移配置
```
// goose 相关配置
// 需要配置 config 目录中 migration.yml
# 迁移文件目录
migration_dir: /migration/scripts

# 是否允许迁移缺失
allow_missing: true

# 是否自动执行数据迁移
auto_migrate: false
```
#### goose - 命令行操作
```
// 使用 goose 命令行
./goose up      // 执行迁移
./goose dwon    // 回退迁移
./goose status  // 查看迁移状态
./goose --help  // 可以查看相关命令格式
```
---
### 其它说明
对一些其它功能内容的使用进行简单说明，主要涉及内容：自定义路由、swagger 使用、前后端分离方式和 view 方式。

---
#### 自定义路由
自定义路由，可分为应用内部自定义路由和外部自定义路由。
** 其 controller 层的路由为动态路由，无需手动维护。
```
* 内部自定义路由
可以在根目录中 app 目录下 app.go 文件中进行应用内部自定义路由的设置。

* 外部自定义路由
可以在根目录 route 目录中 route.go 文件中进行外部自定义路由的设置。
```

---
#### 生成 swagger 文档
通过编写 swagger 规范生成 api 的接口文档
```
## swagger 范例：app/controller/userController.go

// PostLogin user login operate
// @Tags 用户功能模块
// @Summary 用户登录
// @Description 用于用户登录操作
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "auth token"
// @Param data body dto.LoginReq true "用户信息"
// @Success 200 {object} dto.APISuccess	"请求成功"
// @Failure 400 {object} dto.APIError "请求错误"
// @Router /app/user/login [post]
```
swagger 脚本
```
# swag  init
```
---
#### 注意事项
如使用 go build 生成二进制程序(main)，需要同时拷贝 config、migration、public 和 view 文件夹，并确保二进制程序(main)与相关文件夹在同一目录下，如需使用迁移命令行程序(goose)也需与其在同一目录。
```
Iris-ridge 架构目录
├─config         配置目录
├─migration      迁移功能目录（迁移脚本 .go 或 .sql ）          
├─public         WEB目录（对外访问目录）
│
├─goose          迁移二进制程序文件
├─main           入口启动二进制程序文件
│
├─...            更多程序或目录
```



