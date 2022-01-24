# goping-pkg

- 基于 [https://github.com/go-ping/ping] 修改而来  
- 修改记录可以参考 [https://github.com/lishouzheng/ping]  
- 只支持IPv4, IPv6支持在计划中  
- 注释与测试用例缺少  

## 适用场景

应用于一个拥有数千台设备的集群中, 对于每一台设备部署一个客户端, 客户端的其中一个功能就是持续不断的进行ping操作, goping-pkg是其中的一部分底层工具包

## 性能

每台设备每分钟ping 30000个IP, 消耗CPU为7%, 内存为200MB左右

## 相对于go-ping特点

1. 取消命令行操作
2. 分离listen和send操作, 适用于大量ping的场景
3. 使用双层校验, 对埋点UUID和回包ID进行双层校验, 优化性能
4. 使用原子锁解决双层校验加锁频繁问题
5. 支持Reset操作, 可以复用pingTask, 优化性能

## 使用方式

```golang
    pinger := Default(NoopLogger{})
    pp := PingIPTask{}
    pp2 := PingIPTask{}
    pp.New("baidu.com", 5, NoopLogger{}, pinger)
    pp2.New("baidu.com", 5, NoopLogger{}, pinger)
    pp.Start()
    pp2.Start()
    fmt.Println(pp.Rst())
    fmt.Println(pp2.Rst())
```
