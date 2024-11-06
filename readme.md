# Gin控制器层级管理系统

## 简介

这是一个基于Gin框架的控制器层级路由管理的Demo，支持多层控制器嵌套，提供灵活的路由注册机制。

## 功能特点

- 支持控制器(路由)的层级嵌套
- 全局路由注册表
- 支持路由组

## 示例结构

系统包含三层控制器示例：
```
/api/father           -> FatherController（父控制器）
    /son             -> SonController（子控制器）
        /grandson    -> GrandsonController（孙控制器）
```

## 快速开始

1. 创建控制器
```go
type YourController struct {
    *BaseController
}

func NewYourController() *YourController {
    return &YourController{
        BaseController: NewBaseController(),
    }
}

func (y *YourController) Init(router gin.IRouter) {
    y.BaseController.Init(router)
    router.GET("/hello", y.Hello)
}
```

2. 注册控制器
```go
func init() {
    Register("/api/your-path", NewYourController())
}
```

3. 启动应用
```go
func main() {
    engine := gin.Default()
    manager := controller.NewControllerManager(engine)
    manager.Init()
    engine.Run(":8080")
}
```

## 许可证

MIT License
