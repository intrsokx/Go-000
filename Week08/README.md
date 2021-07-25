学习笔记

##安装Redis 并且用redis-benchmark测试redis的效率
```
brew install redis
cd /usr/local/Cellar/redis/6.0.9/bin

redis-benchmark -h 127.0.0.1 -p 6379 -t set,get -d 10 -q
...
redis-benchmark -h 127.0.0.1 -p 6379 -t set,get -d 5120 -q

#redis-cli 进入redis-cli客户端，输入info memory 查看Redis的内存占用信息
info memory
```