artillery run artillery/ws.yml

# Test Summary

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