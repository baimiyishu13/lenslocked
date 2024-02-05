# 代码组织
- 有些代码库很容易导航
- 其他人很难效仿
- 代码组织往往是不同的

# 个好的代码结构将
- 更容易猜测bug在哪里
- 更容易添加新功能
- 还有更多

# 扁平结构
- 所有代码都在一个包中
- 文件仍然可以分离代

# 关注点分离分离代码
基于职责: HTML CSS是一个传统的例子
- HTML 注重整体结构
- CSS专注于对其进行样式设置

-模型视图控制器（MVC）是遵循这种策略的一种流行结构。

Ina web应用程序：
- 模型=>数据、逻辑、规则；通常是数据库
- 视图=渲染事物；通常为html
- 控制器=将其全部连接起来。接受用户输入将其传递给模型

做一些事情，然后将数据传递到视图以渲染事情；通常是处理程序

```
myapp/
    controllers/
        user_handler.go
        gallery_handler.go
        ...
    views/
        user_templates.go
        ...
    models/
        user_story.go
        ...
```
不需要命名为模型视图和控制器
[buffalol(https://gobuffalo.io/en/) 使用Mvc的变量，但并不完全正确。