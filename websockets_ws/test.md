# Queue Performance

## Architecture

### Hardware

r4.xlarge x 3

The spec of each r4.xlarge
vCPU:4
Memory:30.5GiB
EBS storage is gp2(100/3000 iops) 20GiB

They are denated below as kz1,kz2,kz3

# Kafka

## Setup

kz1: zookeeper1+kafka1
kz2: zookeeper2+kafka2
kz3: zookeeper3+kafka3

## zookeeper settings

Keeping default settings:
```
maxClientCnxns=0
initLimit=5
syncLimit=2
tickTime=2000
```
(end-point settings omitted)

## server settings extract(for reference)
```
# End-point-specific settings
broker.id=10 (differs in different server)
listeners=PLAINTEXT://kz1:9092 (differs in different server)
zookeeper.connect=kz1:2181,kz2:2181,kz3:2181

# Default settings
num.network.threads=3
num.io.threads=8
socket.send.buffer.bytes=102400
socket.receive.buffer.bytes=102400
socket.request.max.bytes=104857600
num.partitions=1
num.recovery.threads.per.data.dir=1
log.retention.hours=168
log.segment.bytes=1073741824
log.retention.check.interval.ms=300000
zookeeper.connection.timeout.ms=6000

# Settings for testing
group.initial.rebalance.delay.ms=0 #for quick response

# overridden settings (suggested production setting for HA of metadata)
offsets.topic.replication.factor=3 
transaction.state.log.replication.factor=3
transaction.state.log.min.isr=3
```



## Test method

### Built-in kafka-perf suite

#### Test to produce 1000000 records size of 1000 bytes
```
bin/kafka-producer-perf-test.sh \
  --topic test \
  --num-records 1000000 \
  --record-size 1000 \
  --throughput -1 \
  --producer-props acks=1 \
  bootstrap.servers=kz1:9092 \
  buffer.memory=67108864 \
  batch.size=8196
```

##### Result

100465 records sent, 20093.0 records/sec (19.16 MB/sec), 1960.9 ms avg latency, 2850.0 max latency.
179504 records sent, 35900.8 records/sec (34.24 MB/sec), 1988.7 ms avg latency, 2748.0 max latency.
189934 records sent, 37986.8 records/sec (36.23 MB/sec), 1738.4 ms avg latency, 1798.0 max latency.
169802 records sent, 33960.4 records/sec (32.39 MB/sec), 1861.0 ms avg latency, 2083.0 max latency.
178648 records sent, 35729.6 records/sec (34.07 MB/sec), 1883.7 ms avg latency, 2070.0 max latency.
1000000 records sent, 33654.169752 records/sec (32.10 MB/sec), 1848.76 ms avg latency, 2850.00 ms max latency, 1757 ms 50th, 2630 ms 95th, 2745 ms 99th, 2843 ms 99.9th.

#### Test to produce 2000000 records size of 500 bytes
```
bin/kafka-producer-perf-test.sh \
  --topic test \
  --num-records 2000000 \
  --record-size 500 \
  --throughput -1 \
  --producer-props acks=1 \
  bootstrap.servers=kz1:9092 \
  buffer.memory=67108864 \
  batch.size=8196
```

##### Result

179506 records sent, 35901.2 records/sec (17.12 MB/sec), 1829.3 ms avg latency, 2997.0 max latency.
299910 records sent, 59982.0 records/sec (28.60 MB/sec), 2267.6 ms avg latency, 2979.0 max latency.
333450 records sent, 66690.0 records/sec (31.80 MB/sec), 1839.4 ms avg latency, 1933.0 max latency.
336585 records sent, 67317.0 records/sec (32.10 MB/sec), 1832.4 ms avg latency, 1923.0 max latency.
338700 records sent, 67740.0 records/sec (32.30 MB/sec), 1812.2 ms avg latency, 1888.0 max latency.
337725 records sent, 67545.0 records/sec (32.21 MB/sec), 1801.0 ms avg latency, 1843.0 max latency.
2000000 records sent, 61671.292014 records/sec (29.41 MB/sec), 1884.77 ms avg latency, 2997.00 ms max latency, 1819 ms 50th, 2801 ms 95th, 2958 ms 99th, 2993 ms 99.9th.


#### Test to consume

##### 1 thread

`bin/kafka-consumer-perf-test.sh --topic test --broker-list kz1:9092,kz2:9092,kz3:9092 --messages 1000000 --threads 1`

```
start.time            2019-11-29 05:32:45:830
end.time              2019-11-29 05:32:49:047
data.consumed.in.MB   954.0787
MB.sec                296.5740
data.consumed.in.nMsg 1000424
nMsg.sec              310980.4165
rebalance.time.ms     14
fetch.time.ms         3203
fetch.MB.sec          297.8703
fetch.nMsg.sec        312339.6815
```
298 fetch.MB.sec

##### 2 threads
`bin/kafka-consumer-perf-test.sh --topic test --broker-list kz1:9092,kz2:9092,kz3:9092 --messages 1000000 --threads 2`
```
start.time              2019-11-2905:33:29:432
end.time                2019-11-2905:33:32:595
data.consumed.in.MB	    954.0787
MB.sec	                301.6373
data.consumed.in.nMsg   1000424
nMsg.sec                316289.5985
rebalance.time.ms       12
fetch.time.ms           3151
fetch.MB.sec            302.7860
fetch.nMsg.sec          317494.1288
```
302 fetch.MB.sec

##### 10 threads

`bin/kafka-consumer-perf-test.sh --topic test --broker-list kz1:9092,kz2:9092,kz3:9092 --messages 1000000 --threads 10`
```
start.time             2019-11-2905:35:07:375
end.time               2019-11-2905:35:10:470
data.consumed.in.MB    954.0787
MB.sec                 308.2645
data.consumed.in.nMsg  1000424
nMsg.sec               323238.7722
rebalance.time.ms      13
fetch.time.ms          3082
fetch.MB.sec           309.5648
fetch.nMsg.sec         324602.2064
```
309 fetch.MB.sec

-> Number of threads to consume is not affecting much to read performance.


### Test for end-to-end latency with Python Script

#### Code
Library : https://github.com/dpkp/kafka-python
(Work in process, See kafka_consumer.py and kafka_producer.py below)

#### Result


Alternative: https://github.com/confluentinc/confluent-kafka-python


# Pulsar


## Setup
kz1: zookeeper1+bookie1+broker1
kz2: zookeeper2+bookie2+broker3
kz3: zookeeper3+bookie2+broker3
(proxy ommited to simulate the same environment as Kafka)

## zookeeper settings
Keeping default settings:
```
initLimit=10
syncLimit=5
tickTime=2000
autopurge.snapRetainCount=3
autopurge.purgeInterval=1
forceSync=yes
```
(end-point settings omitted)


## Bookie Settings
`zkServers=kz1:2181,kz2:2181,kz3:2181`

## Pulsar Brokers

```
#Platform specific
zookeeperServers=kz1:2181,kz2:2181,kz3:2181
configurationStoreServers=kz1:2181,kz2:2181,kz3:2181
clusterName=pulsar-cluster-1
brokerServicePort=6650
brokerServicePortTls=
webServicePort=8080
webServicePortTls=
```
others keeps default to: https://pulsar.apache.org/docs/en/reference-configuration/


PULSAR_EXTRA_OPTS="-Dstats_server_port=8001" bin/pulsar-daemon start zookeeper
bin/bookkeeper bookie


## Cluster init

```
bin/pulsar initialize-cluster-metadata \
  --cluster pulsar-cluster-1 \
  --zookeeper kz1:2181 \
  --configuration-store kz1:2181 \
  --web-service-url http://kz1:8080 \
  --web-service-url-tls https://kz1:8443 \
  --broker-service-url pulsar://kz1:6650 \
  --broker-service-url-tls pulsar+ssl://kz1:6651
```

## Test method

### Built-in pulsar-perf suite

#### Producing

`./bin/pulsar-perf produce  persistent://public/default/test1 -s 1000 -r 200000`

##### Produce 1000-byte message

```
Throughput produced:  37553.7  msg/s ---    286.5 Mbit/s --- 
Latency: mean:  23.967 ms - med:  22.925 - 95pct:  28.053 - 99pct:  35.567 - 99.9pct: 141.108 - 99.99pct: 142.199 - Max: 150.105

```

##### Produce 500-byte message

`./bin/pulsar-perf produce  persistent://public/default/test1 -s 500 -r 200000`
```
 Throughput produced:  38951.2  msg/s ---    148.6 Mbit/s --- 
 Latency: mean:  23.462 ms - med:  21.525 - 95pct:  26.235 - 99pct: 128.789 - 99.9pct: 166.911 - 99.99pct: 168.838 - Max: 176.004

```

#### Consuming

Tried to pile up messages and being consumed later:

```
.PerformanceConsumer - Throughput received: 77991.839  msg/s -- 595.031 Mbit/s --- Latency: mean: 63466.319 ms - med: 62734 - 95pct: 71514 - 99pct: 71657 - 99.9pct: 72936 - 99.99pct: 72978 - Max: 72979
07:51:23.351 [main] INFO  org.apache.pulsar.testclient.PerformanceConsumer - Throughput received: 86224.553  msg/s -- 657.841 Mbit/s --- Latency: mean: 45653.488 ms - med: 45126 - 95pct: 53074 - 99pct: 55004 - 99.9pct: 55108 - 99.99pct: 55122 - Max: 55122
07:51:33.364 [main] INFO  org.apache.pulsar.testclient.PerformanceConsumer - Throughput received: 88383.454  msg/s -- 674.312 Mbit/s --- Latency: mean: 27658.546 ms - med: 29003 - 95pct: 37446 - 99pct: 38040 - 99.9pct: 38150 - 99.99pct: 38156 - Max: 38156
07:51:43.374 [main] INFO  org.apache.pulsar.testclient.PerformanceConsumer - Throughput received: 68825.974  msg/s -- 525.101 Mbit/s --- Latency: mean: 6608.969 ms - med: 6760 - 95pct: 14509 - 99pct: 15034 - 99.9pct: 15162 - 99.99pct: 15193 - Max: 15193
07:51:53.381 [main] INFO  org.apache.pulsar.testclient.PerformanceConsumer - Throughput received: 33162.197  msg/s -- 253.007 Mbit/s --- Latency: mean: 37.874 ms - med: 30 - 95pct: 36 - 99pct: 297 - 99.9pct: 911 - 99.99pct: 912 - Max: 912
```


It experiences about 600Mbit/s and chase back to the write rate of the producer.


### Latency

Kafka: End-to-End about 20ms (for one consumer and one producer connecting different brokers)
Pulsar: 