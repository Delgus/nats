# Example 1

start with one subscriber, one queuesubscriber and publisher
```bash
docker-compose up -d --build
```

### Subscribe

scale subscribers to 3

```bash
docker-compose up -d --scale sub_order=3
```

logs for all containers `sub_order`  

```bash
docker-compose logs -f sub_order
```

publish message

```bash
docker-compose exec pub_order ./app
```

all services get message

```
sub_order_2          | 2021/09/09 11:21:22 get data  "Hello World" 
sub_order_1          | 2021/09/09 11:21:22 get data  "Hello World" 
sub_order_3          | 2021/09/09 11:21:22 get data  "Hello World" 
```

### QueueSubscribe

scale queue subscribers to 3

```bash
docker-compose up -d --scale queue_sub_order=3
```

logs for all containers `queue_sub_order`

```bash
docker-compose logs -f queue_sub_order
```

publish message

```bash
docker-compose exec pub_order ./app
```

only one instance get message

### At most once delivery

stopped service

```bash
docker-compose stop sub_order
```

publish message

```bash
docker-compose exec pub_order ./app
```

started service

```bash
docker-compose start sub_order
```

message lost

```bash
docker-compose logs -f sub_order
```



