listeners: 学名叫监听器，其实就是告诉外部连接者要通过什么协议访问指定主机名和端口开放的 Kafka 服务。
advertised.listeners：和 listeners 相比多了个 advertised。Advertised 的含义表示宣称的、公布的，就是说这组监听器是 Broker 用于对外发布的。

当我们对 172.17.0.10:9092 请求建立连接，kafka 服务器会通过 zookeeper 中注册的监听器，找到 INSIDE 监听器，然后通过 listeners 中找到对应的 通讯 ip 和 端口；

同理，当我们对 <公网 ip>:端口 请求建立连接，kafka 服务器会通过 zookeeper 中注册的监听器，找到 OUTSIDE 监听器，然后通过 listeners 中找到对应的 通讯 ip 和 端口 172.17.0.10:9094；

总结：advertised_listeners 是对外暴露的服务端口，真正建立连接用的是 listeners。

https://segmentfault.com/a/1190000020715650

https://blog.csdn.net/weixin_38251332/article/details/105638535
https://www.jianshu.com/p/71b295e1df4f?tdsourcetag=s_pctim_aiomsg

"PLAINTEXT"表示协议，可选的值有PLAINTEXT和SSL

https://blog.51cto.com/13231454/2457090