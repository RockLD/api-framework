# api-framework
使用gin实现的 restful 框架

### 一、项目结构
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
### 二、代码组织方式
___________
使用go module组织代码

### 三、启动方法
___________
1、将代码克隆到本地：

`git clone https://github.com/RockLD/api-framework.git`  

2、导入sql

`db_api.sql`

3、构建项目

```
cd api-frameword
make
```

4、管理项目

`启动项目：sh apiserver.sh start`

`停止项目：sh apiserver.sh stop`

`重启项目：sh apiserver.sh restart`

`查看状态：sh apiserver.sh status`

### 四、开发过程
___________

