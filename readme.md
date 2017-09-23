# Werewolf （通用接口服务器框架）

> 为了快速构建后端接口，遂设计了此通用框架。框架底层基于 Echo (高性能的web框架)。 Model层推荐采用Gorm(一个优雅的对象关系映射库)
当然，你也可以通过框架内部的中间件，引入自己喜欢的 ORM 库。

## 依赖

* github.com/labstack/echo
* github.com/jinzhu/orm
* gopkg.in/yaml.v2

## 若你想在你的项目里面使用它，请...

```
go get -u github.com/alixez/werewolf
```

**如果你经常使用glide作为依赖管理工具**

```
glide get -u github.com/alixez/werewolf
```

> 哈哈，当然现在项目还处于初始阶段。但是随着我们 异样（一个专属年轻人的，潮流文化社区） 项目的开发里程的不断推进，这个框架也会不断的完善。
最终，当它可公开使用时，我们就开始写文档，并正式发布这个框架。waiting...


> 最后，感谢下这个叫 Ningxin的童鞋。因为差一点名字就变成 ningxin 了。想想，这个是一件多么恐怖的事情 ^(o^^o)!


