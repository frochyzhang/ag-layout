# aif-go Project

## **1.安装*aif-go***

```
go install github.com/frochyzhang/ag-layout/cmd/aif-go@latest

```

## **2.基于模版创建服务**

```
# 初始化工程目录
aif-go new server

cd server
# 添加proto模版
aif-go proto add api/server/server.proto
# 生成grpc-server、http-server、vo、swagger API文档等
make api
# 生成grpc&http的module注入
make server
# 基于proto生成service、service的module注入、service层中间件（如事务）
make service
# 运行代码（需要修改Nacos地址）
aif-go run server

```

## **3.在Docker中运行**

```bash
# build
docker build -t <your-docker-image-name> .

# run
docker run --rm -p 8000:8000 -p 9000:9000  <your-docker-image-name>

```

## **4.ag-core功能点**

### **4.1 ag_app**

```
程序启动入口，负责格协议（http/grpc/socket)服务的启动和停止

```

### **4.2 ag_conf**

```
配置处理模块，类spring-boot config
本地配置、环境变量加载
配置解析注入，实体数据的注入构建
配置源扩展支持
配置源的优先级

```

### **4.3 ag_db**

```
基于gorm的封装
提供事务功能（事务待优化）
db2数据库的支持

```

### **4.4 ag_nacos**

```
配置构造NamingClient、ConfigClient
nacos远程配置的扩展实现

```

### **4.5 ag_log**

```
日志配置模块
集成zap日志框架
包装zap为slog提供handler，项目中使用slog接口实现日志实现的解耦，类似slf4j

```

### **4.6 ag_hertz**

```
http服务模块，集成字节的hertz框架，由ag_app统一维护服务启停
hertz服务初始化
服务注册

```

### **4.7 ag_ext**

```
部分扩展功能
kitex客户端resolver识别grpc-spring注册的服务
ipRange扩展
CopyOnWriteSlice的实现，目前用于ag_conf对sources的维护

```

### **4.8 fxs**

```
使用fx依赖注入框架对各组件进行注入整合，类构造注入

```

### **4.9 ag_cache**

```
基于allegro的bigcache缓存框架封装的本地缓存框架

```

### **4.10 ag_crypto**

```
加解密模块
提供统一接口抽象，目前仅实现了base64的加解密

```

### **4.11 ag_error**

```
异常处理模块（待完全实现）

```

### **4.12 ag_kitex**

```
grpc服务模块，集成字节的kitex框架，由ag_app统一维护服务启停
kitex服务初始化
服务注册
自定义扩展，如ipRange（sofarpc）功能，grpc服务的配置
简单的中间件逻辑的注入

```

### **4.13 ag_netty**

```
socket模块，集成字节的netpoll框架，server由ag_app统一维护服务启停,client通过fx注入后使用
自定义连接生命周期事件
自定义编解码器
处理器统一维护

```

## **5.ag-core缺失功能点**

### **5.1 service层**

- service层代码生成注入Repository
- 基于中间件处理service层事务
- 事务传播

### **5.2 db层**

- 分页查询
- 多数据源

### **5.3 conf层**

- env配置刷新到结构体，目前已支持nacos配置更新后刷新到env。

### **5.4 观测**

- 日志脱敏
- 日志服务平台接入
- 链路追踪（skywalking）

### **5.5 grpc**

- 服务发现：基于现有外卡的网关多注册功能，单元服务需要识别调用来源网关，需自定义服务发现的LB层

### **5.6 redis & kafka**

- 待实现

### **5.7 固定线程数&连接数**

### **5.8 gateway（关键）**
- url路由
- sink机制
- 通讯日志