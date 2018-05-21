### Prometheus压力测试

* 访问 [BenchMarkExporter](https://github.com/ArugakiWei/BenchMarkExporter "Title") 执行 build.sh

* 执行下面的命令, 访问 `http://ip:19120/metrics` 出现一堆乱七八糟的数据说明可用

```
docker run -d -e n=48 -e m=100 -p 19120:9120 daocloud.io/daocloud/benchmark-exporter:v1

m为Metrics的总数 n为每个Metrics的name长度
```

* 启动多个exporter,执行start.sh,会开启50000+x个数的端口号，确保不会发生端口冲突

```
sh start.sh x n m 
ex: sh start.sh 100 24 100
x 代表exporter个数
n,m 同上 
```

* 在etcd中加入启动的exporter地址,执行insert.sh 

```
sh insert.sh etcdaddr ip x
ex: sh insert.sh 127.0.0.1:2379 192.168.1.1 100
etcdaddr 为etcd地址,填一个即可
ip 为本机ip,不要填127.0.0.1
x 同上
```

* 查看prometheus的target中是否出现新加入的exporter
* 调整x,n,m参数,查看promethues的状态
* 如果出现数据量过大导致prometheus抓取超时,调大prometheus的scrape_time与scrape_timeout继续观察
