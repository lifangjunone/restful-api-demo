# restful-api-demo
go restful api demo

**组织结构：**
```text
|-cmd                         # 程序Cli工具包
|-conf                        # 程序配置文件
|-protocol                    # 程序监听的协议
|-version                     # 程序自身的版本信息
|-app                         # 业务包
| |-host
|   |-model                   # 业务需要的数据模型
|   |-interface               # 业务接口
|   |-impl                    # 业务具体实现
| |-mysql                     # 数据库
| |-lb
|-main                        # 程序入口文件


```

