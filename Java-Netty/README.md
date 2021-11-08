# Echo
A simple websocket application with echo function only. this websocket application is developed by netty.

# Usage

setup local kafka:
```
docker run --rm -p 2181:2181 -p 3030:3030 -p 8081-8083:8081-8083 \
        -p 9581-9585:9581-9585 -p 9092:9092 -e ADV_HOST=localhost \
        landoop/fast-data-dev:latest
```

setup java:
```
mvn compile
mvn exec:java
```
