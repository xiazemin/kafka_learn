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