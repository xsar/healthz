# healthz

这是一个很简单的域名health检查程序，可以用来监控进程http接口是否工作正常，进而得悉进程是否挂掉。如果挂掉了，就调用报警发送接口发送报警

## 报警发送接口

因为各个公司的报警发送接口可能不一样，我们只能制定规范，然后让各个公司编写适配接口。我们的接口规范是，一旦发生报警，就会调用sender接口：

```
method: post
params:
  - tos: 在配置文件中配置的报警接收人，通常是逗号分隔的手机号，当然了，您也可以配置成邮箱，sender接口来解析，healthz会原封不动的传递参数
  - content: 报警内容，包括调用失败的url和可选的备注信息
```

## 规范约定

什么样的域名被认为是工作正常呢？

- http response code = 200
- http response body 包含配置的特定字符串

## 使用方法

```bash
./control status
./control start
./control stop
./control restart
./control tail
./control reload
```
