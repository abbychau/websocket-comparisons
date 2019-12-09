





# Scenarios
1. long Server-Pushes : 5000 consumer.
2. Echo Server: 5000 consumer.

# Library to be tested

1. 


# Test Settings
```
+-------------------+      +----------------------------------+
|                   |      |                                  |
| AWS EC2 r4.xlarge |2500  |                                  |
|                   +------>                                  |
| (Client)          |      |                                  |
+-------------------+      |                                  |
                           |      AWS EC2 r4.xlarge           |
+-------------------+      |      (Server)                    |
|                   |      |                                  |
| AWS EC2 r4.xlarge |2500  |                                  |
|                   +------>                                  |
| (Client)          |      |                                  |
+-------------------+      +----------------------------------+
```

# Test Summary


## long Server-Pushes 

This test is done by mimicing the following setting:

1. There are 5000 ws connections invoked from server p1.
2. p2 is the server, after recving a heartbeat signal, it will pushes 10 websocket send to the invoker.
3. p1 is invoking p2 every second and check if there are 10 messages.
4. Every message is 1000-byte in size excluding websocket headers.
5. Capture the upper limit of clients that could be supported without runtime error


```

+--Every Second--------------------+
|                                  |
|  +---------+ ack    +---------+  |
|  |         +-------->         |  |
|  | client  |        | Server  |  |
|  |         <--------+         |  |
|  |         |1KB*10  |         |  |
|  |         <--------+         |  |
|  |         <--------|         |  |
|  |         <--------|         |  |
|  |         <--------|         |  |
|  |         <--------|         |  |
|  |         <--------|         |  |
|  |         <--------|         |  |
|  |         <--------|         |  |
|  |         <--------+         |  |
|  |         |        |         |  |
|  |         |        |         |  |
|  +---------+        +---------+  |
|                                  |
+----------------------------------+

```


- Gorilla (Go): 7432

- ws (node)
- Java-WebSocket
- Rust-WebSocket


## Echo Server


This test is done with the following setting:

1. 5000 virtual clients are invoked to server p2 by performance tool `artillery`
2. check if all codes are 0, if yes, take the Request latency
3. Server only return one message

> npm install -g artillery

artillery run artillery/ws.yml


## Gorilla
> go run server.go
```
Summary report @ 11:12:25(+0900) 2019-12-05
  Scenarios launched:  5000
  Scenarios completed: 5000
  Requests completed:  5000
  RPS sent: 897.67
  Request latency:
    min: 0
    max: 2.7
    median: 0
    p95: 0.1
    p99: 0.1
  Scenario counts:
    0: 5000 (100%)
  Codes:
    0: 5000
```

## ws
> node server.js

```
Summary report @ 11:08:42(+0900) 2019-12-05
  Scenarios launched:  5000
  Scenarios completed: 5000
  Requests completed:  5000
  RPS sent: 897.67
  Request latency:
    min: 0
    max: 1.4
    median: 0
    p95: 0.1
    p99: 0.1
  Scenario counts:
    0: 5000 (100%)
  Codes:
    0: 5000
```

## Java-WebSocket
> mvn exec:java -Dexec.mainClass="md.abby.testapp.App"

```
All virtual users finished
Summary report @ 13:03:13(+0900) 2019-12-05
  Scenarios launched:  5000
  Scenarios completed: 5000
  Requests completed:  5000
  RPS sent: 868.06
  Request latency:
    min: 0
    max: 3.4
    median: 0.1
    p95: 0.1
    p99: 0.2
  Scenario counts:
    0: 5000 (100%)
  Codes:
    0: 5000
```

## Rust-WebSocket
> cargo run

```
Summary report @ 14:53:47(+0900) 2019-12-05
  Scenarios launched:  5000
  Scenarios completed: 5000
  Requests completed:  5000
  RPS sent: 897.67
  Request latency:
    min: 0
    max: 1.9
    median: 0
    p95: 0.1
    p99: 0.2
  Scenario counts:
    0: 5000 (100%)
  Codes:
    0: 5000
```

## Netty


`mvn exec:java -Dexec.mainClass="net.netty.websocket.echo.EchoServer"`
