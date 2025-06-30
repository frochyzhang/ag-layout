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

- env配置刷新到结构体，目前已支持acos配置更新后刷新到env。

### **5.4 观测**

- 日志脱敏
- 日志服务平台接入
- 链路追踪（skywalking）

### **5.5 grpc**

- 服务发现：基于现有外卡的网关多注册功能，单元服务需要识别调用来源网关，需自定义服务发现的LB层

### **5.6 redis & kafka**

- 待实现

### **5.7 固定线程数&连接数**

## **6.外卡转GO**

### 6.1 当前服务模块

- 网关服务

| cybersource网关 | tgwc-service |
| --- | --- |
| 商户通知网关 | mgwi-service |
| 核验网关 | mvgw-service |
| 商户网关 | mgwo-service |
| 文件下载网关 | mgwf-service |
| VISA网关 | visagw-service |
| MasterCard网关 | mcgw-service |
| JCB网关 | jcbgw-service |
| 线下交易网关 | mgwp-service |
- 单元服务

| 线上联机服务 | online-service |
| --- | --- |
| 超时处理服务 | tmout-service |
| 卡组路由服务 | cors-service |
| 文件下载服务 | fltsf-service |
| 风险服务 | risk-service |
| 核验服务 | mvs-service |
| 参数管理、加密机服务 | syspm-service |
| 网管服务 | netmgr-service |
| 线下联机服务 | pos-service |
- 辅助服务

| 内部交易网关 | web-internal-service |
| --- | --- |
| 状态服务 | sss-service |
- 控台服务

| cas后端服务 | cas-management |
| --- | --- |
| cas server | cas-overlay |
| 控台网关 | web-gateway |
| 查询服务 | web-query-service |
| 业务服务 | web-business-service |
| 下载服务 | web-download-service |
| 定时服务 | web-quartz-service |
| 控台辅助服务 | web-assist-service |

### 6.3 转go路线

综合影响性、稳定性、改造难度等多方面评估，建议转GO路线如下：

- 按服务类别
    - 辅助服务 → 单元服务 → 网关服务 → 控台服务
- 按服务类别内部细分
    - 辅助服务
        - 内部交易网关：仅用于支持控台发起的手工退货，影响小、稳定、改造难度不大
        - 状态服务： 集群状态管理服务，功能稳定性高
    - 单元服务
        - 超时处理服务：主要接收网关因调用联机服务超时后触发的补偿机制，如冲正，影响小、稳定、改造难度不大
        - 卡组路由服务：仅JCB网关寻找单元服务使用，稳定、改造难度不大
        - 文件下载服务
        - 核验服务
        - 网管服务
        - 参数管理、加密机服务
        - 风险服务
        - 线上联机服务
        - 线下联机服务
    - 网关服务
        - 核验网关
        - 商户通知网关
        - 文件下载网关
        - cybersource网关
        - VISA网关
        - MC网关
        - JCB网关
        - 商户网关
        - 线下交易网关
    - 控台服务
        - 控台辅助服务
        - 定时服务
        - 下载服务
        - 查询服务
        - 业务服务
        - 控台网关
        - CAS：非GRPC服务，为SofaRpc服务

### 7. 排期计划（预估）

1. ag-core基础架构——7月底
    1. 日会机制：同步当前进展，揭示当前风险
    2. PR评审：日会前提交当日任务PR，日会时对PR内容进行评审，评审通过后接受PR。
    3. 功能点设计方案评审，对于新增功能点，需先通过方案评审，重大功能点需通过架构评审方可开始开发。
2. 外卡各模块改造——7月中～9月底（ag-core通用功能开发完成后，即可开始外卡相关改造）
    1. 日会机制：同步当前进展，揭示当前风险
    2. 代码评审：每日日会前提交各负责模块代码，日会时对代码进行评审
    3. 各模块方案设计&评审，重要模块的方案需通过架构评审，评审后方可开始开发。
3. 测试准出（SIT&RC测试预估1个月时间）
4. 上线思路（逐步改造、逐步上线）
    1. 相应模块改造开发完成、符合准出条件后，线上通过流量灰度（复用服务合并时的上线思路）