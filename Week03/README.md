Week03 作业题目：

1. 基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。

## Goutine管理
一定要做好Goroutine的退出机制才能启动
创建goroutine的管理者与执行者