# gin-gorm

本项目是基于 Gin，Gorm 设计的 API 开发框架使用 demo。

项目名称 gin-gorm，以论坛 API 为主题，设计的初衷是将其打造为高性能、功能齐全的 API 框架。基于 gin, cobra, viper, zap, gorm, redis, mysql, sqlite, email, jwt.

程序结构说明，请见 [程序结构](https://learnku.com/courses/go-api/1.17/program-structure/11772)。

## 程序结构
```
|-- README.md
|-- app                                     // 程序具体逻辑代码
|   |-- cmd                                 // 命令
|   |   |-- cache.go
|   |   |-- cmd.go
|   |   |-- key.go
|   |   |-- make                            // make 命令及子命令
|   |   |   |-- make.go
|   |   |   |-- make_apicontroller.go
|   |   |   |-- make_cmd.go
|   |   |   |-- make_factory.go
|   |   |   |-- make_migration.go
|   |   |   |-- make_model.go
|   |   |   |-- make_policy.go
|   |   |   |-- make_request.go
|   |   |   |-- make_seeder.go
|   |   |   `-- stubs                       // make 命令的模板
|   |   |       |-- apicontroller.stub
|   |   |       |-- cmd.stub
|   |   |       |-- factory.stub
|   |   |       |-- migration.stub
|   |   |       |-- model
|   |   |       |   |-- model.stub
|   |   |       |   |-- model_hooks.stub
|   |   |       |   `-- model_util.stub
|   |   |       |-- policy.stub
|   |   |       |-- request.stub
|   |   |       `-- seeder.stub
|   |   |-- migrate.go
|   |   |-- play.go
|   |   |-- seed.go
|   |   `-- serve.go
|   |-- http                                // http 请求处理逻辑
|   |   |-- controllers                     // 控制器，存放 API 和视图控制器
|   |   |   `-- api                         // API 控制器，支持多版本的 API 控制器
|   |   |       `-- v1                      // v1 版本的 API 控制器    
|   |   |           |-- auth
|   |   |           |   |-- login_controller.go
|   |   |           |   |-- password_controller.go
|   |   |           |   |-- signup_controller.go
|   |   |           |   `-- verify_code_controller.go
|   |   |           |-- base_api_controller.go
|   |   |           |-- categories_controller.go
|   |   |           |-- links_controller.go
|   |   |           |-- topics_controller.go
|   |   |           `-- users_controller.go
|   |   `-- middlewares                     // 中间件
|   |       |-- auth_jwt.go
|   |       |-- force_ua.go
|   |       |-- guest_jwt.go
|   |       |-- limit.go
|   |       |-- logger.go
|   |       `-- recovery.go
|   |-- models                              // 数据模型
|   |   |-- category                        // 单独的模型目录
|   |   |   |-- category_hooks.go           // 模型钩子文件
|   |   |   |-- category_model.go           // 模型主文件
|   |   |   `-- category_util.go            // 模型辅助方法
|   |   |-- link
|   |   |   |-- link_hooks.go
|   |   |   |-- link_model.go
|   |   |   `-- link_util.go
|   |   |-- model.go
|   |   |-- topic
|   |   |   |-- topic_hooks.go
|   |   |   |-- topic_model.go
|   |   |   `-- topic_util.go
|   |   `-- user
|   |       |-- user_hooks.go
|   |       |-- user_model.go
|   |       `-- user_util.go
|   |-- policies                            // 授权策略目录
|   |   `-- topic_policy.go
|   `-- requests                            // 请求验证目录（支持表单、标头、Raw JSON、URL Query）
|       |-- category_request.go
|       |-- login_request.go
|       |-- pagination_request.go
|       |-- password_request.go
|       |-- requests.go
|       |-- signup_request.go
|       |-- topic_request.go
|       |-- user_request.go
|       |-- validators                      // 自定的验证规则
|       |   `-- custom_rules.go
|       `-- verify_code_request.go
|-- bootstrap                               // 程序模块初始化目录
|   |-- cache.go
|   |-- database.go
|   |-- logger.go
|   |-- redis.go
|   `-- route.go
|-- config                                  // 配置信息目录
|   |-- app.go
|   |-- captcha.go
|   |-- config.go
|   |-- database.go
|   |-- jwt.go
|   |-- log.go
|   |-- paging.go
|   |-- redis.go
|   |-- sms.go
|   `-- verifycode.go
|-- database                                // 数据库相关目录
|   |-- factories                           // 模型工厂目录
|   |   |-- category_factory.go
|   |   |-- link_factory.go
|   |   |-- topic_factory.go
|   |   `-- user_factory.go
|   |-- migrations                          // 数据库迁移目录
|   |   |-- 2023_01_29_004048_add_users_table.go
|   |   |-- 2023_01_29_135113_add_categories_table.go
|   |   |-- 2023_01_29_144240_add_topics_table.go
|   |   |-- 2023_01_29_225453_add_links_table.go
|   |   |-- 2023_01_29_235040_add_fields_to_user.go
|   |   `-- migrations.go
|   `-- seeders                             // 数据库填充目录
|       |-- categories_seeder.go
|       |-- links_seeder.go
|       |-- seeder.go
|       |-- topics_seeder.go
|       `-- users_seeder.go
|-- docs                                    // api 文档
|   |-- gin-gorm.postman_collection.json    // Postmain Collection
|   `-- gin-gorm.postman_environment.json   // POSTmain Environment
|-- go.mod                                  // Go Module 依赖配置文件
|-- go.sum                                  // Go Module 模块版本锁定文件
|-- main.go                                 // 程序主入口
|-- pkg                                     // 内置辅助包
|   |-- app
|   |   `-- app.go
|   |-- auth
|   |   `-- auth.go
|   |-- cache
|   |   |-- cache.go
|   |   |-- store_interface.go
|   |   `-- store_redis.go
|   |-- captcha
|   |   |-- captcha.go
|   |   `-- store_redis.go
|   |-- config
|   |   `-- config.go
|   |-- console
|   |   `-- console.go
|   |-- database
|   |   `-- database.go
|   |-- file
|   |   `-- file.go
|   |-- hash
|   |   `-- hash.go
|   |-- helpers
|   |   `-- helpers.go
|   |-- jwt
|   |   `-- jwt.go
|   |-- limiter
|   |   `-- limiter.go
|   |-- logger
|   |   |-- gorm_logger.go
|   |   `-- logger.go
|   |-- migrate
|   |   |-- migration_file.go
|   |   `-- migrator.go
|   |-- paginator
|   |   `-- paginator.go
|   |-- redis
|   |   `-- redis.go
|   |-- response
|   |   `-- response.go
|   |-- seed
|   |   `-- seeder.go
|   |-- sms
|   |   |-- driver_aliyun.go
|   |   |-- driver_interface.go
|   |   `-- sms.go
|   |-- str
|   |   `-- str.go
|   `-- verifycode
|       |-- store_interface.go
|       |-- store_redis.go
|       `-- verifycode.go
|-- public                                  // 静态文件存放目录
|   `-- uploads                             // 用户上传文件目录
|-- routes                                  // 路由
|   `-- api.go
`-- storage                                 // 内部存储目录
    `-- logs                                // 日志存储目录
```

## 所有路由

| 请求方法  |  API 地址  | 说明  |
| ------------ | ------------ | ------------ |
| POST |    /api/v1/auth/login/using-phone | 短信 + 手机号登录 |
| POST |    /api/v1/auth/login/using-password | 手机号、用户名 + 密码 |
| POST |    /api/v1/auth/login/refresh-token | 刷下 Token |
| POST |    /api/v1/auth/password-reset/using-phone | 短信验证码密码重置 |
| POST |    /api/v1/auth/signup/using-phone | 使用手机号注册 |
| POST |    /api/v1/auth/signup/phone/exist | 手机号是否已注册 |
| POST |    /api/v1/auth/signup/email/exist | email 是否已支持 |
| POST |    /api/v1/auth/verify-codes/phone | 发送短信验证码 |
| POST |    /api/v1/auth/verify-codes/email | 发送邮件验证码 |
| POST |    /api/v1/auth/verify-codes/captcha | 获取图片验证码 |
| GET |     /api/v1/user              | 获取当前用户 |
| GET |     /api/v1/users             | 用户列表 |
| PUT |     /api/v1/users             | 修改个人资料 |
| PUT |     /api/v1/users/phone       | 修改手机号 |
| PUT |     /api/v1/users/password    | 修改密码 |
| PUT |     /api/v1/users/avatar      | 上传头像 |
| GET |     /api/v1/categories        | 分类列表 |
| POST |    /api/v1/categories        | 创建分类 |
| PUT |     /api/v1/categories/:id    | 更新分类 |
| DELETE |  /api/v1/categories/:id    | 删除分类 |
| GET |     /api/v1/topics            | 话题列表 |
| POST |    /api/v1/topics            | 创建话题 |
| PUT |     /api/v1/topics/:id        | 更新话题  |
| DELETE |  /api/v1/topics/:id        | 删除话题 |
| GET |     /api/v1/topics/:id        | 获取话题 |
| GET |     /api/v1/links             | 友情链接列表 |


## 第三方依赖

使用到的开源库：

- [gin](https://github.com/gin-gonic/gin) —— 路由、路由组、中间件
- [zap](https://github.com/gin-contrib/zap) —— 高性能日志方案
- [gorm](https://github.com/go-gorm/gorm) —— ORM 数据操作
- [cobra](https://github.com/spf13/cobra) —— 命令行结构
- [viper](https://github.com/spf13/viper) —— 配置信息
- [cast](https://github.com/spf13/cast) —— 类型转换
- [redis](https://github.com/go-redis/redis/v8) —— Redis 操作
- [jwt](https://github.com/golang-jwt/jwt) —— JWT 操作
- [base64Captcha](https://github.com/mojocn/base64Captcha) —— 图片验证码
- [govalidator](https://github.com/thedevsaddam/govalidator) —— 请求验证器
- [limiter](https://github.com/ulule/limiter/v3) —— 限流器
- [email](https://github.com/jordan-wright/email) —— SMTP 邮件发送
- [aliyun-communicate](https://github.com/KenmyZhang/aliyun-communicate) —— 发送阿里云短信
- [ansi](https://github.com/mgutz/ansi) —— 终端高亮输出
- [strcase](https://github.com/iancoleman/strcase) —— 字符串大小写操作
- [pluralize](https://github.com/gertd/go-pluralize) —— 英文字符单数复数处理


## 自定义的包

现在来看下我们自建的库：

- app —— 应用对象
- auth —— 用户授权
- cache —— 缓存
- captcha —— 图片验证码
- config —— 配置信息
- console —— 终端
- database —— 数据库操作
- file —— 文件处理
- hash —— 哈希
- helpers —— 辅助方法
- jwt —— JWT 认证
- limiter —— API 限流
- logger —— 日志记录
- migrate —— 数据库迁移
- paginator —— 分页器
- redis —— Redis 数据库操作
- response —— 响应处理
- seed —— 数据填充
- sms —— 发送短信
- str —— 字符串处理
- verifycode —— 数字验证码


## 代码行数

gin-gorm 项目总共有 5300 行代码（工具 [gocloc](https://github.com/hhatto/gocloc)）：

```
$ gocloc .
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                             117           1026            842           4311
JSON                             2              0              0            810
Markdown                         1             37              0            149
XML                              4              0              0            128
-------------------------------------------------------------------------------
TOTAL                          124           1063            842           5398
-------------------------------------------------------------------------------
```

## 所有命令

```
$ go run main.go -h
Default will run "serve" command, you can use "-h" flag to see all subcommands

Usage:
  gin-gorm [command]

Available Commands:
  cache       Cache management
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  key         Generate App Key, will print the generated Key
  make        Generate file and code
  migrate     Run database migration
  play        Likes the Go Playground, but running at our application context
  seed        Insert fake data to the database
  serve       Start web server

Flags:
  -e, --env string   load .env file, example: --env=testing will use .env.testing file
  -h, --help         help for gin-gorm

Use "gin-gorm [command] --help" for more information about a command.
```

make 命令：

```
$ go run main.go make -h
Generate file and code

Usage:
  gin-gorm make [command]

Available Commands:
  apicontroller Create api controller，exmaple: make apicontroller v1/user
  cmd           Create a command, should be snake_case, exmaple: make cmd buckup_database
  factory       Create model's factory file, example: make factory user
  migration     Create a migration file, example: make migration add_users_table
  model         Crate model file, example: make model user
  policy        Create policy file, example: make policy user
  request       Create request file, example make request user
  seeder        Create seeder file, example:  make seeder user

Flags:
  -h, --help   help for make

Global Flags:
  -e, --env string   load .env file, example: --env=testing will use .env.testing file

Use "gin-gorm make [command] --help" for more information about a command.
```

migrate 命令：

```
$ go run main.go migrate -h
Run database migration

Usage:
  gin-gorm migrate [command]

Available Commands:
  down        Reverse the up command
  fresh       Drop all tables and re-run all migrations
  refresh     Reset and re-run all migrations
  reset       Rollback all database migrations
  up          Run unmigrated migrations

Flags:
  -h, --help   help for migrate

Global Flags:
  -e, --env string   load .env file, example: --env=testing will use .env.testing file

Use "gin-gorm migrate [command] --help" for more information about a command.
```