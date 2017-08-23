# logagent use doc


## config file

    $ cd logagent
    $ vim conf/logagent

```conf
[log]

# logagent log config 

log_level = debug
log_path = h:/logs/logagent.log

[kafka]
# kafka server ip and port
server_addr=192.168.11.128:9092 


[collect]
# colect log path
log_path=H:/360_update/oldboy_go/src/go_dev/day11/project/logagent/logs/access.log

# use kafka's topic name
topic=nginx_log

# config log chan size
chan_size=100

```

## 编译并运行

    $ go build go_dev/day11/project/logagent/main
    $ ./main


## 观察logagent日志变化情况，是否出现error情况
    $ tailf logagent.log

## 监控kafka topic通道中的日志
> 搭建kafka文档：http://blog.csdn.net/zhaogaolongsina/article/details/77477197

    $ cd kafka
    $ ./bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic nginx_log --from-beginning


