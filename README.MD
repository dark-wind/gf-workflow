# 工作流

`gf-workflow` 是一个基于GoFrame2框架开发的微型流程引擎，开发目的是学习GoFrame框架并完成课程作业，后续有时间会完善和添加更多功能

# Features
- 基于GoFrame2框架，易扩展
- 仅使用mysql存储流程定义，无需BPML等业务流程建模语言的学习成本
- 结构简单，甚至抛弃了大部分GoFrame的工程目录结构，易于理解
- OpenAPIV3 自动生成可以直接操作的文档
- 类微服务的架构，可以直接调用API进行接入
- 可以实现会签、条件选择等特殊流程节点
- 测试驱动开发，覆盖了大部分的流程操作函数

# 安装使用

## 下载项目并安装依赖
```
git clone https://github.com/dark-wind/gf-workflow.git
cd gf-workflow
go mod tidy
```

## 初始化数据结构
导入 sql/workflow.sql

只有5张表，及其简单
## 运行
```
go run main.go
```

# 演示示例



# Limitation
```
golang version >= 1.15
mysql version >=5.8 (使用了json类型，5.7是不支持的哟)
```

# 架构
现在分为用户管理、流程管理、流程引擎、演示界面，4个模块

# 代理设置
git config --global http.proxy http://127.0.0.1:7890
git config --global --unset http.proxy

#文档（基本等于没写）
## 核心路由

可以调试的文档

演示界面

## task的状态
task的状态现在有3种：
- run 正在流转
- finish 已完成
- fail 被驳回、已经中止，不能再处理

后面准备增加：
- rejected 被驳回到某个节点，可以再继续处理

