https://github.com/confluentinc/confluent-kafka-go

% go mod init kafka_learn
go: creating new go.mod: module kafka_learn

% go get -u github.com/confluentinc/confluent-kafka-go
go: downloading github.com/confluentinc/confluent-kafka-go v1.6.1
go: github.com/confluentinc/confluent-kafka-go upgrade => v1.6.1

https://github.com/confluentinc/confluent-kafka-go/tree/master/examples

https://www.zhihu.com/question/266377862
https://github.com/Shopify/sarama
https://blog.csdn.net/csrh131/article/details/103786796
https://www.cnblogs.com/Qing-840/p/11193238.html
https://blog.csdn.net/weixin_34168700/article/details/92379992

https://github.com/segmentio/kafka-go

http://errornoerror.com/question/13722150641458978435/

https://www.jianshu.com/p/3d4655cd7054

git clone https://github.com/edenhill/librdkafka.git
cd librdkafka
./configure
make
sudo make install

http://kafka.apache.org/

https://blog.csdn.net/oumuv/article/details/84860181
https://www.jianshu.com/p/903c318f71cb

brew install kafka

brew services start kafka

 zookeeper-server-start /usr/local/etc/kafka/zookeeper.properties & kafka-server-start /usr/local/etc/kafka/server.properties


 % cd /usr/local/Cellar/kafka/2.6.0_1/
  % ./bin/kafka-topics --create  --zookeeper localhost:2181 --partitions 1 --replication-factor 3  --topic test

Exception in thread "main" kafka.zookeeper.ZooKeeperClientTimeoutException: Timed out waiting for connection while in state: CONNECTING

 % ./bin/zookeeper-server-start /usr/local/etc/kafka/zookeeper.properties

  % ./bin/kafka-topics --create  --zookeeper localhost:2181 --partitions 1 --replication-factor 3  --topic test
Error while executing topic command : Replication factor: 3 larger than available brokers: 1.


 % ./bin/kafka-topics --create  --zookeeper localhost:2181 --partitions 1 --replication-factor 1  --topic test
Created topic test.

 % ./bin/kafka-topics --list --zookeeper localhost:2181
test

 % ./bin/kafka-topics --describe --zookeeper localhost:2181 --topic test
Topic: test	PartitionCount: 1	ReplicationFactor: 1	Configs:
	Topic: test	Partition: 0	Leader: 0	Replicas: 0	Isr: 0

 %  ./bin/kafka-console-producer --broker-list localhost:9092 --topic test
>This is a message
>aaa
>

 %  ./bin/kafka-console-consumer --bootstrap-server localhost:9092 --topic test --from-beginning
This is a message
aaa

% go run exp1/producer/main.go
%3|1617546888.931|FAIL|rdkafka#producer-1| [thrd:bogon:9092/0]: bogon:9092/0: Failed to resolve 'bogon:9092': nodename nor servname provided, or not known (after 2ms in state CONNECT)

 % go run exp1/consumer/main.go
%3|1617546969.406|FAIL|rdkafka#consumer-1| [thrd:GroupCoordinator]: GroupCoordinator: bogon:9092: Failed to resolve 'bogon:9092': nodename nor servname provided, or not known (after 2ms in state CONNECT)

https://blog.csdn.net/qq_36141369/article/details/108125648


 % vi /usr/local/etc/kafka/server.properties
advertised.listeners=localhost:9092

 %   brew services restart kafka

 % go run exp1/producer/main.go
%3|1617547936.799|FAIL|rdkafka#producer-1| [thrd:localhost:9092/bootstrap]: localhost:9092/bootstrap: Connect to ipv6#[::1]:9092 failed: Connection refused (after 3ms in state CONNECT)


https://blog.csdn.net/fst438060684/article/details/80662305

kafka连接生产者（消费者其实也一样的问题）出现了下面这个报错：（

WARN [Producer clientId=console-producer] Connection to node -1 could not be established. Broker may not be available. (org.apache.kafka.clients.NetworkClient)

在配置plaintext的时候，地址配置的是master（master是我在/etc/hosts文件里面配置的127.0.0.1），而我连接的时候，使用的是：

 bin/kafka-console-producer.sh --broker-list localhost:9092 --topic test

改为：

 bin/kafka-console-producer.sh --broker-list master:9092 --topic test

perfect，完美的解决了（生产者也一样哦）

Consumer error: GroupCoordinator: bogon:9092: Failed to resolve 'bogon:9092': nodename nor servname provided, or not known (after 2ms in state CONNECT, 8 identical error(s) suppressed)

% cd librdkafka
% ./configure 
checking for OS or distribution... ok (osx)

 % make -j 4
Using ctags to generate TAGS

% sudo make install

 % cd /usr/local/Cellar/kafka/2.6.0_1
  %  ./bin/kafka-topics --describe --zookeeper localhost:2181 --topic test
Topic: test	PartitionCount: 1	ReplicationFactor: 1	Configs:
	Topic: test	Partition: 0	Leader: 0	Replicas: 0	Isr: 0

 %  ./bin/kafka-topics --describe --zookeeper localhost:2181 --topic __consumer_offsets
Topic: __consumer_offsets	PartitionCount: 50	ReplicationFactor: 1	Configs: compression.type=producer,cleanup.policy=compact,segment.bytes=104857600
	Topic: __consumer_offsets	Partition: 0	Leader: 0Replicas: 0	Isr: 0

https://stackoverflow.com/questions/63925892/connecting-spark-streaming-from-local-machine-to-kafka-on-gcp

% vi /usr/local/etc/kafka/server.properties
#advertised.listeners=PLAINTEXT://your.host.name:9092
advertised.listeners=PLAINTEXT://localhost:9092


 % brew services  restart kafka
Stopping `kafka`... (might take a while)
==> Successfully stopped `kafka` (label: homebrew.mxcl.kafka)
==> Successfully started `kafka` (label: homebrew.mxcl.kafka)


"bootstrap.servers": "localhost",  

% go run exp1/consumer/main.go
Consumer error: Subscribed topic not available: ^aRegex.*[Tt]opic: Broker: Unknown topic or partition (<nil>)
Message on test[0]@0: This is a message
Message on test[0]@1: aaa
Message on test[0]@2: 123
Message on test[0]@3: 1234

问题解决

% go run exp1/producer/main.go
Delivered message to test[0]@4
Delivered message to test[0]@5
Delivered message to test[0]@6
Delivered message to test[0]@7
Delivered message to test[0]@8
Delivered message to test[0]@9
Delivered message to test[0]@10

% go run exp1/consumer/main.go
Consumer error: Subscribed topic not available: ^aRegex.*[Tt]opic: Broker: Unknown topic or partition (<nil>)
Message on test[0]@4: Welcome
Message on test[0]@5: to
Message on test[0]@6: the
Message on test[0]@7: Confluent
Message on test[0]@8: Kafka
Message on test[0]@9: Golang
Message on test[0]@10: client
^Csignal: interrupt