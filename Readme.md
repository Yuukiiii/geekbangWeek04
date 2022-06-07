# geekbangWeek04

## quiz

1. 按照自己的构想，写一个项目满足基本的目录结构和工程，代码需要包含对数据层、业务层、API 注册，以及 main 函数对于服务的注册和启动，信号处理，使用 Wire 构建依赖。

2. 可以使用自己熟悉的框架。

## introduction

由于对 go 入门较浅，还没有掌握 gRPC 和 kratos 的实际用法，没有很好的实现对数据层、业务层、API 的注册。

参考学习了 beer_shop 的例子，能够运行起来。但不是我个人的劳动成果，故不放在作业提交内容中。

作业基于原生 http 库，完成了 main 函数对于服务的注册和启动，信号处理，使用 Wire 构建依赖，依赖的内容包括一个 http server 和一个简单的 fmt.Println。
