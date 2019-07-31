# api-framework
使用gin实现的 restful 框架

### 项目结构
___________
```
  check/                健康检查    
  conf/                 配置文件    
  config/               加载配置文件方法    
  handler/              请求处理    
  log/                  日志  
  middleware/           中间件 
  model/                数据层 
  pkg/                  一般使用包   
  router/               路由  
  service/              处理负责逻辑，从handler中分离  
  util/                 通用功能    
```
### 代码组织方式
使用go module组织代码

### 启动方法
1、将代码克隆到本地：

`git clone https://github.com/RockLD/api-framework.git`  

2、构建项目

`go build -v`

3、启动项目

`./api`

4、成功界面

```
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] POST   /login                    --> api/handler/demo.Login (7 handlers)
[GIN-debug] POST   /v1/demo                  --> api/handler/demo.Create (8 handlers)
[GIN-debug] DELETE /v1/demo/:id              --> api/handler/demo.Delete (8 handlers)
[GIN-debug] PUT    /v1/demo/:id              --> api/handler/demo.Update (8 handlers)
[GIN-debug] GET    /v1/demo                  --> api/handler/demo.List (8 handlers)
[GIN-debug] GET    /v1/demo/:username        --> api/handler/demo.Get (8 handlers)
[GIN-debug] GET    /check/health             --> api/check.HealthCheck (7 handlers)
{"level":"INFO","timestamp":"2019-07-31 09:43:12.868","file":"api-framework/main.go:63","msg":"Waiting for the router , retry in 1 second"}
{"level":"INFO","timestamp":"2019-07-31 09:43:12.882","file":"api-framework/main.go:50","msg":"Start listen the address: :8080"}
```

