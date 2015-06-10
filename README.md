# healthz

这是一个很简单的域名health检查程序，可以用来监控进程http接口是否工作正常，进而得悉进程是否挂掉

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
